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

	var o [2]*os.File
	for x := 0; x < 2; x++ {
		imgname := fmt.Sprintf("./o%d.jpg", x)
		o[x], err = os.Create(imgname)
		if err != nil {
			fmt.Printf("Failed opening output %d, %v\n", x, err)
		}
		defer o[x].Close()
	}

	var i [2]*image.NRGBA
	for x := 0; x < 2; x++ {
		i[x] = image.NewNRGBA(image.Rect(0, 0, s.X, s.Y))
	}

	for x := 0; x < s.X; x++ {
		for y := 0; y < s.Y; y++ {
			c := img.At(x, y)
			a, b := x%2, y%2
			if a == 0 && b == 0 {
				i[0].Set(x, y, c)
			} else if a != 0 && b != 0 {
				i[1].Set(x, y, c)
			}
		}
	}

	for x := 0; x < 2; x++ {
		err := jpeg.Encode(o[x], i[x], &jpeg.Options{100})
		if err != nil {
			fmt.Printf("Failed writing image %d, %v\n", x, err)
		}
	}
}
