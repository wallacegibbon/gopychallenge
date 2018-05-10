// http://www.pythonchallenge.com/pc/return/mozart.html
// username: huge
// password: file

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
	in, err := os.Open("mozart.gif")
	if err != nil {
		fmt.Println("Failed opening input image")
		return
	}
	defer in.Close()

	out, err := os.Create("out.png")
	if err != nil {
		fmt.Println("Failed opening output image")
		return
	}
	defer out.Close()

	img, err := gif.Decode(in)
	if err != nil {
		fmt.Println("Failed decoding input image")
		return
	}

	io := image.NewNRGBA(img.Bounds())

	xy := img.Bounds().Max
	line := make([]color.Color, xy.X)

	for y := 0; y < xy.Y; y++ {
		for x := 0; x < xy.X; x++ {
			line[x] = img.At(x, y)
		}
		line = shiftLine(line)
		for x := 0; x < xy.X; x++ {
			io.Set(x, y, line[x])
		}
	}

	err = png.Encode(out, io)
	if err != nil {
		fmt.Println("Failed writing output image")
		return
	}
}

func shiftLine(line []color.Color) []color.Color {
	var i int
	for i = 0; i < len(line); i++ {
		r, g, b, _ := line[i].RGBA()
		r >>= 8
		g >>= 8
		b >>= 8
		if r == 255 && g == 0 && b == 255 {
			break
		}
	}
	return append(line[i:], line[:i]...)
}
