package main

import (
	"fmt"
)

// http://www.pythonchallenge.com/pc/def/0.html

func main() {
	var n uint64 = 1
	fmt.Printf("The result of 2 ** 38 if %v\n", n<<38)
}
