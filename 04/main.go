package main

import (
	"fmt"
	"../getutil"
)

// http://www.pythonchallenge.com/pc/def/linkedlist.php

const baseUrl = "http://www.pythonchallenge.com/pc/def/linkedlist.php"

func main() {
	//var next = "12345"
	var next = "8022"
	for {
		d, err := getutil.Get(baseUrl, next)
		if err != nil {
			fmt.Println("Failed get link", next)
			return
		}
		s := string(d)
		fmt.Println("\t", s)

		next, err = getutil.MatchNext(s, `the next nothing is (\d+)`)
		if err != nil {
			fmt.Println("Failed parse:", err)
			return
		}
	}
}

