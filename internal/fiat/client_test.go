package fiat_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	. "github.homedepot.com/cd/cloud-runner/internal/fiat"
)

var _ = Describe("Client", func() {
	var (
		server *ghttp.Server
		client Client
		err    error
		roles  Roles
	)

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = NewClient()
		client.WithURL(server.URL())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("#NewDefaultClient", func() {
		BeforeEach(func() {
			client = NewDefaultClient()
		})

		It("succeeds", func() {
		})
	})

	Describe("#Authorize", func() {
		JustBeforeEach(func() {
			roles, err = client.Roles("fakeAccount")
		})

		When("the uri is invalid", func() {
			BeforeEach(func() {
				client.WithURL("::haha")
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("fiat: error creating request: parse " +
					"\"::haha/authorize/fakeAccount/roles\": missing protocol scheme"))
			})
		})

		When("the server is not reachable", func() {
			BeforeEach(func() {
				server.Close()
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(HaveSuffix("connect: connection refused"))
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
				Expect(err.Error()).To(Equal("fiat: error getting account roles: 500 Internal Server Error"))
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
				Expect(err.Error()).To(Equal("fiat: error unmarshaling roles: " +
					"invalid character ';' looking for beginning of value"))
			})
		})

		When("it succeeds", func() {
			BeforeEach(func() {
				r := `[ {
					  "name" : "gg_cloud_gcp_spinnaker_admins",
					  "source" : "EXTERNAL"
					}, {
					  "name" : "test_group",
					  "source" : "EXTERNAL"
					} ]`

				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/authorize/fakeAccount/roles"),
					ghttp.RespondWith(http.StatusOK, r),
				))
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
				Expect(roles).To(HaveLen(2))
				Expect(roles[0].Name).To(Equal("gg_cloud_gcp_spinnaker_admins"))
				Expect(roles[0].Source).To(Equal("EXTERNAL"))
				Expect(roles[1].Name).To(Equal("test_group"))
				Expect(roles[1].Source).To(Equal("EXTERNAL"))
			})
		})
	})
})
