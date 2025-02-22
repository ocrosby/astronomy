package solar

import (
	"github.com/ocrosby/astronomy/pkg/constants"
	"math"
	"testing"
	"time"
)

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		year     int
		expected bool
	}{
		{2000, true},  // divisible by 400
		{1900, false}, // divisible by 100 but not by 400
		{2004, true},  // divisible by 4 but not by 100
		{2001, false}, // not divisible by 4
	}

	for _, tt := range tests {
		result := IsLeapYear(tt.year)
		if result != tt.expected {
			t.Errorf("IsLeapYear(%v) = %v; want %v", tt.year, result, tt.expected)
		}
	}
}

func TestFractionalYear(t *testing.T) {
	tests := []struct {
		date     time.Time
		expected float64
	}{
		{time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), -0.00860710316051998},    // start of the year
		{time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC), 6.2738609454223555}, // end of the year
		{time.Date(2023, 6, 21, 12, 0, 0, 0, time.UTC), 2.9436292808978335},    // middle of the year
		{time.Date(2020, 2, 29, 0, 0, 0, 0, time.UTC), 1.0042796187705076},     // leap year
		{time.Date(2023, 2, 28, 0, 0, 0, 0, time.UTC), 0.9898168634597978},     // non-leap year
		{time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC), 6.257363997698026},     // last day of the year
	}

	for _, tt := range tests {
		result := FractionalYear(tt.date)
		if math.Abs(result-tt.expected) > 1e-6 {
			t.Errorf("FractionalYear(%v) = %v; want %v", tt.date, result, tt.expected)
		}
	}
}

func TestEquationOfTime(t *testing.T) {
	tests := []struct {
		gamma    float64
		expected float64
	}{
		{1.0, 229.18 * (0.000075 + 0.001868*math.Cos(1.0) - 0.032077*math.Sin(1.0) - 0.014615*math.Cos(2*1.0) - 0.040849*math.Sin(2*1.0))},
	}

	for _, tt := range tests {
		result := EquationOfTime(tt.gamma)
		if math.Abs(result-tt.expected) > 1e-6 {
			t.Errorf("EquationOfTime(%v) = %v; want %v", tt.gamma, result, tt.expected)
		}
	}
}

func TestSolarDeclination(t *testing.T) {
	tests := []struct {
		gamma    float64
		expected float64
	}{
		{1.0, 0.006918 - 0.399912*math.Cos(1.0) + 0.070257*math.Sin(1.0) - 0.006758*math.Cos(2*1.0) + 0.000907*math.Sin(2*1.0) - 0.002697*math.Cos(3*1.0) + 0.00148*math.Sin(3*1.0)},
	}

	for _, tt := range tests {
		result := SolarDeclination(tt.gamma)
		if math.Abs(result-tt.expected) > 1e-6 {
			t.Errorf("SolarDeclination(%v) = %v; want %v", tt.gamma, result, tt.expected)
		}
	}
}

func TestTimeOffset(t *testing.T) {
	tests := []struct {
		eqtime    float64
		longitude float64
		timezone  float64
		expected  float64
	}{
		{3.0, -74.0060, -5.0, 3.0 + 4*(-74.0060) - 60*(-5.0)},
	}

	for _, tt := range tests {
		result := TimeOffset(tt.eqtime, tt.longitude, tt.timezone)
		if math.Abs(result-tt.expected) > 1e-6 {
			t.Errorf("TimeOffset(%v, %v, %v) = %v; want %v", tt.eqtime, tt.longitude, tt.timezone, result, tt.expected)
		}
	}
}

func TestTrueSolarTime(t *testing.T) {
	tests := []struct {
		hour       int
		minute     int
		second     int
		timeOffset float64
		expected   float64
	}{
		{12, 0, 0, 3.0, 12*60 + 0 + 0.0/60 + 3.0},
	}

	for _, tt := range tests {
		result := TrueSolarTime(tt.hour, tt.minute, tt.second, tt.timeOffset)
		if math.Abs(result-tt.expected) > 1e-6 {
			t.Errorf("TrueSolarTime(%v, %v, %v, %v) = %v; want %v", tt.hour, tt.minute, tt.second, tt.timeOffset, result, tt.expected)
		}
	}
}

func TestSolarHourAngle(t *testing.T) {
	tests := []struct {
		tst      float64
		expected float64
	}{
		{720.0, 720.0/4 - 180},   // typical case
		{0.0, 0.0/4 - 180},       // edge case: start of the day
		{1440.0, 1440.0/4 - 180}, // edge case: end of the day
		{-720.0, -720.0/4 - 180}, // error case: negative time
		{2880.0, 2880.0/4 - 180}, // error case: time beyond 24 hours
	}

	for _, tt := range tests {
		result := SolarHourAngle(tt.tst)
		if math.Abs(result-tt.expected) > 1e-6 {
			t.Errorf("SolarHourAngle(%v) = %v; want %v", tt.tst, result, tt.expected)
		}
	}
}

func TestSolarZenithAngle(t *testing.T) {
	tests := []struct {
		lat      float64
		decl     float64
		ha       float64
		expected float64
	}{
		{
			40.7128,
			0.4091,
			0.0,
			math.Acos(math.Sin(40.7128*constants.Rad)*math.Sin(0.4091) + math.Cos(40.7128*constants.Rad)*math.Cos(0.4091)*math.Cos(0.0*constants.Rad)),
		},
	}

	for _, tt := range tests {
		result := SolarZenithAngle(tt.lat, tt.decl, tt.ha)
		if math.Abs(result-tt.expected) > 1e-6 {
			t.Errorf("SolarZenithAngle(%v, %v, %v) = %v; want %v", tt.lat, tt.decl, tt.ha, result, tt.expected)
		}
	}
}

func TestSolarAzimuth(t *testing.T) {
	tests := []struct {
		lat      float64
		decl     float64
		zenith   float64
		expected float64
	}{
		{40.7128, 0.4091, 1.0, math.Acos((math.Sin(40.7128*constants.Rad)*math.Cos(1.0)-math.Sin(0.4091))/(math.Cos(40.7128*constants.Rad)*math.Sin(1.0))) * constants.Deg},
	}

	for _, tt := range tests {
		result := SolarAzimuth(tt.lat, tt.decl, tt.zenith)
		if math.Abs(result-tt.expected) > 1e-6 {
			t.Errorf("SolarAzimuth(%v, %v, %v) = %v; want %v", tt.lat, tt.decl, tt.zenith, result, tt.expected)
		}
	}
}

func TestSunriseSunsetHourAngle(t *testing.T) {
	tests := []struct {
		lat      float64
		decl     float64
		expected float64
	}{
		{40.7128, 0.4091, math.Acos((math.Cos(90.833*constants.Rad)/(math.Cos(40.7128*constants.Rad)*math.Cos(0.4091)) - math.Tan(40.7128*constants.Rad)*math.Tan(0.4091))) * constants.Deg},
	}

	for _, tt := range tests {
		result := SunriseSunsetHourAngle(tt.lat, tt.decl)
		if math.Abs(result-tt.expected) > 1e-6 {
			t.Errorf("SunriseSunsetHourAngle(%v, %v) = %v; want %v", tt.lat, tt.decl, result, tt.expected)
		}
	}
}

func TestSunrise(t *testing.T) {
	tests := []struct {
		longitude float64
		ha        float64
		eqtime    float64
		expected  float64
	}{
		{-74.0060, 1.0, 3.0, 720 - 4*(-74.0060+1.0) - 3.0},
	}

	for _, tt := range tests {
		result := Sunrise(tt.longitude, tt.ha, tt.eqtime)
		if math.Abs(result-tt.expected) > 1e-6 {
			t.Errorf("Sunrise(%v, %v, %v) = %v; want %v", tt.longitude, tt.ha, tt.eqtime, result, tt.expected)
		}
	}
}

func TestSunset(t *testing.T) {
	tests := []struct {
		longitude float64
		ha        float64
		eqtime    float64
		expected  float64
	}{
		{-74.0060, 1.0, 3.0, 720 - 4*(-74.0060-1.0) - 3.0},
	}

	for _, tt := range tests {
		result := Sunset(tt.longitude, tt.ha, tt.eqtime)
		if math.Abs(result-tt.expected) > 1e-6 {
			t.Errorf("Sunset(%v, %v, %v) = %v; want %v", tt.longitude, tt.ha, tt.eqtime, result, tt.expected)
		}
	}
}

func TestSolarNoon(t *testing.T) {
	tests := []struct {
		longitude float64
		eqtime    float64
		expected  float64
	}{
		{-74.0060, 3.0, 720 - 4*(-74.0060) - 3.0},
	}

	for _, tt := range tests {
		result := SolarNoon(tt.longitude, tt.eqtime)
		if math.Abs(result-tt.expected) > 1e-6 {
			t.Errorf("SolarNoon(%v, %v) = %v; want %v", tt.longitude, tt.eqtime, result, tt.expected)
		}
	}
}
