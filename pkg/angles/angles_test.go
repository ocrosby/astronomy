package angles

import (
	"math"
	"strings"

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

	Describe("AngleFormatter", func() {
		Describe("fluent interface with 12.3456", func() {
			It("should format as 12.35 (Dd with precision 2)", func() {
				result := NewFormatter(12.3456).Format(Dd).Precision(2).String()
				Expect(result).To(Equal("12.35"))
			})

			It("should format as '12 20' (DMM)", func() {
				result := NewFormatter(12.3456).Format(DMM).String()
				Expect(result).To(Equal("12 20"))
			})

			It("should format as '12 20.74' (DMMm with precision 2)", func() {
				result := NewFormatter(12.3456).Format(DMMm).Precision(2).String()
				Expect(result).To(Equal("12 20.74"))
			})

			It("should format as '12 20 44' (DMMSS)", func() {
				result := NewFormatter(12.3456).Format(DMMSS).String()
				Expect(result).To(Equal("12 20 44"))
			})

			It("should format as '12 20 44.16' (DMMSSs with precision 2)", func() {
				result := NewFormatter(12.3456).Format(DMMSSs).Precision(2).String()
				Expect(result).To(Equal("12 20 44.16"))
			})
		})

		Describe("width formatting", func() {
			It("should apply left justification with width", func() {
				result := NewFormatter(12.3456).Format(Dd).Precision(2).Width(10).String()
				Expect(result).To(Equal("12.35     "))
				Expect(len(result)).To(Equal(10))
			})

			It("should handle negative values with width", func() {
				result := NewFormatter(-12.3456).Format(Dd).Precision(2).Width(12).String()
				Expect(result).To(Equal("-12.35      "))
				Expect(len(result)).To(Equal(12))
			})
		})

		Describe("negative angle handling", func() {
			It("should handle negative angles in DMM format", func() {
				result := NewFormatter(-0.3456).Format(DMM).String()
				Expect(result).To(Equal("0 -20"))
			})

			It("should handle negative angles in DMMm format", func() {
				result := NewFormatter(-0.3456).Format(DMMm).Precision(2).String()
				Expect(result).To(Equal("0 -20.74"))
			})

			It("should handle negative angles in DMMSS format", func() {
				result := NewFormatter(-0.3456).Format(DMMSS).String()
				Expect(result).To(Equal("0 -20 44"))
			})

			It("should handle negative angles in DMMSSs format", func() {
				result := NewFormatter(-0.3456).Format(DMMSSs).Precision(2).String()
				Expect(result).To(Equal("0 -20 44.16"))
			})
		})

		Describe("chaining methods", func() {
			It("should allow method chaining in any order", func() {
				result1 := NewFormatter(12.3456).Precision(3).Format(DMMSSs).Width(15).String()
				result2 := NewFormatter(12.3456).Width(15).Format(DMMSSs).Precision(3).String()
				Expect(result1).To(Equal(result2))
				Expect(len(result1)).To(Equal(15))
			})
		})
	})

	Describe("ParseAngle", func() {
		Describe("valid formats", func() {
			It("should parse Dd format", func() {
				angle, err := ParseAngle("12.35")
				Expect(err).To(BeNil())
				Expect(angle.alpha).To(BeNumerically("~", 12.35, 1e-6))
				Expect(angle.format).To(Equal(Dd))
			})

			It("should parse DMM format", func() {
				angle, err := ParseAngle("12 20")
				Expect(err).To(BeNil())
				Expect(angle.alpha).To(BeNumerically("~", 12.333333, 1e-6))
				Expect(angle.format).To(Equal(DMM))
			})

			It("should parse DMMm format", func() {
				angle, err := ParseAngle("12 20.74")
				Expect(err).To(BeNil())
				Expect(angle.alpha).To(BeNumerically("~", 12.3456666, 1e-6))
				Expect(angle.format).To(Equal(DMMm))
			})

			It("should parse DMMSS format", func() {
				angle, err := ParseAngle("12 20 44")
				Expect(err).To(BeNil())
				Expect(angle.alpha).To(BeNumerically("~", 12.345555, 1e-6))
				Expect(angle.format).To(Equal(DMMSS))
			})

			It("should parse DMMSSs format", func() {
				angle, err := ParseAngle("12 20 44.16")
				Expect(err).To(BeNil())
				Expect(angle.alpha).To(BeNumerically("~", 12.3456, 1e-6))
				Expect(angle.format).To(Equal(DMMSSs))
			})
		})

		Describe("negative angles", func() {
			It("should parse negative Dd format", func() {
				angle, err := ParseAngle("-12.35")
				Expect(err).To(BeNil())
				Expect(angle.alpha).To(BeNumerically("~", -12.35, 1e-6))
				Expect(angle.format).To(Equal(Dd))
			})

			It("should parse negative DMM format", func() {
				angle, err := ParseAngle("-12 20")
				Expect(err).To(BeNil())
				Expect(angle.alpha).To(BeNumerically("~", -12.333333, 1e-6))
				Expect(angle.format).To(Equal(DMM))
			})

			It("should parse small negative angles", func() {
				angle, err := ParseAngle("0 -20")
				Expect(err).To(BeNil())
				Expect(angle.alpha).To(BeNumerically("~", -0.333333, 1e-6))
				Expect(angle.format).To(Equal(DMM))
			})

			It("should parse negative seconds", func() {
				angle, err := ParseAngle("0 -20 44.16")
				Expect(err).To(BeNil())
				Expect(angle.alpha).To(BeNumerically("~", -0.345600, 1e-6))
				Expect(angle.format).To(Equal(DMMSSs))
			})
		})

		Describe("round-trip compatibility", func() {
			It("should round-trip through formatter", func() {
				original := 12.3456
				formatted := NewFormatter(original).Format(DMMSSs).Precision(2).String()
				parsed, err := ParseAngle(formatted)

				Expect(err).To(BeNil())
				Expect(parsed.alpha).To(BeNumerically("~", original, 1e-6))
				Expect(parsed.format).To(Equal(DMMSSs))
			})

			It("should round-trip negative angles with precision loss", func() {
				original := -0.3456
				formatted := NewFormatter(original).Format(DMM).String() // "0 -20"
				parsed, err := ParseAngle(formatted)

				Expect(err).To(BeNil())
				// DMM format truncates to whole minutes, so -0.3456 becomes -20/60 = -0.3333...
				Expect(parsed.alpha).To(BeNumerically("~", -20.0/60.0, 1e-6))
				Expect(parsed.format).To(Equal(DMM))
			})

			It("should round-trip negative angles precisely with DMMSSs", func() {
				original := -0.3456
				formatted := NewFormatter(original).Format(DMMSSs).Precision(2).String()
				parsed, err := ParseAngle(formatted)

				Expect(err).To(BeNil())
				Expect(parsed.alpha).To(BeNumerically("~", original, 1e-4))
				Expect(parsed.format).To(Equal(DMMSSs))
			})
		})

		Describe("error cases", func() {
			Describe("basic validation", func() {
				It("should reject empty strings", func() {
					_, err := ParseAngle("")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("empty input"))
				})

				It("should reject whitespace-only strings", func() {
					_, err := ParseAngle("   ")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("whitespace"))
				})

				It("should reject too many components", func() {
					_, err := ParseAngle("12 20 44 16")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("expected 1-3"))
				})
			})

			Describe("invalid characters", func() {
				It("should reject alphabetic characters in decimal format", func() {
					_, err := ParseAngle("12abc")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("invalid decimal degrees"))
				})

				It("should reject special symbols", func() {
					_, err := ParseAngle("12@34")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("invalid character"))
				})

				It("should reject mixed invalid characters", func() {
					_, err := ParseAngle("12 20# 44")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("invalid character"))
				})
			})

			Describe("decimal format errors", func() {
				It("should reject multiple decimal points", func() {
					_, err := ParseAngle("12.34.56")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("multiple decimal points"))
				})

				It("should reject multiple signs", func() {
					_, err := ParseAngle("-+12.34")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("multiple signs"))
				})

				It("should reject misplaced signs", func() {
					_, err := ParseAngle("12.-34")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("sign must be at beginning"))
				})
			})

			Describe("component validation errors", func() {
				It("should reject invalid degrees", func() {
					_, err := ParseAngle("abc 20")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("invalid degrees"))
				})

				It("should reject degrees with decimal point in integer format", func() {
					_, err := ParseAngle("12.5 20")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("unexpected decimal point"))
				})

				It("should reject invalid minutes", func() {
					_, err := ParseAngle("12 abc")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("invalid minutes"))
				})

				It("should reject invalid seconds", func() {
					_, err := ParseAngle("12 20 abc")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("invalid seconds"))
				})

				It("should reject non-numeric degrees", func() {
					_, err := ParseAngle("xyz 20")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("invalid degrees"))
				})

				It("should reject non-numeric minutes", func() {
					_, err := ParseAngle("12 xyz")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("invalid minutes"))
				})

				It("should reject non-numeric seconds", func() {
					_, err := ParseAngle("12 20 xyz")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("invalid seconds"))
				})
			})

			Describe("range validation errors", func() {
				It("should reject minutes >= 60 in DMM format", func() {
					_, err := ParseAngle("12 60")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("must be less than 60"))
				})

				It("should reject minutes >= 60 in DMMm format", func() {
					_, err := ParseAngle("12 60.5")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("must be less than 60"))
				})

				It("should reject minutes >= 60 in DMMSS format", func() {
					_, err := ParseAngle("12 60 30")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("must be less than 60"))
				})

				It("should reject seconds >= 60 in DMMSS format", func() {
					_, err := ParseAngle("12 30 60")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("must be less than 60"))
				})

				It("should reject seconds >= 60 in DMMSSs format", func() {
					_, err := ParseAngle("12 30 60.5")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("must be less than 60"))
				})

				It("should reject negative minutes >= 60", func() {
					_, err := ParseAngle("12 -60")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("must be less than 60"))
				})
			})

			Describe("malformed number errors", func() {
				It("should reject multiple decimal points in minutes", func() {
					_, err := ParseAngle("12 20.34.56")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("multiple decimal points"))
				})

				It("should reject multiple signs in components", func() {
					_, err := ParseAngle("12 -+20")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("multiple signs"))
				})

				It("should reject misplaced signs in components", func() {
					_, err := ParseAngle("12 2-0")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("sign must be at beginning"))
				})
			})

			Describe("edge case errors", func() {
				It("should reject infinity", func() {
					_, err := ParseAngle("inf")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("infinite or NaN"))
				})

				It("should reject NaN", func() {
					_, err := ParseAngle("NaN")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("infinite or NaN"))
				})

				It("should reject +inf", func() {
					_, err := ParseAngle("+inf")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("infinite or NaN"))
				})

				It("should reject -inf", func() {
					_, err := ParseAngle("-inf")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("infinite or NaN"))
				})

				It("should handle very long invalid strings gracefully", func() {
					longInvalid := strings.Repeat("xyz", 100) // Use valid chars but invalid format
					_, err := ParseAngle(longInvalid)
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("invalid decimal degrees"))
				})
			})
		})

		Describe("whitespace handling", func() {
			It("should handle leading/trailing whitespace", func() {
				angle, err := ParseAngle("  12.35  ")
				Expect(err).To(BeNil())
				Expect(angle.alpha).To(BeNumerically("~", 12.35, 1e-6))
			})

			It("should handle multiple spaces between components", func() {
				angle, err := ParseAngle("12   20    44.16")
				Expect(err).To(BeNil())
				Expect(angle.alpha).To(BeNumerically("~", 12.3456, 1e-6))
				Expect(angle.format).To(Equal(DMMSSs))
			})
		})
	})
})
