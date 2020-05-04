package features

import (
	"gonum.org/v1/gonum/mat"
)

// NewVecDenseFromTuple : convert Tuple to VecDense
func NewVecDenseFromTuple(t Tuple) *mat.VecDense {
	return mat.NewVecDense(4, []float64{t.x, t.y, t.z, t.w})
}

// NewTupleFromVecDense : convert VecDense to Tuple
func NewTupleFromVecDense(m *mat.VecDense) Tuple {
	return Tuple{
		m.AtVec(0),
		m.AtVec(1),
		m.AtVec(2),
		m.AtVec(3),
	}
}
