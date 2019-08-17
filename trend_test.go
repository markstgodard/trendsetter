package trend_test

import (
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	trend "github.com/markstgodard/trendsetter"
)

var _ = Describe("Trend", func() {

	var (
		data []float64
	)
	Context("when just a single number", func() {
		JustBeforeEach(func() {
			data = []float64{1.0}
		})

		It("returns the single number for slope and y intercept of a trend line", func() {
			m, y := trend.Calc(data)
			Expect(m).To(Equal(1.0))
			Expect(y).To(Equal(0.0))
		})
	})

	Context("when multiple data points", func() {
		JustBeforeEach(func() {
			data = []float64{3, 5, 6}
		})

		It("calculates the slope and y intercept of a trend line", func() {
			m, y := trend.Calc(data)
			Expect(math.Round(m*100) / 100).To(Equal(1.5))
			Expect(math.Round(y*100) / 100).To(Equal(1.67))
		})
	})
})
