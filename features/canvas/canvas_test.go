package canvas

import (
	"sarim-tracer/features/tuples"
	"testing"
)

func TestCanvasCreate(t *testing.T) {
	c := CanvasNew(10, 20)
	got := c.width
	want := 10
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
	got = c.height
	want = 20
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			if c.pixels[i][j] != tuples.ColorNew(0, 0, 0) {
				t.Errorf("got %v want %v", c.pixels[i][j], tuples.ColorNew(0, 0, 0))
			}
		}

	}
}

func TestReadWriteCanvas(t *testing.T) {
	c := CanvasNew(10, 20)
	red := tuples.ColorNew(1, 0, 0)
	c.SetPixel(9, 19, red)
	got := c.GetPixel(9, 19)
	want := tuples.ColorNew(1, 0, 0)
	if !got.Equal(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCanvasToPPM(t *testing.T) {
	c := CanvasNew(10, 20)
	red := tuples.ColorNew(1, 0, 0)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			c.SetPixel(i, j, red)
		}
	}
	c.ToPPM("canvas_test.ppm", false, true)
}
