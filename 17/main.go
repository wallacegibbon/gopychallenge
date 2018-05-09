package main

import (
	"../getutil"
	"../phonebook"
	"compress/bzip2"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// http://www.pythonchallenge.com/pc/return/romance.html
// username: huge
// password: file

const baseUrl = "http://www.pythonchallenge.com/pc/def/linkedlist.php?busynothing="

func main() {
	var err error
	//err = phase1()
	//err = phase2()
	err = phase3()
	if err != nil {
		fmt.Println(err)
	}
}

func phase1() error {
	var next = "12345"
	var r string
	for {
		d, c, err := getutil.Get(baseUrl, next)
		if err != nil {
			return err
		}
		fmt.Println("\t", d)
		r += c

		next, err = getutil.MatchNext(d,
			`the next busynothing is (\d+)`)

		if err != nil {
			return err
		}
	}
	u, err := url.QueryUnescape(r)
	if err != nil {
		return err
	}
	//fmt.Println(u)
	t, err := decode(u)
	if err != nil {
		return err
	}

	fmt.Println(t)
	return nil
}

func phase2() error {
	buf, err := phonebook.Req("phone", "Leopold")
	if err != nil {
		return err
	}
	fmt.Println(buf)
	return nil
}

func phase3() error {
	req, err := http.NewRequest("GET", "http://www.pythonchallenge.com/pc/stuff/violin.php", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Cookie", "info=the flowers are on their way")
	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	resp.Write(os.Stdout)
	return nil
}

func decode(s string) (string, error) {
	r := bzip2.NewReader(strings.NewReader(s))
	t, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(t), nil
}
