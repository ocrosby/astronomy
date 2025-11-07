package angles

import (
	"fmt"
	"github.com/ocrosby/astronomy/pkg/constants"
	"math"
	"strconv"
	"strings"
)

// SOLID Principle Interfaces

// AngleValue represents a value that can be expressed as an angle
type AngleValue interface {
	// Degrees returns the angle value in decimal degrees
	Degrees() float64
	// Radians returns the angle value in radians
	Radians() float64
}

// FluentAngleFormatter provides a complete fluent interface
type FluentAngleFormatter interface {
	// Format sets the output format
	Format(format AngleFormat) FluentAngleFormatter
	// Precision sets the number of decimal places
	Precision(precision int) FluentAngleFormatter
	// Width sets the minimum field width
	Width(width int) FluentAngleFormatter
	// String returns the formatted representation
	String() string
}

// AngleParser represents the ability to parse angle strings
type AngleParser interface {
	// Parse converts a string to an angle value
	Parse(input string) (AngleValue, error)
}

// AngleConverter provides conversion between angle units
type AngleConverter interface {
	// ToRadians converts degrees to radians
	ToRadians(degrees float64) float64
	// ToDegrees converts radians to degrees
	ToDegrees(radians float64) float64
}

// AngleValidator validates angle components
type AngleValidator interface {
	// ValidateMinutes checks if minutes are in valid range
	ValidateMinutes(minutes float64) error
	// ValidateSeconds checks if seconds are in valid range
	ValidateSeconds(seconds float64) error
}

// DMSCalculator performs DMS calculations
type DMSCalculator interface {
	// ConvertToDMS converts decimal degrees to degrees/minutes/seconds
	ConvertToDMS(decimalDegrees float64) (int, int, float64)
	// ConvertFromDMS converts degrees/minutes/seconds to decimal degrees
	ConvertFromDMS(degrees, minutes int, seconds float64) float64
}

// FormatStrategy defines how to format angles (Strategy Pattern for Open/Closed)
type FormatStrategy interface {
	// Format formats an angle value according to the strategy
	Format(value float64, precision int) string
}

// ExtensibleAngleFormatter allows custom formatting strategies
type ExtensibleAngleFormatter struct {
	value    AngleValue
	strategy FormatStrategy
	display  *DisplayOptions
}

// Angle parsing and formatting constants
const (
	MaxMinutes       = 60
	MaxSeconds       = 60.0
	DefaultPrecision = 2
	DefaultWidth     = 0
	SecondsPerMinute = 60.0
	MinutesPerDegree = 60.0
	SecondsPerDegree = 3600.0
	ValidParseChars  = "0123456789.-+ \tabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
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

// Alpha returns the angle value in decimal degrees
func (a *Angle) Alpha() float64 {
	return a.alpha
}

// Degrees returns the angle value in decimal degrees (implements AngleValue)
func (a *Angle) Degrees() float64 {
	return a.alpha
}

// Radians returns the angle value in radians (implements AngleValue)
func (a *Angle) Radians() float64 {
	return DegreesToRadians(a.alpha)
}

// Format returns the current angle format
func (a *Angle) Format() AngleFormat {
	return a.format
}

// String creates a string representation from an Angle reference
func (a *Angle) String() string {
	return formatAngle(a.alpha, a.format, 3, 0, true)
}

// DisplayOptions holds formatting display options
type DisplayOptions struct {
	Precision int
	Width     int
}

// NewDisplayOptions creates default display options
func NewDisplayOptions() *DisplayOptions {
	return &DisplayOptions{
		Precision: DefaultPrecision,
		Width:     DefaultWidth,
	}
}

// ConcreteAngleFormatter provides a fluent interface for angle formatting
// Separated concerns: value, format, and display options
type ConcreteAngleFormatter struct {
	value   AngleValue
	format  AngleFormat
	display *DisplayOptions
}

// NewFormatter creates a new ConcreteAngleFormatter with the given angle value
func NewFormatter(alpha float64) *ConcreteAngleFormatter {
	return &ConcreteAngleFormatter{
		value:   NewAngle(alpha, Dd),
		format:  Dd,
		display: NewDisplayOptions(),
	}
}

// Format sets the angle format and returns the formatter for chaining
func (f *ConcreteAngleFormatter) Format(format AngleFormat) FluentAngleFormatter {
	f.format = format
	return f
}

// Precision sets the decimal precision and returns the formatter for chaining
func (f *ConcreteAngleFormatter) Precision(precision int) FluentAngleFormatter {
	f.display.Precision = precision
	return f
}

// Width sets the field width and returns the formatter for chaining
func (f *ConcreteAngleFormatter) Width(width int) FluentAngleFormatter {
	f.display.Width = width
	return f
}

// String formats the angle according to the configured settings
func (f *ConcreteAngleFormatter) String() string {
	return formatAngle(f.value.Degrees(), f.format, f.display.Precision, f.display.Width, false)
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
	result := math.Abs(float64(degrees)) + math.Abs(float64(minutes))/MinutesPerDegree +
		math.Abs(seconds)/SecondsPerDegree

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
	remainder := (decimalDegrees - float64(*degrees)) * MinutesPerDegree
	*minutes = int(remainder)
	*seconds = (remainder - float64(*minutes)) * SecondsPerMinute

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

// ParseAngle parses a string in fluent output format and returns an Angle
func ParseAngle(input string) (*Angle, error) {
	// Validate input
	if input == "" {
		return nil, fmt.Errorf("empty input string")
	}

	originalInput := input
	input = strings.TrimSpace(input)

	if input == "" {
		return nil, fmt.Errorf("input contains only whitespace")
	}

	// Check for invalid characters that would indicate a malformed angle
	// Allow letters for special values like "inf", "nan", etc.
	for _, char := range input {
		if !strings.ContainsRune(ValidParseChars, char) {
			return nil, fmt.Errorf("invalid character '%c' in input '%s'", char, originalInput)
		}
	}

	// Count spaces to determine format
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil, fmt.Errorf("no valid components found in input '%s'", originalInput)
	}

	// Validate each part is not empty
	for i, part := range parts {
		if part == "" {
			return nil, fmt.Errorf("component %d is empty in input '%s'", i+1, originalInput)
		}
	}

	switch len(parts) {
	case 1:
		// Dd format - single decimal number
		return parseDdFormat(parts[0], originalInput)

	case 2:
		// DMM or DMMm format - degrees and minutes
		return parseDMMFormat(parts[0], parts[1], originalInput)

	case 3:
		// DMMSS or DMMSSs format - degrees, minutes, and seconds
		return parseDMMSSFormat(parts[0], parts[1], parts[2], originalInput)

	default:
		return nil, fmt.Errorf("invalid format: expected 1-3 space-separated components, got %d in input '%s'", len(parts), originalInput)
	}
}

// parseDdFormat handles parsing of decimal degrees format
func parseDdFormat(degreeStr, originalInput string) (*Angle, error) {
	value, err := parseFloatComponent(degreeStr, "decimal degrees", originalInput)
	if err != nil {
		return nil, err
	}

	return NewAngle(value, Dd), nil
}

// parseDMMFormat handles parsing of degrees and minutes formats
func parseDMMFormat(degreeStr, minuteStr, originalInput string) (*Angle, error) {
	// Parse degrees
	degrees, err := parseIntegerComponent(degreeStr, "degrees", originalInput)
	if err != nil {
		return nil, err
	}

	// Check if minutes contains decimal point
	if strings.Contains(minuteStr, ".") {
		// DMMm format
		minutes, err := parseFloatComponent(minuteStr, "minutes", originalInput)
		if err != nil {
			return nil, err
		}

		// Validate minutes range
		if err := validateMinutesFloat(minutes, originalInput); err != nil {
			return nil, err
		}

		decimalDegrees := Ddd(degrees, int(minutes), (minutes-float64(int(minutes)))*SecondsPerMinute)
		return NewAngle(decimalDegrees, DMMm), nil
	} else {
		// DMM format
		minutes, err := parseIntegerComponent(minuteStr, "minutes", originalInput)
		if err != nil {
			return nil, err
		}

		// Validate minutes range
		if err := validateMinutesInt(minutes, originalInput); err != nil {
			return nil, err
		}

		decimalDegrees := Ddd(degrees, minutes, 0.0)
		return NewAngle(decimalDegrees, DMM), nil
	}
}

// parseDMMSSFormat handles parsing of degrees, minutes, and seconds formats
func parseDMMSSFormat(degreeStr, minuteStr, secondStr, originalInput string) (*Angle, error) {
	// Parse degrees
	degrees, err := parseIntegerComponent(degreeStr, "degrees", originalInput)
	if err != nil {
		return nil, err
	}

	// Parse minutes
	minutes, err := parseIntegerComponent(minuteStr, "minutes", originalInput)
	if err != nil {
		return nil, err
	}

	// Validate minutes range
	if err := validateMinutesInt(minutes, originalInput); err != nil {
		return nil, err
	}

	// Handle special case where degrees is "-0" indicating small negative angle
	isNegativeZero := strings.HasPrefix(degreeStr, "-") && degrees == 0

	// Check if seconds contains decimal point
	if strings.Contains(secondStr, ".") {
		// DMMSSs format
		seconds, err := parseFloatComponent(secondStr, "seconds", originalInput)
		if err != nil {
			return nil, err
		}

		// Validate seconds range
		if err := validateSecondsFloat(seconds, originalInput); err != nil {
			return nil, err
		}

		decimalDegrees := Ddd(degrees, minutes, seconds)
		if isNegativeZero {
			decimalDegrees = -decimalDegrees
		}
		return NewAngle(decimalDegrees, DMMSSs), nil
	} else {
		// DMMSS format
		seconds, err := parseIntegerComponent(secondStr, "seconds", originalInput)
		if err != nil {
			return nil, err
		}

		// Validate seconds range
		if err := validateSecondsInt(seconds, originalInput); err != nil {
			return nil, err
		}

		decimalDegrees := Ddd(degrees, minutes, float64(seconds))
		if isNegativeZero {
			decimalDegrees = -decimalDegrees
		}
		return NewAngle(decimalDegrees, DMMSS), nil
	}
}

// Common validation patterns for parsing

// validateNumericString performs common string validation for numeric components
func validateNumericString(str, componentName, originalInput string) error {
	// Check for multiple signs
	signCount := strings.Count(str, "+") + strings.Count(str, "-")
	if signCount > 1 {
		return fmt.Errorf("invalid %s: multiple signs in '%s'", componentName, originalInput)
	}

	// Check for sign not at beginning
	if len(str) > 1 && (strings.Contains(str[1:], "+") || strings.Contains(str[1:], "-")) {
		return fmt.Errorf("invalid %s: sign must be at beginning in '%s'", componentName, originalInput)
	}

	return nil
}

// parseIntegerComponent parses an integer component with validation
func parseIntegerComponent(str, componentName, originalInput string) (int, error) {
	if err := validateNumericString(str, componentName, originalInput); err != nil {
		return 0, err
	}

	// Check for decimal point in integer component
	if strings.Contains(str, ".") {
		return 0, fmt.Errorf("invalid %s: unexpected decimal point in integer value '%s'", componentName, originalInput)
	}

	value, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("invalid %s value in '%s': %v", componentName, originalInput, err)
	}

	return value, nil
}

// parseFloatComponent parses a float component with validation
func parseFloatComponent(str, componentName, originalInput string) (float64, error) {
	if err := validateNumericString(str, componentName, originalInput); err != nil {
		return 0, err
	}

	// Check for multiple decimal points
	if strings.Count(str, ".") > 1 {
		return 0, fmt.Errorf("invalid %s: multiple decimal points in '%s'", componentName, originalInput)
	}

	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s value in '%s': %v", componentName, originalInput, err)
	}

	// Check for reasonable values
	if math.IsInf(value, 0) || math.IsNaN(value) {
		return 0, fmt.Errorf("invalid %s: value is infinite or NaN in '%s'", componentName, originalInput)
	}

	return value, nil
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// ConcreteAngleParser implements AngleParser interface
type ConcreteAngleParser struct{}

// NewAngleParser creates a new parser instance
func NewAngleParser() AngleParser {
	return &ConcreteAngleParser{}
}

// Parse implements the AngleParser interface
func (p *ConcreteAngleParser) Parse(input string) (AngleValue, error) {
	return ParseAngle(input)
}

// StandardAngleConverter implements AngleConverter interface
type StandardAngleConverter struct{}

// NewAngleConverter creates a new angle converter
func NewAngleConverter() AngleConverter {
	return &StandardAngleConverter{}
}

// ToRadians converts degrees to radians
func (c *StandardAngleConverter) ToRadians(degrees float64) float64 {
	return DegreesToRadians(degrees)
}

// ToDegrees converts radians to degrees
func (c *StandardAngleConverter) ToDegrees(radians float64) float64 {
	return RadiansToDegrees(radians)
}

// StandardAngleValidator implements AngleValidator interface
type StandardAngleValidator struct{}

// NewAngleValidator creates a new angle validator
func NewAngleValidator() AngleValidator {
	return &StandardAngleValidator{}
}

// ValidateMinutes checks if minutes are in valid range
func (v *StandardAngleValidator) ValidateMinutes(minutes float64) error {
	return validateMinutesFloat(minutes, "validation")
}

// ValidateSeconds checks if seconds are in valid range
func (v *StandardAngleValidator) ValidateSeconds(seconds float64) error {
	return validateSecondsFloat(seconds, "validation")
}

// StandardDMSCalculator implements DMSCalculator interface
type StandardDMSCalculator struct{}

// NewDMSCalculator creates a new DMS calculator
func NewDMSCalculator() DMSCalculator {
	return &StandardDMSCalculator{}
}

// ConvertToDMS converts decimal degrees to degrees/minutes/seconds
func (d *StandardDMSCalculator) ConvertToDMS(decimalDegrees float64) (int, int, float64) {
	var degrees int
	var minutes int
	var seconds float64
	DMS(decimalDegrees, &degrees, &minutes, &seconds)
	return degrees, minutes, seconds
}

// ConvertFromDMS converts degrees/minutes/seconds to decimal degrees
func (d *StandardDMSCalculator) ConvertFromDMS(degrees, minutes int, seconds float64) float64 {
	return Ddd(degrees, minutes, seconds)
}

// Concrete Format Strategies (Open/Closed Principle - extensible without modification)

// DecimalFormatStrategy formats as decimal degrees
type DecimalFormatStrategy struct{}

// Format implements FormatStrategy for decimal degrees
func (s *DecimalFormatStrategy) Format(value float64, precision int) string {
	return fmt.Sprintf("%.*f", precision, value)
}

// DMMSSFormatStrategy formats as degrees/minutes/seconds
type DMMSSFormatStrategy struct {
	calculator DMSCalculator
}

// NewDMMSSFormatStrategy creates a new DMMSS format strategy
func NewDMMSSFormatStrategy(calc DMSCalculator) FormatStrategy {
	return &DMMSSFormatStrategy{calculator: calc}
}

// Format implements FormatStrategy for DMMSS format
func (s *DMMSSFormatStrategy) Format(value float64, precision int) string {
	degrees, minutes, seconds := s.calculator.ConvertToDMS(value)
	components := getDMSComponents(value)

	if components.isNegativeZero {
		return fmt.Sprintf("%d %d %.*f", degrees, minutes, precision, seconds)
	}
	return fmt.Sprintf("%d %d %.*f", degrees, int(math.Abs(float64(minutes))), precision, math.Abs(seconds))
}

// NewExtensibleFormatter creates a formatter with custom strategy
func NewExtensibleFormatter(value AngleValue, strategy FormatStrategy) *ExtensibleAngleFormatter {
	return &ExtensibleAngleFormatter{
		value:    value,
		strategy: strategy,
		display:  NewDisplayOptions(),
	}
}

// String formats using the configured strategy
func (f *ExtensibleAngleFormatter) String() string {
	result := f.strategy.Format(f.value.Degrees(), f.display.Precision)

	// Apply width formatting
	if f.display.Width > 0 {
		result = fmt.Sprintf("%-*s", f.display.Width, result)
	}

	return result
}

// WithStrategy allows changing the format strategy (Open/Closed)
func (f *ExtensibleAngleFormatter) WithStrategy(strategy FormatStrategy) *ExtensibleAngleFormatter {
	f.strategy = strategy
	return f
}

// WithPrecision sets precision for extensible formatter
func (f *ExtensibleAngleFormatter) WithPrecision(precision int) *ExtensibleAngleFormatter {
	f.display.Precision = precision
	return f
}

// WithWidth sets width for extensible formatter
func (f *ExtensibleAngleFormatter) WithWidth(width int) *ExtensibleAngleFormatter {
	f.display.Width = width
	return f
}

// Validation functions to eliminate duplicated validation logic

// validateMinutesInt validates integer minutes are within valid range
func validateMinutesInt(minutes int, originalInput string) error {
	if abs(minutes) >= MaxMinutes {
		return fmt.Errorf("invalid minutes value: must be less than %d, got %d in '%s'", MaxMinutes, minutes, originalInput)
	}
	return nil
}

// validateMinutesFloat validates float minutes are within valid range
func validateMinutesFloat(minutes float64, originalInput string) error {
	if math.Abs(minutes) >= MaxMinutes {
		return fmt.Errorf("invalid minutes value: must be less than %d, got %.2f in '%s'", MaxMinutes, minutes, originalInput)
	}
	return nil
}

// validateSecondsInt validates integer seconds are within valid range
func validateSecondsInt(seconds int, originalInput string) error {
	if abs(seconds) >= int(MaxSeconds) {
		return fmt.Errorf("invalid seconds value: must be less than %d, got %d in '%s'", int(MaxSeconds), seconds, originalInput)
	}
	return nil
}

// validateSecondsFloat validates float seconds are within valid range
func validateSecondsFloat(seconds float64, originalInput string) error {
	if math.Abs(seconds) >= MaxSeconds {
		return fmt.Errorf("invalid seconds value: must be less than %.0f, got %.2f in '%s'", MaxSeconds, seconds, originalInput)
	}
	return nil
}

// DMSComponents holds the components of a DMS angle with formatting context
type DMSComponents struct {
	degrees        int
	minutes        int
	seconds        float64
	isNegativeZero bool
}

// getDMSComponents extracts DMS components and handles negative zero case
func getDMSComponents(alpha float64) DMSComponents {
	var degrees int
	var minutes int
	var seconds float64

	DMS(alpha, &degrees, &minutes, &seconds)

	return DMSComponents{
		degrees:        degrees,
		minutes:        minutes,
		seconds:        seconds,
		isNegativeZero: alpha < 0 && degrees == 0,
	}
}

// formatAngle provides unified formatting logic for both Angle and AngleFormatter
func formatAngle(alpha float64, format AngleFormat, precision int, width int, useSymbols bool) string {
	components := getDMSComponents(alpha)
	var result string

	switch format {
	case Dd:
		if useSymbols {
			result = fmt.Sprintf("%.5f°", alpha)
		} else {
			result = fmt.Sprintf("%.*f", precision, alpha)
		}
	case DMM:
		if useSymbols {
			if components.isNegativeZero {
				result = fmt.Sprintf("%d°%02d'", components.degrees, components.minutes)
			} else {
				result = fmt.Sprintf("%d°%02d'", components.degrees, int(math.Abs(float64(components.minutes))))
			}
		} else {
			if components.isNegativeZero {
				result = fmt.Sprintf("%d %d", components.degrees, components.minutes)
			} else {
				result = fmt.Sprintf("%d %d", components.degrees, int(math.Abs(float64(components.minutes))))
			}
		}
	case DMMm:
		minutesDecimal := math.Abs(float64(components.minutes)) + math.Abs(components.seconds)/SecondsPerMinute
		if components.isNegativeZero {
			minutesDecimal = -minutesDecimal
		}
		if useSymbols {
			result = fmt.Sprintf("%d°%.3f'", components.degrees, minutesDecimal)
		} else {
			result = fmt.Sprintf("%d %.*f", components.degrees, precision, minutesDecimal)
		}
	case DMMSS:
		if useSymbols {
			if components.isNegativeZero {
				result = fmt.Sprintf("%d°%02d'%02d\"", components.degrees, components.minutes, int(components.seconds))
			} else {
				result = fmt.Sprintf("%d°%02d'%02d\"", components.degrees, int(math.Abs(float64(components.minutes))), int(math.Abs(components.seconds)))
			}
		} else {
			if components.isNegativeZero {
				result = fmt.Sprintf("%d %d %d", components.degrees, components.minutes, int(components.seconds))
			} else {
				result = fmt.Sprintf("%d %d %d", components.degrees, int(math.Abs(float64(components.minutes))), int(math.Abs(components.seconds)))
			}
		}
	case DMMSSs:
		if useSymbols {
			if components.isNegativeZero {
				result = fmt.Sprintf("%d°%02d'%.3f\"", components.degrees, components.minutes, components.seconds)
			} else {
				result = fmt.Sprintf("%d°%02d'%.3f\"", components.degrees, int(math.Abs(float64(components.minutes))), math.Abs(components.seconds))
			}
		} else {
			if components.isNegativeZero {
				result = fmt.Sprintf("%d %d %.*f", components.degrees, components.minutes, precision, components.seconds)
			} else {
				result = fmt.Sprintf("%d %d %.*f", components.degrees, int(math.Abs(float64(components.minutes))), precision, math.Abs(components.seconds))
			}
		}
	default:
		if useSymbols {
			result = fmt.Sprintf("%.5f°", alpha)
		} else {
			result = fmt.Sprintf("%.*f", precision, alpha)
		}
	}

	// Apply width formatting with left justification
	if width > 0 {
		result = fmt.Sprintf("%-*s", width, result)
	}

	return result
}
