package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
	"github.homedepot.com/cd/cloud-runner/pkg/api"
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

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("[CLOUD RUNNER] WARNING: API_KEY not set")
	}

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

	api.Init(r)

	err := r.Run(":80")
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
