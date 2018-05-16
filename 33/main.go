// http://www.pythonchallenge.com/pc/rock/beer.html
// username: kohsamui
// password: thailand

package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {
	f, err := os.Open("beer2.png")
	if err != nil {
		fmt.Println("failed opening input image")
		return
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		fmt.Println("failed decoding input image")
		return
	}

	img1, ok := img.(*image.Gray)
	if !ok {
		fmt.Println("failed trans input image")
		return
	}

	pixels := img1.Pix

	fmt.Printf("(%d)%v\n", len(pixels), pixels[:100])
}
