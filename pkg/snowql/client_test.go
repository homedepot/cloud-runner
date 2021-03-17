package snowql_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	. "github.homedepot.com/cd/cloud-runner/pkg/snowql"
)

var _ = Describe("Client", func() {
	var (
		server *ghttp.Server
		client Client
		err    error
		lcp    string
	)

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = NewClient()
		client.WithURL(server.URL())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("#GetLCP", func() {
		JustBeforeEach(func() {
			lcp, err = client.GetLCP("fake-project-id", "fake.bearer.token")
		})

		When("the uri is invalid", func() {
			BeforeEach(func() {
				client.WithURL("::haha")
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("snowql: error creating request: parse \"::haha/graphql\": missing protocol scheme"))
			})
		})

		When("the server is not reachable", func() {
			BeforeEach(func() {
				server.Close()
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})

		When("the response is not 2XX", func() {
			BeforeEach(func() {
				server.AppendHandlers(
					ghttp.RespondWith(http.StatusInternalServerError, nil),
				)
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("error getting LCP: 500 Internal Server Error"))
			})
		})

		When("the server returns bad data", func() {
			BeforeEach(func() {
				server.AppendHandlers(
					ghttp.RespondWith(http.StatusOK, ";{["),
				)
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("snowql: error unmarshaling response body: " +
					"invalid character ';' looking for beginning of value"))
			})
		})

		When("the project ID cannot be found", func() {
			BeforeEach(func() {
				res := `{
					"data": {
						"application": null
					},
					"extensions": {
						"runTime": 3225,
						"hasCachedData": false,
						"lengths": {
							"application": 0
						},
						"records": {
							"Total": 0,
							"FromCache": 0
						},
						"dbQueries": {
							"Total": 1,
							"v_cmdb_ci_appl": 1
						}
					}
				}`

				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPost, "/graphql"),
					ghttp.RespondWith(http.StatusOK, res),
				))
			})

			It("succeeds", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("project ID 'fake-project-id' not found"))
			})
		})

		When("SnowQL returns an LCP", func() {
			BeforeEach(func() {
				res := `{
					"data": {
						"application": {
							"lifecyclePhase": "PR"
						}
					},
					"extensions": {
						"runTime": 924,
						"hasCachedData": false,
						"lengths": {
							"application": 1
						},
						"records": {
							"Total": 1,
							"FromCache": 0,
							"Application": 1
						},
						"dbQueries": {
							"Total": 1,
							"v_cmdb_ci_appl": 1
						}
					}
				}`

				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPost, "/graphql"),
					ghttp.RespondWith(http.StatusOK, res),
				))
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
				Expect(lcp).To(Equal("PR"))
			})
		})
	})
})
