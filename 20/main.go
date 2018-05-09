package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

// http://www.pythonchallenge.com/pc/hex/idiot2.html
// username: butter
// password: fly

// After find the final position, I will just use curl to download it:
// curl http://www.pythonchallenge.com/pc/hex/unreal.jpg -H "Range: bytes=1152983631-" -H "Authorization: Basic YnV0dGVyOmZseQ==" > out

func main() {
	var err error

	// This will lead you to find the password for the target file
	//start := 30203
	//useLeft := false

	// This will lead you to find the target file
	start := 2123456789
	useLeft := true

	for {
		start, err = doreq(start, useLeft)
		if err != nil {
			fmt.Println("**err:", err)
			return
		}
		if start == 0 {
			break
		}
	}
}

func doreq(start int, useLeft bool) (int, error) {
	const url = "http://www.pythonchallenge.com/pc/hex/unreal.jpg"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Authorization", "Basic YnV0dGVyOmZseQ==")

	cli := &http.Client{}

	pattern, err := regexp.Compile(`(\d+)\s*-\s*(\d+)`)
	if err != nil {
		return 0, err
	}

	rangeStr := fmt.Sprintf("bytes=%d-", start)
	req.Header.Set("Range", rangeStr)
	resp, err := cli.Do(req)
	if err != nil {
		return 0, err
	}

	respRange := resp.Header.Get("Content-Range")
	v := pattern.FindStringSubmatch(respRange)
	if len(v) == 0 {
		return 0, nil
	}

	p, err := strconv.ParseInt(v[1], 10, 64)
	if err != nil {
		return 0, err
	}
	if p == 0 {
		return 0, nil
	}
	n, err := strconv.ParseInt(v[2], 10, 64)
	if err != nil {
		return 0, err
	}
	resp.Write(os.Stdout)

	if useLeft {
		return int(p - 1), nil
	} else {
		return int(n + 1), nil
	}
}
