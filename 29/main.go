// http://www.pythonchallenge.com/pc/ring/guido.html
// username: repeat
// password: switch

package main

import (
	"bytes"
	"compress/bzip2"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	tgt := "http://www.pythonchallenge.com/pc/ring/guido.html"
	req, err := http.NewRequest("GET", tgt, nil)
	if err != nil {
		fmt.Println("Failed constructing request")
		return
	}
	req.Header.Set("Authorization", "Basic cmVwZWF0OnN3aXRjaA==")
	cli := http.Client{}

	resp, err := cli.Do(req)
	if err != nil {
		fmt.Println("Failed sending request")
		return
	}
	defer resp.Body.Close()

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed reading body")
		return
	}

	lst := bytes.Split(d, []byte{'\n'})[12:]
	//fmt.Println(lst)

	//Using go will left one more blank line than using python
	blks := make([]byte, len(lst)-1)
	for i := 0; i < len(lst)-1; i++ {
		blks[i] = byte(len(lst[i]))
	}

	r := bzip2.NewReader(bytes.NewReader(blks))
	v, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println("Failed decoding bzip2:", err)
		return
	}

	fmt.Println(string(v))
}
