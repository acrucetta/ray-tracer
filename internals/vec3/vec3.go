package vec3

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

type Point3 = Vec3

func New(x, y, z float64) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}

func (v Vec3) Neg() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

// Add returns the sum of two vectors.
func (v Vec3) Add(u Vec3) Vec3 {
	return Vec3{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

// AddAssign adds u to v in place.
func (v *Vec3) AddAssign(u Vec3) {
	v.X += u.X
	v.Y += u.Y
	v.Z += u.Z
}

// Scale returns the vector scaled by t.
func (v Vec3) Scale(t float64) Vec3 {
	return Vec3{v.X * t, v.Y * t, v.Z * t}
}

// ScaleAssign multiplies the vector by t in place.
func (v *Vec3) ScaleAssign(t float64) {
	v.X *= t
	v.Y *= t
	v.Z *= t
}

// Div returns the vector divided by t.
func (v Vec3) Div(t float64) Vec3 {
	return v.Scale(1 / t)
}

// DivAssign divides the vector by t in place.
func (v *Vec3) DivAssign(t float64) {
	v.ScaleAssign(1 / t)
}

// Length returns the vector's length.
func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// LengthSquared returns the squared length of the vector.
func (v Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// String implements the fmt.Stringer interface for printing.
func (v Vec3) String() string {
	return fmt.Sprintf("%g %g %g", v.X, v.Y, v.Z)
}

// Add returns the vector sum of u and v.
func Add(u, v Vec3) Vec3 {
	return Vec3{u.X + v.X, u.Y + v.Y, u.Z + v.Z}
}

// Sub returns the vector difference u - v.
func Sub(u, v Vec3) Vec3 {
	return Vec3{u.X - v.X, u.Y - v.Y, u.Z - v.Z}
}

// Mul returns the element-wise product of u and v.
func Mul(u, v Vec3) Vec3 {
	return Vec3{u.X * v.X, u.Y * v.Y, u.Z * v.Z}
}

// MulScalar returns the product of vector v and scalar t.
func MulScalar(v Vec3, t float64) Vec3 {
	return Vec3{v.X * t, v.Y * t, v.Z * t}
}

// DivScalar divides vector v by scalar t.
func DivScalar(v Vec3, t float64) Vec3 {
	invT := 1.0 / t
	return MulScalar(v, invT)
}

// Dot returns the dot product of u and v.
func Dot(u, v Vec3) float64 {
	return u.X*v.X + u.Y*v.Y + u.Z*v.Z
}

// Cross returns the cross product of u and v.
func Cross(u, v Vec3) Vec3 {
	return Vec3{
		u.Y*v.Z - u.Z*v.Y,
		u.Z*v.X - u.X*v.Z,
		u.X*v.Y - u.Y*v.X,
	}
}

// Unit returns a unit vector in the direction of v.
func Unit(v Vec3) Vec3 {
	return DivScalar(v, v.Length())
}
