package features

import (
	"math"
	"testing"
)

func TestTuple(t *testing.T) {
	got := Tuple{4.3, -4.2, 3.1, 1.0}
	if got.x != 4.3 {
		t.Errorf("got %f want %f", got.x, 4.3)
	}
	if got.y != -4.2 {
		t.Errorf("got %f want %f", got.y, -4.2)
	}
	if got.z != 3.1 {
		t.Errorf("got %f want %f", got.z, 3.1)
	}
	if got.w != 1.0 {
		t.Errorf("got %f want %f", got.w, 1.0)
	}
}

func TestPoint(t *testing.T) {
	got := PointNew(4, -4, 3)
	want := Tuple{4, -4, 3, 1}
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestVector(t *testing.T) {
	got := VectorNew(4, -4, 3)
	want := Tuple{4, -4, 3, 0}
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestTupleAdd(t *testing.T) {
	a1 := Tuple{3, -2, 5, 1}
	a2 := Tuple{-2, 3, 1, 0}
	got := a1.Add(a2)
	want := Tuple{1, 1, 6, 1}
	if !TupleEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestTupleSubtract(t *testing.T) {
	a1 := PointNew(3, 2, 1)
	a2 := PointNew(5, 6, 7)
	got := a1.Subtract(a2)
	want := VectorNew(-2, -4, -6)
	if !TupleEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSubtractVectorFromPoint(t *testing.T) {
	p := PointNew(3, 2, 1)
	v := VectorNew(5, 6, 7)
	got := p.Subtract(v)
	want := PointNew(-2, -4, -6)
	if !TupleEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSubtractTwoVectors(t *testing.T) {
	v1 := VectorNew(3, 2, 1)
	v2 := VectorNew(5, 6, 7)
	got := v1.Subtract(v2)
	want := VectorNew(-2, -4, -6)
	if !TupleEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNegateTuple(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	got := a.Negate()
	want := Tuple{-1, 2, -3, 4}
	if !TupleEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestScaleTuple(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	got := a.ScalarMultiply(3.5)
	want := Tuple{3.5, -7, 10.5, -14}
	if !TupleEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDivideTuple(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	got := a.ScalarDivide(2)
	want := Tuple{0.5, -1, 1.5, -2}
	if !TupleEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestMagnitudeVector(t *testing.T) {
	v := VectorNew(1, 0, 0)
	got := v.Magnitude()
	want := 1.0
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	v = VectorNew(0, 1, 0)
	got = v.Magnitude()
	want = 1.0
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	v = VectorNew(0, 0, 1)
	got = v.Magnitude()
	want = 1.0
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	v = VectorNew(1, 2, 3)
	got = v.Magnitude()
	want = math.Sqrt(14)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	v = VectorNew(-1, -2, -3)
	got = v.Magnitude()
	want = math.Sqrt(14)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNormalizeVector(t *testing.T) {
	v := VectorNew(4, 0, 0)
	got := v.Normalize()
	want := VectorNew(1, 0, 0)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	v = VectorNew(1, 2, 3)
	got = v.Normalize()
	want = VectorNew(0.26726, 0.53452, 0.80178)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}

	gotmag := got.Magnitude()
	wantmag := 1.0
	if gotmag != wantmag {
		t.Errorf("got %v want %v", gotmag, wantmag)
	}
}

func TestTupleDotProduct(t *testing.T) {
	v1 := VectorNew(1, 2, 3)
	v2 := VectorNew(2, 3, 4)
	got := v1.DotProduct(v2)
	want := 20.0
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestVectorCrossProduct(t *testing.T) {
	v1 := VectorNew(1, 2, 3)
	v2 := VectorNew(2, 3, 4)
	got := v1.CrossProduct(v2)
	want := VectorNew(-1, 2, -1)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
	got = v2.CrossProduct(v1)
	want = VectorNew(1, -2, 1)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestTupleHadamardProduct(t *testing.T) {
	c1 := ColorNew(1, 0.2, 0.4)
	c2 := ColorNew(0.9, 1.0, 0.1)
	got := c1.HadamardProduct(c2)
	want := ColorNew(0.9, 0.2, 0.04)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}
