package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// http://www.pythonchallenge.com/pc/return/disproportional.html
// username: huge
// password: file

const url = "http://www.pythonchallenge.com/pc/phonebook.php"

const listmethod = `<?xml version="1.0"?>
<methodCall>
	<methodName>system.listMethods</methodName>
	<params></params>
</methodCall>
`

const methodhelp = `<?xml version="1.0"?>
<methodCall>
	<methodName>system.methodHelp</methodName>
	<params>
		<param>
			<value><string>phone</string></value>
		</param>
	</params>
</methodCall>
`

const phonecall = `<?xml version="1.0"?>
<methodCall>
	<methodName>phone</methodName>
	<params>
		<param>
			<value><string>%s</string></value>
		</param>
	</params>
</methodCall>
`

func main() {
	//buf, err := req(listmethod)
	//buf, err := req(methodhelp)
	buf, err := req(fmt.Sprintf(phonecall, "Bert"))
	if err != nil {
		fmt.Println("Failed request:", err)
		return
	}
	fmt.Println(string(buf))
}

func req(reqString string) (string, error) {
	strReader := strings.NewReader(reqString)
	resp, err := http.Post(url, "text/xml", strReader)
	if err != nil {
		return "", err
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	return string(buf), nil
}
