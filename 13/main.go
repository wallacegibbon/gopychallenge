// http://www.pythonchallenge.com/pc/return/disproportional.html
// username: huge
// password: file

package main

import (
	"../phonebook"
	"fmt"
)

func main() {
	//buf, err := phonebook.Req("listmethod", "")
	//buf, err := phonebook.Req("methodhelp", "phone")
	buf, err := phonebook.Req("phone", "Bert")
	if err != nil {
		fmt.Println("Failed request:", err)
		return
	}
	fmt.Println(buf)
}
