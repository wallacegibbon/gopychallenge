package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// http://www.pythonchallenge.com/pc/def/linkedlist.php

const baseUrl = "http://www.pythonchallenge.com/pc/def/linkedlist.php"

var httpcli = http.Client{Timeout: time.Duration(time.Second * 3)}

func main() {
	//var curLink = "12345"
	var curLink = "8022"
	for {
		d, err := getLinkTry(curLink)
		if err != nil {
			fmt.Println("Failed get link", curLink)
			return
		}
		s := string(d)
		fmt.Println("\t", s)

		curLink, err = matchLink(s)
		if err != nil {
			fmt.Println("Failed parse:", err)
			return
		}
	}
}

func matchLink(content string) (string, error) {
	r, err := regexp.Compile(`and the next nothing is (\d+)`)
	if err != nil {
		return "", err
	}
	v := r.FindStringSubmatch(content)
	if len(v) == 0 {
		return "", errors.New("string mismatch")
	} else {
		return v[1], nil
	}
}

func getLinkTry(linkTail string) ([]byte, error) {
	var cnt = 1
	for {
		r, err := getLink(linkTail)
		if err != nil {
			fmt.Println("**HTTP Get error:", err)
			if cnt < 3 {
				cnt++
			} else {
				return nil, err
			}
		} else {
			return r, nil
		}
	}
}

func getLink(linkTail string) ([]byte, error) {
	url := fmt.Sprintf("%s?nothing=%s", baseUrl, linkTail)
	fmt.Println("Fetching", url)

	resp, err := httpcli.Get(url)
	if err != nil {
		return nil, err
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	return buf, nil
}
