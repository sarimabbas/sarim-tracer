package transformations

import (
	"math"
	"sarim-tracer/features/tuples"
	"testing"
)

func TestTranslationPoint(t *testing.T) {
	transform := TranslationNew(5, -3, 2)
	point := tuples.NewVecDenseFromTuple(tuples.PointNew(-3, 4, 5))
	point.MulVec(transform, point)
	got := tuples.NewTupleFromVecDense(point)
	want := tuples.PointNew(2, 1, 7)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInverseTranslationPoint(t *testing.T) {
	transform := TranslationNew(5, -3, 2)
	transform.Inverse(transform)
	point := tuples.NewVecDenseFromTuple(tuples.PointNew(-3, 4, 5))
	point.MulVec(transform, point)
	got := tuples.NewTupleFromVecDense(point)
	want := tuples.PointNew(-8, 7, 3)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// Translation should have no effect on a vector
func TestTranslationVector(t *testing.T) {
	transform := TranslationNew(5, -3, 2)
	transform.Inverse(transform)
	v := tuples.NewVecDenseFromTuple(tuples.VectorNew(-3, 4, 5))
	v.MulVec(transform, v)
	got := tuples.NewTupleFromVecDense(v)
	want := tuples.VectorNew(-3, 4, 5)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestScalingPoint(t *testing.T) {
	transform := ScalingNew(2, 3, 4)
	point := tuples.NewVecDenseFromTuple(tuples.PointNew(-4, 6, 8))
	point.MulVec(transform, point)
	got := tuples.NewTupleFromVecDense(point)
	want := tuples.PointNew(-8, 18, 32)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestScalingVector(t *testing.T) {
	transform := ScalingNew(2, 3, 4)
	point := tuples.NewVecDenseFromTuple(tuples.VectorNew(-4, 6, 8))
	point.MulVec(transform, point)
	got := tuples.NewTupleFromVecDense(point)
	want := tuples.VectorNew(-8, 18, 32)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInverseScalingVector(t *testing.T) {
	transform := ScalingNew(2, 3, 4)
	transform.Inverse(transform)
	point := tuples.NewVecDenseFromTuple(tuples.VectorNew(-4, 6, 8))
	point.MulVec(transform, point)
	got := tuples.NewTupleFromVecDense(point)
	want := tuples.VectorNew(-2, 2, 2)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// Reflection is scaling by a negative value
func TestReflectionByScaling(t *testing.T) {
	transform := ScalingNew(-1, 1, 1)
	point := tuples.NewVecDenseFromTuple(tuples.PointNew(2, 3, 4))
	point.MulVec(transform, point)
	got := tuples.NewTupleFromVecDense(point)
	want := tuples.PointNew(-2, 3, 4)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRotationX(t *testing.T) {
	halfQuarter := RotationXNew(math.Pi / 4)
	p1 := tuples.NewVecDenseFromTuple(tuples.PointNew(0, 1, 0))
	p1.MulVec(halfQuarter, p1)
	got := tuples.NewTupleFromVecDense(p1)
	want := tuples.PointNew(0, math.Sqrt2/2, math.Sqrt2/2)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}

	fullQuarter := RotationXNew(math.Pi / 2)
	p1 = tuples.NewVecDenseFromTuple(tuples.PointNew(0, 1, 0))
	p1.MulVec(fullQuarter, p1)
	got = tuples.NewTupleFromVecDense(p1)
	want = tuples.PointNew(0, 0, 1)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRotationXInverse(t *testing.T) {
	halfQuarter := RotationXNew(math.Pi / 4)
	halfQuarter.Inverse(halfQuarter)
	p1 := tuples.NewVecDenseFromTuple(tuples.PointNew(0, 1, 0))
	p1.MulVec(halfQuarter, p1)
	got := tuples.NewTupleFromVecDense(p1)
	want := tuples.PointNew(0, math.Sqrt2/2, -math.Sqrt2/2)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRotationY(t *testing.T) {
	halfQuarter := RotationYNew(math.Pi / 4)
	p1 := tuples.NewVecDenseFromTuple(tuples.PointNew(0, 0, 1))
	p1.MulVec(halfQuarter, p1)
	got := tuples.NewTupleFromVecDense(p1)
	want := tuples.PointNew(math.Sqrt2/2, 0, math.Sqrt2/2)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}

	fullQuarter := RotationYNew(math.Pi / 2)
	p1 = tuples.NewVecDenseFromTuple(tuples.PointNew(0, 0, 1))
	p1.MulVec(fullQuarter, p1)
	got = tuples.NewTupleFromVecDense(p1)
	want = tuples.PointNew(1, 0, 0)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRotationZ(t *testing.T) {
	p1 := tuples.PointNew(0, 1, 0)
	halfQuarter := RotationZNew(math.Pi / 4)
	got := tuples.TupleTransform(p1, halfQuarter)
	want := tuples.PointNew(-math.Sqrt2/2, math.Sqrt2/2, 0)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}

	fullQuarter := RotationZNew(math.Pi / 2)
	got = tuples.TupleTransform(p1, fullQuarter)
	want = tuples.PointNew(-1, 0, 0)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestShear(t *testing.T) {
	transform := ShearNew(1, 0, 0, 0, 0, 0)
	p := tuples.PointNew(2, 3, 4)
	got := tuples.TupleTransform(p, transform)
	want := tuples.PointNew(5, 3, 4)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}

	transform = ShearNew(0, 1, 0, 0, 0, 0)
	p = tuples.PointNew(2, 3, 4)
	got = tuples.TupleTransform(p, transform)
	want = tuples.PointNew(6, 3, 4)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}

	transform = ShearNew(0, 0, 1, 0, 0, 0)
	p = tuples.PointNew(2, 3, 4)
	got = tuples.TupleTransform(p, transform)
	want = tuples.PointNew(2, 5, 4)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestChain(t *testing.T) {
	// translation is applied last (right to left)
	transform := ChainTransform(TranslationNew(10, 5, 7), ScalingNew(5, 5, 5), RotationXNew(math.Pi/2))
	p := tuples.PointNew(1, 0, 1)
	got := tuples.TupleTransform(p, transform)
	want := tuples.PointNew(15, 0, 7)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}
