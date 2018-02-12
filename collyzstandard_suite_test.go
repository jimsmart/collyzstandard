package collyzstandard_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCollyzstandard(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CollyZstandard Suite")
}
