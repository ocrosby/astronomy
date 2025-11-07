package vectors

import (
	"fmt"
	"math"
)

// Vector interfaces for Liskov Substitution Principle

// Vector represents a mathematical vector
type Vector interface {
	// Magnitude returns the magnitude of the vector
	Magnitude() float64
	// String returns string representation
	String() string
}

// Vector2DOperations defines 2D vector operations
type Vector2DOperations interface {
	Add(v Vector2D) Vector2D
	Subtract(v Vector2D) Vector2D
	ScalarMultiply(s float64) Vector2D
	DotProduct(v Vector2D) float64
	Normalize() Vector2D
}

// Vector3DOperations defines 3D vector operations
type Vector3DOperations interface {
	Add(v Vector3D) Vector3D
	Subtract(v Vector3D) Vector3D
	ScalarMultiply(s float64) Vector3D
	DotProduct(v Vector3D) float64
	CrossProduct(v Vector3D) Vector3D
	Normalize() Vector3D
}

// Vector2D represents a 2-dimensional vector
type Vector2D struct {
	X, Y float64
}

// String returns string representation of Vector2D
func (v Vector2D) String() string {
	return fmt.Sprintf("(%.3f, %.3f)", v.X, v.Y)
}

// Magnitude calculates the magnitude of a 2D vector (implements Vector interface)
func (v Vector2D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Add adds two 2D vectors (implements Vector2DOperations)
func (v Vector2D) Add(other Vector2D) Vector2D {
	return Vector2D{v.X + other.X, v.Y + other.Y}
}

// Subtract subtracts two 2D vectors (implements Vector2DOperations)
func (v Vector2D) Subtract(other Vector2D) Vector2D {
	return Vector2D{v.X - other.X, v.Y - other.Y}
}

// ScalarMultiply multiplies vector by scalar (implements Vector2DOperations)
func (v Vector2D) ScalarMultiply(s float64) Vector2D {
	return Vector2D{v.X * s, v.Y * s}
}

// DotProduct calculates dot product (implements Vector2DOperations)
func (v Vector2D) DotProduct(other Vector2D) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Normalize normalizes the vector (implements Vector2DOperations)
func (v Vector2D) Normalize() Vector2D {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector2D{0, 0}
	}
	return Vector2D{v.X / mag, v.Y / mag}
}

// Vector3D represents a 3-dimensional vector
type Vector3D struct {
	X, Y, Z float64
}

// String returns string representation of Vector3D
func (v Vector3D) String() string {
	return fmt.Sprintf("(%.3f, %.3f, %.3f)", v.X, v.Y, v.Z)
}

// Magnitude calculates the magnitude of a 3D vector (implements Vector interface)
func (v Vector3D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Add adds two 3D vectors (implements Vector3DOperations)
func (v Vector3D) Add(other Vector3D) Vector3D {
	return Vector3D{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

// Subtract subtracts two 3D vectors (implements Vector3DOperations)
func (v Vector3D) Subtract(other Vector3D) Vector3D {
	return Vector3D{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

// ScalarMultiply multiplies vector by scalar (implements Vector3DOperations)
func (v Vector3D) ScalarMultiply(s float64) Vector3D {
	return Vector3D{v.X * s, v.Y * s, v.Z * s}
}

// DotProduct calculates dot product (implements Vector3DOperations)
func (v Vector3D) DotProduct(other Vector3D) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// CrossProduct calculates cross product (implements Vector3DOperations)
func (v Vector3D) CrossProduct(other Vector3D) Vector3D {
	return Vector3D{
		v.Y*other.Z - v.Z*other.Y,
		v.Z*other.X - v.X*other.Z,
		v.X*other.Y - v.Y*other.X,
	}
}

// Normalize normalizes the vector (implements Vector3DOperations)
func (v Vector3D) Normalize() Vector3D {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector3D{0, 0, 0}
	}
	return Vector3D{v.X / mag, v.Y / mag, v.Z / mag}
}

// CrossProduct calculates the cross product of two vectors
func CrossProduct(v1, v2 Vector2D) float64 {
	return v1.X*v2.Y - v1.Y*v2.X
}

// Magnitude calculates the magnitude of a vector
func Magnitude(v Vector2D) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize normalizes a vector
func Normalize(v Vector2D) Vector2D {
	mag := Magnitude(v)
	return Vector2D{v.X / mag, v.Y / mag}
}

// DotProduct calculates dot product of two 2D vectors
func DotProduct(v1, v2 Vector2D) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

// ScalarMultiply multiplies a 2D vector by a scalar
func ScalarMultiply(v Vector2D, s float64) Vector2D {
	return Vector2D{v.X * s, v.Y * s}
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
	rSquared := v.X*v.X + v.Y*v.Y
	return math.Sqrt(rSquared), math.Atan2(v.Y, v.X)
}

// PolarToVector converts polar coordinates to a 2D vector
func PolarToVector(r, theta float64) Vector2D {
	return Vector2D{r * math.Cos(theta), r * math.Sin(theta)}
}

// VectorToCylindrical converts a 3D vector to cylindrical coordinates
func VectorToCylindrical(v Vector3D) (r, theta, z float64) {
	rSquared := v.X*v.X + v.Y*v.Y
	r = math.Sqrt(rSquared)
	theta = math.Atan2(v.Y, v.X)
	return r, theta, v.Z
}

// CylindricalToVector converts cylindrical coordinates to a 3D vector
func CylindricalToVector(r, theta, z float64) Vector3D {
	return Vector3D{r * math.Cos(theta), r * math.Sin(theta), z}
}

// VectorToSpherical converts a 3D vector to spherical coordinates
func VectorToSpherical(v Vector3D) (r, theta, phi float64) {
	rSquared := v.X*v.X + v.Y*v.Y + v.Z*v.Z
	if rSquared == 0 {
		return 0, 0, 0
	}
	r = math.Sqrt(rSquared)
	theta = math.Atan2(v.Y, v.X)
	phi = math.Acos(v.Z / r)
	return r, theta, phi
}

// SphericalToVector converts spherical coordinates to a 3D vector
func SphericalToVector(r, theta, phi float64) Vector3D {
	sinPhi := math.Sin(phi)
	cosPhi := math.Cos(phi)
	cosTheta := math.Cos(theta)
	sinTheta := math.Sin(theta)
	return Vector3D{r * sinPhi * cosTheta, r * sinPhi * sinTheta, r * cosPhi}
}

// fastInvSqrt implements fast inverse square root approximation (Quake algorithm)
// Use with caution - trades precision for speed
func fastInvSqrt(x float64) float64 {
	if x <= 0 {
		return 0
	}
	const threehalfs = 1.5
	x2 := x * 0.5
	i := math.Float64bits(x)
	i = 0x5fe6eb50c7b537a9 - (i >> 1)
	y := math.Float64frombits(i)
	y = y * (threehalfs - (x2 * y * y))
	return y
}

// VectorToPolarFast converts a 2D vector to polar coordinates using fast inverse sqrt
func VectorToPolarFast(v Vector2D) (r, theta float64) {
	rSquared := v.X*v.X + v.Y*v.Y
	if rSquared == 0 {
		return 0, 0
	}
	r = rSquared * fastInvSqrt(rSquared)
	theta = math.Atan2(v.Y, v.X)
	return r, theta
}

// VectorToSphericalFast converts a 3D vector to spherical coordinates using fast inverse sqrt
func VectorToSphericalFast(v Vector3D) (r, theta, phi float64) {
	rSquared := v.X*v.X + v.Y*v.Y + v.Z*v.Z
	if rSquared == 0 {
		return 0, 0, 0
	}
	r = rSquared * fastInvSqrt(rSquared)
	theta = math.Atan2(v.Y, v.X)
	phi = math.Acos(v.Z / r)
	return r, theta, phi
}

// BulkVectorToPolar converts multiple 2D vectors to polar coordinates
func BulkVectorToPolar(vectors []Vector2D) ([]float64, []float64) {
	n := len(vectors)
	radii := make([]float64, n)
	angles := make([]float64, n)

	for i, v := range vectors {
		rSquared := v.X*v.X + v.Y*v.Y
		radii[i] = math.Sqrt(rSquared)
		angles[i] = math.Atan2(v.Y, v.X)
	}

	return radii, angles
}

// BulkVectorToSpherical converts multiple 3D vectors to spherical coordinates
func BulkVectorToSpherical(vectors []Vector3D) ([]float64, []float64, []float64) {
	n := len(vectors)
	radii := make([]float64, n)
	thetas := make([]float64, n)
	phis := make([]float64, n)

	for i, v := range vectors {
		rSquared := v.X*v.X + v.Y*v.Y + v.Z*v.Z
		if rSquared == 0 {
			radii[i] = 0
			thetas[i] = 0
			phis[i] = 0
		} else {
			radii[i] = math.Sqrt(rSquared)
			thetas[i] = math.Atan2(v.Y, v.X)
			phis[i] = math.Acos(v.Z / radii[i])
		}
	}

	return radii, thetas, phis
}

// BulkPolarToVector converts multiple polar coordinates to 2D vectors
func BulkPolarToVector(radii, angles []float64) []Vector2D {
	n := len(radii)
	if n > len(angles) {
		n = len(angles)
	}

	vectors := make([]Vector2D, n)
	for i := 0; i < n; i++ {
		cosTheta := math.Cos(angles[i])
		sinTheta := math.Sin(angles[i])
		vectors[i] = Vector2D{radii[i] * cosTheta, radii[i] * sinTheta}
	}

	return vectors
}

// BulkSphericalToVector converts multiple spherical coordinates to 3D vectors
func BulkSphericalToVector(radii, thetas, phis []float64) []Vector3D {
	n := len(radii)
	if n > len(thetas) {
		n = len(thetas)
	}
	if n > len(phis) {
		n = len(phis)
	}

	vectors := make([]Vector3D, n)
	for i := 0; i < n; i++ {
		sinPhi := math.Sin(phis[i])
		cosPhi := math.Cos(phis[i])
		cosTheta := math.Cos(thetas[i])
		sinTheta := math.Sin(thetas[i])
		vectors[i] = Vector3D{
			radii[i] * sinPhi * cosTheta,
			radii[i] * sinPhi * sinTheta,
			radii[i] * cosPhi,
		}
	}

	return vectors
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
	rSquared := v.X*v.X + v.Y*v.Y
	r = math.Sqrt(rSquared)
	theta = math.Atan2(v.Y, v.X)
	return r, theta, v.Z
}

// CylindricalToVector3D converts cylindrical coordinates to a 3D vector
func CylindricalToVector3D(r, theta, z float64) Vector3D {
	return Vector3D{r * math.Cos(theta), r * math.Sin(theta), z}
}

// VectorToSpherical3D converts a 3D vector to spherical coordinates
func VectorToSpherical3D(v Vector3D) (r, theta, phi float64) {
	rSquared := v.X*v.X + v.Y*v.Y + v.Z*v.Z
	if rSquared == 0 {
		return 0, 0, 0
	}
	r = math.Sqrt(rSquared)
	theta = math.Atan2(v.Y, v.X)
	phi = math.Acos(v.Z / r)
	return r, theta, phi
}

// SphericalToVector3D converts spherical coordinates to a 3D vector
func SphericalToVector3D(r, theta, phi float64) Vector3D {
	sinPhi := math.Sin(phi)
	cosPhi := math.Cos(phi)
	cosTheta := math.Cos(theta)
	sinTheta := math.Sin(theta)
	return Vector3D{r * sinPhi * cosTheta, r * sinPhi * sinTheta, r * cosPhi}
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
