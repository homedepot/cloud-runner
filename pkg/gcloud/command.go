package gcloud

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

const (
	defaultBase = `gcloud run deploy`
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
	// Service name must use only lowercase alphanumeric characters and dashes.
	// Cannot begin or end with a dash, and cannot be longer than 63 characters.
	regexLowerCaseAlphanumericMax63Chars = `^[a-z]([a-z0-9-]{0,61})?[a-z0-9]$`
	// The project ID must be a unique string of 6 to 30 lowercase letters, digits, or hyphens.
	// It must start with a letter, and cannot have a trailing hyphen.
	// See https://cloud.google.com/resource-manager/docs/creating-managing-projects#before_you_begin
	regexLowerCaseAlphanumericMax30CharsMin6Chars = `^[a-z]([a-z0-9-]{4,28})?[a-z0-9]$`
	regexNoSingleQuotes                           = `^[^']+$`
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
	base                 string
	image                string
	maxInstances         int
	memory               string
	projectID            string
	region               string
	service              string
	vpcConnector         string
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
	return &cloudRunCommandBuilder{
		base: defaultBase,
	}
}

// Build builds, validates, and sets the command.
func (c *cloudRunCommandBuilder) Build() (CloudRunCommand, error) {
	command := c.base

	err := c.validateRequiredFlags()
	if err != nil {
		return nil, err
	}

	// Set the service name.
	command = appendString(command, c.service)
	// Set the --project flag
	command = appendString(command, flagProject)
	command = appendString(command, c.projectID)
	// Set the --image '' flag.
	command = appendString(command, flagImage)
	command = appendWrappedString(command, c.image)
	// Set the --platform managed flag.
	command = appendString(command, flagPlatform)
	command = appendString(command, "managed")
	// Set the --region flag.
	command = appendString(command, flagRegion)
	command = appendString(command, c.region)

	if c.allowUnauthenticated {
		command = appendString(command, flagAllowUnauthenticated)
	} else {
		command = appendString(command, flagNoAllowUnauthenticated)
	}

	// Optional flags.
	if c.maxInstances > 0 {
		command = appendString(command, flagMaxInstances)
		command = appendInt(command, c.maxInstances)
	}

	if c.memory != "" {
		// Memory is validated server-side in GCP.
		// Invalid memory will throw the error 'ERROR: (gcloud.run.deploy) Could not parse Quantity: <MEMORY>'
		err = validate(regexNoSingleQuotes, c.memory)
		if err != nil {
			return nil, fmt.Errorf("error validating memory: %w", err)
		}

		command = appendString(command, flagMemory)
		command = appendWrappedString(command, c.memory)
	}

	if c.vpcConnector != "" {
		// VPC connector is validated server-side in GCP.
		err = validate(regexNoSingleQuotes, c.vpcConnector)
		if err != nil {
			return nil, fmt.Errorf("error validating VPC connector: %w", err)
		}

		command = appendString(command, flagVPCConnector)
		command = appendWrappedString(command, c.vpcConnector)
	}

	cmdSlice := strings.Split(command, " ")
	cmd := exec.Command(cmdSlice[0], cmdSlice[1:]...)

	return &cloudRunCommand{cmd: cmd}, nil
}

func (c *cloudRunCommandBuilder) validateRequiredFlags() error {
	err := validate(regexNoSingleQuotes, c.image)
	if err != nil {
		return fmt.Errorf("error validating image: %w", err)
	}

	err = validate(regexLowerCaseAlphanumericMax30CharsMin6Chars, c.projectID)
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

func appendString(command, s string) string {
	return fmt.Sprintf("%s %s", command, s)
}

func appendWrappedString(command, s string) string {
	return fmt.Sprintf("%s '%s'", command, s)
}

func appendInt(command string, i int) string {
	return fmt.Sprintf("%s %d", command, i)
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
