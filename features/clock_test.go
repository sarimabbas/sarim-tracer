package features

import (
	"math"
	"testing"
)

func TestDrawClock(t *testing.T) {
	canvasWidth := 100
	c := CanvasNew(canvasWidth, canvasWidth)
	twelve := PointNew(0, 0, 1)
	translate := TranslationNew(float64(canvasWidth/2), 0, float64(canvasWidth/2))
	clockRadius := (3.0 / 8.0) * float64(canvasWidth)
	scale := ScalingNew(clockRadius, 0, clockRadius)
	for i := 0; i < 12; i++ {
		// construct point by rotating the twelve
		r := RotationYNew(float64(i) * math.Pi / 6.0)
		hour := TransformTuple(twelve, r)
		// scale
		hour = TransformTuple(hour, scale)
		// translate hour to center
		hour = TransformTuple(hour, translate)
		// set pixel
		c.SetPixel(int(hour.x), int(hour.z), ColorNew(1, 0, 0))
	}
	c.ToPPM("clock_test.ppm", false, true)
}
