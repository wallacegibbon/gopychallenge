package phonebook

import (
	"io/ioutil"
	"net/http"
	"strings"
	"fmt"
)

const url = "http://www.pythonchallenge.com/pc/phonebook.php"

var methods = map[string]string{
	"listmethod": `<?xml version="1.0"?>
<methodCall>
	<methodName>system.listMethods</methodName>
	<params></params>
</methodCall>%s
`,
	"methodhelp": `<?xml version="1.0"?>
<methodCall>
	<methodName>system.methodHelp</methodName>
	<params>
		<param>
			<value><string>%s</string></value>
		</param>
	</params>
</methodCall>
`,
	"phone": `<?xml version="1.0"?>
<methodCall>
	<methodName>phone</methodName>
	<params>
		<param>
			<value><string>%s</string></value>
		</param>
	</params>
</methodCall>
`}

func Req(method, value string) (string, error) {
	queryDoc := fmt.Sprintf(methods[method], value)
	//fmt.Println(queryDoc)
	strReader := strings.NewReader(queryDoc)
	resp, err := http.Post(url, "text/xml", strReader)
	if err != nil {
		return "", err
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	return string(buf), nil
}
