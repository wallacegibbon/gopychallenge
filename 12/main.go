// http://www.pythonchallenge.com/pc/return/evil.html
// username: huge
// password: file

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

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

	var o [5]*os.File
	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("o%d.jpg", i)
		o[i], err = os.Create(name)
		if err != nil {
			fmt.Printf("Failed openning o%d\n", i)
			return
		}
	}

	for i := 0; i < len(b); i += 5 {
		for j := 0; j < 5; j++ {
			o[j].Write(b[i+j : i+j+1])
		}
	}
}
