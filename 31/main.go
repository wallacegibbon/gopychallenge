// http://www.pythonchallenge.com/pc/ring/grandpa.html
// username: repeat
// password: switch

// http://www.pythonchallenge.com/pc/rock/grandpa.html
// username: kohsamui
// password: thailand

// This level is not solved yet.
// I am not sure about the reason, but there are 3255 different points in the
// result(which is wrong, it should be 1679(23*73) different points)
// Even thought the mandelbrot shape looks the same.

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"math/cmplx"
	"os"
)

const maxIteration = 128

func main() {
	f, err := os.Open("mandelbrot.gif")
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

	img, err := gif.Decode(f)
	if err != nil {
		fmt.Println("Failed decoding image")
		return
	}

	img1, ok := img.(*image.Paletted)
	if !ok {
		fmt.Println("Failed get input paletted image")
		return
	}

	oi := image.NewNRGBA(img.Bounds())

	l1 := img1.Pix
	l2 := draw(0.34, 0.57, 0.036, 0.027, oi)

	//fmt.Printf("(%d)\n=%x\n>%x\n", len(l1), l1[:50], l2[:50])

	diff := make([]uint8, 0)
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			diff = append(diff, l1[i]-l2[i])
		}
	}

	//fmt.Printf("(%d)%x\n", len(diff), diff[:500])

	//the output image is just for checking the shape. the color
	//won't be the same as the gif because that is paletted image
	png.Encode(o, oi)
}

func draw(left, top, width, height float64, img *image.NRGBA) []uint8 {
	xy := img.Bounds().Max
	xs := width / float64(xy.X)
	ys := height / float64(xy.Y)
	r := make([]uint8, xy.X*xy.Y)
	for y := xy.Y - 1; y >= 0; y-- {
		for x := 0; x < xy.X; x++ {
			c := complex(left+float64(x)*xs, top+float64(y)*ys)
			cnt := uint8(calc(c))
			img.Set(x, xy.Y-1-y, color.NRGBA{0, 0, cnt, 255})
			r[(xy.Y-1-y)*xy.X+x] = cnt
		}
	}
	return r
}

func calc(c complex128) int {
	z := 0 + 0i
	var cnt int
	for ; cnt < maxIteration; cnt++ {
		z = z*z + c
		if cmplx.Abs(z) > 2 {
			break
		}
	}
	return cnt
}
