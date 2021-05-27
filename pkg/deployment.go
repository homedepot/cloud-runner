package cloudrunner

import "time"

// DeploymentDescription is a Cloud Run deployment request that describes
// the `gcloud run deploy` command to build.
type DeploymentDescription struct {
	Account              string `json:"account"`
	AllowUnauthenticated bool   `json:"allowUnauthenticated"`
	ID                   string `json:"id"`
	Image                string `json:"image"`
	MaxInstances         int    `json:"maxInstances"`
	Memory               string `json:"memory"`
	Region               string `json:"region"`
	Service              string `json:"service"`
	VPCConnector         string `json:"vpcConnector"`
}

// Deployment describes a Cloud Run deployment.
//
// I was a bit confused as to why creating this table was failing in GCP.
// There are two timestamp columns defined below, and in GCP it is required
// that timestamp columns MUST have a DEFAUL:CURRENT_TIMESTAMP. It was failing
// to create because StartTime was not defining its default to CURRENT_TIMESTAMP,
// but was not throwing the same error for EndTime.
//
// That is because...
//
// The first TIMESTAMP column in a table, if not declared with the NULL
// attribute or an explicit DEFAULT or ON UPDATE clause,
// is automatically assigned the DEFAULT CURRENT_TIMESTAMP and
// ON UPDATE CURRENT_TIMESTAMP attributes.
//
// See https://stackoverflow.com/a/39547074/3878752.
type Deployment struct {
	EndTime   *time.Time `json:"endTime,omitempty" gorm:"type:timestamp;DEFAULT:current_timestamp"`
	ID        string     `json:"id"`
	StartTime *time.Time `json:"startTime,omitempty" gorm:"type:timestamp;DEFAULT:current_timestamp"`
	Status    string     `json:"status"`
	// Column type MEDIUMTEXT holds ~16MB of data.
	Output  string `json:"output,omitempty" gorm:"type:mediumtext"`
	Command string `json:"command,omitempty" gorm:"-"`
}
