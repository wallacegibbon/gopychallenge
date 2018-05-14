// http://www.pythonchallenge.com/pc/hex/lake.html
// username: butter
// password: fly

package main

import (
	"fmt"
	"github.com/wallacegibbon/gowav"
	"image"
	"image/color"
	"image/png"
	"io"
	"net/http"
	"os"
)

func main() {
	/*
		err := downloadWaves()
		if err != nil {
			fmt.Println("Failed download wav files:", err)
			return
		}
	*/
	err := handleWaves()
	if err != nil {
		fmt.Println("Failed handling wav files:", err)
		return
	}
}

func downloadWaves() error {
	for i := 1; i <= 25; i++ {
		resp, err := doReq(i)
		if err != nil {
			return err
		}
		tgtFile := fmt.Sprintf("pieces/%02d.wav", i)
		out, err := os.Create(tgtFile)
		if err != nil {
			return err
		}
		io.Copy(out, resp.Body)
		out.Close()
	}
	return nil
}

func handleWaves() error {
	img := image.NewNRGBA(image.Rect(0, 0, 300, 300))
	o, err := os.Create("out.png")
	if err != nil {
		return err
	}
	defer o.Close()
	for i := 0; i < 25; i++ {
		raw, err := getWaveBytes(i + 1)
		if err != nil {
			return err
		}
		writeSquare(img, zeroPoint(i), raw)
	}
	png.Encode(o, img)
	return nil
}

func getWaveBytes(idx int) ([]byte, error) {
	tgt := fmt.Sprintf("pieces/%02d.wav", idx)
	fmt.Println("Getting bytes from " + tgt)
	w, err := gowav.NewWavFile(tgt)
	if err != nil {
		return nil, err
	}
	t, err := w.GetAllFrames()
	if err != nil {
		return nil, err
	}
	return t, nil
}

func writeSquare(img *image.NRGBA, base *image.Point, raw []byte) {
	for y := 0; y < 60; y++ {
		for x := 0; x < 60; x++ {
			offset := y*60*3 + x*3
			c := bytesToColor(raw[offset : offset+3])
			img.Set(base.X+x, base.Y+y, c)
		}
	}
}

func bytesToColor(raw []byte) color.Color {
	r := uint8(raw[0])
	g := uint8(raw[1])
	b := uint8(raw[2])
	return &color.NRGBA{r, g, b, 255}
}

func zeroPoint(idx int) *image.Point {
	line := idx / 5
	column := idx % 5
	return &image.Point{column * 60, line * 60}
}

func doReq(idx int) (*http.Response, error) {
	const url = "http://www.pythonchallenge.com/pc/hex/lake%d.wav"
	tgt := fmt.Sprintf(url, idx)
	fmt.Println("Fetching " + tgt)

	req, err := http.NewRequest("GET", tgt, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Basic YnV0dGVyOmZseQ==")

	cli := &http.Client{}
	return cli.Do(req)
}
