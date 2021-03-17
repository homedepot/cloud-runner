package thd_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestThd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Thd Suite")
}
