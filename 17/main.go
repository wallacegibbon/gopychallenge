package main

import (
	"../getutil"
	"compress/bzip2"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

// http://www.pythonchallenge.com/pc/return/romance.html
// username: huge
// password: file

const baseUrl = "http://www.pythonchallenge.com/pc/def/linkedlist.php?busynothing="

func main() {
	var next = "12345"
	var r string
	for {
		d, c, err := getutil.Get(baseUrl, next)
		if err != nil {
			fmt.Println("Failed get link", next)
			break
		}
		fmt.Println("\t", d)
		r += c

		next, err = getutil.MatchNext(d,
			`the next busynothing is (\d+)`)

		if err != nil {
			fmt.Println("Failed parse:", err)
			break
		}
	}
	u, err := url.QueryUnescape(r)
	if err != nil {
		fmt.Println("unescape error:", err)
		return
	}
	//fmt.Println(u)
	t, err := decode(u)
	if err != nil {
		fmt.Println("bzip2 decode error:", err)
		return
	}

	fmt.Println(t)
}

func decode(s string) (string, error) {
	r := bzip2.NewReader(strings.NewReader(s))
	t, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(t), nil
}
