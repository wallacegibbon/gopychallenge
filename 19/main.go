package main

import (
	"encoding/base64"
	"os"
	"fmt"
	"io"
)

// http://www.pythonchallenge.com/pc/hex/bin.html
// username: butter
// password: fly

func main() {
	f, err := os.Open("wav.txt")
	if err != nil {
		fmt.Println("Failed opening wav.txt")
		return
	}
	defer f.Close()

	r := base64.NewDecoder(base64.StdEncoding, f)

	w, err := os.Create("o.wav")
	if err != nil {
		fmt.Println("Failed opening output file")
		return
	}
	defer w.Close()

	io.Copy(w, r)
}
