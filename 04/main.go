// http://www.pythonchallenge.com/pc/def/linkedlist.php

package main

import (
	"../getutil"
	"fmt"
)

const baseUrl = "http://www.pythonchallenge.com/pc/def/linkedlist.php?nothing="

func main() {
	//var next = "12345"
	var next = "8022"
	for {
		d, _, err := getutil.Get(baseUrl, next)
		if err != nil {
			fmt.Println("Failed get link", next)
			return
		}
		fmt.Println("\t", d)

		next, err = getutil.MatchNext(d, `the next nothing is (\d+)`)
		if err != nil {
			fmt.Println("Failed parse:", err)
			return
		}
	}
}
