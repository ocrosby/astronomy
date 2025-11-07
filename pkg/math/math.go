package math

import "math"

// Frac returns the fractional part of a floating point number
func Frac(x float64) float64 {
	return x - math.Floor(x)
}

// Mod returns the floating-point remainder of x/y
func Mod(x, y float64) float64 {
	return x - y*math.Floor(x/y)
}

// Fabs returns the absolute value of a floating point number
func Fabs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// Sign returns the sign of degrees, minutes, seconds as a floating point number
// Returns -1.0 if any component is negative, otherwise 1.0
func Sign(degrees, minutes int, seconds float64) float64 {
	if degrees < 0 || minutes < 0 || seconds < 0 {
		return -1.0
	}
	return 1.0
}
