package main

import (
	"fmt"
)

// http://www.pythonchallenge.com/pc/return/bull.html
// username: huge
// password: file

// a = [1, 11, 21, 1211, 111221,

func main() {
	cur := "1"
	for i := 0; i <= 30; i++ {
		//fmt.Printf("a[%d] = %s\n", i, cur)
		fmt.Printf("len(a[%d]) = %d\n", i, len(cur))
		cur = getNext(cur)

	}
}

func getNext(origStr string) string {
	curCh := int(origStr[0])
	prevCh := curCh
	result := ""

	cnt := 0
	origStr += "X"

	for _, ch := range origStr {
		prevCh = curCh
		curCh = int(ch)
		if prevCh != curCh {
			result += fmt.Sprintf("%d%c", cnt, prevCh)
			cnt = 1
		} else {
			cnt++
		}
	}

	return result
}
