package getutil

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

var httpcli = http.Client{Timeout: time.Duration(time.Second * 3)}

func Get(baseUrl, next string) ([]byte, error) {
	var cnt = 1
	for {
		r, err := get(baseUrl, next)
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

func get(baseUrl, next string) ([]byte, error) {
	url := fmt.Sprintf("%s?nothing=%s", baseUrl, next)
	fmt.Println("Fetching", url)

	resp, err := httpcli.Get(url)
	if err != nil {
		return nil, err
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	return buf, nil
}

func MatchNext(content, regstr string) (string, error) {
	r, err := regexp.Compile(regstr)
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
