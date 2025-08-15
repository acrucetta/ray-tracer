// Package vec3 provides a minimal 3D vector type with
// idiomatic, value-receiver methods suitable for ray tracing
// and graphics math.
package vec3

import (
	"fmt"
	"math"
)

// Vec3 is a 3D vector.
//
// All methods use value receivers and return new values.
// For hot paths, this is usually fine with Go's escape analysis;
// switch to pointer receivers only if profiling shows a need.
type Vec3 struct {
	X, Y, Z float64
}

// Point3 is an alias for a 3D point (semantic clarity in APIs).
type Point3 = Vec3

// V constructs a new Vec3.
func V(x, y, z float64) Vec3 { return Vec3{X: x, Y: y, Z: z} }

// Add returns v + u.
func (v Vec3) Add(u Vec3) Vec3 {
	return Vec3{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

// Sub returns v - u.
func (v Vec3) Sub(u Vec3) Vec3 {
	return Vec3{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

// Mul returns the element-wise (Hadamard) product v ⊙ u.
func (v Vec3) Mul(u Vec3) Vec3 {
	return Vec3{v.X * u.X, v.Y * u.Y, v.Z * u.Z}
}

// Scale returns v scaled by scalar t.
func (v Vec3) Scale(t float64) Vec3 {
	return Vec3{v.X * t, v.Y * t, v.Z * t}
}

// Div returns v divided by scalar t.
// (Caller is responsible for ensuring t != 0.)
func (v Vec3) Div(t float64) Vec3 {
	inv := 1.0 / t
	return Vec3{v.X * inv, v.Y * inv, v.Z * inv}
}

// Neg returns -v.
func (v Vec3) Neg() Vec3 { return Vec3{-v.X, -v.Y, -v.Z} }

// Dot returns v ⋅ u.
func (v Vec3) Dot(u Vec3) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

// Cross returns v × u.
func (v Vec3) Cross(u Vec3) Vec3 {
	return Vec3{
		v.Y*u.Z - v.Z*u.Y,
		v.Z*u.X - v.X*u.Z,
		v.X*u.Y - v.Y*u.X,
	}
}

// Len2 returns |v|^2 (squared length).
func (v Vec3) Len2() float64 { return v.X*v.X + v.Y*v.Y + v.Z*v.Z }

// Len returns |v|.
func (v Vec3) Len() float64 { return math.Sqrt(v.Len2()) }

// Normalize returns v scaled to unit length.
// If v is near zero, it returns v unchanged to avoid NaNs.
func (v Vec3) Normalize() Vec3 {
	l := v.Len()
	if l == 0 {
		return v
	}
	return v.Div(l)
}

// NearZero reports whether all components are ~0.
// Useful to avoid artifacts when reflecting/refracting.
func (v Vec3) NearZero() bool {
	const eps = 1e-12
	return math.Abs(v.X) < eps && math.Abs(v.Y) < eps && math.Abs(v.Z) < eps
}

// String implements fmt.Stringer.
func (v Vec3) String() string { return fmt.Sprintf("%g %g %g", v.X, v.Y, v.Z) }
