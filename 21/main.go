// The package.pack and readme.txt file come from the previous level's result.
// bzip2 format starts with "BZh", and zlib starts with 0x78 0x9c

package main

import (
	"bytes"
	"compress/bzip2"
	"compress/zlib"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("package.pack")
	if err != nil {
		fmt.Println("failed opening input file")
		return
	}
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("failed reading input file")
		return
	}
	for {
		if !itIsZlib(buf) && !itIsBzip2(buf) {
			tmp := reverseBytes(buf)
			if !itIsZlib(tmp) && !itIsBzip2(tmp) {
				fmt.Println("decompress:", string(buf))
				return
			}
			buf = tmp
			fmt.Println()
		}
		buf, err = decBytes(buf)
		if err != nil {
			fmt.Println("**Err: %v\n", err)
			return
		}
		//inspect(buf)
	}
}

func decBytes(raw []byte) ([]byte, error) {
	if itIsBzip2(raw) {
		fmt.Print("@")
		return decBzip2(raw)
	}
	if itIsZlib(raw) {
		fmt.Print(" ")
		return decZlib(raw)
	}
	l := len(raw)
	ei := fmt.Sprintf("unknown type: %X...%X", raw[:8], raw[l-8:])
	return nil, errors.New(ei)
}

func itIsBzip2(raw []byte) bool {
	return bytes.Equal(raw[:3], []byte{'B', 'Z', 'h'})
}

func itIsZlib(raw []byte) bool {
	return bytes.Equal(raw[:2], []byte{0x78, 0x9c})
}

func decZlib(raw []byte) ([]byte, error) {
	reader, err := zlib.NewReader(bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func decBzip2(raw []byte) ([]byte, error) {
	reader := bzip2.NewReader(bytes.NewReader(raw))
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func inspect(raw []byte) {
	fmt.Printf("Head:%X, Tail:%X\n", raw[:3], raw[len(raw)-3:])
}

func reverseBytes(raw []byte) []byte {
	result := make([]byte, len(raw))
	cnt := 0
	for i := len(raw) - 1; i >= 0; i-- {
		result[cnt] = raw[i]
		cnt++
	}
	return result
}
