package features

import (
	"math"
	"testing"
)

func TestTranslationPoint(t *testing.T) {
	transform := TranslationNew(5, -3, 2)
	point := NewVecDenseFromTuple(PointNew(-3, 4, 5))
	point.MulVec(transform, point)
	got := NewTupleFromVecDense(point)
	want := PointNew(2, 1, 7)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInverseTranslationPoint(t *testing.T) {
	transform := TranslationNew(5, -3, 2)
	transform.Inverse(transform)
	point := NewVecDenseFromTuple(PointNew(-3, 4, 5))
	point.MulVec(transform, point)
	got := NewTupleFromVecDense(point)
	want := PointNew(-8, 7, 3)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// Translation should have no effect on a vector
func TestTranslationVector(t *testing.T) {
	transform := TranslationNew(5, -3, 2)
	transform.Inverse(transform)
	v := NewVecDenseFromTuple(VectorNew(-3, 4, 5))
	v.MulVec(transform, v)
	got := NewTupleFromVecDense(v)
	want := VectorNew(-3, 4, 5)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestScalingPoint(t *testing.T) {
	transform := ScalingNew(2, 3, 4)
	point := NewVecDenseFromTuple(PointNew(-4, 6, 8))
	point.MulVec(transform, point)
	got := NewTupleFromVecDense(point)
	want := PointNew(-8, 18, 32)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestScalingVector(t *testing.T) {
	transform := ScalingNew(2, 3, 4)
	point := NewVecDenseFromTuple(VectorNew(-4, 6, 8))
	point.MulVec(transform, point)
	got := NewTupleFromVecDense(point)
	want := VectorNew(-8, 18, 32)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInverseScalingVector(t *testing.T) {
	transform := ScalingNew(2, 3, 4)
	transform.Inverse(transform)
	point := NewVecDenseFromTuple(VectorNew(-4, 6, 8))
	point.MulVec(transform, point)
	got := NewTupleFromVecDense(point)
	want := VectorNew(-2, 2, 2)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// Reflection is scaling by a negative value
func TestReflectionByScaling(t *testing.T) {
	transform := ScalingNew(-1, 1, 1)
	point := NewVecDenseFromTuple(PointNew(2, 3, 4))
	point.MulVec(transform, point)
	got := NewTupleFromVecDense(point)
	want := PointNew(-2, 3, 4)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRotationX(t *testing.T) {
	halfQuarter := RotationXNew(math.Pi / 4)
	p1 := NewVecDenseFromTuple(PointNew(0, 1, 0))
	p1.MulVec(halfQuarter, p1)
	got := NewTupleFromVecDense(p1)
	want := PointNew(0, math.Sqrt2/2, math.Sqrt2/2)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}

	fullQuarter := RotationXNew(math.Pi / 2)
	p1 = NewVecDenseFromTuple(PointNew(0, 1, 0))
	p1.MulVec(fullQuarter, p1)
	got = NewTupleFromVecDense(p1)
	want = PointNew(0, 0, 1)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRotationXInverse(t *testing.T) {
	halfQuarter := RotationXNew(math.Pi / 4)
	halfQuarter.Inverse(halfQuarter)
	p1 := NewVecDenseFromTuple(PointNew(0, 1, 0))
	p1.MulVec(halfQuarter, p1)
	got := NewTupleFromVecDense(p1)
	want := PointNew(0, math.Sqrt2/2, -math.Sqrt2/2)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRotationY(t *testing.T) {
	halfQuarter := RotationYNew(math.Pi / 4)
	p1 := NewVecDenseFromTuple(PointNew(0, 0, 1))
	p1.MulVec(halfQuarter, p1)
	got := NewTupleFromVecDense(p1)
	want := PointNew(math.Sqrt2/2, 0, math.Sqrt2/2)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}

	fullQuarter := RotationYNew(math.Pi / 2)
	p1 = NewVecDenseFromTuple(PointNew(0, 0, 1))
	p1.MulVec(fullQuarter, p1)
	got = NewTupleFromVecDense(p1)
	want = PointNew(1, 0, 0)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRotationZ(t *testing.T) {
	p1 := PointNew(0, 1, 0)
	halfQuarter := RotationZNew(math.Pi / 4)
	got := TransformTuple(p1, halfQuarter)
	want := PointNew(-math.Sqrt2/2, math.Sqrt2/2, 0)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}

	fullQuarter := RotationZNew(math.Pi / 2)
	got = TransformTuple(p1, fullQuarter)
	want = PointNew(-1, 0, 0)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestShear(t *testing.T) {
	transform := ShearNew(1, 0, 0, 0, 0, 0)
	p := PointNew(2, 3, 4)
	got := TransformTuple(p, transform)
	want := PointNew(5, 3, 4)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}

	transform = ShearNew(0, 1, 0, 0, 0, 0)
	p = PointNew(2, 3, 4)
	got = TransformTuple(p, transform)
	want = PointNew(6, 3, 4)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}

	transform = ShearNew(0, 0, 1, 0, 0, 0)
	p = PointNew(2, 3, 4)
	got = TransformTuple(p, transform)
	want = PointNew(2, 5, 4)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestChain(t *testing.T) {
	// translation is applied last (right to left)
	transform := ChainTransform(TranslationNew(10, 5, 7), ScalingNew(5, 5, 5), RotationXNew(math.Pi/2))
	p := PointNew(1, 0, 1)
	got := TransformTuple(p, transform)
	want := PointNew(15, 0, 7)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}
