package v1

import (
	"strings"

	"github.homedepot.com/cd/cloud-runner/internal/fiat"
	cloudrunner "github.homedepot.com/cd/cloud-runner/pkg"
)

// filterCredentials takes in a list of credentials and roles. It
// returns a filtered list of credentials that contain both a read and write
// group from the respective roles slice passed in.
func (cc *Controller) filterCredentials(credentials []cloudrunner.Credentials,
	roles fiat.Roles) []cloudrunner.Credentials {
	c := []cloudrunner.Credentials{}
	g := map[string]bool{}

	for _, role := range roles {
		if contains(cc.AdminRoles,role.Name){
			return credentials
		}

		g[strings.ToLower(role.Name)] = true
	}

	for _, creds := range credentials {
		var read, write bool

		for _, readGroup := range creds.ReadGroups {
			if g[strings.ToLower(readGroup)] {
				read = true

				break
			}
		}

		for _, writeGroup := range creds.WriteGroups {
			if g[strings.ToLower(writeGroup)] {
				write = true

				break
			}
		}

		if read && write {
			c = append(c, creds)
		}
	}

	return c
}
