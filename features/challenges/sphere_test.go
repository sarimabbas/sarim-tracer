package challenges

import (
	"math"
	"sarim-tracer/features/canvas"
	"sarim-tracer/features/rays"
	"sarim-tracer/features/shapes"
	"sarim-tracer/features/transformations"
	"sarim-tracer/features/tuples"
	"testing"
)

func TestDrawSphere(t *testing.T) {
	// start the ray at z = -5
	rayOrigin := tuples.PointNew(0, 0, -5)
	// put the wall at z = 10
	wallZ := 10.0
	wallSize := 7.0
	// pixel size
	canvasPixels := 100
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2.0

	canvas := canvas.CanvasNew(canvasPixels, canvasPixels)
	s := shapes.SphereNew(transformations.ScalingNew(1, 0.5, 1), transformations.RotationXNew(math.Pi/4))

	// for each row of pixels in canvas
	for y := 0; y < canvasPixels-1; y++ {
		// compute the world y coordinate
		worldY := half - pixelSize*float64(y)
		// for each pixel in the row
		for x := 0; x < canvasPixels-1; x++ {
			// compute the world x coordinate
			worldX := -half + pixelSize*float64(x)
			// describe point on wall that ray will target
			position := tuples.PointNew(worldX, worldY, wallZ)

			r := rays.RayNew(rayOrigin, position.Subtract(rayOrigin).Normalize())
			xs := s.Intersect(r)

			// if there is a hit, write pixel
			if len(xs) > 0 {
				canvas.SetPixel(x, y, tuples.ColorNew(1, 0, 0))
			}
		}
	}

	// dump to image
	canvas.ToPPM("sphere_test.ppm", false, true)

}
