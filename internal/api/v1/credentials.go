package v1

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/homedepot/cloud-runner/internal/fiat"
	"github.com/homedepot/cloud-runner/internal/sql"
	cloudrunner "github.com/homedepot/cloud-runner/pkg"
)

// CreateCredentials creates a new account for Cloud Run.
// If the account field it not provided it is generated in the format
// `cr-<PROJECT_ID>`.
func (cc *Controller) CreateCredentials(c *gin.Context) {
	creds := cloudrunner.Credentials{}

	err := c.ShouldBindJSON(&creds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	// If an account name was not provided, generate one.
	if creds.Account == "" {
		creds.Account = fmt.Sprintf("cr-%s", creds.ProjectID)
	}

	_, err = cc.SqlClient.GetCredentials(creds.Account)
	if err != gorm.ErrRecordNotFound && err != sql.ErrCredentialsNotFound {
		if err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "credentials already exists"})

			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}
	}

	// Make sure credentials read groups contain all write groups.
	for _, wg := range creds.WriteGroups {
		if !contains(creds.ReadGroups, wg) {
			creds.ReadGroups = append(creds.ReadGroups, wg)
		}
	}

	err = cc.SqlClient.CreateCredentials(creds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, creds)
}

// contains returns true if slice s contains element e (case insensitive).
func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.EqualFold(a, e) {
			return true
		}
	}

	return false
}

// DeleteCredentials deletes credentials from the DB by account name.
func (cc *Controller) DeleteCredentials(c *gin.Context) {
	account := c.Param("account")

	user := c.GetHeader("X-Spinnaker-User")
	if user == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "X-Spinnaker-User header not set"})

		return
	}

	roles, err := cc.FiatClient.Roles(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	credentials, err := cc.SqlClient.GetCredentials(account)
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
		c.JSON(http.StatusForbidden,
			gin.H{"error": fmt.Sprintf("user %s does not have access to delete account %s", user, account)})

		return
	}

	err = cc.SqlClient.DeleteCredentials(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetCredentials gets credentials by account name.
func (cc *Controller) GetCredentials(c *gin.Context) {
	account := c.Param("account")

	creds, err := cc.SqlClient.GetCredentials(account)
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
func (cc *Controller) ListCredentials(c *gin.Context) {
	creds, err := cc.SqlClient.ListCredentials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	// If onlyForUser is true, filter the credentials according to what the user
	// has read and write access to.
	if c.Query("onlyForUser") == "true" {
		user := c.GetHeader("X-Spinnaker-User")
		if user == "" {
			c.JSON(http.StatusUnauthorized,
				gin.H{"error": "requested credentials only for user, but X-Spinnaker-User header was not set"})

			return
		}

		roles, err := cc.FiatClient.Roles(user)
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
