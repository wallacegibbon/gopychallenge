package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

// http://www.pythonchallenge.com/pc/return/italy.html
// username: huge
// password: file

func main() {
	f, err := os.Open("./wire.png")
	if err != nil {
		fmt.Println("Failed opening wire.png")
		return
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		fmt.Println("Failed decode image")
		return
	}

	s := img.Bounds().Max
	//fmt.Println(s)

	oi := image.NewNRGBA(image.Rect(0, 0, 100, 100))

	o, err := os.Create("./o.png")
	if err != nil {
		fmt.Println("Failed create output image")
		return
	}
	defer o.Close()

	err = png.Encode(o, oi)
	if err != nil {
		fmt.Println("Failed encode output image")
	}
}
