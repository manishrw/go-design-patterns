package singleton_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSingletonSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Singleton Pattern Suite")
}
