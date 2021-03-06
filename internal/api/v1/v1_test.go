package v1_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	internal "github.com/homedepot/cloud-runner/internal"
	"github.com/homedepot/cloud-runner/internal/api"
	"github.com/homedepot/cloud-runner/internal/fiat"
	"github.com/homedepot/cloud-runner/internal/fiat/fiatfakes"
	"github.com/homedepot/cloud-runner/internal/gcloud/gcloudfakes"
	"github.com/homedepot/cloud-runner/internal/sql/sqlfakes"
	cloudrunner "github.com/homedepot/cloud-runner/pkg"

	// . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	err            error
	svr            *httptest.Server
	uri            string
	req            *http.Request
	body           *bytes.Buffer
	res            *http.Response
	fakeBuilder    *gcloudfakes.FakeCloudRunCommandBuilder
	fakeCommand    *gcloudfakes.FakeCloudRunCommand
	fakeFiatClient *fiatfakes.FakeClient
	fakeSQLClient  *sqlfakes.FakeClient
)

func setup() {
	// Ignore any logs.
	log.SetOutput(ioutil.Discard)

	// Setup fake builder.
	fakeBuilder = &gcloudfakes.FakeCloudRunCommandBuilder{}
	fakeCommand = &gcloudfakes.FakeCloudRunCommand{}
	// This is a not so great side-effect of using the builder method.
	fakeBuilder.AllowUnauthenticatedReturns(fakeBuilder)
	fakeBuilder.ImageReturns(fakeBuilder)
	fakeBuilder.MaxInstancesReturns(fakeBuilder)
	fakeBuilder.MemoryReturns(fakeBuilder)
	fakeBuilder.ProjectIDReturns(fakeBuilder)
	fakeBuilder.RegionReturns(fakeBuilder)
	fakeBuilder.ServiceReturns(fakeBuilder)
	fakeBuilder.VPCConnectorReturns(fakeBuilder)
	fakeBuilder.BuildReturns(fakeCommand, nil)

	// Setup fake Fiat client.
	fakeFiatClient = &fiatfakes.FakeClient{}
	fakeFiatClient.RolesReturns(fiat.Roles{{Name: "gg_test"}}, nil)

	// Setup fake SQL client.
	fakeSQLClient = &sqlfakes.FakeClient{}
	fakeSQLClient.GetCredentialsReturns(
		cloudrunner.Credentials{
			Account:   "cr-test-project-id",
			ProjectID: "test-project-id",
			ReadGroups: []string{
				"gg_test",
			},
			WriteGroups: []string{
				"gg_test",
			},
		},
		nil,
	)
	t := internal.CurrentTimeUTC()
	fakeDeployment := cloudrunner.Deployment{
		EndTime:   &t,
		ID:        "fake-id",
		StartTime: &t,
		Status:    "fake-status",
		Output:    "fake-output",
	}
	fakeSQLClient.GetDeploymentReturns(
		fakeDeployment,
		nil,
	)
	fakeSQLClient.ListCredentialsReturns(
		[]cloudrunner.Credentials{
			{
				Account:   "cr-test-project-id",
				ProjectID: "test-project-id",
				ReadGroups: []string{
					"gg_test",
				},
				WriteGroups: []string{
					"gg_test",
				},
			},
			{
				Account:   "cr-test-project-id2",
				ProjectID: "test-project-id2",
				ReadGroups: []string{
					"gg_test2",
				},
				WriteGroups: []string{
					"gg_test2",
				},
			},
		},
		nil,
	)

	// Disable debug logging.
	gin.SetMode(gin.ReleaseMode)

	// Create new gin instead of using gin.Default().
	// This disables request logging which we don't want for tests.
	e := gin.New()
	e.Use(gin.Recovery())

	s := api.NewServer()
	s.WithAPIKey("test")
	s.WithEngine(e)
	s.WithBuilder(fakeBuilder)
	s.WithFiatClient(fakeFiatClient)
	s.WithSQLClient(fakeSQLClient)
	s.Setup()

	// Create server.
	svr = httptest.NewServer(e)
	body = &bytes.Buffer{}
}

func teardown() {
	svr.Close()

	mt, mtp, _ := mime.ParseMediaType(res.Header.Get("content-type"))
	Expect(mt).To(Equal("application/json"), "content-type")
	Expect(mtp["charset"]).To(Equal("utf-8"), "charset")
	res.Body.Close()
}

func createRequest(method string) {
	req, _ = http.NewRequest(method, uri, ioutil.NopCloser(body))
	req.Header.Set("API-Key", "test")
	req.Header.Set("X-Spinnaker-User", "test-account")
}

func doRequest() {
	res, err = http.DefaultClient.Do(req)
}

func validateResponse(expected string) {
	actual, _ := ioutil.ReadAll(res.Body)
	Expect(actual).To(MatchJSON(expected), "correct body")
}
