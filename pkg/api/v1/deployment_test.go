package v1_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	cloudrunner "github.homedepot.com/cd/cloud-runner/pkg"
	"github.homedepot.com/cd/cloud-runner/pkg/fiat"
	"github.homedepot.com/cd/cloud-runner/pkg/sql"
)

var _ = Describe("Deployment", func() {
	const (
		uuidLength = 36
	)

	Describe("#CreateDeployment", func() {
		BeforeEach(func() {
			setup()
			uri = svr.URL + "/v1/deployments"
			body.Write([]byte(payloadRequestDeployment))
			createRequest(http.MethodPost)
		})

		AfterEach(func() {
			teardown()
		})

		JustBeforeEach(func() {
			doRequest()
		})

		When("the request body is bad data", func() {
			BeforeEach(func() {
				body = &bytes.Buffer{}
				body.Write([]byte("dasdf[]dsf;;"))
				createRequest(http.MethodPost)
			})

			It("returns status bad request", func() {
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				validateResponse(payloadBadRequest)
			})
		})

		When("the X-Spinnaker-User header is not set", func() {
			BeforeEach(func() {
				req.Header.Set("X-Spinnaker-User", "")
			})

			It("returns status unauthorized", func() {
				Expect(res.StatusCode).To(Equal(http.StatusUnauthorized))
				validateResponse(payloadUnauthorized)
			})
		})

		When("getting the roles from fiat returns an error", func() {
			BeforeEach(func() {
				fakeFiatClient.RolesReturns(nil, errors.New("error getting roles"))
			})

			It("returns status unauthorized", func() {
				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				validateResponse(payloadErrorGettingRoles)
			})
		})

		When("the record is not found", func() {
			BeforeEach(func() {
				fakeSQLClient.GetCredentialsReturns(cloudrunner.Credentials{}, sql.ErrCredentialsNotFound)
			})

			It("returns an error", func() {
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
				validateResponse(payloadCredentialsNotFound)
			})
		})

		When("getting the credentials returns a generic error", func() {
			BeforeEach(func() {
				fakeSQLClient.GetCredentialsReturns(cloudrunner.Credentials{}, errors.New("error getting credentials"))
			})

			It("returns an error", func() {
				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				validateResponse(payloadCredentialsGetGenericError)
			})
		})

		When("the user does not have access to use the account", func() {
			BeforeEach(func() {
				fakeFiatClient.RolesReturns(fiat.Roles{{Name: "not_a_good_group"}}, nil)
			})

			It("returns an error", func() {
				Expect(res.StatusCode).To(Equal(http.StatusForbidden))
				validateResponse(payloadDeploymentsForbiddenError)
			})
		})

		When("building the command returns an error", func() {
			BeforeEach(func() {
				fakeBuilder.BuildReturns(nil, errors.New("error building command"))
			})

			It("returns an error", func() {
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				validateResponse(payloadDeploymentsErrorBuildingCommand)
			})
		})

		When("creating the deployment returns an error", func() {
			BeforeEach(func() {
				fakeSQLClient.CreateDeploymentReturns(errors.New("error creating deployment"))
			})

			It("returns an error", func() {
				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				validateResponse(payloadDeploymentsErrorCreatingDeployment)
			})
		})

		Context("concurrent command run", func() {
			When("getting the combined output returns an error", func() {
				BeforeEach(func() {
					fakeCommand.CombinedOutputReturns(nil, errors.New("error getting combined output"))
				})

				It("stores status failed", func() {
					Expect(fakeSQLClient.UpdateDeploymentCallCount()).To(Equal(1))
					d := fakeSQLClient.UpdateDeploymentArgsForCall(0)
					Expect(d.Status).To(Equal("FAILED"))
				})
			})

			When("updating the deployment returns an error", func() {
				BeforeEach(func() {
					fakeSQLClient.UpdateDeploymentReturns(errors.New("error updating deployment"))
				})

				It("does nothing", func() {
				})
			})
		})

		When("it creates the deployment", func() {
			It("returns the deployment", func() {
				Expect(res.StatusCode).To(Equal(http.StatusCreated))
				var d cloudrunner.Deployment
				b, _ := ioutil.ReadAll(res.Body)
				json.Unmarshal(b, &d)
				Expect(d.ID).To(HaveLen(uuidLength))
				Expect(d.Status).To(Equal("RUNNING"))
				Expect(*d.StartTime).To(BeTemporally("~", cloudrunner.CurrentTimeUTC(), time.Second))
			})
		})
	})

	Describe("#GetDeployment", func() {
		BeforeEach(func() {
			setup()
			uri = svr.URL + "/v1/deployments/test-id"
			createRequest(http.MethodGet)
		})

		AfterEach(func() {
			teardown()
		})

		JustBeforeEach(func() {
			doRequest()
		})

		When("the record is not found", func() {
			BeforeEach(func() {
				fakeSQLClient.GetDeploymentReturns(cloudrunner.Deployment{}, gorm.ErrRecordNotFound)
			})

			It("returns an error", func() {
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
				validateResponse(payloadDeploymentNotFound)
			})
		})

		When("getting the deployment returns a generic error", func() {
			BeforeEach(func() {
				fakeSQLClient.GetDeploymentReturns(cloudrunner.Deployment{}, errors.New("error getting deployment"))
			})

			It("returns an error", func() {
				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				validateResponse(payloadDeploymentsGetGenericError)
			})
		})

		When("the deployment is running", func() {
			BeforeEach(func() {
				t := cloudrunner.CurrentTimeUTC()
				fakeSQLClient.GetDeploymentReturns(cloudrunner.Deployment{
					EndTime:   &t,
					ID:        "fake-id",
					StartTime: &t,
					Status:    "RUNNING",
				}, nil)
			})

			It("returns the deployment without an endtime", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				var d cloudrunner.Deployment
				b, _ := ioutil.ReadAll(res.Body)
				json.Unmarshal(b, &d)
				Expect(d.ID).To(Equal("fake-id"))
				Expect(d.Status).To(Equal("RUNNING"))
				Expect(*d.StartTime).To(BeTemporally("~", cloudrunner.CurrentTimeUTC(), time.Second))
				Expect(d.EndTime).To(BeNil())
				Expect(d.Output).To(BeEmpty())
			})
		})

		When("it gets the deployment", func() {
			It("returns status OK", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				var d cloudrunner.Deployment
				b, _ := ioutil.ReadAll(res.Body)
				json.Unmarshal(b, &d)
				Expect(d.ID).To(Equal("fake-id"))
				Expect(d.Status).To(Equal("fake-status"))
				Expect(*d.StartTime).To(BeTemporally("~", cloudrunner.CurrentTimeUTC(), time.Second))
				Expect(*d.EndTime).To(BeTemporally("~", cloudrunner.CurrentTimeUTC(), time.Second))
				Expect(d.Output).To(Equal("fake-output"))
			})
		})
	})
})
