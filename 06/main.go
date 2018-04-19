package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
)

// http://www.pythonchallenge.com/pc/def/channel.html

func main() {
	r, err := zip.OpenReader("./channel.zip")
	if err != nil {
		fmt.Println("Failed openning channel.zip")
		return
	}
	defer r.Close()

	//fmt.Println("Zip Comment:", r.Comment)

	//showReadableInfo(r.File)
	collectInfo(r.File)
}

func collectInfo(files []*zip.File) {
	var link = "90052"
	var cmt string
	var err error

	for {
		cmt, link, err = getCmtAndRetNext(link, files)
		if err != nil {
			fmt.Println("Failed getting link", link, ":", err)
			return
		}
		fmt.Print(cmt)
	}
}

func getCmtAndRetNext(link string, files []*zip.File) (string, string, error) {
	for _, f := range files {
		if f.Name == link+".txt" {
			t, err := readContent(f)
			if err != nil {
				return "", "", err
			}
			n, err := matchLink(t)
			if err != nil {
				return "", "", err
			}
			return f.Comment, n, nil
		}
	}
	return "", "", errors.New("link not found:" + link)
}

func showReadableInfo(files []*zip.File) {
	for _, f := range files {
		fmt.Println("File comment:", f.Name, ",", f.Comment)
		t, err := readContent(f)
		if err != nil {
			fmt.Println("Failed reading", f.Name)
			return
		}
		fmt.Println("\t", t)
	}
}

func readContent(file *zip.File) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}

	t, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(t), nil
}

func matchLink(content string) (string, error) {
	r, err := regexp.Compile(`Next nothing is (\d+)`)
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
