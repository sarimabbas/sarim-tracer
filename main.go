package main

import (
	"sarim-tracer/features/canvas"
	"sarim-tracer/features/tuples"
)

func main() {
	c := canvas.CanvasNew(900, 500)
	for i := 0; i < 400; i++ {
		for j := 0; j < 200; j++ {
			c.SetPixel(i, j, tuples.ColorNew(1, 0, 0))
		}
	}
	c.ToPPM("main.ppm", false, true)
}
