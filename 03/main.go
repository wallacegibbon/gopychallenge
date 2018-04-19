package main

import (
	"fmt"
	"io/ioutil"
	"unicode"
)

// http://www.pythonchallenge.com/pc/def/equality.html

func main() {
	raw, err := ioutil.ReadFile("./raw.txt")
	if err != nil {
		fmt.Println("Failed reading raw.txt:", err)
	}
	if len(raw) >= 7 {
		printMatched(raw)
	}
}

func printMatched(raw []byte) {
	for i := 0; i < len(raw)-7; {
		d := raw[i:]
		if i == 0 {
			if chkPatternCommon(d, true) {
				fmt.Printf("%c", raw[i+3])
				i += 7
			} else {
				i++
			}
		} else {
			if chkPatternCommon(d, false) {
				fmt.Printf("%c", raw[i+4])
				i += 8
			} else {
				i++
			}
		}
	}
	fmt.Println()
}

func chkPatternCommon(d []byte, isStart bool) bool {
	var a bool
	if isStart {
		a = chkPattern(d)
	} else {
		a = l(d[0]) && chkPattern(d[1:])
	}
	if len(d) > 8 {
		return a && l(d[8])
	} else {
		return a
	}

}

func chkPattern(d []byte) bool {
	return u(d[0]) && u(d[1]) && u(d[2]) && l(d[3]) &&
		u(d[4]) && u(d[5]) && u(d[6])
}

func u(v byte) bool {
	return unicode.IsUpper(rune(v))
}

func l(v byte) bool {
	return unicode.IsLower(rune(v))
}
