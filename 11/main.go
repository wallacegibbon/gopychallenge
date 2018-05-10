// http://www.pythonchallenge.com/pc/return/5808.html
// username: huge
// password: file

package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func main() {
	f, err := os.Open("./cave.jpg")
	if err != nil {
		fmt.Println("Failed openning image")
		return
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		fmt.Println("Failed decoding image")
		return
	}

	s := img.Bounds().Max

	o, err := os.Create("o.jpg")
	if err != nil {
		fmt.Println("Failed openning output file")
		return
	}

	i := image.NewNRGBA(image.Rect(0, 0, s.X, s.Y))

	for x := 0; x < s.X; x++ {
		for y := 0; y < s.Y; y++ {
			c := img.At(x, y)
			a, b := x%2, y%2
			if a == 0 && b == 0 || a != 0 && b != 0 {
				i.Set(x, y, c)
			}
		}
	}

	err = jpeg.Encode(o, i, &jpeg.Options{100})
	if err != nil {
		fmt.Println("Failed writing image")
	}
}
