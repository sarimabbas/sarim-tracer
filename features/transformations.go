package features

import (
	"gonum.org/v1/gonum/mat"
	"math"
)

// TranslationNew : construct a translation matrix
func TranslationNew(x, y, z float64) *mat.Dense {
	m := mat.NewDense(4, 4, nil)
	m.Zero()
	m.SetRow(0, []float64{1, 0, 0, x})
	m.SetRow(1, []float64{0, 1, 0, y})
	m.SetRow(2, []float64{0, 0, 1, z})
	m.SetRow(3, []float64{0, 0, 0, 1})
	return m
}

// ScalingNew : construct a scaling matrix
func ScalingNew(x, y, z float64) *mat.Dense {
	m := mat.NewDense(4, 4, nil)
	m.Zero()
	m.SetRow(0, []float64{x, 0, 0, 0})
	m.SetRow(1, []float64{0, y, 0, 0})
	m.SetRow(2, []float64{0, 0, z, 0})
	m.SetRow(3, []float64{0, 0, 0, 1})
	return m
}

// RotationXNew : construct a rotation matrix around x axis
func RotationXNew(deg float64) *mat.Dense {
	m := mat.NewDense(4, 4, nil)
	m.Zero()
	m.SetRow(0, []float64{1, 0, 0, 0})
	m.SetRow(1, []float64{0, math.Cos(deg), -math.Sin(deg), 0})
	m.SetRow(2, []float64{0, math.Sin(deg), math.Cos(deg), 0})
	m.SetRow(3, []float64{0, 0, 0, 1})
	return m
}

// RotationYNew : construct a rotation matrix around y axis
func RotationYNew(deg float64) *mat.Dense {
	m := mat.NewDense(4, 4, nil)
	m.Zero()
	m.SetRow(0, []float64{math.Cos(deg), 0, math.Sin(deg), 0})
	m.SetRow(1, []float64{0, 1, 0, 0})
	m.SetRow(2, []float64{-math.Sin(deg), 0, math.Cos(deg), 0})
	m.SetRow(3, []float64{0, 0, 0, 1})
	return m
}

// RotationZNew : construct a rotation matrix around z axis
func RotationZNew(deg float64) *mat.Dense {
	m := mat.NewDense(4, 4, nil)
	m.Zero()
	m.SetRow(0, []float64{math.Cos(deg), -math.Sin(deg), 0, 0})
	m.SetRow(1, []float64{math.Sin(deg), math.Cos(deg), 0, 0})
	m.SetRow(2, []float64{0, 0, 0, 0})
	m.SetRow(3, []float64{0, 0, 0, 1})
	return m
}

// ShearNew : construct a shear matrix
//
// “When applied to a tuple, a shearing transformation changes each component of
// the tuple in proportion to the other two components. So the x component
// changes in proportion to y and z, y changes in proportion to x and z, and z
// changes in proportion to x and y.”
func ShearNew(Xy, Xz, Yx, Yz, Zx, Zy float64) *mat.Dense {
	m := mat.NewDense(4, 4, nil)
	m.Zero()
	m.SetRow(0, []float64{1, Xy, Xz, 0})
	m.SetRow(1, []float64{Yx, 1, Yz, 0})
	m.SetRow(2, []float64{Zx, Zy, 1, 0})
	m.SetRow(3, []float64{0, 0, 0, 1})
	return m
}

// TransformTuple : conveniently transform tuples
func TransformTuple(t Tuple, transform *mat.Dense) Tuple {
	v := NewVecDenseFromTuple(t)
	v.MulVec(transform, v)
	return NewTupleFromVecDense(v)
}

// IdentityNew : return identity matrix
func IdentityNew(n int) *mat.Dense {
	iden := mat.NewDense(n, n, nil)
	iden.Zero()
	for i := 0; i < n; i++ {
		iden.Set(i, i, 1)
	}
	return iden
}

// ChainTransform : join multiple transforms, right to left
func ChainTransform(transforms ...*mat.Dense) *mat.Dense {
	final := IdentityNew(4)
	for _, t := range transforms {
		final.Mul(final, t)
	}
	return final
}
