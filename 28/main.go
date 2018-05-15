// http://www.pythonchallenge.com/pc/ring/bell.html
// username: repeat
// password: switch

package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
)

func main() {
	f, err := os.Open("bell.png")
	if err != nil {
		fmt.Println("Failed opening input image")
		return
	}
	defer f.Close()

	o, err := os.Create("out.png")
	if err != nil {
		fmt.Println("Failed opening output image")
		return
	}
	defer o.Close()

	img, err := png.Decode(f)
	if err != nil {
		fmt.Println("Failed decoding image")
		return
	}

	xy := img.Bounds().Max
	lst := make([]byte, 0)

	for y := 0; y < xy.Y; y++ {
		for x := 0; x < xy.X; x += 2 {
			v := handleTwoPoint(img.At(x, y), img.At(x+1, y))
			if v != 42 {
				lst = append(lst, byte(v))
			}
		}
	}

	fmt.Println(string(lst))
}

func handleTwoPoint(c1, c2 color.Color) byte {
	_, g1, _, _ := c1.RGBA()
	_, g2, _, _ := c2.RGBA()
	g1 >>= 8
	g2 >>= 8
	v := int8(g2 - g1)
	if v < 0 {
		v = -v
	}
	return byte(v)
}
