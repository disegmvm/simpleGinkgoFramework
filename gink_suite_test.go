package gink_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGink(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gink Suite")
}
