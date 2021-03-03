package gcloud_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGcloud(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gcloud Suite")
}
