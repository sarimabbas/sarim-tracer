package rays_test

import (
	"sarim-tracer/features/rays"
	"sarim-tracer/features/shapes"
	"sarim-tracer/features/transformations"
	"sarim-tracer/features/tuples"
	"testing"
)

func TestRayPointCompute(t *testing.T) {
	ray := rays.RayNew(tuples.PointNew(2, 3, 4), tuples.VectorNew(1, 0, 0))
	got := ray.Position(1)
	want := tuples.PointNew(3, 3, 4)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestIntersectionHit(t *testing.T) {
	s := shapes.SphereNew()
	intersections := []shapes.Intersection{
		shapes.IntersectionNew(5, s),
		shapes.IntersectionNew(7, s),
		shapes.IntersectionNew(-3, s),
		shapes.IntersectionNew(2, s),
	}
	i, err := shapes.IntersectionHit(intersections)
	if err != nil {
		panic(err)
	}
	got := i.IntersectionValue
	want := 2.0
	if !tuples.FloatEqual(got, want) {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestTranslateRay(t *testing.T) {
	r1 := rays.RayNew(tuples.PointNew(1, 2, 3), tuples.VectorNew(0, 1, 0))
	m := transformations.ScalingNew(2, 3, 4)
	r2 := r1.Transform(m)
	got := r2.Origin
	want := tuples.PointNew(2, 6, 12)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
	got = r2.Direction
	want = tuples.VectorNew(0, 3, 0)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}
