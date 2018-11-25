package goerkin_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoerkin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goerkin Suite")
}
