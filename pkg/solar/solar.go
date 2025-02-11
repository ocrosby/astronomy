package solar

import (
	"github.com/ocrosby/astronomy/pkg/constants"
	"math"
	"time"
)

// IsLeapYear checks if a year is a leap year
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// DaysInYear returns the number of days in a year
func DaysInYear(year int) int {
	if IsLeapYear(year) {
		return 366
	}
	return 365
}

// FractionalYear calculates the fractional year in radians
func FractionalYear(t time.Time) float64 {
	daysInYear := float64(DaysInYear(t.Year()))
	zeroBasedDayOfYear := float64(t.YearDay() - 1)
	hour := float64(t.Hour())
	return 2 * math.Pi / daysInYear * (zeroBasedDayOfYear + (hour-12.0)/24.0)
}

// EquationOfTime calculates the equation of time in minutes
func EquationOfTime(gamma float64) float64 {
	return 229.18 * (0.000075 + 0.001868*math.Cos(gamma) - 0.032077*math.Sin(gamma) - 0.014615*math.Cos(2*gamma) - 0.040849*math.Sin(2*gamma))
}

// SolarDeclination calculates the solar declination angle in radians
func SolarDeclination(gamma float64) float64 {
	return 0.006918 - 0.399912*math.Cos(gamma) + 0.070257*math.Sin(gamma) - 0.006758*math.Cos(2*gamma) + 0.000907*math.Sin(2*gamma) - 0.002697*math.Cos(3*gamma) + 0.00148*math.Sin(3*gamma)
}

// TimeOffset calculates the time offset in minutes
func TimeOffset(eqtime, longitude, timezone float64) float64 {
	return eqtime + 4*longitude - 60*timezone
}

// TrueSolarTime calculates the true solar time in minutes
func TrueSolarTime(hour, minute, second int, timeOffset float64) float64 {
	return float64(hour*60+minute) + float64(second)/60 + timeOffset
}

// SolarHourAngle calculates the solar hour angle in degrees
func SolarHourAngle(tst float64) float64 {
	return tst/4 - 180
}

// SolarZenithAngle calculates the solar zenith angle in radians
func SolarZenithAngle(lat, decl, ha float64) float64 {
	return math.Acos(math.Sin(lat*constants.Rad)*math.Sin(decl) + math.Cos(lat*constants.Rad)*math.Cos(decl)*math.Cos(ha*constants.Rad))
}

// SolarAzimuth calculates the solar azimuth angle in degrees
func SolarAzimuth(lat, decl, zenith float64) float64 {
	return math.Acos((math.Sin(lat*constants.Rad)*math.Cos(zenith)-math.Sin(decl))/(math.Cos(lat*constants.Rad)*math.Sin(zenith))) * constants.Deg
}

// SunriseSunsetHourAngle calculates the hour angle for sunrise or sunset
func SunriseSunsetHourAngle(lat, decl float64) float64 {
	return math.Acos((math.Cos(90.833*constants.Rad)/(math.Cos(lat*constants.Rad)*math.Cos(decl)) - math.Tan(lat*constants.Rad)*math.Tan(decl))) * constants.Deg
}

// Sunrise calculates the UTC time of sunrise in minutes
func Sunrise(longitude, ha, eqtime float64) float64 {
	return 720 - 4*(longitude+ha) - eqtime
}

// Sunset calculates the UTC time of sunset in minutes
func Sunset(longitude, ha, eqtime float64) float64 {
	return 720 - 4*(longitude-ha) - eqtime
}

// SolarNoon calculates the solar noon in minutes
func SolarNoon(longitude, eqtime float64) float64 {
	return 720 - 4*longitude - eqtime
}
