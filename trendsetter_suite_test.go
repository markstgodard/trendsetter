package trend_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTrendsetter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Trendsetter Suite")
}
