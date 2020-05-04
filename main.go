package main

import (
	"sarim-tracer/features"
)

func main() {
	c := features.CanvasNew(900, 500)
	for i := 0; i < 400; i++ {
		for j := 0; j < 200; j++ {
			c.SetPixel(i, j, features.ColorNew(1, 0, 0))
		}
	}
	c.ToPPM("main.ppm", false, true)
}
