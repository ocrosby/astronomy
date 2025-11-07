package angles

import (
	"fmt"
	"github.com/ocrosby/astronomy/pkg/constants"
	"math"
)

// AngleFormat represents different angle representation formats
type AngleFormat int

const (
	Dd     AngleFormat = iota // decimal representation
	DMM                       // degrees and whole minutes of arc
	DMMm                      // degrees and minutes of arc in decimal representation
	DMMSS                     // degrees, minutes of arc and whole seconds of arc
	DMMSSs                    // degrees, minutes, and seconds of arc in decimal representation
)

// String returns the string representation of AngleFormat
func (af AngleFormat) String() string {
	return [...]string{"Dd", "DMM", "DMMm", "DMMSS", "DMMSSs"}[af]
}

// Angle represents a sexagesimal angle output
type Angle struct {
	alpha  float64
	format AngleFormat
}

// NewAngle creates a new Angle instance with default format Dd
func NewAngle(alpha float64, format ...AngleFormat) *Angle {
	f := Dd // default format
	if len(format) > 0 {
		f = format[0]
	}
	return &Angle{
		alpha:  alpha,
		format: f,
	}
}

// Set sets the angle format, defaulting to Dd
func (a *Angle) Set(format ...AngleFormat) {
	f := Dd // default format
	if len(format) > 0 {
		f = format[0]
	}
	a.format = f
}

// String creates a string representation from an Angle reference
func (a *Angle) String() string {
	var degrees int
	var minutes int
	var seconds float64

	switch a.format {
	case Dd:
		return fmt.Sprintf("%.5f°", a.alpha)
	case DMM:
		DMS(a.alpha, &degrees, &minutes, &seconds)
		if a.alpha < 0 && degrees == 0 {
			return fmt.Sprintf("%d°%02d'", degrees, minutes)
		}
		return fmt.Sprintf("%d°%02d'", degrees, int(math.Abs(float64(minutes))))
	case DMMm:
		DMS(a.alpha, &degrees, &minutes, &seconds)
		minutesDecimal := math.Abs(float64(minutes)) + math.Abs(seconds)/60.0
		if a.alpha < 0 && degrees == 0 {
			minutesDecimal = -minutesDecimal
		}
		return fmt.Sprintf("%d°%.3f'", degrees, minutesDecimal)
	case DMMSS:
		DMS(a.alpha, &degrees, &minutes, &seconds)
		if a.alpha < 0 && degrees == 0 {
			return fmt.Sprintf("%d°%02d'%02d\"", degrees, minutes, int(seconds))
		}
		return fmt.Sprintf("%d°%02d'%02d\"", degrees, int(math.Abs(float64(minutes))), int(math.Abs(seconds)))
	case DMMSSs:
		DMS(a.alpha, &degrees, &minutes, &seconds)
		if a.alpha < 0 && degrees == 0 {
			return fmt.Sprintf("%d°%02d'%.3f\"", degrees, minutes, seconds)
		}
		return fmt.Sprintf("%d°%02d'%.3f\"", degrees, int(math.Abs(float64(minutes))), math.Abs(seconds))
	default:
		return fmt.Sprintf("%.5f°", a.alpha)
	}
}

// DegreesToRadians converts degrees to radians
func DegreesToRadians(degrees float64) float64 {
	return degrees * constants.Rad
}

// RadiansToDegrees converts radians to degrees
func RadiansToDegrees(radians float64) float64 {
	return radians * constants.Deg
}

// NormalizeDegrees normalizes degrees to the range [0, 360)
func NormalizeDegrees(degrees float64) float64 {
	return degrees - 360.0*math.Floor(degrees/360.0)
}

// Ddd converts degrees, minutes, seconds to decimal degrees
func Ddd(degrees, minutes int, seconds float64) float64 {
	negative := degrees < 0 || (degrees == 0 && minutes < 0) ||
		(degrees == 0 && minutes == 0 && seconds < 0)
	result := math.Abs(float64(degrees)) + math.Abs(float64(minutes))/60.0 +
		math.Abs(seconds)/3600.0

	if negative {
		result = -result
	}
	return result
}

// DMS converts decimal degrees to degrees, minutes, seconds using pointers
func DMS(decimalDegrees float64, degrees *int, minutes *int, seconds *float64) {
	negative := decimalDegrees < 0
	if negative {
		decimalDegrees = -decimalDegrees
	}

	*degrees = int(decimalDegrees)
	remainder := (decimalDegrees - float64(*degrees)) * 60.0
	*minutes = int(remainder)
	*seconds = (remainder - float64(*minutes)) * 60.0

	if negative {
		if *degrees != 0 {
			*degrees = -*degrees
		} else if *minutes != 0 {
			*minutes = -*minutes
		} else {
			*seconds = -*seconds
		}
	}
}
