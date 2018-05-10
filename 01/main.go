// http://www.pythonchallenge.com/pc/def/map.html

package main

import (
	"fmt"
)

const raw = `g fmnc wms bgblr rpylqjyrc gr zw fylb. rfyrq ufyr amknsrcpq ypc dmp. bmgle gr gl zw fylb gq glcddgagclr ylb rfyr'q ufw rfgq rcvr gq qm jmle. sqgle qrpgle.kyicrpylq() gq pcamkkclbcb. lmu ynnjw ml rfc spj.`

func main() {
	fmt.Println("Translating the raw string...")
	translateStr(raw)
	fmt.Println("Translating url...")
	translateStr("map")
}

func translateStr(str string) {
	for _, ch := range str {
		fmt.Printf("%c", translate(ch))
	}
	fmt.Println()
}

func translate(ch rune) rune {
	switch {
	case ch >= 'a' && ch <= 'x':
		return ch + 2
	case ch >= 'A' && ch <= 'X':
		return ch + 2
	case ch == 'y':
		return 'a'
	case ch == 'Y':
		return 'A'
	case ch == 'z':
		return 'b'
	case ch == 'Z':
		return 'B'
	default:
		return ch
	}
}
