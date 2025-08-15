package ray

import "ray-tracer/internals/vec3"

type Ray struct {
	Origin, Direction vec3.Vec3
}

// NewRay creates a new Ray given an origin and direction.
func NewRay(origin vec3.Vec3, direction vec3.Point3) Ray {
	return Ray{
		Origin:    origin,
		Direction: direction,
	}
}

// At returns the point along the ray at parameter t.
func (r Ray) At(t float64) vec3.Vec3 {
	return vec3.Vec3{
		X: r.Origin.X + t*r.Direction.X,
		Y: r.Origin.Y + t*r.Direction.Y,
		Z: r.Origin.Z + t*r.Direction.Z,
	}
}
