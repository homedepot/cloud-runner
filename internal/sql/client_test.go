package sql_test

import (
	"database/sql"
	"errors"
	"time"

	internal "github.com/homedepot/cloud-runner/internal"
	. "github.com/homedepot/cloud-runner/internal/sql"
	cloudrunner "github.com/homedepot/cloud-runner/pkg"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sql", func() {
	var (
		mock sqlmock.Sqlmock
		d    *sql.DB
		c    Client
		err  error
	)

	BeforeEach(func() {
		// Mock DB.
		d, mock, _ = sqlmock.New()
		c = NewClient()
		err = c.Connect("sqlite3", d)
		// Enable DB logging.
		// c.DB().LogMode(true)
	})

	AfterEach(func() {
		c.DB().Close()
	})

	Describe("#Connect", func() {
		When("it fails to connect", func() {
			BeforeEach(func() {
				err = c.Connect("mysql", "mysql")
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("error opening connection to DB: " +
					"invalid DSN: missing the slash separating the database name"))
			})
		})
	})

	Describe("#Connection", func() {
		var driver, connection string

		When("no vars are set for the connection", func() {
			BeforeEach(func() {
				driver, connection = c.Connection()
			})

			It("returns a disk db", func() {
				Expect(connection).To(Equal("cloud-runner.db"))
			})
		})

		When("the sql config is set", func() {
			BeforeEach(func() {
				c.WithHost("host")
				c.WithName("name")
				c.WithPass("pass")
				c.WithUser("user")
				driver, connection = c.Connection()
			})

			It("returns a mysql db", func() {
				Expect(connection).To(Equal("user:pass@tcp(host)/name?timeout=5s&charset=utf8&parseTime=True&loc=UTC"))
				Expect(driver).To(Equal("mysql"))
			})
		})
	})

	Describe("#CreateCredentials", func() {
		var credentials cloudrunner.Credentials

		BeforeEach(func() {
			credentials = cloudrunner.Credentials{
				Account: "test-account",
			}
		})

		JustBeforeEach(func() {
			err = c.CreateCredentials(credentials)
		})

		When("it creates the credentials", func() {
			BeforeEach(func() {
				mock.ExpectBegin()
				mock.ExpectExec(`(?i)^INSERT INTO "credentials" \(` +
					`"account",` +
					`"project_id"` +
					`\) VALUES \(\?,\?\)$`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#CreateDeployment", func() {
		var deployment cloudrunner.Deployment

		BeforeEach(func() {
			t := internal.CurrentTimeUTC()
			deployment = cloudrunner.Deployment{
				EndTime:   &t,
				ID:        "test-id",
				StartTime: &t,
				Status:    "RUNNING",
				Output:    "test-output",
			}
		})

		JustBeforeEach(func() {
			err = c.CreateDeployment(deployment)
		})

		When("it creates the deployment", func() {
			BeforeEach(func() {
				mock.ExpectBegin()
				mock.ExpectExec(`(?i)^INSERT INTO "deployments" \(` +
					`"end_time",` +
					`"id",` +
					`"start_time",` +
					`"status",` +
					`"output"` +
					`\) VALUES \(\?,\?,\?,\?,\?\)$`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#CreateReadPermission", func() {
		var rp cloudrunner.CredentialsReadPermission

		BeforeEach(func() {
			rp = cloudrunner.CredentialsReadPermission{
				ID:        "test-id",
				Account:   "test-account-name",
				ReadGroup: "test-write-group",
			}
		})

		JustBeforeEach(func() {
			err = c.CreateReadPermission(rp)
		})

		When("it creates the read permissions", func() {
			BeforeEach(func() {
				mock.ExpectBegin()
				mock.ExpectExec(`(?i)^INSERT INTO "credentials_read_permissions" \(` +
					`"account",` +
					`"id",` +
					`"read_group"` +
					`\) VALUES \(\?,\?,\?\)$`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#CreateWritePermission", func() {
		var wp cloudrunner.CredentialsWritePermission

		BeforeEach(func() {
			wp = cloudrunner.CredentialsWritePermission{
				ID:         "test-id",
				Account:    "test-account-name",
				WriteGroup: "test-write-group",
			}
		})

		JustBeforeEach(func() {
			err = c.CreateWritePermission(wp)
		})

		When("it creates the write permissions", func() {
			BeforeEach(func() {
				mock.ExpectBegin()
				mock.ExpectExec(`(?i)^INSERT INTO "credentials_write_permissions" \(` +
					`"account",` +
					`"id",` +
					`"write_group"` +
					`\) VALUES \(\?,\?,\?\)$`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#DeleteCredentials", func() {
		var account string

		BeforeEach(func() {
			account = "test-account"
		})

		JustBeforeEach(func() {
			err = c.DeleteCredentials(account)
		})

		When("deleting the credentials returns an error", func() {
			BeforeEach(func() {
				mock.ExpectBegin()
				mock.ExpectExec(`(?i)^DELETE FROM "credentials" WHERE
				"credentials"."account" = \?$`).
					WillReturnError(errors.New("error deleting credentials"))
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("error deleting credentials"))
			})
		})

		When("deleting the read permissions returns an error", func() {
			BeforeEach(func() {
				mock.ExpectBegin()
				mock.ExpectExec(`(?i)^DELETE FROM "credentials" WHERE
				"credentials"."account" = \?$`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				mock.ExpectBegin()
				mock.ExpectExec(`(?i)^DELETE FROM "credentials_read_permissions" WHERE
				\(account = \?\)$`).
					WillReturnError(errors.New("error deleting read permissions"))
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("error deleting read permissions"))
			})
		})

		When("deleting the write permissions returns an error", func() {
			BeforeEach(func() {
				mock.ExpectBegin()
				mock.ExpectExec(`(?i)^DELETE FROM "credentials" WHERE
				"credentials"."account" = \?$`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				mock.ExpectBegin()
				mock.ExpectExec(`(?i)^DELETE FROM "credentials_read_permissions" WHERE
				\(account = \?\)$`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				mock.ExpectBegin()
				mock.ExpectExec(`(?i)^DELETE FROM "credentials_write_permissions" WHERE
				\(account = \?\)$`).
					WillReturnError(errors.New("error deleting write permissions"))
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("error deleting write permissions"))
			})
		})

		When("it deletes the creds", func() {
			BeforeEach(func() {
				mock.ExpectBegin()
				mock.ExpectExec(`(?i)^DELETE FROM "credentials" WHERE
				"credentials"."account" = \?$`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				mock.ExpectBegin()
				mock.ExpectExec(`(?i)^DELETE FROM "credentials_read_permissions" WHERE
				\(account = \?\)$`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				mock.ExpectBegin()
				mock.ExpectExec(`(?i)^DELETE FROM "credentials_write_permissions" WHERE
				\(account = \?\)$`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#GetCredentials", func() {
		var credentials cloudrunner.Credentials
		var account string

		BeforeEach(func() {
			account = "test-account"
		})

		JustBeforeEach(func() {
			credentials, err = c.GetCredentials(account)
		})

		When("getting the rows returns an error", func() {
			BeforeEach(func() {
				c.DB().Close()
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("sql: database is closed"))
			})
		})

		When("scanning a row returns an error", func() {
			BeforeEach(func() {
				sqlRows := sqlmock.NewRows([]string{"account", "read_group", "write_group"}).
					AddRow("account1", "read_group1", "write_group1")
				mock.ExpectQuery(`(?i)^SELECT ` +
					`a.account, ` +
					`a.project_id, ` +
					`b.read_group, ` +
					`c.write_group ` +
					`FROM credentials a ` +
					`left join credentials_read_permissions b on a.account = b.account ` +
					`left join credentials_write_permissions c on a.account = c.account ` +
					`WHERE \(a.account = \?\)`).
					WillReturnRows(sqlRows)
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("sql: expected 3 destination arguments in Scan, not 4"))
			})
		})

		When("no rows are returned", func() {
			BeforeEach(func() {
				sqlRows := sqlmock.NewRows([]string{"account", "read_group", "write_group"})
				mock.ExpectQuery(`(?i)^SELECT ` +
					`a.account, ` +
					`a.project_id, ` +
					`b.read_group, ` +
					`c.write_group ` +
					`FROM credentials a ` +
					`left join credentials_read_permissions b on a.account = b.account ` +
					`left join credentials_write_permissions c on a.account = c.account ` +
					`WHERE \(a.account = \?\)`).
					WillReturnRows(sqlRows)
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("credentials not found"))
			})
		})

		When("it gets the creds", func() {
			BeforeEach(func() {
				sqlRows := sqlmock.NewRows([]string{"account", "project_id", "read_group", "write_group"}).
					AddRow("test-account", "project_id1", "read_group1", "write_group1").
					AddRow("test-account", "project_id1", "read_group2", "write_group1")
				mock.ExpectQuery(`(?i)^SELECT ` +
					`a.account, ` +
					`a.project_id, ` +
					`b.read_group, ` +
					`c.write_group ` +
					`FROM credentials a ` +
					`left join credentials_read_permissions b on a.account = b.account ` +
					`left join credentials_write_permissions c on a.account = c.account ` +
					`WHERE \(a.account = \?\)`).
					WillReturnRows(sqlRows)
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
				Expect(credentials.Account).To(Equal("test-account"))
				Expect(credentials.ProjectID).To(Equal("project_id1"))
				Expect(credentials.ReadGroups).To(HaveLen(2))
				Expect(credentials.WriteGroups).To(HaveLen(1))
			})
		})
	})

	Describe("#GetDeployment", func() {
		var deployment cloudrunner.Deployment
		var id string
		var t time.Time

		BeforeEach(func() {
			id = "e22a8932-7b2a-472d-8858-4690c2e5bf5c"
			t = internal.CurrentTimeUTC()
		})

		JustBeforeEach(func() {
			deployment, err = c.GetDeployment(id)
		})

		When("it gets the deployment", func() {
			BeforeEach(func() {
				sqlRows := sqlmock.NewRows([]string{"end_time", "id", "start_time", "status", "output"}).
					AddRow(t, id, t, "RUNNING", "test-output")
				mock.ExpectQuery(`(?i)^SELECT ` +
					`\* ` +
					`FROM "deployments" ` +
					`WHERE \(id = \?\)$`).
					WillReturnRows(sqlRows)
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
				Expect(*deployment.EndTime).To(BeTemporally("~", t, time.Second))
				Expect(deployment.ID).To(Equal(id))
				Expect(*deployment.StartTime).To(BeTemporally("~", t, time.Second))
				Expect(deployment.Status).To(Equal("RUNNING"))
				Expect(deployment.Output).To(Equal("test-output"))
			})
		})
	})

	Describe("#ListCredentials", func() {
		var credentials []cloudrunner.Credentials

		JustBeforeEach(func() {
			credentials, err = c.ListCredentials()
		})

		When("getting the rows returns an error", func() {
			BeforeEach(func() {
				c.DB().Close()
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("sql: database is closed"))
			})
		})

		When("scanning a row returns an error", func() {
			BeforeEach(func() {
				sqlRows := sqlmock.NewRows([]string{"account", "read_group", "write_group"}).
					AddRow("account1", "read_group1", "write_group1")
				mock.ExpectQuery(`(?i)^SELECT ` +
					`a.account, ` +
					`a.project_id, ` +
					`b.read_group, ` +
					`c.write_group ` +
					`FROM credentials a ` +
					`left join credentials_read_permissions b on a.account = b.account ` +
					`left join credentials_write_permissions c on a.account = c.account$`).
					WillReturnRows(sqlRows)
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("sql: expected 3 destination arguments in Scan, not 4"))
			})
		})

		When("it lists the creds", func() {
			BeforeEach(func() {
				sqlRows := sqlmock.NewRows([]string{"account", "project_id", "read_group", "write_group"}).
					AddRow("account1", "project_id1", "read_group1", "write_group1").
					AddRow("account1", "project_id1", "read_group2", "write_group1").
					AddRow("account2", "project_id2", "read_group2", "write_group2").
					AddRow("account2", "project_id2", "read_group2", "write_group3")
				mock.ExpectQuery(`(?i)^SELECT ` +
					`a.account, ` +
					`a.project_id, ` +
					`b.read_group, ` +
					`c.write_group ` +
					`FROM credentials a ` +
					`left join credentials_read_permissions b on a.account = b.account ` +
					`left join credentials_write_permissions c on a.account = c.account$`).
					WillReturnRows(sqlRows)
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
				Expect(credentials).To(HaveLen(2))
				Expect(credentials[0].ReadGroups).To(HaveLen(2))
				Expect(credentials[0].WriteGroups).To(HaveLen(1))
				Expect(credentials[1].ReadGroups).To(HaveLen(1))
				Expect(credentials[1].WriteGroups).To(HaveLen(2))
			})
		})
	})

	Describe("#UpdateDeployment", func() {
		var deployment cloudrunner.Deployment

		BeforeEach(func() {
			t := internal.CurrentTimeUTC()
			deployment = cloudrunner.Deployment{
				StartTime: &t,
				EndTime:   &t,
				ID:        "test-id",
				Status:    "SUCCEEDED",
				Output:    "test-output",
			}
		})

		JustBeforeEach(func() {
			err = c.UpdateDeployment(deployment)
		})

		When("it updates the deployment", func() {
			BeforeEach(func() {
				mock.ExpectBegin()
				mock.ExpectExec(`(?i)^UPDATE "deployments" SET ` +
					`"end_time" = \?, ` +
					`"start_time" = \?, ` +
					`"status" = \?, ` +
					`"output" = \? ` +
					`WHERE "deployments"."id" = \?$`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
			})
		})
	})
})
