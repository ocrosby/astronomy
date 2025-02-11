package angles

import (
	"math"
	"testing"
)

func TestDegreesToRadians(t *testing.T) {
	tests := []struct {
		degrees float64
		want    float64
	}{
		{0, 0},
		{90, math.Pi / 2},
		{180, math.Pi},
		{360, 2 * math.Pi},
		{-90, -math.Pi / 2},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := DegreesToRadians(tt.degrees); got != tt.want {
				t.Errorf("DegreesToRadians(%v) = %v, want %v", tt.degrees, got, tt.want)
			}
		})
	}
}

func TestRadiansToDegrees(t *testing.T) {
	tests := []struct {
		radians float64
		want    float64
	}{
		{0, 0},
		{math.Pi / 2, 90},
		{math.Pi, 180},
		{2 * math.Pi, 360},
		{-math.Pi / 2, -90},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := RadiansToDegrees(tt.radians); got != tt.want {
				t.Errorf("RadiansToDegrees(%v) = %v, want %v", tt.radians, got, tt.want)
			}
		})
	}
}

func TestNormalizeDegrees(t *testing.T) {
	tests := []struct {
		degrees float64
		want    float64
	}{
		{0, 0},
		{360, 0},
		{720, 0},
		{-360, 0},
		{450, 90},
		{-450, 270},
		{1080, 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := NormalizeDegrees(tt.degrees); got != tt.want {
				t.Errorf("NormalizeDegrees(%v) = %v, want %v", tt.degrees, got, tt.want)
			}
		})
	}
}
