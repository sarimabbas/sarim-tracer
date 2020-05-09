package rays

import (
	"gonum.org/v1/gonum/mat"
	"sarim-tracer/features/tuples"
)

// Ray : Ray structure
type Ray struct {
	Origin    tuples.Tuple
	Direction tuples.Tuple
}

// RayNew : create a Ray
func RayNew(origin, direction tuples.Tuple) Ray {
	return Ray{origin, direction}
}

// RayPosition : get new point from moving along the ray
func RayPosition(r Ray, t float64) tuples.Tuple {
	return r.Origin.Add(r.Direction.ScalarMultiply(t))
}

// Position : get new point from moving along the ray
func (r Ray) Position(t float64) tuples.Tuple {
	return RayPosition(r, t)
}

// RayTransform : conveniently transform rays
func RayTransform(ray Ray, transform *mat.Dense) Ray {
	o := ray.Origin.Transform(transform)
	d := ray.Direction.Transform(transform)
	return RayNew(o, d)
}

// Transform : conveniently transform rays
func (r Ray) Transform(transform *mat.Dense) Ray {
	return RayTransform(r, transform)
}
