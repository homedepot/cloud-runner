package v1

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	cloudrunner "github.homedepot.com/cd/cloud-runner/pkg"
	"github.homedepot.com/cd/cloud-runner/pkg/fiat"
	"github.homedepot.com/cd/cloud-runner/pkg/snowql"
	"github.homedepot.com/cd/cloud-runner/pkg/sql"
	"github.homedepot.com/cd/cloud-runner/pkg/thd"
)

// CreateCredentials generates a new account name for Cloud Run
// in the format `cr-<PROJECT_ID>-<LIFECYCLE>` and inserts it into the DB.
func CreateCredentials(c *gin.Context) {
	creds := cloudrunner.Credentials{}
	snowQLClient := snowql.Instance(c)
	sqlClient := sql.Instance(c)
	thdIdentityClient := thd.Instance(c)

	err := c.ShouldBindJSON(&creds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := thdIdentityClient.Token()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	lcp, err := snowQLClient.GetLCP(creds.ProjectID, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	creds.Lifecycle = strings.ToLower(lcp)
	creds.Account = fmt.Sprintf("cr-%s-%s", creds.ProjectID, creds.Lifecycle)

	_, err = sqlClient.GetCredentials(creds.Account)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "credentials already exists"})
		return
	}

	err = sqlClient.CreateCredentials(creds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, group := range creds.ReadGroups {
		rp := cloudrunner.CredentialsReadPermission{
			ID:        uuid.New().String(),
			Account:   creds.Account,
			ReadGroup: group,
		}

		err = sqlClient.CreateReadPermission(rp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	for _, group := range creds.WriteGroups {
		wp := cloudrunner.CredentialsWritePermission{
			ID:         uuid.New().String(),
			Account:    creds.Account,
			WriteGroup: group,
		}

		err = sqlClient.CreateWritePermission(wp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, creds)
}

// DeleteCredentials deletes credentials from the DB by account name.
func DeleteCredentials(c *gin.Context) {
	account := c.Param("account")
	fc := fiat.Instance(c)
	sc := sql.Instance(c)

	user := c.GetHeader("X-Spinnaker-User")
	if user == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "X-Spinnaker-User header not set"})
		return
	}

	roles, err := fc.Roles(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	credentials, err := sc.GetCredentials(account)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "credentials not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Check if the user has access to delete the account. If 'credentials'
	// gets filtered down to an empty slice, they do not.
	creds := filterCredentials([]cloudrunner.Credentials{credentials}, roles)
	if len(creds) == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("user %s does not have access to delete account %s", user, account)})
		return
	}

	err = sc.DeleteCredentials(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetCredentials gets credentials by account name.
func GetCredentials(c *gin.Context) {
	sc := sql.Instance(c)
	account := c.Param("account")

	creds, err := sc.GetCredentials(account)
	if err != nil {
		if err == gorm.ErrRecordNotFound || err == sql.ErrCredentialsNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "credentials not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, creds)
}

// ListCredentials lists all credentials. If the query param 'onlyForUser' is true,
// then grab the user from the `X-SPINNAKER-USER` header, get their groups,
// and filter accounts by read/write groups that are contained within the user's groups.
func ListCredentials(c *gin.Context) {
	fc := fiat.Instance(c)
	sc := sql.Instance(c)

	creds, err := sc.ListCredentials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If onlyForUser is true, filter the credentials according to what the user
	// has read and write access to.
	if c.Query("onlyForUser") == "true" {
		user := c.GetHeader("X-Spinnaker-User")
		if user == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "requested credentials only for user, but X-Spinnaker-User header was not set"})
			return
		}

		roles, err := fc.Roles(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		creds = filterCredentials(creds, roles)
	}

	c.JSON(http.StatusOK, gin.H{"credentials": creds})
}

// filterCredentials takes in a list of credenitals and roles. It
// returns a filtered list of credentials that contain both a read and write
// group from the respective roles slice passed in.
func filterCredentials(credentials []cloudrunner.Credentials, roles fiat.Roles) []cloudrunner.Credentials {
	c := []cloudrunner.Credentials{}

	g := map[string]bool{}

	for _, role := range roles {
		g[strings.ToLower(role.Name)] = true
	}

	for _, creds := range credentials {
		var read, write bool

		for _, readGroup := range creds.ReadGroups {
			if g[strings.ToLower(readGroup)] {
				read = true
				break
			}
		}

		for _, writeGroup := range creds.WriteGroups {
			if g[strings.ToLower(writeGroup)] {
				write = true
				break
			}
		}

		if read && write {
			c = append(c, creds)
		}
	}

	return c
}
