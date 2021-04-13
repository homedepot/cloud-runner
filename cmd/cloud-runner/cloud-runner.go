package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/homedepot/cloud-runner/internal/api"
	"github.com/homedepot/cloud-runner/internal/fiat"
	"github.com/homedepot/cloud-runner/internal/gcloud"
	"github.com/homedepot/cloud-runner/internal/sql"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
)

const (
	banner = `        __                __
  _____/ /___  __  ______/ /  _______  ______  ____  ___  _____
 / ___/ / __ \/ / / / __  /  / ___/ / / / __ \/ __ \/ _ \/ ___/
/ /__/ / /_/ / /_/ / /_/ /  / /  / /_/ / / / / / / /  __/ /
\___/_/\____/\__,_/\__,_/  /_/   \__,_/_/ /_/_/ /_/\___/_/
                                                               `
)

func main() {
	fmt.Println(banner)

	r := gin.New()
	// Make the logs use color.
	gin.ForceConsoleColor()

	// Set the API key for create/delete operations.
	api.WithKey(os.Getenv("API_KEY"))

	// Setup metrics.
	p := ginprometheus.NewPrometheus("cloud_runner")
	p.MetricsPath = "/metrics"
	p.Use(r)

	// Preserve low cardinality for the request counter.
	// See https://github.com/zsais/go-gin-prometheus#preserving-a-low-cardinality-for-the-request-counter.
	p.ReqCntURLLabelMappingFn = reqCntURLLabelMappingFn

	// Ignore logging of certain endpoints.
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/healthz"}}))
	r.Use(gin.Recovery())

	// Setup Fiat Client.
	fiatClient := fiat.NewDefaultClient()

	// Setup SQL Client.
	sqlClient := sql.NewClient()
	sqlClient.WithHost(os.Getenv("SQL_HOST"))
	sqlClient.WithName(os.Getenv("SQL_NAME"))
	sqlClient.WithPass(os.Getenv("SQL_PASS"))
	sqlClient.WithUser(os.Getenv("SQL_USER"))

	err := sqlClient.Connect(sqlClient.Connection())
	if err != nil {
		panic(err)
	}

	// Setup the server.
	server := api.NewServer()
	server.WithEngine(r)
	server.WithBuilder(gcloud.NewCloudRunCommandBuilder())
	server.WithFiatClient(fiatClient)
	server.WithSQLClient(sqlClient)
	server.Setup()

	err = r.Run(":80")
	if err != nil {
		panic(err)
	}
}

func reqCntURLLabelMappingFn(c *gin.Context) string {
	url := c.Request.URL.Path

	for _, p := range c.Params {
		if p.Key == "account" || p.Key == "deploymentID" {
			url = strings.Replace(url, p.Value, ":"+p.Key, 1)
		}
	}

	return url
}
