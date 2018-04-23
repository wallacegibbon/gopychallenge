package main

import (
	"compress/bzip2"
	"fmt"
	"io/ioutil"
	"strings"
)

// http://www.pythonchallenge.com/pc/def/integrity.html

var name = "BZh91AY&SYA\xaf\x82\r\x00\x00\x01\x01\x80\x02\xc0\x02\x00 \x00!\x9ah3M\x07<]\xc9\x14\xe1BA\x06\xbe\x084"

var pass = "BZh91AY&SY\x94$|\x0e\x00\x00\x00\x81\x00\x03$ \x00!\x9ah3M\x13<]\xc9\x14\xe1BBP\x91\xf08"

func main() {
	n, err := decode(name)
	if err != nil {
		fmt.Println("Failed decode name")
	}
	p, err := decode(pass)
	if err != nil {
		fmt.Println("Failed decode pass")
	}

	fmt.Println("username:", n)
	fmt.Println("password:", p)
}

func decode(s string) (string, error) {
	r := bzip2.NewReader(strings.NewReader(s))
	t, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(t), nil
}
