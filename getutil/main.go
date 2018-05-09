package getutil

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

type MyJar struct {
	cookies []*http.Cookie
}

func (jar *MyJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.cookies = cookies
}

func (jar *MyJar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies
}

var httpcli = &http.Client{Jar: new(MyJar), Timeout: time.Duration(time.Second * 3)}

func Get(baseUrl, next string) (content, cookie string, err error) {
	var cnt = 1
	for {
		content, cookie, err = get(baseUrl, next)
		if err != nil {
			fmt.Println("**HTTP Get error:", err)
			if cnt < 3 {
				cnt++
			} else {
				return
			}
		} else {
			return
		}
	}
}

func get(baseUrl, next string) (content, cookie string, err error) {
	//url := fmt.Sprintf("%s?nothing=%s", baseUrl, next)
	url := fmt.Sprintf("%s%s", baseUrl, next)
	fmt.Println("Fetching", url)

	resp, err := httpcli.Get(url)
	if err != nil {
		return
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	content = string(buf)

	c := resp.Cookies()
	//fmt.Println("cookies:", c)
	if len(c) > 0 {
		cookie = c[0].Value
	}
	return
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
