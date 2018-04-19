package main

import (
	"fmt"
	"io/ioutil"
	"unicode"
)

// http://www.pythonchallenge.com/pc/def/ocr.html

func main() {
	d, err := ioutil.ReadFile("./raw.txt")
	if err != nil {
		fmt.Println("Failed reading raw.txt:", err)
		return
	}
	for _, c := range d {
		if unicode.IsLetter(rune(c)) {
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()
}
