package shapes

import (
	"errors"
	"gonum.org/v1/gonum/mat"
	"math"
	"sarim-tracer/features/rays"
	"sarim-tracer/features/transformations"
	"sarim-tracer/features/tuples"
	"sort"
)

// Shape interface
type Shape interface {
	Intersect(r rays.Ray) []Intersection
}

// Sphere : a Shape
type Sphere struct {
	transform *mat.Dense
}

// SphereNew : sphere constructor
//
// using variadic function to make transform optional
func SphereNew(transform ...*mat.Dense) Sphere {
	if len(transform) == 0 {
		return Sphere{transformations.IdentityNew(4)}
	}
	return Sphere{transform[0]}
}

// Intersect sphere with ray
func (s Sphere) Intersect(r rays.Ray) []Intersection {
	// inverse transform the ray
	transform := transformations.IdentityNew(4)
	transform.Inverse(s.transform)
	r = r.Transform(transform)
	// assume unit sphere at global origin
	// create ray from sphere center to ray origin
	sphereToRay := r.Origin.Subtract(tuples.PointNew(0, 0, 0))
	// compute discriminant
	a := r.Direction.DotProduct(r.Direction)
	b := 2.0 * r.Direction.DotProduct(sphereToRay)
	c := sphereToRay.DotProduct(sphereToRay) - 1.0
	discriminant := (b * b) - (4 * a * c)
	// ray misses the sphere
	if discriminant < 0 {
		return []Intersection{}
	}
	// ray is either crossing, tangent, inside or in front of the sphere
	t1 := (-b - math.Sqrt(discriminant)) / (2.0 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2.0 * a)
	return []Intersection{IntersectionNew(t1, s),
		IntersectionNew(t2, s)}
}

// Intersection : intersection result of ray with shape
type Intersection struct {
	IntersectionValue float64
	Shape             Shape
}

// IntersectionNew : constructor for Intersect type
func IntersectionNew(intersectionValue float64, shape Shape) Intersection {
	return Intersection{intersectionValue, shape}
}

// IntersectionHit : the hit will always be the intersection with the lowest nonnegative t
// value.
func IntersectionHit(intersections []Intersection) (Intersection, error) {
	// sort the intersections in ascending order by intersectionValue
	sort.Slice(intersections, func(i, j int) bool {
		return intersections[i].IntersectionValue < intersections[j].IntersectionValue
	})
	// starting from the beginning, find the first non-neg intersection and return
	for _, intersection := range intersections {
		if !(intersection.IntersectionValue < 0.0) {
			return intersection, nil
		}
	}
	err := errors.New("No non-negative intersection found")
	return IntersectionNew(0.0, nil), err
}
