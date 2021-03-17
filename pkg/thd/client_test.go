package thd_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	. "github.homedepot.com/cd/cloud-runner/pkg/thd"
)

var _ = Describe("IdentityClient", func() {
	var (
		server *ghttp.Server
		client IdentityClient
		err    error
		token  string
	)

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = NewIdentityClient()
		client.WithTokenEndpoint(server.URL())
		client.WithClientID("fake-client-id")
		client.WithClientSecret("fake-client-secret")
		client.WithResource("fake-resource")
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("#Token", func() {
		JustBeforeEach(func() {
			token, err = client.Token()
		})

		When("the uri is invalid", func() {
			BeforeEach(func() {
				client.WithTokenEndpoint("::haha")
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("error making request for THD IDP: parse \"::haha\": missing protocol scheme"))
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
				Expect(err.Error()).To(Equal("error getting token from THD IDP: 500 Internal Server Error"))
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
				Expect(err.Error()).To(Equal("error unmarshaling body from THD IDP: " +
					"invalid character ';' looking for beginning of value"))
			})
		})

		When("the server returns a descriptive error", func() {
			BeforeEach(func() {
				res := `{
						"error_description": "Error - requested resource not allowed",
						"error": "invalid_grant"
					}`

				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPost, "/"),
					ghttp.RespondWith(http.StatusBadRequest, res),
				))
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("error getting token from THD IDP: Error - requested resource not allowed"))
			})
		})

		When("the token is cached", func() {
			BeforeEach(func() {
				res := `{
						"access_token": "fake.bearer.token.cached",
						"token_type": "Bearer",
						"expires_in": 3599
				  }`

				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPost, "/"),
					ghttp.RespondWith(http.StatusOK, res),
				))
			})

			JustBeforeEach(func() {
				token, _ = client.Token()
			})

			It("returns the cached token", func() {
				Expect(err).To(BeNil())
				Expect(token).To(Equal("fake.bearer.token.cached"))
			})
		})

		When("THD IDP returns a token", func() {
			BeforeEach(func() {
				res := `{
						"access_token": "fake.bearer.token",
						"token_type": "Bearer",
						"expires_in": 3599
				  }`

				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPost, "/"),
					ghttp.RespondWith(http.StatusOK, res),
				))
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
				Expect(token).To(Equal("fake.bearer.token"))
			})
		})
	})
})
