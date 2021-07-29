package v1_test

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/homedepot/cloud-runner/internal/fiat"
	"github.com/homedepot/cloud-runner/internal/sql"
	cloudrunner "github.com/homedepot/cloud-runner/pkg"
)

var _ = Describe("Credentials", func() {
	Describe("#CreateCredentials", func() {
		BeforeEach(func() {
			setup()
			uri = svr.URL + "/v1/credentials"
			body.Write([]byte(payloadRequestCredentials))
			createRequest(http.MethodPost)
			// This is required for the success case.
			fakeSQLClient.GetCredentialsReturns(cloudrunner.Credentials{}, sql.ErrCredentialsNotFound)
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

		When("the project ID is not defined", func() {
			BeforeEach(func() {
				body = &bytes.Buffer{}
				body.Write([]byte("{}"))
				createRequest(http.MethodPost)
			})

			It("returns status bad request", func() {
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				validateResponse(payloadProjectIDRequired)
			})
		})

		When("the credentials already exists", func() {
			BeforeEach(func() {
				fakeSQLClient.GetCredentialsReturns(cloudrunner.Credentials{}, nil)
			})

			It("returns status conflict", func() {
				Expect(res.StatusCode).To(Equal(http.StatusConflict))
				validateResponse(payloadConflictRequest)
			})
		})

		When("getting the credentials returns a generic error", func() {
			BeforeEach(func() {
				fakeSQLClient.GetCredentialsReturns(cloudrunner.Credentials{}, errors.New("generic error"))
			})

			It("returns status internal server error", func() {
				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				validateResponse(payloadGenericError)
			})
		})

		When("creating the credentials returns an error", func() {
			BeforeEach(func() {
				fakeSQLClient.CreateCredentialsReturns(errors.New("error creating credentials"))
			})

			It("returns status internal server error", func() {
				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				validateResponse(payloadErrorCreatingCredentials)
			})
		})

		When("the account is not defined", func() {
			BeforeEach(func() {
				body = &bytes.Buffer{}
				body.Write([]byte(`{"projectID": "test-project-id"}`))
				createRequest(http.MethodPost)
			})

			It("generates the account name", func() {
				Expect(res.StatusCode).To(Equal(http.StatusCreated))
				validateResponse(payloadCredentialsCreatedNoAccountProvided)
			})
		})

		When("the request contains write groups that are not present in read groups", func() {
			BeforeEach(func() {
				body = &bytes.Buffer{}
				body.Write([]byte(payloadRequestCredentialsMismatchedGroups))
				createRequest(http.MethodPost)
			})

			It("merges the write groups into the read groups", func() {
				Expect(res.StatusCode).To(Equal(http.StatusCreated))
				validateResponse(payloadCredentialsCreatedMergedGroups)
			})
		})

		When("it creates the credentials", func() {
			It("returns status created", func() {
				Expect(res.StatusCode).To(Equal(http.StatusCreated))
				validateResponse(payloadCredentialsCreated)
			})
		})
	})

	Describe("#DeleteCredentials", func() {
		BeforeEach(func() {
			setup()
			uri = svr.URL + "/v1/credentials/test-name"
			createRequest(http.MethodDelete)
		})

		AfterEach(func() {
			teardown()
		})

		JustBeforeEach(func() {
			doRequest()
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
				fakeSQLClient.GetCredentialsReturns(cloudrunner.Credentials{}, gorm.ErrRecordNotFound)
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

		When("the user does not have access to delete the account", func() {
			BeforeEach(func() {
				fakeFiatClient.RolesReturns(fiat.Roles{{Name: "not_a_good_group"}}, nil)
			})

			It("returns an error", func() {
				Expect(res.StatusCode).To(Equal(http.StatusForbidden))
				validateResponse(payloadCredentialsDeleteForbiddentError)
			})
		})

		When("deleting the credentials returns an error", func() {
			BeforeEach(func() {
				fakeSQLClient.DeleteCredentialsReturns(errors.New("error deleting credentials"))
			})

			It("returns an error", func() {
				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				validateResponse(payloadCredentialsDeleteGenericError)
			})
		})

		When("it deletes the credentials", func() {
			It("returns status no content", func() {
				Expect(res.StatusCode).To(Equal(http.StatusNoContent))
			})
		})
	})

	Describe("#GetCredentials", func() {
		BeforeEach(func() {
			setup()
			uri = svr.URL + "/v1/credentials/test-name"
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

		When("it gets the credentials", func() {
			It("returns status OK", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				validateResponse(payloadCredentials)
			})
		})
	})

	Describe("#ListCredentials", func() {
		BeforeEach(func() {
			setup()
			uri = svr.URL + "/v1/credentials"
			createRequest(http.MethodGet)
		})

		AfterEach(func() {
			teardown()
		})

		JustBeforeEach(func() {
			doRequest()
		})

		When("getting the credentials returns a generic error", func() {
			BeforeEach(func() {
				fakeSQLClient.ListCredentialsReturns(nil, errors.New("error getting credentials"))
			})

			It("returns an error", func() {
				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				validateResponse(payloadCredentialsGetGenericError)
			})
		})

		Context("when onlyForUser is true", func() {
			BeforeEach(func() {
				setup()
				uri = svr.URL + "/v1/credentials?onlyForUser=true"
				createRequest(http.MethodGet)
			})

			When("the X-Spinnaker-User header is not set", func() {
				BeforeEach(func() {
					req.Header.Set("X-Spinnaker-User", "")
				})

				It("returns status unauthorized", func() {
					Expect(res.StatusCode).To(Equal(http.StatusUnauthorized))
					validateResponse(payloadCredentialsUnauthorized)
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

			When("the user is not part of any roles", func() {
				BeforeEach(func() {
					fakeFiatClient.RolesReturns(nil, nil)
				})

				It("returns no credentials", func() {
					Expect(res.StatusCode).To(Equal(http.StatusOK))
					validateResponse(payloadCredentialsListEmpty)
				})
			})

			When("the user is part of a role", func() {
				BeforeEach(func() {
					fakeFiatClient.RolesReturns(fiat.Roles{{Name: "gg_test2"}}, nil)
				})

				It("returns one credential", func() {
					Expect(res.StatusCode).To(Equal(http.StatusOK))
					validateResponse(payloadCredentialsListFiltered)
				})
			})
		})

		When("it lists the credentials", func() {
			It("returns status OK", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				validateResponse(payloadCredentialsList)
			})
		})
	})
})
