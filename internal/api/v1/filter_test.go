package v1

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.homedepot.com/cd/cloud-runner/internal/fiat"
	cloudrunner "github.homedepot.com/cd/cloud-runner/pkg"
)

var _ = Describe("Filter", func() {
	var (
		c                   *Controller
		adminRoles          []string
		credentials         []cloudrunner.Credentials
		filteredCredentials []cloudrunner.Credentials
		roles               fiat.Roles
	)

	BeforeEach(func() {
		credentials = []cloudrunner.Credentials{
			{
				Account:   "cr-test",
				ProjectID: "test",
				ReadGroups: []string{
					"group1",
					"group2",
				},
				WriteGroups: []string{
					"group1",
				},
			},
			{
				Account:   "cr-test2",
				ProjectID: "test2",
				ReadGroups: []string{
					"group2",
				},
				WriteGroups: []string{
					"group2",
				},
			},
		}
		roles = fiat.Roles{
			fiat.Role{
				Name: "group1",
			},
		}
		adminRoles = []string{
			"admin_group1",
			"admin_group2",
		}
		c = &Controller{
			AdminRoles: adminRoles,
		}
	})

	JustBeforeEach(func() {
		filteredCredentials = c.filterCredentials(credentials, roles)
	})

	When("the user does not have access to any credentials", func() {
		BeforeEach(func() {
			roles = fiat.Roles{}
		})

		It("returns an empty slice of credentials", func() {
			Expect(filteredCredentials).To(HaveLen(0))
		})
	})

	When("the user is an admin", func() {
		BeforeEach(func() {
			roles = fiat.Roles{
				fiat.Role{
					Name: "group3",
				},
				fiat.Role{
					Name: "admin_group1",
				},
			}
		})

		It("returns all credentials", func() {
			Expect(filteredCredentials).To(HaveLen(2))
		})
	})

	When("it succeeds", func() {
		It("returns an empty slice of credentials", func() {
			Expect(filteredCredentials).To(HaveLen(1))
			Expect(filteredCredentials[0].Account).To(Equal("cr-test"))
			Expect(filteredCredentials[0].ProjectID).To(Equal("test"))
			Expect(filteredCredentials[0].ReadGroups).To(HaveLen(2))
			Expect(filteredCredentials[0].WriteGroups).To(HaveLen(1))
		})
	})
})
