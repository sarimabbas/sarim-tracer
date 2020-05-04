package features

import (
	"fmt"
	"os"
)

// Canvas : image canvas to save to PPM
type Canvas struct {
	width, height int
	pixels        [][]Tuple
}

// CanvasNew : create new canvas
func CanvasNew(width, height int) Canvas {
	// create pixels
	pixels := make([][]Tuple, width)
	for i := range pixels {
		pixels[i] = make([]Tuple, height)
	}

	// init pixels
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			pixels[i][j] = ColorNew(0, 0, 0)
		}
	}
	canvas := Canvas{width: width, height: height, pixels: pixels}
	return canvas
}

// SetPixel : write pixel to canvas
func (c *Canvas) SetPixel(x, y int, pixel Tuple) {
	// if flipY {
	// 	y = c.height - 1 - y
	// }
	// c.pixels[(y*c.width)+x] = pixel
	c.pixels[x][y] = pixel
}

// GetPixel : get pixel from canvas
func (c Canvas) GetPixel(x, y int) Tuple {
	// return c.pixels[(y*c.width)+x]
	return c.pixels[x][y]
}

// ToPPM : dump canvas to file
func (c Canvas) ToPPM(path string, flipX, flipY bool) {
	// check path
	if path == "" {
		return
	}
	// open file
	os.Remove(path)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	// header
	header := fmt.Sprintf("P3\n%d %d\n255\n", c.width, c.height)
	if _, err = f.WriteString(header); err != nil {
		panic(err)
	}
	// body
	pixelCount := 0
	for j := 0; j < c.height; j++ {
		for i := 0; i < c.width; i++ {
			// clamp the pixel and write it
			modI := i
			modJ := j
			if flipY {
				modJ = c.height - j - 1
			}
			if flipX {
				modI = c.width - i - 1
			}
			clamped := c.pixels[modI][modJ].Clamp(0.0, 1.0)
			clampedString := fmt.Sprintf("%d %d %d ",
				int(clamped.x*255), int(clamped.y*255), int(clamped.z*255))
			if _, err = f.WriteString(clampedString); err != nil {
				panic(err)
			}
			// add a line break every five pixels (roughly 70 chars)
			if pixelCount%5 == 0 {
				if _, err = f.WriteString("\n"); err != nil {
					panic(err)
				}
			}
			// increment count
			pixelCount++
		}
	}
	// footer
	if _, err = f.WriteString("\n"); err != nil {
		panic(err)
	}
	// close file
	f.Close()
}
