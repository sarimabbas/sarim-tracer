package features

import (
	"gonum.org/v1/gonum/mat"
	"sarim-tracer/features/tuples"
	"testing"
)

func Test4X4Matrix(t *testing.T) {
	m := mat.NewDense(4, 4, nil)
	m.Set(0, 3, 4)
	got := m.At(0, 3)
	want := 4.0
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestMatrixEquality(t *testing.T) {
	dat := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}
	m1 := mat.NewDense(4, 4, dat)
	m2 := mat.NewDense(4, 4, dat)
	got := mat.EqualApprox(m1, m2, tuples.EPSILON)
	want := true
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestMatrixMultiply(t *testing.T) {
	m1 := mat.NewDense(4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2})
	m2 := mat.NewDense(4, 4, []float64{-2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 1, 2, 7, 8})
	m1.Mul(m1, m2)
	got := m1
	want := mat.NewDense(4, 4, []float64{20, 22, 50, 48, 44, 54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42})
	if !mat.EqualApprox(got, want, tuples.EPSILON) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestMatrixTupleMultiply(t *testing.T) {
	m1 := mat.NewDense(4, 4, []float64{1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1})
	m2 := tuples.NewVecDenseFromTuple(tuples.Tuple{X: 1, Y: 2, Z: 3, W: 1})

	m2.MulVec(m1, m2)
	got := tuples.NewTupleFromVecDense(m2)
	want := tuples.Tuple{X: 18, Y: 24, Z: 33, W: 1}
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestIdentityMatrix(t *testing.T) {
	m1 := mat.NewDense(4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2})
	dat := []float64{}
	for i := 0; i < 16; i++ {
		dat = append(dat, 1.0)
	}
	iden := mat.NewDense(4, 4, dat)
	iden.Mul(m1, iden)
	got := iden
	want := m1
	if mat.EqualApprox(got, want, tuples.EPSILON) {
		t.Errorf("got %v want %v", got, want)
	}
}
