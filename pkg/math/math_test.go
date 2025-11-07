package math

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Math", func() {
	Describe("Frac", func() {
		DescribeTable("returns fractional part of number",
			func(input, expected float64) {
				Expect(Frac(input)).To(BeNumerically("~", expected, 1e-10))
			},
			Entry("positive decimal", 3.14, 0.14),
			Entry("negative decimal", -3.14, 0.86),
			Entry("whole number", 5.0, 0.0),
			Entry("negative whole number", -5.0, 0.0),
			Entry("zero", 0.0, 0.0),
			Entry("small positive", 0.25, 0.25),
			Entry("small negative", -0.75, 0.25),
		)
	})

	Describe("Mod", func() {
		DescribeTable("returns floating-point modulo",
			func(x, y, expected float64) {
				Expect(Mod(x, y)).To(BeNumerically("~", expected, 1e-10))
			},
			Entry("positive x, positive y", 7.5, 3.0, 1.5),
			Entry("negative x, positive y", -7.5, 3.0, 1.5),
			Entry("positive x, negative y", 7.5, -3.0, -1.5),
			Entry("negative x, negative y", -7.5, -3.0, -1.5),
			Entry("x smaller than y", 2.5, 5.0, 2.5),
			Entry("exact division", 6.0, 3.0, 0.0),
			Entry("fractional result", 10.7, 3.2, 1.1),
		)
	})

	Describe("Fabs", func() {
		DescribeTable("returns absolute value",
			func(input, expected float64) {
				Expect(Fabs(input)).To(Equal(expected))
			},
			Entry("positive number", 5.5, 5.5),
			Entry("negative number", -5.5, 5.5),
			Entry("zero", 0.0, 0.0),
			Entry("negative zero", -0.0, 0.0),
			Entry("large positive", 1234.5678, 1234.5678),
			Entry("large negative", -1234.5678, 1234.5678),
		)
	})

	Describe("Sign", func() {
		DescribeTable("returns sign of DMS components",
			func(degrees, minutes int, seconds, expected float64) {
				Expect(Sign(degrees, minutes, seconds)).To(Equal(expected))
			},
			Entry("all positive", 12, 34, 56.0, 1.0),
			Entry("all zero", 0, 0, 0.0, 1.0),
			Entry("negative degrees", -12, 34, 56.0, -1.0),
			Entry("negative minutes", 12, -34, 56.0, -1.0),
			Entry("negative seconds", 12, 34, -56.0, -1.0),
			Entry("all negative", -12, -34, -56.0, -1.0),
			Entry("mixed negative", -12, 34, -56.0, -1.0),
		)
	})
})
