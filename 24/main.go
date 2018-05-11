// http://www.pythonchallenge.com/pc/hex/ambiguity.html
// username: butter
// password: fly

package main

import (
	"fmt"
	"image/png"
	"os"
)

func main() {
	f, err := os.Open("maze.png")
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

	png.Encode(o, img)
}
