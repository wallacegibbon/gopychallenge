// http://www.pythonchallenge.com/pc/hex/ambiguity.html
// username: butter
// password: fly

package main

import (
	"fmt"
	"image"
	"image/color"
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

	o1, err := os.Create("out.png")
	if err != nil {
		fmt.Println("Failed opening output image")
		return
	}
	defer o1.Close()

	o2, err := os.Create("out.zip")
	if err != nil {
		fmt.Println("Failed opening output image")
		return
	}
	defer o2.Close()

	rawimg, err := png.Decode(f)
	if err != nil {
		fmt.Println("Failed decoding image")
		return
	}

	img := rawimg.(*image.NRGBA)

	xy := img.Bounds().Max

	entry := image.Point{xy.X - 2, 0}
	exit := image.Point{1, xy.Y - 1}

	pmap := make(map[image.Point]image.Point)

	stack := make([]image.Point, 1)
	stack[0] = exit

	dirs := [4]image.Point{}
	dirs[0].X, dirs[0].Y = 0, 1
	dirs[1].X, dirs[1].Y = 0, -1
	dirs[2].X, dirs[2].Y = 1, 0
	dirs[3].X, dirs[3].Y = -1, 0

	var cur image.Point

	for {
		slen := len(stack)
		cur = stack[slen-1]
		stack = stack[:slen-1]

		if eql(cur, entry) {
			break
		}
		for _, v := range dirs {
			tmp := image.Point{cur.X + v.X, cur.Y + v.Y}
			if _, ok := pmap[tmp]; ok {
				continue
			}
			if tmp.X < 0 || tmp.X >= xy.X {
				continue
			}
			if tmp.Y < 0 || tmp.Y >= xy.Y {
				continue
			}
			if isWhite(img.At(tmp.X, tmp.Y)) {
				continue
			}
			pmap[tmp] = cur
			stack = append(stack, tmp)
		}
	}

	//fmt.Println("result route find")

	path := make([]byte, 10)
	for !eql(cur, exit) {
		c := img.At(cur.X, cur.Y)
		r, _, _, _ := c.RGBA()
		path = append(path, byte(r>>8))

		img.Set(cur.X, cur.Y, color.RGBA{0, 255, 255, 255})

		cur = pmap[cur]
	}

	p1 := make([]byte, len(path) / 2 + 1)
	for i := 0; i < len(path); i += 2 {
		p1[i/2] = path[i + 1]
	}

	png.Encode(o1, img)
	o2.Write(p1)
}

func eql(p1, p2 image.Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

// The wall is white
func isWhite(c color.Color) bool {
	r, g, b, a := c.RGBA()
	x := uint32(65535)
	return r == x && g == x && b == x && a == x
}
