package main

import (
	"fmt"
	"net/http"
	"os"
)

// http://www.pythonchallenge.com/pc/hex/idiot2.html
// username: butter
// password: fly

const url = "http://www.pythonchallenge.com/pc/hex/unreal.jpg"

func main() {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed creating request")
		return
	}

	cli := &http.Client{}
}
