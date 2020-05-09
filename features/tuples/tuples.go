package tuples

import (
	"gonum.org/v1/gonum/mat"
	"math"
)

// EPSILON used for floating point comparison
var EPSILON float64 = 0.00001

// Tuple type
type Tuple struct {
	X, Y, Z, W float64
}

// TupleNew : initialize and create a new Tuple
func TupleNew(x, y, z, w float64) Tuple {
	var t Tuple
	t.X = x
	t.Y = y
	t.Z = z
	t.W = w
	return t
}

// PointNew : initialize and create a new Tuple
func PointNew(x, y, z float64) Tuple {
	var t Tuple
	t.X = x
	t.Y = y
	t.Z = z
	t.W = 1.0
	return t
}

// VectorNew : initialize and create a new Tuple
func VectorNew(x, y, z float64) Tuple {
	var t Tuple
	t.X = x
	t.Y = y
	t.Z = z
	t.W = 0.0
	return t
}

func (t Tuple) isVector() bool {
	return t.W == 0.0
}

func (t Tuple) isPoint() bool {
	return t.W == 1.0
}

// FloatEqual : compare two float64 as equal
func FloatEqual(a, b float64) bool {
	if math.Abs(a-b) < EPSILON {
		return true
	}
	return false
}

// TupleEqual : compare two Tuple as equal
func TupleEqual(a, b Tuple) bool {
	if FloatEqual(a.X, b.X) &&
		FloatEqual(a.Y, b.Y) &&
		FloatEqual(a.Z, b.Z) &&
		FloatEqual(a.W, b.W) {
		return true
	}
	return false
}

// Equal : compare two tuples
func (t Tuple) Equal(u Tuple) bool {
	return TupleEqual(t, u)
}

// TupleAdd : add two tuples and return the result
func TupleAdd(a, b Tuple) Tuple {
	var res Tuple
	res.X = a.X + b.X
	res.Y = a.Y + b.Y
	res.Z = a.Z + b.Z
	res.W = a.W + b.W
	return res
}

// Add : add two tuples
func (t Tuple) Add(u Tuple) Tuple {
	return TupleAdd(t, u)
}

// TupleSubtract : subtract b from a and return the result
func TupleSubtract(a, b Tuple) Tuple {
	var res Tuple
	res.X = a.X - b.X
	res.Y = a.Y - b.Y
	res.Z = a.Z - b.Z
	res.W = a.W - b.W
	return res
}

// Subtract : subtract two tuples
func (t Tuple) Subtract(u Tuple) Tuple {
	return TupleSubtract(t, u)
}

// TupleNegate : negate a tuple (flip it)
func TupleNegate(t Tuple) Tuple {
	t.X = -t.X
	t.Y = -t.Y
	t.Z = -t.Z
	t.W = -t.W
	return t
}

// Negate : negate a tuple (flip it)
func (t Tuple) Negate() Tuple {
	return TupleNegate(t)
}

// TupleScalarMultiply : scale each component
func TupleScalarMultiply(t Tuple, s float64) Tuple {
	t.X = s * t.X
	t.Y = s * t.Y
	t.Z = s * t.Z
	t.W = s * t.W
	return t
}

// ScalarMultiply : scale a tuple
func (t Tuple) ScalarMultiply(s float64) Tuple {
	return TupleScalarMultiply(t, s)
}

// TupleScalarDivide : divide each component
func TupleScalarDivide(t Tuple, s float64) Tuple {
	t.X = t.X / 2
	t.Y = t.Y / 2
	t.Z = t.Z / 2
	t.W = t.W / 2
	return t
}

// ScalarDivide : divide a tuple
func (t Tuple) ScalarDivide(s float64) Tuple {
	return TupleScalarDivide(t, s)
}

// TupleMagnitude : get length of tuple
//
// Remember, at the start of this chapter, when you read that a vector was a
// value that encoded direction and distance? The distance represented by a
// vector is called its magnitude, or length. It’s how far you would travel in a
// straight line if you were to walk from one end of the vector to the other
func TupleMagnitude(t Tuple) float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

// Magnitude : get length of tuple
func (t Tuple) Magnitude() float64 {
	return TupleMagnitude(t)
}

// TupleNormalize : normalize tuple
func TupleNormalize(t Tuple) Tuple {
	m := t.Magnitude()
	t.X /= m
	t.Y /= m
	t.Z /= m
	t.W /= m
	return t
}

// Normalize : normalize tuple
func (t Tuple) Normalize() Tuple {
	return TupleNormalize(t)
}

// TupleDotProduct : dot two tuples
//
// “The dot product can feel pretty abstract, but here’s one quick way to
// internalize it: the smaller the dot product, the larger the angle between the
// vectors. For example, given two unit vectors, a dot product of 1 means the
// vectors are identical, and a dot product of -1 means they point in opposite
// directions. More specifically, and again if the two vectors are unit vectors,
// the dot product is actually the cosine of the angle between them”
func TupleDotProduct(t Tuple, u Tuple) float64 {
	var res Tuple
	res.X = t.X * u.X
	res.Y = t.Y * u.Y
	res.Z = t.Z * u.Z
	res.W = t.W * u.W
	return res.X + res.Y + res.Z + res.W
}

// DotProduct : dot two tuples
//
// “The dot product can feel pretty abstract, but here’s one quick way to
// internalize it: the smaller the dot product, the larger the angle between the
// vectors. For example, given two unit vectors, a dot product of 1 means the
// vectors are identical, and a dot product of -1 means they point in opposite
// directions. More specifically, and again if the two vectors are unit vectors,
// the dot product is actually the cosine of the angle between them”
func (t Tuple) DotProduct(u Tuple) float64 {
	return TupleDotProduct(t, u)
}

// VectorCrossProduct : take the cross product of two vectors
//
// “Note that this is specifically testing vectors, not tuples. This is because
// the four-dimensional cross product is significantly more complicated than the
// three-dimensional cross product, and your ray tracer really only needs the
// three-dimensional version anyway.”
func VectorCrossProduct(a Tuple, b Tuple) Tuple {
	return VectorNew(a.Y*b.Z-a.Z*b.Y, a.Z*b.X-a.X*b.Z, a.X*b.Y-a.Y*b.X)
}

// CrossProduct : take the cross product of two vectors
//
// “Note that this is specifically testing vectors, not tuples. This is because
// the four-dimensional cross product is significantly more complicated than the
// three-dimensional cross product, and your ray tracer really only needs the
// three-dimensional version anyway.”
func (t Tuple) CrossProduct(u Tuple) Tuple {
	return VectorCrossProduct(t, u)
}

// TupleHadamardProduct : take the component-wise product
func TupleHadamardProduct(a Tuple, b Tuple) Tuple {
	return Tuple{a.X * b.X, a.Y * b.Y, a.Z * b.Z, a.W * b.W}
}

// HadamardProduct : take the component-wise product
//
// c1 = (1, 0.2, 0.4); c2 = (0.9, 0.2, 0.04)
//
// res = (0.9, 0.2, 0.04)
//
// “Consider this test again. It says that if you were to view that yellow-green
// surface (c2) under a reddish-purple light (c1), the result color will seem
// red (because the result's red component, 0.9, is largest).”
func (t Tuple) HadamardProduct(u Tuple) Tuple {
	return TupleHadamardProduct(t, u)
}

// ColorNew : create color tuple
func ColorNew(r, g, b float64) Tuple {
	var c Tuple
	c.X = r
	c.Y = g
	c.Z = b
	c.W = 1.0
	return c
}

// FloatClamp : clamp a float between low and high
func FloatClamp(f, low, high float64) float64 {
	if f < low {
		return low
	}
	if f > high {
		return high
	}
	return f
}

// Clamp : clamp tuple components
func (t Tuple) Clamp(low, high float64) Tuple {
	t.X = FloatClamp(t.X, low, high)
	t.Y = FloatClamp(t.Y, low, high)
	t.Z = FloatClamp(t.Z, low, high)
	t.W = FloatClamp(t.W, low, high)
	return t
}

// NewVecDenseFromTuple : convert Tuple to VecDense
func NewVecDenseFromTuple(t Tuple) *mat.VecDense {
	return mat.NewVecDense(4, []float64{t.X, t.Y, t.Z, t.W})
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

// TupleTransform : conveniently transform tuples
func TupleTransform(t Tuple, transform *mat.Dense) Tuple {
	v := NewVecDenseFromTuple(t)
	v.MulVec(transform, v)
	return NewTupleFromVecDense(v)
}

// Transform : conveniently transform tuples
func (t Tuple) Transform(transform *mat.Dense) Tuple {
	return TupleTransform(t, transform)
}
