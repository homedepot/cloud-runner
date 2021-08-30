package gcloud

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

const (
	BuilderInstanceKey = `CloudRunCommandBuilder`

	// Flags.
	flagAllowUnauthenticated   = `--allow-unauthenticated`
	flagNoAllowUnauthenticated = `--no-allow-unauthenticated`
	flagImage                  = `--image`
	flagMaxInstances           = `--max-instances`
	flagMemory                 = `--memory`
	flagPlatform               = `--platform`
	flagProject                = `--project`
	flagRegion                 = `--region`
	flagVPCConnector           = `--vpc-connector`
	flagServiceAccount         = `--service-account`
	// Service name must use only lowercase alphanumeric characters and dashes.
	// Cannot begin or end with a dash, and cannot be longer than 63 characters.
	regexLowerCaseAlphanumericMax63Chars = `^[a-z]([a-z0-9-]{0,61})?[a-z0-9]$`
	// The project ID must be a unique string of 6 to 30 lowercase letters, digits, or hyphens.
	// It must start with a letter, and cannot have a trailing hyphen.
	// See https://cloud.google.com/resource-manager/docs/creating-managing-projects#before_you_begin
	regexLowerCaseAlphanumericMax30CharsMin6Chars = `^[a-z]([a-z0-9-]{4,28})?[a-z0-9]$`
)

var (
	defaultBase = []string{"gcloud", "run", "deploy"}
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . CloudRunCommandBuilder

// CloudRunCommand holds the data to build and run a `gcloud run deploy` command.
type CloudRunCommandBuilder interface {
	Build() (CloudRunCommand, error)
	AllowUnauthenticated(bool) CloudRunCommandBuilder
	Image(string) CloudRunCommandBuilder
	MaxInstances(int) CloudRunCommandBuilder
	Memory(string) CloudRunCommandBuilder
	ProjectID(string) CloudRunCommandBuilder
	Region(string) CloudRunCommandBuilder
	Service(string) CloudRunCommandBuilder
	VPCConnector(string) CloudRunCommandBuilder
}

type cloudRunCommandBuilder struct {
	allowUnauthenticated bool
	image                string
	maxInstances         int
	memory               string
	projectID            string
	region               string
	service              string
	vpcConnector         string
	serviceAccount       string
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . CloudRunCommand

// CloudRunCommand wraps the executable command.
type CloudRunCommand interface {
	CombinedOutput() ([]byte, error)
	String() string
}

type cloudRunCommand struct {
	cmd *exec.Cmd
}

// NewCloudRunCommand returns an implementation of a `gcloud run deploy` command.
func NewCloudRunCommandBuilder() CloudRunCommandBuilder {
	return &cloudRunCommandBuilder{}
}

// Build builds, validates, and sets the command.
func (c *cloudRunCommandBuilder) Build() (CloudRunCommand, error) {
	command := defaultBase

	err := c.validateRequiredFlags()
	if err != nil {
		return nil, err
	}

	// Set the service name.
	command = append(command, c.service)
	// Set the --project flag
	command = append(command, flagProject)
	command = append(command, c.projectID)
	// Set the --image flag.
	// I originally wanted to wrap this in single quotes to avoid shell injection, but it kept throwing the error:
	//
	// Expected a Container Registry image path like [region.]gcr.io/repo-path[:tag or @digest]
	// or an Artifact Registry image path like [region-]docker.pkg.dev/repo-path[:tag or @digest],
	// but obtained 'gcr.io/github-replication-sandbox/rf:1.0.9'
	//
	// While the command runs fine directly in the shell, Go is definitely doing something to make these
	// arguments safe to not require wrapping (I attempted to shell inject and it failed).
	//
	// For further reading see https://docs.guardrails.io/docs/en/vulnerabilities/go/insecure_use_of_dangerous_function
	command = append(command, flagImage)
	command = append(command, c.image)
	// Set the --platform managed flag.
	command = append(command, flagPlatform)
	command = append(command, "managed")
	// Set the --region flag.
	command = append(command, flagRegion)
	command = append(command, c.region)

	if c.allowUnauthenticated {
		command = append(command, flagAllowUnauthenticated)
	} else {
		command = append(command, flagNoAllowUnauthenticated)
	}

	// Optional flags.
	if c.maxInstances > 0 {
		command = append(command, flagMaxInstances)
		command = append(command, strconv.Itoa(c.maxInstances))
	}

	if c.memory != "" {
		// Memory is validated server-side in GCP.
		// Invalid memory will throw the error 'ERROR: (gcloud.run.deploy) Could not parse Quantity: <MEMORY>'
		command = append(command, flagMemory)
		command = append(command, c.memory)
	}

	if c.vpcConnector != "" {
		// VPC connector is validated server-side in GCP.
		command = append(command, flagVPCConnector)
		command = append(command, c.vpcConnector)
	}

	if c.serviceAccount != "" {
		// Service Account for container service
		command = append(command, flagServiceAccount)
		command = append(command, c.serviceAccount)
	}

	cmd := exec.Command(command[0], command[1:]...)

	return &cloudRunCommand{cmd: cmd}, nil
}

func (c *cloudRunCommandBuilder) validateRequiredFlags() error {
	err := validate(regexLowerCaseAlphanumericMax30CharsMin6Chars, c.projectID)
	if err != nil {
		return fmt.Errorf("error validating project ID: %w", err)
	}

	err = validate(regexLowerCaseAlphanumericMax63Chars, c.service)
	if err != nil {
		return fmt.Errorf("error validating service name: %w", err)
	}

	err = validate(regexLowerCaseAlphanumericMax63Chars, c.region)
	if err != nil {
		return fmt.Errorf("error validating region: %w", err)
	}

	return nil
}

func validate(regex, s string) error {
	r := regexp.MustCompile(regex)
	if !r.MatchString(s) {
		return fmt.Errorf("%s failed validation", s)
	}

	return nil
}

// String returns the underlying command.
func (c *cloudRunCommand) String() string {
	return c.cmd.String()
}

// CombinedOutput runs the command and returns the output.
func (c *cloudRunCommand) CombinedOutput() ([]byte, error) {
	return c.cmd.CombinedOutput()
}

// AllowUnauthenticated sets the allowUnauthenticated boolean.
func (c *cloudRunCommandBuilder) AllowUnauthenticated(allowUnauthenticated bool) CloudRunCommandBuilder {
	c.allowUnauthenticated = allowUnauthenticated

	return c
}

// Image sets the --image flag.
func (c *cloudRunCommandBuilder) Image(image string) CloudRunCommandBuilder {
	c.image = image

	return c
}

// MaxInstances sets the --max-instances flag (optional).
func (c *cloudRunCommandBuilder) MaxInstances(maxInstances int) CloudRunCommandBuilder {
	c.maxInstances = maxInstances

	return c
}

// Memory sets the --memory flag (optional).
func (c *cloudRunCommandBuilder) Memory(memory string) CloudRunCommandBuilder {
	c.memory = memory

	return c
}

// ProjectID sets the --project flag.
func (c *cloudRunCommandBuilder) ProjectID(projectID string) CloudRunCommandBuilder {
	c.projectID = projectID

	return c
}

// Region sets the --region flag.
func (c *cloudRunCommandBuilder) Region(region string) CloudRunCommandBuilder {
	c.region = region

	return c
}

// Service sets the service name for the command.
func (c *cloudRunCommandBuilder) Service(service string) CloudRunCommandBuilder {
	c.service = service

	return c
}

// VPCConnector sets the --vpc-connector flag (optional).
func (c *cloudRunCommandBuilder) VPCConnector(vpcConnector string) CloudRunCommandBuilder {
	c.vpcConnector = vpcConnector

	return c
}

// ServiceAccount sets the --service-account flag (optional).
func (c *cloudRunCommandBuilder) ServiceAccount(serviceAccount string) CloudRunCommandBuilder {
	c.serviceAccount = serviceAccount

	return c
}

