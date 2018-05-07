package main

import (
	"bufio"
	"fmt"
	"os"
)

// http://www.pythonchallenge.com/pc/return/balloons.html
// http://www.pythonchallenge.com/pc/return/brightness.html
// username: huge
// password: file

func main() {
	f, err := os.Open("./deltas")
	if err != nil {
		fmt.Println("Failed opening input file:", err)
		return
	}
	defer f.Close()

	o1, err := os.Create("./o1.txt")
	if err != nil {
		fmt.Println("Failed opening output file 1:", err)
		return
	}
	defer o1.Close()

	o2, err := os.Create("./o2.txt")
	if err != nil {
		fmt.Println("Failed opening output file 2:", err)
		return
	}
	defer o2.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		l := scanner.Bytes()
		o1.Write(l[:53])
		o1.Write([]byte{'\n'})
		o2.Write(l[56:])
		o2.Write([]byte{'\n'})
	}
}

// This level heavily relies on python's difflib. There is third-party go
// library can do this, too. But it's not interesting, I will just drop it here
