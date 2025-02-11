package vectors

import "math"

// Vector2D represents a 2-dimensional vector
type Vector2D struct {
	X, Y float64
}

// Vector3D represents a 3-dimensional vector
type Vector3D struct {
	X, Y, Z float64
}

// ScalarMultiply multiplies a vector by a scalar
func ScalarMultiply(v Vector2D, s float64) Vector2D {
	return Vector2D{v.X * s, v.Y * s}
}

// Add adds two vectors
func Add(v1, v2 Vector2D) Vector2D {
	return Vector2D{v1.X + v2.X, v1.Y + v2.Y}
}

// Subtract subtracts two vectors
func Subtract(v1, v2 Vector2D) Vector2D {
	return Vector2D{v1.X - v2.X, v1.Y - v2.Y}
}

// DotProduct calculates the dot product of two vectors
func DotProduct(v1, v2 Vector2D) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

// CrossProduct calculates the cross product of two vectors
func CrossProduct(v1, v2 Vector2D) float64 {
	return v1.X*v2.Y - v1.Y*v2.X
}

// Magnitude calculates the magnitude of a vector
func Magnitude(v Vector2D) float64 {
	return (v.X*v.X + v.Y*v.Y)
}

// Normalize normalizes a vector
func Normalize(v Vector2D) Vector2D {
	mag := Magnitude(v)
	return Vector2D{v.X / mag, v.Y / mag}
}

// Rotate rotates a vector by an angle in radians
func Rotate(v Vector2D, angle float64) Vector2D {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return Vector2D{v.X*cos - v.Y*sin, v.X*sin + v.Y*cos}
}

// Angle calculates the angle between two vectors in radians
func Angle(v1, v2 Vector2D) float64 {
	return math.Acos(DotProduct(v1, v2) / (Magnitude(v1) * Magnitude(v2)))
}

// Project projects a vector onto another vector
func Project(v1, v2 Vector2D) Vector2D {
	return ScalarMultiply(v2, DotProduct(v1, v2)/Magnitude(v2))
}

// VectorToPolar converts a 2D vector to polar coordinates
func VectorToPolar(v Vector2D) (r, theta float64) {
	return math.Sqrt(v.X*v.X + v.Y*v.Y), math.Atan2(v.Y, v.X)
}

// PolarToVector converts polar coordinates to a 2D vector
func PolarToVector(r, theta float64) Vector2D {
	return Vector2D{r * math.Cos(theta), r * math.Sin(theta)}
}

// VectorToCylindrical converts a 3D vector to cylindrical coordinates
func VectorToCylindrical(v Vector3D) (r, theta, z float64) {
	r = math.Sqrt(v.X*v.X + v.Y*v.Y)
	theta = math.Atan2(v.Y, v.X)
	return r, theta, v.Z
}

// CylindricalToVector converts cylindrical coordinates to a 3D vector
func CylindricalToVector(r, theta, z float64) Vector3D {
	return Vector3D{r * math.Cos(theta), r * math.Sin(theta), z}
}

// VectorToSpherical converts a 3D vector to spherical coordinates
func VectorToSpherical(v Vector3D) (r, theta, phi float64) {
	r = math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	theta = math.Atan2(v.Y, v.X)
	phi = math.Acos(v.Z / r)
	return r, theta, phi
}

// SphericalToVector converts spherical coordinates to a 3D vector
func SphericalToVector(r, theta, phi float64) Vector3D {
	return Vector3D{r * math.Sin(phi) * math.Cos(theta), r * math.Sin(phi) * math.Sin(theta), r * math.Cos(phi)}
}

// Add3D adds two 3D vectors
func Add3D(v1, v2 Vector3D) Vector3D {
	return Vector3D{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

// Subtract3D subtracts two 3D vectors
func Subtract3D(v1, v2 Vector3D) Vector3D {
	return Vector3D{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

// DotProduct3D calculates the dot product of two 3D vectors
func DotProduct3D(v1, v2 Vector3D) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// CrossProduct3D calculates the cross product of two 3D vectors
func CrossProduct3D(v1, v2 Vector3D) Vector3D {
	return Vector3D{v1.Y*v2.Z - v1.Z*v2.Y, v1.Z*v2.X - v1.X*v2.Z, v1.X*v2.Y - v1.Y*v2.X}
}

// Magnitude3D calculates the magnitude of a 3D vector
func Magnitude3D(v Vector3D) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalize3D normalizes a 3D vector
func Normalize3D(v Vector3D) Vector3D {
	mag := Magnitude3D(v)
	return Vector3D{v.X / mag, v.Y / mag, v.Z / mag}
}

// Rotate3D rotates a 3D vector by an angle in radians about the x-axis
func Rotate3Dx(v Vector3D, angle float64) Vector3D {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return Vector3D{v.X, v.Y*cos - v.Z*sin, v.Y*sin + v.Z*cos}
}

// Rotate3Dy rotates a 3D vector by an angle in radians about the y-axis
func Rotate3Dy(v Vector3D, angle float64) Vector3D {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return Vector3D{v.Z*sin + v.X*cos, v.Y, v.Z*cos - v.X*sin}
}

// Rotate3Dz rotates a 3D vector by an angle in radians about the z-axis
func Rotate3Dz(v Vector3D, angle float64) Vector3D {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return Vector3D{v.X*cos - v.Y*sin, v.X*sin + v.Y*cos, v.Z}
}

// Angle3D calculates the angle between two 3D vectors in radians
func Angle3D(v1, v2 Vector3D) float64 {
	return math.Acos(DotProduct3D(v1, v2) / (Magnitude3D(v1) * Magnitude3D(v2)))
}

// ScalarMultiply3D multiplies a 3D vector by a scalar
func ScalarMultiply3D(v Vector3D, s float64) Vector3D {
	return Vector3D{v.X * s, v.Y * s, v.Z * s}
}

// Project3D projects a 3D vector onto another 3D vector
func Project3D(v1, v2 Vector3D) Vector3D {
	return ScalarMultiply3D(v2, DotProduct3D(v1, v2)/Magnitude3D(v2))
}

// VectorToCylindrical3D converts a 3D vector to cylindrical coordinates
func VectorToCylindrical3D(v Vector3D) (r, theta, z float64) {
	r = math.Sqrt(v.X*v.X + v.Y*v.Y)
	theta = math.Atan2(v.Y, v.X)
	return r, theta, v.Z
}

// CylindricalToVector3D converts cylindrical coordinates to a 3D vector
func CylindricalToVector3D(r, theta, z float64) Vector3D {
	return Vector3D{r * math.Cos(theta), r * math.Sin(theta), z}
}

// VectorToSpherical3D converts a 3D vector to spherical coordinates
func VectorToSpherical3D(v Vector3D) (r, theta, phi float64) {
	r = math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	theta = math.Atan2(v.Y, v.X)
	phi = math.Acos(v.Z / r)
	return r, theta, phi
}

// SphericalToVector3D converts spherical coordinates to a 3D vector
func SphericalToVector3D(r, theta, phi float64) Vector3D {
	return Vector3D{r * math.Sin(phi) * math.Cos(theta), r * math.Sin(phi) * math.Sin(theta), r * math.Cos(phi)}
}

// Rotate3D rotates a 3D vector by an angle in radians about an arbitrary axis
func Rotate3D(v Vector3D, axis Vector3D, angle float64) Vector3D {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	cos1 := 1 - cos
	x := axis.X
	y := axis.Y
	z := axis.Z
	return Vector3D{
		(cos+cos1*x*x)*v.X + (cos1*x*y-sin*z)*v.Y + (cos1*x*z+sin*y)*v.Z,
		(cos1*x*y+sin*z)*v.X + (cos+cos1*y*y)*v.Y + (cos1*y*z-sin*x)*v.Z,
		(cos1*x*z-sin*y)*v.X + (cos1*y*z+sin*x)*v.Y + (cos+cos1*z*z)*v.Z,
	}
}

// AngleBetweenPlanes calculates the angle between two planes in radians
func AngleBetweenPlanes(n1, n2 Vector3D) float64 {
	return math.Acos(DotProduct3D(n1, n2) / (Magnitude3D(n1) * Magnitude3D(n2)))
}

// AngleBetweenLines calculates the angle between two lines in radians
func AngleBetweenLines(v1, v2 Vector3D, n1, n2 Vector3D) float64 {
	return math.Acos(math.Abs(DotProduct3D(v1, v2)) / (Magnitude3D(v1) * Magnitude3D(v2)))
}

// LineOfIntersection calculates the line of intersection between two planes
func LineOfIntersection(n1, n2 Vector3D, d1, d2 float64) (Vector3D, Vector3D) {
	line := CrossProduct3D(n1, n2)
	return line, Add3D(ScalarMultiply3D(n1, d2), ScalarMultiply3D(n2, d1))
}

// DistanceBetweenLines calculates the distance between two lines
func DistanceBetweenLines(v1, v2 Vector3D, n1, n2 Vector3D, d1, d2 float64) float64 {
	return math.Abs(DotProduct3D(Subtract3D(v1, v2), CrossProduct3D(n1, n2))) / Magnitude3D(CrossProduct3D(n1, n2))
}

// DistanceToLine calculates the distance from a point to a line
func DistanceToLine(p, v Vector3D, n Vector3D, d float64) float64 {
	return math.Abs(DotProduct3D(Subtract3D(p, v), n)) / Magnitude3D(n)
}

// DistanceToPlane calculates the distance from a point to a plane
func DistanceToPlane(p Vector3D, n Vector3D, d float64) float64 {
	return math.Abs(DotProduct3D(p, n)-d) / Magnitude3D(n)
}

// LineOfIntersectionBetweenPlanes calculates the line of intersection between three planes
func LineOfIntersectionBetweenPlanes(n1, n2, n3 Vector3D, d1, d2, d3 float64) (Vector3D, Vector3D) {
	line := CrossProduct3D(n1, n2)
	return line, Add3D(Add3D(ScalarMultiply3D(n1, d2), ScalarMultiply3D(n2, d1)), ScalarMultiply3D(n3, d3))
}

// PointOfIntersectionBetweenLines calculates the point of intersection between two lines
func PointOfIntersectionBetweenLines(v1, v2 Vector3D, n1, n2 Vector3D, d1, d2 float64) Vector3D {
	line, _ := LineOfIntersection(n1, n2, d1, d2)
	return Add3D(v1, Project3D(Subtract3D(v2, v1), line))
}

// PointOfIntersectionBetweenPlaneAndLine calculates the point of intersection between a plane and a line
func PointOfIntersectionBetweenPlaneAndLine(v, n Vector3D, d float64, p, q Vector3D) Vector3D {
	t := (d - DotProduct3D(v, n)) / DotProduct3D(Subtract3D(p, q), n)
	return Add3D(p, ScalarMultiply3D(Subtract3D(q, p), t))
}

// PointOfIntersectionBetweenPlanes calculates the point of intersection between three planes
func PointOfIntersectionBetweenPlanes(n1, n2, n3 Vector3D, d1, d2, d3 float64) Vector3D {
	line, point := LineOfIntersectionBetweenPlanes(n1, n2, n3, d1, d2, d3)
	return Add3D(point, ScalarMultiply3D(line, d1))
}
