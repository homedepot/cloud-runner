package sql

import (
	"database/sql"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	cloudrunner "github.com/homedepot/cloud-runner/pkg"

	// Needed for connection.
	_ "github.com/go-sql-driver/mysql"

	// Needed for connection.
	_ "github.com/mattn/go-sqlite3"
)

const (
	ClientInstanceKey = `SQLClient`
	maxOpenConns      = 5
	connMaxLifetime   = time.Second * 30
)

var (
	ErrCredentialsNotFound = errors.New("credentials not found")
)

//go:generate counterfeiter . Client

// Client holds the DB connection and makes all queries to the DB.
type Client interface {
	Connect(string, interface{}) error
	Connection() (string, string)
	CreateCredentials(cloudrunner.Credentials) error
	CreateDeployment(cloudrunner.Deployment) error
	DB() *gorm.DB
	DeleteCredentials(string) error
	GetCredentials(string) (cloudrunner.Credentials, error)
	GetDeployment(string) (cloudrunner.Deployment, error)
	ListCredentials() ([]cloudrunner.Credentials, error)
	UpdateDeployment(cloudrunner.Deployment) error
	WithHost(string)
	WithName(string)
	WithPass(string)
	WithUser(string)
}

// NewClient returns a new instance of Client.
func NewClient() Client {
	return &client{}
}

type client struct {
	db   *gorm.DB
	host string
	name string
	pass string
	user string
}

// Connect sets up the database connection and creates tables.
//
// Connection is of type interface{} - this allows for tests to
// pass in a sqlmock connection and for main to connect given a
// connection string.
func (c *client) Connect(driver string, connection interface{}) error {
	db, err := gorm.Open(driver, connection)
	if err != nil {
		return fmt.Errorf("error opening connection to DB: %w", err)
	}

	db.LogMode(false)
	db.AutoMigrate(
		&cloudrunner.Credentials{},
		&cloudrunner.CredentialsReadPermission{},
		&cloudrunner.CredentialsWritePermission{},
		&cloudrunner.Deployment{},
	)

	db.DB().SetMaxOpenConns(maxOpenConns)
	db.DB().SetMaxIdleConns(1)
	db.DB().SetConnMaxLifetime(connMaxLifetime)

	c.db = db

	return nil
}

// Connection returns a connection string to the DB.
func (c *client) Connection() (string, string) {
	if c.user == "" || c.pass == "" || c.host == "" || c.name == "" {
		return "sqlite3", "cloud-runner.db"
	}

	return "mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=5s&charset=utf8&parseTime=True&loc=UTC",
		c.user, c.pass, c.host, c.name)
}

// CreateCredentials inserts a new set of credentials into the DB along with
// its associated read/write groups. It uses a transaction, so that if any
// operation fails the operations are rolled back.
func (c *client) CreateCredentials(credentials cloudrunner.Credentials) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		for _, group := range credentials.ReadGroups {
			r := cloudrunner.CredentialsReadPermission{
				ID:        uuid.New().String(),
				Account:   credentials.Account,
				ReadGroup: group,
			}

			err := tx.Create(&r).Error
			if err != nil {
				return err
			}
		}

		for _, group := range credentials.WriteGroups {
			w := cloudrunner.CredentialsWritePermission{
				ID:         uuid.New().String(),
				Account:    credentials.Account,
				WriteGroup: group,
			}

			err := tx.Create(&w).Error
			if err != nil {
				return err
			}
		}

		return tx.Create(&credentials).Error
	})
}

// CreateDeployment inserts a new deployment into the DB.
func (c *client) CreateDeployment(d cloudrunner.Deployment) error {
	return c.db.Create(&d).Error
}

// DB returns the underlying db.
func (c *client) DB() *gorm.DB {
	return c.db
}

// DeleteCredentials deletes credentials by account name,
// then deletes read and write permissions by the same account name.
//
// TODO this should be a transaction that we can roll back.
func (c *client) DeleteCredentials(account string) error {
	err := c.db.Delete(&cloudrunner.Credentials{Account: account}).Error
	if err != nil {
		return err
	}

	err = c.db.Where("account = ?", account).Delete(&cloudrunner.CredentialsReadPermission{}).Error
	if err != nil {
		return err
	}

	err = c.db.Where("account = ?", account).Delete(&cloudrunner.CredentialsWritePermission{}).Error
	if err != nil {
		return err
	}

	return nil
}

// GetCredentials selects credentials and permissions from the DB by account name.
func (c *client) GetCredentials(account string) (cloudrunner.Credentials, error) {
	rows, err := c.db.Table("credentials a").
		Select("a.account, "+
			"a.project_id, "+
			"b.read_group, "+
			"c.write_group").
		Joins("LEFT JOIN credentials_read_permissions b ON a.account = b.account").
		Joins("LEFT JOIN credentials_write_permissions c ON a.account = c.account").
		Where("a.account = ?", account).
		Rows()
	if err != nil {
		return cloudrunner.Credentials{}, err
	}
	defer rows.Close()

	cs, err := mergeRows(rows)
	if err != nil {
		return cloudrunner.Credentials{}, err
	}

	if len(cs) == 0 {
		return cloudrunner.Credentials{}, ErrCredentialsNotFound
	}

	return cs[0], nil
}

// GetDeployment selects a deployment status by ID from the DB.
func (c *client) GetDeployment(id string) (cloudrunner.Deployment, error) {
	var d cloudrunner.Deployment
	db := c.db.Find(&d, "id = ?", id)

	return d, db.Error
}

// ListCredentials selects all credentials and their corresponding
// permissions from the DB.
func (c *client) ListCredentials() ([]cloudrunner.Credentials, error) {
	rows, err := c.db.Table("credentials a").
		Select("a.account, " +
			"a.project_id, " +
			"b.read_group, " +
			"c.write_group").
		Joins("LEFT JOIN credentials_read_permissions b ON a.account = b.account").
		Joins("LEFT JOIN credentials_write_permissions c ON a.account = c.account").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cs, err := mergeRows(rows)
	if err != nil {
		return nil, err
	}

	return cs, nil
}

// UpdateDeployment updates a given deployment.
func (c *client) UpdateDeployment(d cloudrunner.Deployment) error {
	return c.db.Save(d).Error
}

func mergeRows(rows *sql.Rows) ([]cloudrunner.Credentials, error) {
	cs := []cloudrunner.Credentials{}
	credentials := map[string]cloudrunner.Credentials{}
	readGroups := map[string][]string{}
	writeGroups := map[string][]string{}

	for rows.Next() {
		var r struct {
			Account    string
			ProjectID  string
			ReadGroup  *string
			WriteGroup *string
		}

		err := rows.Scan(&r.Account, &r.ProjectID, &r.ReadGroup, &r.WriteGroup)
		if err != nil {
			return nil, err
		}

		if _, ok := credentials[r.Account]; !ok {
			creds := cloudrunner.Credentials{
				Account:   r.Account,
				ProjectID: r.ProjectID,
			}
			credentials[r.Account] = creds
		}

		if r.ReadGroup != nil {
			if _, ok := readGroups[r.Account]; !ok {
				readGroups[r.Account] = []string{}
			}

			if !contains(readGroups[r.Account], *r.ReadGroup) {
				readGroups[r.Account] = append(readGroups[r.Account], *r.ReadGroup)
			}
		}

		if r.WriteGroup != nil {
			if _, ok := writeGroups[r.Account]; !ok {
				writeGroups[r.Account] = []string{}
			}

			if !contains(writeGroups[r.Account], *r.WriteGroup) {
				writeGroups[r.Account] = append(writeGroups[r.Account], *r.WriteGroup)
			}
		}
	}

	for name, creds := range credentials {
		creds.ReadGroups = readGroups[name]
		creds.WriteGroups = writeGroups[name]
		cs = append(cs, creds)
	}

	// Sort ascending by name.
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Account < cs[j].Account
	})

	return cs, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

// WithHost set the host for the connection.
func (c *client) WithHost(host string) {
	c.host = host
}

// WithName set the DB name for the connection.
func (c *client) WithName(name string) {
	c.name = name
}

// WithPassword set the password for the connection.
func (c *client) WithPass(pass string) {
	c.pass = pass
}

// WithUser set the user for the connection.
func (c *client) WithUser(user string) {
	c.user = user
}
