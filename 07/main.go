package main

import (
	"fmt"
	"image/png"
	"os"
)

// http://www.pythonchallenge.com/pc/def/oxygen.html

func main() {
	f, err := os.Open("./oxygen.png")
	if err != nil {
		fmt.Println("Failed opening image")
		return
	}

	img, err := png.Decode(f)
	if err != nil {
		fmt.Println("Failed decoding image")
		return
	}

	s := img.Bounds().Max
	fmt.Println(s)
	y := s.Y / 2
	for x := 0; x < s.X; x += 7 {
		r, _, _, _ := img.At(x, y).RGBA()
		fmt.Printf("%c", r>>8)
	}
	fmt.Println()

	x := []int{105, 110, 116, 101, 103, 114, 105, 116, 121}
	for _, c := range x {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}
