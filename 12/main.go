package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// http://www.pythonchallenge.com/pc/return/evil.html
// username: huge
// password: file

func main() {
	f, err := os.Open("evil2.gfx")
	if err != nil {
		fmt.Println("Failed openning input file")
		return
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("Failed reading input file")
		return
	}

	o1, err := os.Create("o1.jpg")
	if err != nil {
		fmt.Println("Failed openning o1")
		return
	}
	defer o1.Close()

	o2, err := os.Create("o2.jpg")
	if err != nil {
		fmt.Println("Failed openning o2")
		return
	}
	defer o2.Close()

	o3, err := os.Create("o3.jpg")
	if err != nil {
		fmt.Println("Failed openning o3")
		return
	}
	defer o3.Close()

	o4, err := os.Create("o4.jpg")
	if err != nil {
		fmt.Println("Failed openning o4")
		return
	}
	defer o4.Close()

	o5, err := os.Create("o5.jpg")
	if err != nil {
		fmt.Println("Failed openning o5")
		return
	}
	defer o5.Close()

	for i := 0; i < len(b); i += 5 {
		o1.Write(b[i+0 : i+1])
		o2.Write(b[i+1 : i+2])
		o3.Write(b[i+2 : i+3])
		o4.Write(b[i+3 : i+4])
		o5.Write(b[i+4 : i+5])
	}
}
