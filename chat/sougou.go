package chat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"

	xmlpath "gopkg.in/xmlpath.v2"
)

// find chat by query string
func FindChatUrl(query string) (firstUrl string) {
	// use sogou wexin search
	var url = "http://weixin.sogou.com/weixin?type=1&ie=utf8&query=" + query

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	node, _ := xmlpath.ParseHTML(res.Body)

	// match first chat, get url
	firstUrlPath := xmlpath.MustCompile(`//*[@id="sogou_vr_11002301_box_0"]/div/div[2]/p[1]/a/@href`)

	firstUrl, _ = firstUrlPath.String(node)

	return
}

// article list struct
type ArticleList struct {
	List []Article
}

type Article struct {
	App_msg_ext_info App_msg_ext_info
	Comm_msg_info    Comm_msg_info
}

type App_msg_ext_info struct {
	Author                  string
	Content                 string
	Content_url             string
	Copyright_stat          int
	Cover                   string
	Digest                  string
	Fileid                  int
	Is_multi                int
	Multi_app_msg_item_list []Multi_app_msg_item_list
	Source_url              string
	Subtype                 int
	Title                   string
}

type Multi_app_msg_item_list struct {
	author         string
	content        string
	content_url    string
	copyright_stat int
	cover          string
	digest         int
	fileid         int
	source_url     string
	title          string
}

type Comm_msg_info struct {
	Content  string
	Datetime int
	Fakeid   int
	Id       int
	Status   int
	Type     int
}

// get article list from chat base url
func GetArticleList(url string) (result ArticleList) {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) )Chrome/57.0.2987.133 Safari/537.36")

	res, err := client.Do(req)
	defer res.Body.Close()

	html, _ := ioutil.ReadAll(res.Body)

	// match msgList from js
	re := regexp.MustCompile(`var msgList = (\{.*\})`)
	msgList := re.FindSubmatch(html)

	var articleJsonStr []byte

	// if match result > 1
	if len(msgList) > 1 {
		articleJsonStr = msgList[1]
	}

	// parse result to list struct
	json.Unmarshal(articleJsonStr, &result)

	return
}
