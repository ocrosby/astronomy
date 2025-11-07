package solar

import (
	"math"
	"time"

	"github.com/ocrosby/astronomy/pkg/constants"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Solar", func() {
	Describe("IsLeapYear", func() {
		DescribeTable("correctly identifies leap years",
			func(year int, expected bool) {
				Expect(IsLeapYear(year)).To(Equal(expected))
			},
			Entry("year 2000 (divisible by 400)", 2000, true),
			Entry("year 1900 (divisible by 100 but not 400)", 1900, false),
			Entry("year 2004 (divisible by 4 but not 100)", 2004, true),
			Entry("year 2001 (not divisible by 4)", 2001, false),
		)
	})

	Describe("FractionalYear", func() {
		DescribeTable("calculates fractional year correctly",
			func(date time.Time, expected float64) {
				result := FractionalYear(date)
				Expect(math.Abs(result - expected)).To(BeNumerically("<", 1e-6))
			},
			Entry("start of year 2023", time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), -0.00860710316051998),
			Entry("end of year 2023", time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC), 6.2738609454223555),
			Entry("middle of year 2023", time.Date(2023, 6, 21, 12, 0, 0, 0, time.UTC), 2.9436292808978335),
			Entry("leap year date", time.Date(2020, 2, 29, 0, 0, 0, 0, time.UTC), 1.0042796187705076),
			Entry("non-leap year Feb 28", time.Date(2023, 2, 28, 0, 0, 0, 0, time.UTC), 0.9898168634597978),
			Entry("last day of year", time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC), 6.257363997698026),
		)
	})

	Describe("EquationOfTime", func() {
		It("calculates equation of time correctly", func() {
			gamma := 1.0
			expected := 229.18 * (0.000075 + 0.001868*math.Cos(gamma) - 0.032077*math.Sin(gamma) - 0.014615*math.Cos(2*gamma) - 0.040849*math.Sin(2*gamma))
			result := EquationOfTime(gamma)
			Expect(math.Abs(result - expected)).To(BeNumerically("<", 1e-6))
		})
	})

	Describe("SolarDeclination", func() {
		It("calculates solar declination correctly", func() {
			gamma := 1.0
			expected := 0.006918 - 0.399912*math.Cos(gamma) + 0.070257*math.Sin(gamma) - 0.006758*math.Cos(2*gamma) + 0.000907*math.Sin(2*gamma) - 0.002697*math.Cos(3*gamma) + 0.00148*math.Sin(3*gamma)
			result := SolarDeclination(gamma)
			Expect(math.Abs(result - expected)).To(BeNumerically("<", 1e-6))
		})
	})

	Describe("TimeOffset", func() {
		It("calculates time offset correctly", func() {
			eqtime := 3.0
			longitude := -74.0060
			timezone := -5.0
			expected := eqtime + 4*longitude - 60*timezone
			result := TimeOffset(eqtime, longitude, timezone)
			Expect(math.Abs(result - expected)).To(BeNumerically("<", 1e-6))
		})
	})

	Describe("TrueSolarTime", func() {
		It("calculates true solar time correctly", func() {
			hour := 12
			minute := 0
			second := 0
			timeOffset := 3.0
			expected := float64(hour*60+minute) + float64(second)/60 + timeOffset
			result := TrueSolarTime(hour, minute, second, timeOffset)
			Expect(math.Abs(result - expected)).To(BeNumerically("<", 1e-6))
		})
	})

	Describe("SolarHourAngle", func() {
		DescribeTable("calculates solar hour angle correctly",
			func(tst, expected float64) {
				result := SolarHourAngle(tst)
				Expect(math.Abs(result - expected)).To(BeNumerically("<", 1e-6))
			},
			Entry("typical case", 720.0, 720.0/4-180),
			Entry("start of day", 0.0, 0.0/4-180),
			Entry("end of day", 1440.0, 1440.0/4-180),
			Entry("negative time", -720.0, -720.0/4-180),
			Entry("time beyond 24 hours", 2880.0, 2880.0/4-180),
		)
	})

	Describe("SolarZenithAngle", func() {
		It("calculates solar zenith angle correctly", func() {
			lat := 40.7128
			decl := 0.4091
			ha := 0.0
			expected := math.Acos(math.Sin(lat*constants.Rad)*math.Sin(decl) + math.Cos(lat*constants.Rad)*math.Cos(decl)*math.Cos(ha*constants.Rad))
			result := SolarZenithAngle(lat, decl, ha)
			Expect(math.Abs(result - expected)).To(BeNumerically("<", 1e-6))
		})
	})

	Describe("SolarAzimuth", func() {
		It("calculates solar azimuth correctly", func() {
			lat := 40.7128
			decl := 0.4091
			zenith := 1.0
			expected := math.Acos((math.Sin(lat*constants.Rad)*math.Cos(zenith)-math.Sin(decl))/(math.Cos(lat*constants.Rad)*math.Sin(zenith))) * constants.Deg
			result := SolarAzimuth(lat, decl, zenith)
			Expect(math.Abs(result - expected)).To(BeNumerically("<", 1e-6))
		})
	})

	Describe("SunriseSunsetHourAngle", func() {
		It("calculates sunrise/sunset hour angle correctly", func() {
			lat := 40.7128
			decl := 0.4091
			expected := math.Acos((math.Cos(90.833*constants.Rad)/(math.Cos(lat*constants.Rad)*math.Cos(decl)) - math.Tan(lat*constants.Rad)*math.Tan(decl))) * constants.Deg
			result := SunriseSunsetHourAngle(lat, decl)
			Expect(math.Abs(result - expected)).To(BeNumerically("<", 1e-6))
		})
	})

	Describe("Sunrise", func() {
		It("calculates sunrise time correctly", func() {
			longitude := -74.0060
			ha := 1.0
			eqtime := 3.0
			expected := 720 - 4*(longitude+ha) - eqtime
			result := Sunrise(longitude, ha, eqtime)
			Expect(math.Abs(result - expected)).To(BeNumerically("<", 1e-6))
		})
	})

	Describe("Sunset", func() {
		It("calculates sunset time correctly", func() {
			longitude := -74.0060
			ha := 1.0
			eqtime := 3.0
			expected := 720 - 4*(longitude-ha) - eqtime
			result := Sunset(longitude, ha, eqtime)
			Expect(math.Abs(result - expected)).To(BeNumerically("<", 1e-6))
		})
	})

	Describe("SolarNoon", func() {
		It("calculates solar noon correctly", func() {
			longitude := -74.0060
			eqtime := 3.0
			expected := 720 - 4*longitude - eqtime
			result := SolarNoon(longitude, eqtime)
			Expect(math.Abs(result - expected)).To(BeNumerically("<", 1e-6))
		})
	})
})
