package main

import (
	"fmt"
	"time"
)

// http://www.pythonchallenge.com/pc/return/uzi.html
// username: huge
// password: file

func main() {
	//loc, err := time.LoadLocation("Asia/Shanghai")
	//loc, err := time.LoadLocation("Local")
	loc, err := time.LoadLocation("")
	if err != nil {
		fmt.Println("Failed loading timezone")
		return
	}
	for y := 1006; y <= 1996; y += 10 {
		if !isLeapYear(y) {
			continue
		}
		t := time.Date(y, 1, 26, 0, 0, 0, 0, loc)
		if t.Weekday() == time.Monday {
			fmt.Println(t)
		}
	}
}

func isLeapYear(year int) bool {
	if year % 100 != 0 {
		return year % 4 == 0
	} else {
		return year % 400 == 0
	}
}
