package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// http://www.pythonchallenge.com/pc/return/disproportional.html
// username: huge
// password: file

const url = "http://www.pythonchallenge.com/pc/phonebook.php"

func main() {
	buf, err := req()
	if err != nil {
		fmt.Println("Failed request:", err)
		return
	}
	fmt.Println(string(buf))
}

func req() ([]byte, error) {
	strReader := strings.NewReader("")
	resp, err := http.Post(url, "text/xml", strReader)
	if err != nil {
		return nil, err
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	return buf, nil
}
