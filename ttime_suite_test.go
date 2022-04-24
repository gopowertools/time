package ttime_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTtime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ttime Suite")
}
