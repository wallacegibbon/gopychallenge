// http://www.pythonchallenge.com/pc/hex/decent.html
// username: butter
// password: fly

// the broken.zip comes from level 24
// should send email to leopold.moz@pythonchallenge.com, and get the md5 of
// the target file
// but where is the email address come from

package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	expect, err := hex.DecodeString("bbb8b499a0eef99b52c7f13f4e78c24b")
	if err != nil {
		fmt.Println("Failed decoding md5 string")
		return
	}
	f, err := os.Open("mybroken.zip")
	if err != nil {
		fmt.Println("Failed opening input file:")
		return
	}
	defer f.Close()

	d, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("Failed reading input file")
		return
	}

	o, err := os.Create("out.zip")
	if err != nil {
		fmt.Println("Failed opening output file")
		return
	}
	defer o.Close()

outter:
	for n := 0; n < len(d); n++ {
		for x := 0; x <= 0xff; x++ {
			tmp := d[n]
			d[n] = byte(x)
			tgtstr := md5.Sum(d)
			if bytes.Equal(tgtstr[:], expect) {
				fmt.Println("find!")
				o.Write(d)
				break outter
			} else {
				d[n] = tmp
			}
		}
	}
}
