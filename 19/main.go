package main

import (
	"encoding/base64"
	"fmt"
	"github.com/wallacegibbon/gowav"
	"io"
	"os"
)

// http://www.pythonchallenge.com/pc/hex/bin.html
// username: butter
// password: fly

func main() {
	// Giving the Base64 stream to NewWav will cause error for some reason
	err := makeRawAudio("wav.txt", "raw.wav")
	if err != nil {
		fmt.Println("Failed parsing raw file")
		return
	}

	in, err := os.Open("raw.wav")
	if err != nil {
		fmt.Println("Failed opening raw.wav")
		return
	}

	out, err := os.Create("out.wav")
	if err != nil {
		fmt.Println("Failed opening output file")
		return
	}

	w, err := gowav.NewWav(in)
	if err != nil {
		fmt.Println("Failed loading wav:", err)
		return
	}

	w.WriteParams(out)
	for {
		frm, err := w.GetFrame()
		if err != nil {
			fmt.Println("GetFrame error:", err)
			return
		}
		reverseSlice(frm)
		if frm != nil {
			out.Write(frm)
		} else {
			break
		}
	}

}

func makeRawAudio(rawfile, outfile string) error {
	f, err := os.Open(rawfile)
	if err != nil {
		return err
	}
	defer f.Close()

	r := base64.NewDecoder(base64.StdEncoding, f)

	w, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}
	return nil
}

func reverseSlice(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		tgt := len(s) - 1 - i
		tmp := s[i]
		s[i] = s[tgt]
		s[tgt] = tmp
	}
}
