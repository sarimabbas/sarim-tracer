package challenges

import (
	"math"
	"sarim-tracer/features/canvas"
	"sarim-tracer/features/transformations"
	"sarim-tracer/features/tuples"
	"testing"
)

func TestDrawClock(t *testing.T) {
	canvasWidth := 100
	c := canvas.CanvasNew(canvasWidth, canvasWidth)

	// “Multiplying a point in object space by a transformation matrix converts
	// that point to world space—scaling it, translating, rotating it, or
	// whatever.”
	twelve := tuples.PointNew(0, 0, 1)

	translate := transformations.TranslationNew(float64(canvasWidth/2), 0, float64(canvasWidth/2))
	clockRadius := (3.0 / 8.0) * float64(canvasWidth)
	scale := transformations.ScalingNew(clockRadius, 0, clockRadius)
	for i := 0; i < 12; i++ {
		// construct point by rotating the twelve
		r := transformations.RotationYNew(float64(i) * math.Pi / 6.0)
		hour := tuples.TupleTransform(twelve, r)
		// scale
		hour = tuples.TupleTransform(hour, scale)
		// translate hour to center
		hour = tuples.TupleTransform(hour, translate)
		// set pixel
		c.SetPixel(int(hour.X), int(hour.Z), tuples.ColorNew(1, 0, 0))
	}
	c.ToPPM("clock_test.ppm", false, true)
}
