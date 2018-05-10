// http://www.pythonchallenge.com/pc/return/italy.html
// username: huge
// password: file

package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

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

	oi := image.NewNRGBA(image.Rect(0, 0, 100, 100))

	var pixcnt int
	for i := 0; i <= 49; i++ {
		for n := i; n <= 99-i; n++ {
			p := img.At(pixcnt, 0)
			pixcnt++
			oi.Set(i, n, p)
		}
		if i == 49 {
			break
		}
		for n := i + 1; n <= 99-i; n++ {
			p := img.At(pixcnt, 0)
			pixcnt++
			oi.Set(n, 99-i, p)
		}
		for n := 99 - i - 1; n >= i; n-- {
			p := img.At(pixcnt, 0)
			pixcnt++
			oi.Set(99-i, n, p)
		}
		for n := 99 - i - 1; n > i; n-- {
			p := img.At(pixcnt, 0)
			pixcnt++
			oi.Set(n, i, p)
		}
	}

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
