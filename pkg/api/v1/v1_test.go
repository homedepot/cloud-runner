package v1_test

import (
	"bytes"
	"io/ioutil"
	"mime"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	cloudrunner "github.homedepot.com/cd/cloud-runner/pkg"
	"github.homedepot.com/cd/cloud-runner/pkg/api"
	"github.homedepot.com/cd/cloud-runner/pkg/fiat"
	"github.homedepot.com/cd/cloud-runner/pkg/fiat/fiatfakes"
	"github.homedepot.com/cd/cloud-runner/pkg/snowql/snowqlfakes"
	"github.homedepot.com/cd/cloud-runner/pkg/sql/sqlfakes"
	"github.homedepot.com/cd/cloud-runner/pkg/thd/thdfakes"

	// . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	err                   error
	svr                   *httptest.Server
	uri                   string
	req                   *http.Request
	body                  *bytes.Buffer
	res                   *http.Response
	fakeFiatClient        *fiatfakes.FakeClient
	fakeSnowQLClient      *snowqlfakes.FakeClient
	fakeSQLClient         *sqlfakes.FakeClient
	fakeTHDIdentityClient *thdfakes.FakeIdentityClient
)

func setup() {
	// Setup fake Fiat client.
	fakeFiatClient = &fiatfakes.FakeClient{}
	fakeFiatClient.RolesReturns(fiat.Roles{{Name: "gg_test"}}, nil)

	// Setup fake SnowQL Client.
	fakeSnowQLClient = &snowqlfakes.FakeClient{}
	fakeSnowQLClient.GetLCPReturns("PR", nil)

	// Setup fake THD Identity Client.
	fakeTHDIdentityClient = &thdfakes.FakeIdentityClient{}

	// Setup fake SQL client.
	fakeSQLClient = &sqlfakes.FakeClient{}
	fakeSQLClient.GetCredentialsReturns(
		cloudrunner.Credentials{
			Account:   "cr-test-project-id-pr",
			Lifecycle: "pr",
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
	fakeSQLClient.ListCredentialsReturns(
		[]cloudrunner.Credentials{
			{
				Account:   "cr-test-project-id-pr",
				Lifecycle: "pr",
				ProjectID: "test-project-id",
				ReadGroups: []string{
					"gg_test",
				},
				WriteGroups: []string{
					"gg_test",
				},
			},
			{
				Account:   "cr-test-project-id2-pr",
				Lifecycle: "pr",
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

	api.WithKey("test")

	// Create new gin instead of using gin.Default().
	// This disables request logging which we don't want for tests.
	e := gin.New()
	e.Use(gin.Recovery())

	s := api.NewServer()
	s.WithEngine(e)
	s.WithFiatClient(fakeFiatClient)
	s.WithSnowQLClient(fakeSnowQLClient)
	s.WithSQLClient(fakeSQLClient)
	s.WithTHDIdentityClient(fakeTHDIdentityClient)
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
