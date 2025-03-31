package go_design_patterns_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoDesignPatterns(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoDesignPatterns Suite")
}
