// http://www.pythonchallenge.com/pc/hex/speedboat.html
// username: butter
// password: fly

package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"os"
)

func main() {
	f, err := os.Open("zigzag.gif")
	if err != nil {
		fmt.Println("Failed opening input file")
		return
	}
	defer f.Close()

	oi, err := os.Create("out.png")
	if err != nil {
		fmt.Println("Failed opening output file")
		return
	}
	defer oi.Close()

	of, err := os.Create("out.bz2")
	if err != nil {
		fmt.Println("Failed opening output file")
		return
	}
	defer of.Close()

	img, err := gif.Decode(f)
	if err != nil {
		fmt.Println("Failed decoding gif image")
		return
	}

	ipalette, ok := img.(*image.Paletted)
	if !ok {
		fmt.Println("Convert to paletted failed")
		return
	}

	p1s := ipalette.Pix
	palette := ipalette.Palette
	p2s := make([]uint8, len(p1s))

	//fmt.Printf("palette length: %d, palette: %v\n",
	//	len(palette), palette[:10])

	for i := 0; i < len(p1s); i++ {
		rawr, _, _, _ := palette[p1s[i]].RGBA()
		r := uint8(rawr >> 8)
		p2s[i] = r
	}

	//fmt.Printf("%x <-> %x\n", p1s[0:10], p2s[0:10])

	//this is one of the most important operation
	p1s = p1s[1:]

	diff := make([]uint8, 0)

	for i := 0; i < len(p1s); i++ {
		if p2s[i] != p1s[i] {
			diff = append(diff, p1s[i])
		}
	}

	fmt.Printf("Diff length: %d\n", len(diff))
	of.Write(diff)

	for i := 0; i < len(p1s); i++ {
		if p2s[i] != p1s[i] {
			p2s[i] = 255
		} else {
			p2s[i] = 0
		}
	}

	imgo := image.NewGray(img.Bounds())
	imgo.Pix = p2s

	png.Encode(oi, imgo)
}
