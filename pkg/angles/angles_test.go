package angles

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Angles", func() {
	Describe("AngleFormat", func() {
		DescribeTable("String representation",
			func(format AngleFormat, expected string) {
				Expect(format.String()).To(Equal(expected))
			},
			Entry("Dd format", Dd, "Dd"),
			Entry("DMM format", DMM, "DMM"),
			Entry("DMMm format", DMMm, "DMMm"),
			Entry("DMMSS format", DMMSS, "DMMSS"),
			Entry("DMMSSs format", DMMSSs, "DMMSSs"),
		)

		It("should have correct iota values", func() {
			Expect(int(Dd)).To(Equal(0))
			Expect(int(DMM)).To(Equal(1))
			Expect(int(DMMm)).To(Equal(2))
			Expect(int(DMMSS)).To(Equal(3))
			Expect(int(DMMSSs)).To(Equal(4))
		})
	})

	Describe("Angle", func() {
		Describe("NewAngle", func() {
			It("should create angle with default Dd format", func() {
				angle := NewAngle(15.5)
				Expect(angle.String()).To(Equal("15.50000°"))
			})

			It("should create angle with specified format", func() {
				angle := NewAngle(15.5, DMMSS)
				Expect(angle.String()).To(Equal("15°30'00\""))
			})
		})

		Describe("Set", func() {
			It("should set format to default Dd", func() {
				angle := NewAngle(15.5, DMMSS)
				angle.Set()
				Expect(angle.String()).To(Equal("15.50000°"))
			})

			It("should set format to specified value", func() {
				angle := NewAngle(15.5)
				angle.Set(DMM)
				Expect(angle.String()).To(Equal("15°30'"))
			})
		})

		Describe("String", func() {
			DescribeTable("formats angles correctly",
				func(alpha float64, format AngleFormat, expected string) {
					angle := NewAngle(alpha, format)
					Expect(angle.String()).To(Equal(expected))
				},
				Entry("Dd format", 15.5, Dd, "15.50000°"),
				Entry("DMM format", 15.5, DMM, "15°30'"),
				Entry("DMMm format", 15.516667, DMMm, "15°31.000'"),
				Entry("DMMSS format", 15.516667, DMMSS, "15°31'00\""),
				Entry("DMMSSs format", 15.516667, DMMSSs, "15°31'0.001\""),
				Entry("negative Dd", -8.15278, Dd, "-8.15278°"),
				Entry("negative DMM", -8.15278, DMM, "-8°09'"),
				Entry("negative DMMm", -8.15278, DMMm, "-8°9.167'"),
				Entry("negative DMMSS", -8.15278, DMMSS, "-8°09'10\""),
				Entry("negative DMMSSs", -8.15278, DMMSSs, "-8°09'10.008\""),
			)
		})
	})
	Describe("DegreesToRadians", func() {
		DescribeTable("converts degrees to radians correctly",
			func(degrees, expected float64) {
				Expect(DegreesToRadians(degrees)).To(Equal(expected))
			},
			Entry("0 degrees", 0.0, 0.0),
			Entry("90 degrees", 90.0, math.Pi/2),
			Entry("180 degrees", 180.0, math.Pi),
			Entry("360 degrees", 360.0, 2*math.Pi),
			Entry("-90 degrees", -90.0, -math.Pi/2),
		)
	})

	Describe("RadiansToDegrees", func() {
		DescribeTable("converts radians to degrees correctly",
			func(radians, expected float64) {
				Expect(RadiansToDegrees(radians)).To(Equal(expected))
			},
			Entry("0 radians", 0.0, 0.0),
			Entry("π/2 radians", math.Pi/2, 90.0),
			Entry("π radians", math.Pi, 180.0),
			Entry("2π radians", 2*math.Pi, 360.0),
			Entry("-π/2 radians", -math.Pi/2, -90.0),
		)
	})

	Describe("NormalizeDegrees", func() {
		DescribeTable("normalizes degrees to 0-360 range",
			func(degrees, expected float64) {
				Expect(NormalizeDegrees(degrees)).To(Equal(expected))
			},
			Entry("0 degrees", 0.0, 0.0),
			Entry("360 degrees", 360.0, 0.0),
			Entry("720 degrees", 720.0, 0.0),
			Entry("-360 degrees", -360.0, 0.0),
			Entry("450 degrees", 450.0, 90.0),
			Entry("-450 degrees", -450.0, 270.0),
			Entry("1080 degrees", 1080.0, 0.0),
		)
	})

	Describe("Ddd", func() {
		DescribeTable("converts DMS to decimal degrees",
			func(degrees, minutes int, seconds, expected float64) {
				Expect(Ddd(degrees, minutes, seconds)).To(BeNumerically("~", expected, 1e-5))
			},
			Entry("15°30'0\"", 15, 30, 0.00, 15.50000),
			Entry("-8°09'10\"", -8, 9, 10.0, -8.15278),
			Entry("0°01'0\"", 0, 1, 0.0, 0.01667),
			Entry("0°-05'0\"", 0, -5, 0.0, -0.08334),
		)
	})

	Describe("DMS", func() {
		DescribeTable("converts decimal degrees to DMS using pointers",
			func(decimal float64, expectedDeg, expectedMin int, expectedSec float64) {
				var degrees int
				var minutes int
				var seconds float64
				DMS(decimal, &degrees, &minutes, &seconds)
				Expect(degrees).To(Equal(expectedDeg))
				Expect(minutes).To(Equal(expectedMin))
				Expect(seconds).To(BeNumerically("~", expectedSec, 0.1))
			},
			Entry("15.50000", 15.50000, 15, 30, 0.00),
			Entry("-8.15278", -8.15278, -8, 9, 10.0),
			Entry("0.01667", 0.01667, 0, 1, 0.012),
			Entry("-0.08334", -0.08334, 0, -5, -0.024),
		)
	})
})
