package testutil

import (
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

// Test precision constants for numerical comparisons
const (
	// High precision for mathematical operations
	HighPrecision = 1e-10
	// Standard precision for most calculations
	StandardPrecision = 1e-6
	// Medium precision for approximate calculations
	MediumPrecision = 1e-4
	// Low precision for rough comparisons
	LowPrecision = 1e-2
)

// BeNearStandard provides standard numerical comparison with 1e-6 tolerance
func BeNearStandard(expected float64) types.GomegaMatcher {
	return gomega.BeNumerically("~", expected, StandardPrecision)
}

// BeNearHigh provides high precision numerical comparison with 1e-10 tolerance
func BeNearHigh(expected float64) types.GomegaMatcher {
	return gomega.BeNumerically("~", expected, HighPrecision)
}

// BeNearMedium provides medium precision numerical comparison with 1e-4 tolerance
func BeNearMedium(expected float64) types.GomegaMatcher {
	return gomega.BeNumerically("~", expected, MediumPrecision)
}

// BeLessStandard provides standard "less than" comparison with 1e-6 tolerance
func BeLessStandard() types.GomegaMatcher {
	return gomega.BeNumerically("<", StandardPrecision)
}
