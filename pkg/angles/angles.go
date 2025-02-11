package angles

import (
	"github.com/ocrosby/astronomy/pkg/constants"
	"math"
)

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
