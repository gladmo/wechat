package notice

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

var SendType string = "ding-talk"

func Send(str string, ch chan int) {
	switch SendType {
	case "ding-talk":
		dingTalk(str)
	default:
		dingTalk(str)
	}

	ch <- 0
}

func dingTalk(str string) string {
	client := http.Client{}

	url := "https://oapi.dingtalk.com/robot/send?access_token=42b39c41c8d1692e05fa6ad5104464cecdc9519ce3c03570c4b26051d6f953b7"

	data := `
	{
	    "msgtype": "text",
	    "text": {
	        "content": "%s"
	    }
	}
	`

	data = fmt.Sprintf(data, str)

	body := bytes.NewBuffer([]byte(data))

	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	result, _ := ioutil.ReadAll(res.Body)
	return string(result)
}
