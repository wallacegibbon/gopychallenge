package draw

import (
	"image"
	"image/color"
	"errors"
)

// The Bresenham algorithm
func DrawLine(img *image.NRGBA, x0, y0, x1, y1 int) {
	dx, dy := iAbs(x1-x0), iAbs(y1-y0)
	sx, sy := -1, -1
	if x0 < x1 {
		sx = 1
	}
	if y0 < y1 {
		sy = 1
	}

	var e, e2 int
	if dx > dy {
		e = dx / 2
	} else {
		e = -dy / 2
	}

	for {
		img.Set(x0, y0, color.RGBA{255,0,0,255})
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 = e
		if e2 > -dx {
			e -= dy
			x0 += sx
		}
		if e2 < dy {
			e += dx
			y0 += sy
		}
	}
}

func Polygon(img *image.NRGBA, ps []int) error {
	if len(ps) < 4 || len(ps) % 2 != 0 {
		return errors.New("polygon point argument odd")
	}
	for x := 0; x < len(ps) - 2; x += 2 {
		DrawLine(img, ps[x+0], ps[x+1], ps[x+2], ps[x+3])
	}
	return nil
}

func iAbs(num int) int {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}
