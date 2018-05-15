// http://www.pythonchallenge.com/pc/ring/yankeedoodle.html
// username: repeat
// password: switch

package main

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("yankeedoodle.csv")
	if err != nil {
		fmt.Println("Failed opening input file")
		return
	}
	defer f.Close()

	o, err := os.Create("out.png")
	if err != nil {
		fmt.Println("Failed opening output file")
		return
	}
	defer o.Close()

	t, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("Failed reading input file")
		return
	}

	raw := strings.Split(string(t), ",")
	numstrs := make([]string, 0)
	for i := 0; i < len(raw); i++ {
		tmp := strings.Trim(raw[i], " \r\n")
		if len(tmp) > 0 {
			numstrs = append(numstrs, tmp)
		}
	}

	nums := collectFloats(numstrs)

	//fmt.Printf("nums len: %d, nums: %v\n", len(nums), nums[:50])

	// 7367 = 53 * 139
	img := image.NewGray(image.Rect(0, 0, 53, 139))
	img.Pix = floatLstToUint8Lst(nums)

	png.Encode(o, img)

	r, err := decodeRawStr(numstrs)
	if err != nil {
		fmt.Println("decodeRawStr error:", err)
		return
	}

	fmt.Println(string(r[:200]))
}

func decodeRawStr(raw []string) ([]byte, error) {
	r := make([]byte, 0)
	for i := 0; i < len(raw); i += 3 {
		if i+3 >= len(raw) {
			break
		}
		v, err := formularExecute(raw[i : i+3])
		if err != nil {
			return nil, err
		}
		r = append(r, v)
	}
	return r, nil
}

func formularExecute(s []string) (byte, error) {
	numstr := fmt.Sprintf("%c%c%c", s[0][5], s[1][5], s[2][6])
	v, err := strconv.ParseInt(numstr, 10, 64)
	if err != nil {
		return 0, err
	}
	return byte(v), nil
}

func floatLstToUint8Lst(flst []float64) []uint8 {
	r := make([]uint8, len(flst))
	for i := 0; i < len(flst); i++ {
		r[i] = uint8(flst[i] * 0xff)
	}
	return r
}

func collectFloats(numstrs []string) []float64 {
	r := make([]float64, 0)
	for _, v := range numstrs {
		f, err := strconv.ParseFloat(v, 64)
		if err == nil {
			r = append(r, f)
		}
	}
	return r
}
