// http://www.pythonchallenge.com/pc/hex/copper.html
// username: butter
// password: fly

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"os"
)

func main() {
	f, err := os.Open("white.gif")
	if err != nil {
		fmt.Println("Failed opening input image")
		return
	}
	defer f.Close()

	img, err := gif.DecodeAll(f)
	if err != nil {
		fmt.Println("Failed decode gif")
		return
	}

	o, err := os.Create("out.png")
	if err != nil {
		fmt.Println("Failed opening output file")
		return
	}
	defer o.Close()

	oi := image.NewNRGBA(image.Rect(0, 0, 500, 200))
	for _, p := range collectPoints(img.Image) {
		oi.Set(p.X, p.Y, color.RGBA{255, 0, 0, 255})
	}

	png.Encode(o, oi)
}

func collectPoints(imgs []*image.Paletted) []image.Point {
	points := make([]image.Point, 10)
	cur := image.Point{10, 10}
	for _, i := range imgs {
		x, y, found := findWhiteSpot(i)
		if !found {
			continue
		}
		dx, dy := x-100, y-100
		if dx == 0 && dy == 0 {
			cur.X += 50
		} else {
			cur.X += dx
			cur.Y += dy
		}
		points = append(points, cur)
	}
	return points
}

func findWhiteSpot(img *image.Paletted) (int, int, bool) {
	xy := img.Bounds().Max
	for x := 0; x < xy.X; x++ {
		for y := 0; y < xy.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if r != 0 && g != 0 && b != 0 {
				return x, y, true
			}
		}
	}
	return 0, 0, false
}
