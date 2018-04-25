package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

// http://www.pythonchallenge.com/pc/return/5808.html
// username: huge
// password: file

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

	o1, err := os.Create("./o1.jpg")
	if err != nil {
		fmt.Println("Failed opening o1:", err)
	}
	defer o1.Close()

	o2, err := os.Create("./o2.jpg")
	if err != nil {
		fmt.Println("Failed opening o2:", err)
	}
	defer o2.Close()

	i1 := image.NewNRGBA(image.Rect(0, 0, s.X, s.Y))
	i2 := image.NewNRGBA(image.Rect(0, 0, s.X, s.Y))

	for x := 0; x < s.X; x++ {
		for y := 0; y < s.Y; y++ {
			c := img.At(x, y)
			a, b := x%2, y%2
			if a == 0 && b == 0 {
				i1.Set(x, y, c)
			} else if a != 0 && b != 0 {
				i2.Set(x, y, c)
			}
		}
	}

	err = jpeg.Encode(o1, i1, &jpeg.Options{100})
	if err != nil {
		fmt.Println("Failed writing image:", err)
	}

	err = jpeg.Encode(o2, i2, &jpeg.Options{100})
	if err != nil {
		fmt.Println("Failed writing image:", err)
	}
}
