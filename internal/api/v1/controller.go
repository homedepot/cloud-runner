package v1

import (
	"github.homedepot.com/cd/cloud-runner/internal/fiat"
	"github.homedepot.com/cd/cloud-runner/internal/gcloud"
	"github.homedepot.com/cd/cloud-runner/internal/sql"
)

type Controller struct {
	FiatClient fiat.Client
	SqlClient  sql.Client
	Builder    gcloud.CloudRunCommandBuilder
}
