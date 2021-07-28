package v1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	internal "github.com/homedepot/cloud-runner/internal"
	"github.com/homedepot/cloud-runner/internal/gcloud"
	"github.com/homedepot/cloud-runner/internal/sql"
	cloudrunner "github.com/homedepot/cloud-runner/pkg"
)

var (
	statusFailed    = "FAILED"
	statusRunning   = "RUNNING"
	statusSucceeded = "SUCCEEDED"
)

// CreateDeployment generates and runs a `gcloud run deploy` command.
func (cc *Controller) CreateDeployment(c *gin.Context) {
	dd := cloudrunner.DeploymentDescription{}

	err := c.ShouldBindJSON(&dd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	// This could be a middleware. If it bothers you create a story and do so :).
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

	credentials, err := cc.SqlClient.GetCredentials(dd.Account)
	if err != nil {
		if err == gorm.ErrRecordNotFound || err == sql.ErrCredentialsNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "credentials not found"})

			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}
	}

	// Check if the user has r/w access to use the account. If 'credentials'
	// gets filtered down to an empty slice, they do not.
	creds := filterCredentials([]cloudrunner.Credentials{credentials}, roles)
	if len(creds) == 0 {
		c.JSON(http.StatusForbidden,
			gin.H{"error": fmt.Sprintf("user %s does not have access to use account %s", user, dd.Account)})

		return
	}

	// Build the command.
	cmd, err := cc.buildCommand(dd, credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	t := internal.CurrentTimeUTC()
	d := cloudrunner.Deployment{
		Command:   cmd.String(),
		ID:        uuid.New().String(),
		StartTime: &t,
		Status:    statusRunning,
	}

	err = cc.SqlClient.CreateDeployment(d)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	// We need to run the command concurrently and immediately return the
	// deployment status to the user.
	go cc.run(cmd, d)

	c.JSON(http.StatusCreated, d)
}

// buildCommand builds the `gcloud run deploy` command. We can pass all fields in here, the command
// builder will only set flags for valid inputs and return an error or ignore
// on invalid inputs.
func (cc *Controller) buildCommand(dd cloudrunner.DeploymentDescription,
	credentials cloudrunner.Credentials) (gcloud.CloudRunCommand, error) {
	return cc.Builder.
		AllowUnauthenticated(dd.AllowUnauthenticated).
		Image(dd.Image).
		MaxInstances(dd.MaxInstances).
		Memory(dd.Memory).
		ProjectID(credentials.ProjectID).
		Region(dd.Region).
		Service(dd.Service).
		VPCConnector(dd.VPCConnector).
		Build()
}

func (cc *Controller) run(cmd gcloud.CloudRunCommand, d cloudrunner.Deployment) {
	co, err := cmd.CombinedOutput()
	if err != nil {
		d.Status = statusFailed
	} else {
		d.Status = statusSucceeded
	}

	d.Output = string(co)
	t := internal.CurrentTimeUTC()
	d.EndTime = &t

	err = cc.SqlClient.UpdateDeployment(d)
	if err != nil {
		// Nothing to really do here besides add a log.
		log.Printf("error updating deployment with ID %s: %s", d.ID, err.Error())
	}
}

// GetDeployment gets a deployment from the DB by a given deployment ID.
func (cc *Controller) GetDeployment(c *gin.Context) {
	id := c.Param("deploymentID")

	d, err := cc.SqlClient.GetDeployment(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "deployment not found"})

			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}
	}

	// MySQL requires that timestamp columns have a default of CURRENT_TIMESTAMP
	// in GCP, so if the deployment is still in a running state, set the EndTime to be nil.
	if d.Status == statusRunning {
		d.EndTime = nil
	}

	c.JSON(http.StatusOK, d)
}
