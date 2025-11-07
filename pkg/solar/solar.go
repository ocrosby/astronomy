package solar

import (
	"github.com/ocrosby/astronomy/pkg/constants"
	"math"
	"time"
)

// Solar calculation constants
const (
	// Equation of Time coefficients
	EqTimeBase   = 229.18
	EqTimeCoeff1 = 0.000075
	EqTimeCoeff2 = 0.001868
	EqTimeCoeff3 = -0.032077
	EqTimeCoeff4 = -0.014615
	EqTimeCoeff5 = -0.040849

	// Solar declination coefficients
	DeclCoeff1 = 0.006918
	DeclCoeff2 = -0.399912
	DeclCoeff3 = 0.070257
	DeclCoeff4 = -0.006758
	DeclCoeff5 = 0.000907
	DeclCoeff6 = -0.002697
	DeclCoeff7 = 0.00148

	// Solar calculations
	SunriseAngle    = 90.833 // degrees
	HourAngleDiv    = 4.0    // divisor for hour angle calculation
	HourAngleOffset = 180.0  // offset for hour angle
	TimeBase        = 720.0  // base time in minutes
	LongitudeFactor = 4.0    // minutes per degree of longitude
	TimezoneFactor  = 60.0   // minutes per hour of timezone
	NoonHour        = 12.0   // solar noon reference hour
	HoursPerDay     = 24.0   // hours per day
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
	return 2 * math.Pi / daysInYear * (zeroBasedDayOfYear + (hour-NoonHour)/HoursPerDay)
}

// EquationOfTime calculates the equation of time in minutes
func EquationOfTime(gamma float64) float64 {
	return EqTimeBase * (EqTimeCoeff1 + EqTimeCoeff2*math.Cos(gamma) + EqTimeCoeff3*math.Sin(gamma) + EqTimeCoeff4*math.Cos(2*gamma) + EqTimeCoeff5*math.Sin(2*gamma))
}

// SolarDeclination calculates the solar declination angle in radians
func SolarDeclination(gamma float64) float64 {
	return DeclCoeff1 + DeclCoeff2*math.Cos(gamma) + DeclCoeff3*math.Sin(gamma) + DeclCoeff4*math.Cos(2*gamma) + DeclCoeff5*math.Sin(2*gamma) + DeclCoeff6*math.Cos(3*gamma) + DeclCoeff7*math.Sin(3*gamma)
}

// TimeOffset calculates the time offset in minutes
func TimeOffset(eqtime, longitude, timezone float64) float64 {
	return eqtime + LongitudeFactor*longitude - TimezoneFactor*timezone
}

// TrueSolarTime calculates the true solar time in minutes
func TrueSolarTime(hour, minute, second int, timeOffset float64) float64 {
	return float64(hour*60+minute) + float64(second)/60 + timeOffset
}

// SolarHourAngle calculates the solar hour angle in degrees
func SolarHourAngle(tst float64) float64 {
	return tst/HourAngleDiv - HourAngleOffset
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
	return math.Acos((math.Cos(SunriseAngle*constants.Rad)/(math.Cos(lat*constants.Rad)*math.Cos(decl)) - math.Tan(lat*constants.Rad)*math.Tan(decl))) * constants.Deg
}

// Sunrise calculates the UTC time of sunrise in minutes
func Sunrise(longitude, ha, eqtime float64) float64 {
	return TimeBase - LongitudeFactor*(longitude+ha) - eqtime
}

// Sunset calculates the UTC time of sunset in minutes
func Sunset(longitude, ha, eqtime float64) float64 {
	return TimeBase - LongitudeFactor*(longitude-ha) - eqtime
}

// SolarNoon calculates the solar noon in minutes
func SolarNoon(longitude, eqtime float64) float64 {
	return TimeBase - LongitudeFactor*longitude - eqtime
}
