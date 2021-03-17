package snowql_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSnowql(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Snowql Suite")
}
