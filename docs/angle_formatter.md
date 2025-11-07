# AngleFormatter Fluent Interface

The `AngleFormatter` provides a fluent interface for converting and formatting angles into various sexagesimal representations with customizable precision and width settings.

## Table of Contents

- [Overview](#overview)
- [Quick Start](#quick-start)
- [Angle Formats](#angle-formats)
- [Methods](#methods)
- [Examples](#examples)
- [Advanced Usage](#advanced-usage)
- [Technical Details](#technical-details)

## Overview

The AngleFormatter allows you to convert decimal degrees into multiple formatted representations:

- **Decimal degrees** (Dd): Standard floating-point representation
- **Degrees and whole minutes** (DMM): Integer degrees and minutes
- **Degrees and decimal minutes** (DMMm): Integer degrees with fractional minutes
- **Degrees, minutes, and whole seconds** (DMMSS): Integer degrees, minutes, and seconds
- **Degrees, minutes, and decimal seconds** (DMMSSs): Integer degrees and minutes with fractional seconds

## Quick Start

```go
import "github.com/ocrosby/astronomy/pkg/angles"

// Convert 12.3456 degrees to "12.35"
result := angles.NewFormatter(12.3456).Format(angles.Dd).Precision(2).String()

// Convert to "12 20 44.16"
result := angles.NewFormatter(12.3456).Format(angles.DMMSSs).Precision(2).String()
```

## Angle Formats

### Dd - Decimal Degrees
Represents angles as floating-point decimal values.

```go
angles.NewFormatter(12.3456).Format(angles.Dd).Precision(2).String()
// Output: "12.35"
```

### DMM - Degrees and Whole Minutes
Shows integer degrees and integer minutes of arc.

```go
angles.NewFormatter(12.3456).Format(angles.DMM).String()
// Output: "12 20"
```

### DMMm - Degrees and Decimal Minutes
Shows integer degrees with fractional minutes.

```go
angles.NewFormatter(12.3456).Format(angles.DMMm).Precision(2).String()
// Output: "12 20.74"
```

### DMMSS - Degrees, Minutes, and Whole Seconds
Shows integer degrees, integer minutes, and integer seconds of arc.

```go
angles.NewFormatter(12.3456).Format(angles.DMMSS).String()
// Output: "12 20 44"
```

### DMMSSs - Degrees, Minutes, and Decimal Seconds
Shows integer degrees, integer minutes, with fractional seconds.

```go
angles.NewFormatter(12.3456).Format(angles.DMMSSs).Precision(2).String()
// Output: "12 20 44.16"
```

## Methods

### NewFormatter(alpha float64) *AngleFormatter
Creates a new formatter instance with the specified angle value.

**Parameters:**
- `alpha`: The angle value in decimal degrees

**Returns:**
- `*AngleFormatter`: A new formatter instance

**Default Settings:**
- Format: `Dd` (decimal degrees)
- Precision: `2` decimal places
- Width: `0` (no width constraint)

### Format(format AngleFormat) *AngleFormatter
Sets the output format for the angle.

**Parameters:**
- `format`: One of `Dd`, `DMM`, `DMMm`, `DMMSS`, or `DMMSSs`

**Returns:**
- `*AngleFormatter`: The formatter instance for method chaining

### Precision(precision int) *AngleFormatter
Sets the number of decimal places for fractional components.

**Parameters:**
- `precision`: Number of decimal places (0 or greater)

**Returns:**
- `*AngleFormatter`: The formatter instance for method chaining

**Applies to:**
- `Dd`: Decimal places for the angle value
- `DMMm`: Decimal places for the minutes component
- `DMMSSs`: Decimal places for the seconds component
- `DMM` and `DMMSS`: No effect (integer values only)

### Width(width int) *AngleFormatter
Sets the minimum field width with left justification.

**Parameters:**
- `width`: Minimum field width (0 for no width constraint)

**Returns:**
- `*AngleFormatter`: The formatter instance for method chaining

**Behavior:**
- If the formatted string is shorter than the specified width, it's padded with spaces on the right
- If the formatted string is longer than the width, it's not truncated
- Negative signs are properly left-justified

### String() string
Formats the angle according to the current settings and returns the result.

**Returns:**
- `string`: The formatted angle representation

## Examples

### Basic Formatting

```go
angle := 12.3456

// Different formats for the same angle
fmt.Println(angles.NewFormatter(angle).Format(angles.Dd).Precision(2).String())
// Output: "12.35"

fmt.Println(angles.NewFormatter(angle).Format(angles.DMM).String())
// Output: "12 20"

fmt.Println(angles.NewFormatter(angle).Format(angles.DMMm).Precision(2).String())
// Output: "12 20.74"

fmt.Println(angles.NewFormatter(angle).Format(angles.DMMSS).String())
// Output: "12 20 44"

fmt.Println(angles.NewFormatter(angle).Format(angles.DMMSSs).Precision(2).String())
// Output: "12 20 44.16"
```

### Precision Control

```go
angle := 12.3456789

// Different precision levels
fmt.Println(angles.NewFormatter(angle).Format(angles.Dd).Precision(1).String())
// Output: "12.3"

fmt.Println(angles.NewFormatter(angle).Format(angles.Dd).Precision(3).String())
// Output: "12.346"

fmt.Println(angles.NewFormatter(angle).Format(angles.DMMSSs).Precision(4).String())
// Output: "12 20 44.1604"
```

### Width Formatting

```go
angle := 12.35

// Left-justified formatting with width
fmt.Printf("'%s'\n", angles.NewFormatter(angle).Format(angles.Dd).Precision(2).Width(10).String())
// Output: "'12.35     '"

fmt.Printf("'%s'\n", angles.NewFormatter(angle).Format(angles.DMM).Width(8).String())
// Output: "'12 20   '"
```

### Negative Angles

```go
angle := -12.3456

fmt.Println(angles.NewFormatter(angle).Format(angles.Dd).Precision(2).String())
// Output: "-12.35"

fmt.Println(angles.NewFormatter(angle).Format(angles.DMM).String())
// Output: "-12 20"

// Small negative angles
smallAngle := -0.3456
fmt.Println(angles.NewFormatter(smallAngle).Format(angles.DMM).String())
// Output: "0 -20"

fmt.Println(angles.NewFormatter(smallAngle).Format(angles.DMMSSs).Precision(2).String())
// Output: "0 -20 44.16"
```

### Method Chaining

```go
// Methods can be chained in any order
result1 := angles.NewFormatter(12.3456).Precision(3).Format(angles.DMMSSs).Width(15).String()
result2 := angles.NewFormatter(12.3456).Width(15).Format(angles.DMMSSs).Precision(3).String()
result3 := angles.NewFormatter(12.3456).Format(angles.DMMSSs).Width(15).Precision(3).String()

// All produce the same result: "12 20 44.160   "
```

## Advanced Usage

### Coordinate Formatting

```go
// Format latitude and longitude pairs
lat := 40.7589
lon := -73.9851

fmt.Printf("Coordinates: %s N, %s W\n",
    angles.NewFormatter(lat).Format(angles.DMMSSs).Precision(2).String(),
    angles.NewFormatter(-lon).Format(angles.DMMSSs).Precision(2).String())
// Output: "Coordinates: 40 45 32.04 N, 73 59 6.36 W"
```

### Astronomical Applications

```go
// Right ascension in hours, minutes, seconds
ra := 83.6333  // degrees
raHours := ra / 15.0  // convert to hours

fmt.Printf("RA: %sh %sm %ss\n",
    fmt.Sprintf("%.0f", raHours),
    fmt.Sprintf("%.0f", (raHours - float64(int(raHours))) * 60),
    angles.NewFormatter((raHours - float64(int(raHours))) * 3600).Format(angles.Dd).Precision(1).String())
```

### Table Formatting

```go
angles_list := []float64{12.3456, -45.6789, 123.4567, -0.1234}

fmt.Println("Angle       | Dd      | DMM     | DMMSSs")
fmt.Println("------------|---------|---------|----------")

for _, angle := range angles_list {
    dd := angles.NewFormatter(angle).Format(angles.Dd).Precision(2).Width(7).String()
    dmm := angles.NewFormatter(angle).Format(angles.DMM).Width(7).String()
    dmmss := angles.NewFormatter(angle).Format(angles.DMMSSs).Precision(2).Width(9).String()
    
    fmt.Printf("%-11.4f | %s | %s | %s\n", angle, dd, dmm, dmmss)
}
```

## Technical Details

### Precision Behavior

- **Integer formats** (`DMM`, `DMMSS`): Precision setting is ignored
- **Decimal formats** (`Dd`, `DMMm`, `DMMSSs`): Precision applies to the fractional component
- **Default precision**: 2 decimal places
- **Minimum precision**: 0 (no decimal places)

### Width Behavior

- **Left justification**: Padding added to the right
- **Minimum width**: If content is longer than specified width, no truncation occurs
- **Default width**: 0 (no width constraint)
- **Negative signs**: Properly positioned at the leftmost position

### Negative Angle Handling

The formatter handles negative angles according to standard conventions:

1. **Large negative angles**: The degrees component carries the negative sign
   - `-12.3456°` → `"-12 20 44.16"`

2. **Small negative angles**: When degrees is zero, the first non-zero component carries the negative sign
   - `-0.3456°` → `"0 -20 44.16"`

### Error Handling

The formatter is designed to be robust:

- **Invalid precision**: Negative precision values are treated as 0
- **Invalid width**: Negative width values are treated as 0
- **Infinite/NaN values**: Formatted as Go's standard string representation

### Performance Considerations

- **Method chaining**: Each method returns the same instance, so chaining is efficient
- **String formatting**: Uses Go's standard `fmt` package for consistent behavior
- **Memory allocation**: Minimal allocations during formatting operations

## Integration with Other Types

The AngleFormatter works seamlessly with other angle types in the package:

```go
// Convert from Angle struct
angle := angles.NewAngle(12.3456, angles.DMMSS)
formatted := angles.NewFormatter(12.3456).Format(angles.DMMSSs).Precision(2).String()

// Use with DMS function
var degrees int
var minutes int
var seconds float64
angles.DMS(12.3456, &degrees, &minutes, &seconds)
reconstructed := angles.Ddd(degrees, minutes, seconds)
formatted := angles.NewFormatter(reconstructed).Format(angles.Dd).Precision(4).String()
```