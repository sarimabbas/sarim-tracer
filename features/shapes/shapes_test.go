package shapes

import (
	"gonum.org/v1/gonum/mat"
	"sarim-tracer/features/rays"
	"sarim-tracer/features/transformations"
	"sarim-tracer/features/tuples"
	"testing"
)

func TestSphereIntersect(t *testing.T) {
	r := rays.RayNew(tuples.PointNew(0, 0, -5), tuples.VectorNew(0, 0, 1))
	s := SphereNew()
	intersections := s.Intersect(r)
	if !(len(intersections) == 2) {
		t.Errorf("got %v want %v", len(intersections), 2)
	}
	if !(tuples.FloatEqual(intersections[0].IntersectionValue, 4.0)) {
		t.Errorf("got %v want %v", intersections[0].IntersectionValue, 4.0)
	}
	if !(tuples.FloatEqual(intersections[1].IntersectionValue, 6.0)) {
		t.Errorf("got %v want %v", intersections[1].IntersectionValue, 6.0)
	}
}

func TestSphereTransform(t *testing.T) {
	s := SphereNew()
	if !mat.EqualApprox(s.Transform, transformations.IdentityNew(4), tuples.EPSILON) {
		t.Errorf("got %v want %v", s.Transform, transformations.IdentityNew(4))
	}
	s.Transform = transformations.TranslationNew(2, 3, 4)
	if !mat.EqualApprox(s.Transform, transformations.TranslationNew(2, 3, 4), tuples.EPSILON) {
		t.Errorf("got %v want %v", s.Transform, transformations.TranslationNew(2, 3, 4))
	}
}

func TestScaledSphereIntersect(t *testing.T) {
	r := rays.RayNew(tuples.PointNew(0, 0, -5), tuples.VectorNew(0, 0, 1))
	s := SphereNew()
	s.Transform = transformations.ScalingNew(2, 2, 2)
	intersections := s.Intersect(r)
	if len(intersections) != 2 {
		t.Errorf("got %d want %d", len(intersections), 2)
	}
	if !tuples.FloatEqual(intersections[0].IntersectionValue, 3) {
		t.Errorf("got %f want %f", intersections[0].IntersectionValue, 3.0)
	}
	if !tuples.FloatEqual(intersections[1].IntersectionValue, 7) {
		t.Errorf("got %f want %f", intersections[1].IntersectionValue, 7.0)
	}
}

func TestTranslatedSphereIntersect(t *testing.T) {
	r := rays.RayNew(tuples.PointNew(0, 0, -5), tuples.VectorNew(0, 0, 1))
	s := SphereNew()
	s.Transform = transformations.TranslationNew(5, 0, 0)
	intersections := s.Intersect(r)
	if len(intersections) != 0 {
		t.Errorf("got %d want %d", len(intersections), 0)
	}
}
