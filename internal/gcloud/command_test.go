package gcloud_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/homedepot/cloud-runner/internal/gcloud"
)

var _ = Describe("Command", func() {
	var (
		c   CloudRunCommandBuilder
		cmd CloudRunCommand
		err error
	)

	BeforeEach(func() {
		c = NewCloudRunCommandBuilder()
		c.AllowUnauthenticated(true).
			Image("gcr.io/not-a-project13765/not-an-image:9.0.4").
			ProjectID("fake-project-id").
			Region("fake-region").
			Service("fake-service")
	})

	Describe("#Build", func() {
		JustBeforeEach(func() {
			cmd, err = c.Build()
		})

		When("validating the project ID fails", func() {
			BeforeEach(func() {
				c.ProjectID("sdfSDFfwefwe_")
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("error validating project ID: sdfSDFfwefwe_ failed validation"))
			})
		})

		When("validating the service fails", func() {
			BeforeEach(func() {
				c.Service("welkjJKFWE-")
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("error validating service name: welkjJKFWE- failed validation"))
			})
		})

		When("validating the region fails", func() {
			BeforeEach(func() {
				c.Region("jwflkWEJFKWEJ-_")
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("error validating region: jwflkWEJFKWEJ-_ failed validation"))
			})
		})

		When("allow unauthenticated is false", func() {
			BeforeEach(func() {
				c.AllowUnauthenticated(false)
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
				Expect(cmd.String()).To(HaveSuffix("gcloud run deploy fake-service " +
					"--project fake-project-id " +
					"--image gcr.io/not-a-project13765/not-an-image:9.0.4 " +
					"--platform managed " +
					"--region fake-region " +
					"--no-allow-unauthenticated"))
			})
		})

		When("max instances is set", func() {
			BeforeEach(func() {
				c.MaxInstances(4)
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
				Expect(cmd.String()).To(HaveSuffix("gcloud run deploy fake-service " +
					"--project fake-project-id " +
					"--image gcr.io/not-a-project13765/not-an-image:9.0.4 " +
					"--platform managed " +
					"--region fake-region " +
					"--allow-unauthenticated " +
					"--max-instances 4"))
			})
		})

		Context("when memory is set", func() {
			BeforeEach(func() {
				c.Memory("1G")
			})

			When("the memory is valid", func() {
				It("succeeds", func() {
					Expect(err).To(BeNil())
					Expect(cmd.String()).To(HaveSuffix("gcloud run deploy fake-service " +
						"--project fake-project-id " +
						"--image gcr.io/not-a-project13765/not-an-image:9.0.4 " +
						"--platform managed " +
						"--region fake-region " +
						"--allow-unauthenticated " +
						"--memory 1G"))
				})
			})
		})

		Context("when vpc connector is set", func() {
			BeforeEach(func() {
				c.VPCConnector("fake-vpc-connector")
			})

			When("the vpc connector is valid", func() {
				It("succeeds", func() {
					Expect(err).To(BeNil())
					Expect(cmd.String()).To(HaveSuffix("gcloud run deploy fake-service " +
						"--project fake-project-id " +
						"--image gcr.io/not-a-project13765/not-an-image:9.0.4 " +
						"--platform managed " +
						"--region fake-region " +
						"--allow-unauthenticated " +
						"--vpc-connector fake-vpc-connector"))
				})
			})
		})

		When("the command is valid", func() {
			It("succeeds", func() {
				Expect(err).To(BeNil())
				Expect(cmd.String()).To(HaveSuffix("gcloud run deploy fake-service " +
					"--project fake-project-id " +
					"--image gcr.io/not-a-project13765/not-an-image:9.0.4 " +
					"--platform managed " +
					"--region fake-region " +
					"--allow-unauthenticated"))
			})
		})
	})
})
