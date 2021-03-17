package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
	"github.homedepot.com/cd/cloud-runner/pkg/api"
	"github.homedepot.com/cd/cloud-runner/pkg/fiat"
	"github.homedepot.com/cd/cloud-runner/pkg/snowql"
	"github.homedepot.com/cd/cloud-runner/pkg/sql"
	"github.homedepot.com/cd/cloud-runner/pkg/thd"
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

	// Set the API key for create/delete operations.
	api.WithKey(os.Getenv("API_KEY"))

	// Setup metrics.
	p := ginprometheus.NewPrometheus("cloud_runner")
	p.MetricsPath = "/metrics"
	p.Use(r)

	// Preserve low cardinality for the request counter.
	// See https://github.com/zsais/go-gin-prometheus#preserving-a-low-cardinality-for-the-request-counter.
	p.ReqCntURLLabelMappingFn = reqCntURLLabelMappingFn

	// Make the logs use color.
	gin.ForceConsoleColor()
	// Ignore logging of certain endpoints.
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/healthz"}}))
	r.Use(gin.Recovery())

	// Setup Fiat Client.
	fiatClient := fiat.NewDefaultClient()

	// Setup SnowQL Client.
	snowQLClient := snowql.NewClient()
	snowQLClient.WithURL(os.Getenv("SNOWQL_URL"))

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

	// Setup THD Identity Client.
	thdIdenityClient := thd.NewIdentityClient()
	thdIdenityClient.WithClientID(os.Getenv("THD_IDENTITY_CLIENT_ID"))
	thdIdenityClient.WithClientSecret(os.Getenv("THD_IDENTITY_CLIENT_SECRET"))
	thdIdenityClient.WithResource(os.Getenv("THD_IDENTITY_RESOURCE"))
	thdIdenityClient.WithTokenEndpoint(os.Getenv("THD_IDENTITY_TOKEN_ENDPOINT"))

	// Setup the server.
	server := api.NewServer()
	server.WithEngine(r)
	server.WithFiatClient(fiatClient)
	server.WithSnowQLClient(snowQLClient)
	server.WithSQLClient(sqlClient)
	server.WithTHDIdentityClient(thdIdenityClient)

	err = server.Setup()
	if err != nil {
		panic(err)
	}

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
