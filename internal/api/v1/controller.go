package v1

import (
	"github.com/homedepot/cloud-runner/internal/fiat"
	"github.com/homedepot/cloud-runner/internal/gcloud"
	"github.com/homedepot/cloud-runner/internal/sql"
)

type Controller struct {
	FiatClient fiat.Client
	SqlClient  sql.Client
	Builder    gcloud.CloudRunCommandBuilder
}
