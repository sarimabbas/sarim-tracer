package main

import (
	"math"
)

var epsilon float64 = 0.00001

// Tuple type
type Tuple struct {
	x, y, z, w float64
}

// PointNew : initialize and create a new Tuple
func PointNew(x, y, z float64) Tuple {
	var t Tuple
	t.x = x
	t.y = y
	t.z = z
	t.w = 1.0
	return t
}

// VectorNew : initialize and create a new Tuple
func VectorNew(x, y, z float64) Tuple {
	var t Tuple
	t.x = x
	t.y = y
	t.z = z
	t.w = 0.0
	return t
}

func (t Tuple) isVector() bool {
	return t.w == 0.0
}

func (t Tuple) isPoint() bool {
	return t.w == 1.0
}

// FloatEqual : compare two float64 as equal
func FloatEqual(a, b float64) bool {
	if math.Abs(a-b) < epsilon {
		return true
	}
	return false
}

// TupleEqual : compare two Tuple as equal
func TupleEqual(a, b Tuple) bool {
	if FloatEqual(a.x, b.x) &&
		FloatEqual(a.y, b.y) &&
		FloatEqual(a.z, b.z) &&
		FloatEqual(a.w, b.w) {
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
	res.x = a.x + b.x
	res.y = a.y + b.y
	res.z = a.z + b.z
	res.w = a.w + b.w
	return res
}

// Add : add two tuples
func (t Tuple) Add(u Tuple) Tuple {
	return TupleAdd(t, u)
}

// TupleSubtract : subtract b from a and return the result
func TupleSubtract(a, b Tuple) Tuple {
	var res Tuple
	res.x = a.x - b.x
	res.y = a.y - b.y
	res.z = a.z - b.z
	res.w = a.w - b.w
	return res
}

// Subtract : subtract two tuples
func (t Tuple) Subtract(u Tuple) Tuple {
	return TupleSubtract(t, u)
}

// TupleNegate : negate a tuple (flip it)
func TupleNegate(t Tuple) Tuple {
	t.x = -t.x
	t.y = -t.y
	t.z = -t.z
	t.w = -t.w
	return t
}

// Negate : negate a tuple (flip it)
func (t Tuple) Negate() Tuple {
	return TupleNegate(t)
}

// TupleScalarMultiply : scale each component
func TupleScalarMultiply(t Tuple, s float64) Tuple {
	t.x = s * t.x
	t.y = s * t.y
	t.z = s * t.z
	t.w = s * t.w
	return t
}

// ScalarMultiply : scale a tuple
func (t Tuple) ScalarMultiply(s float64) Tuple {
	return TupleScalarMultiply(t, s)
}

// TupleScalarDivide : divide each component
func TupleScalarDivide(t Tuple, s float64) Tuple {
	t.x = t.x / 2
	t.y = t.y / 2
	t.z = t.z / 2
	t.w = t.w / 2
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
	return math.Sqrt(t.x*t.x + t.y*t.y + t.z*t.z + t.w*t.w)
}

// Magnitude : get length of tuple
func (t Tuple) Magnitude() float64 {
	return TupleMagnitude(t)
}

// TupleNormalize : normalize tuple
func TupleNormalize(t Tuple) Tuple {
	m := t.Magnitude()
	t.x /= m
	t.y /= m
	t.z /= m
	t.w /= m
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
	res.x = t.x * u.x
	res.y = t.y * u.y
	res.z = t.z * u.z
	res.w = t.w * u.w
	return res.x + res.y + res.z + res.w
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
	return VectorNew(a.y*b.z-a.z*b.y, a.z*b.x-a.x*b.z, a.x*b.y-a.y*b.x)
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
