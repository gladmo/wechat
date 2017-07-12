package lengtoo

import (
	"github.com/gladmo/notice"
	vsc "github.com/gladmo/verifyCodeServer"
	"github.com/gladmo/wechat/chat"
	"github.com/gladmo/wechat/image"
	"github.com/gladmo/wechat/models"
	xmlpath "gopkg.in/xmlpath.v2"

	"encoding/json"
	"html"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

const SOURCE string = "lengtoo"

const HOST string = "https://mp.weixin.qq.com"

func Spider() {

	println("start " + SOURCE + " spider")

	// find chat by name(lengtoo)
	chatUrl := chat.FindChatUrl("冷兔")
	println("Get chat home url: " + chatUrl)

	// Have Verification code
	haveVerifyCode := vsc.HaveVerifyCode(chatUrl)
	if haveVerifyCode {
		// start web server, wait to input and submit
		vsc.StartWebServer()
	}

	// get article url list by chat url
	list := chat.GetArticleList(chatUrl)
	print("Get data list, length: ", len(list.List), "\n")

	// make len(list.List) chans run goroutine
	// ch := make(chan int, len(list.List))

	// Each list result
	for _, v := range list.List {

		title := v.App_msg_ext_info.Title

		// if crawled or other break
		no, ok := unique(title)
		if !ok {
			continue
		}

		info := &models.Crawl{
			Source:    SOURCE,
			Create_at: time.Now().Unix(),
			Is_del:    0,
			Unique:    no,
			Url:       v.App_msg_ext_info.Content_url,
		}

		// insert base crawl info
		info.Save()

		// crawl one page(info.Url)
		// go crawlOne(info.Url, info.C_id, ch)
		crawlOne(title, info.Url, info.C_id)
	}

	// read all chan then break
	// for i := 0; i < cap(ch); i++ {
	// 	<-ch
	// }
}

func crawlOne(title, url string, c_id int64) {
	// write to chan
	// ch <- 1

	url = HOST + html.UnescapeString(url)

	println("--------------------------------------------")
	println("start crawl: " + url)
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	print("crawl code: ", res.StatusCode, "\n")

	defer res.Body.Close()

	node, _ := xmlpath.ParseHTML(res.Body)

	SaveHtml(node.Bytes(), title)

	// match js get data, get all tag p
	contentsPath := xmlpath.MustCompile("//*[@id='js_content']/p")

	it := contentsPath.Iter(node)

	// first tag br in all tag p
	firstBr := false

	// match first text repalce ad
	firstText := false

	// text slice
	var jokeText []string

	// img slice
	var jokeImg [][]string

	// next tag p have tag img?
	var inImg = false

	var preBr = true
	for it.Next() {
		// match br, No matter what position in tag p
		brPath := xmlpath.MustCompile("descendant-or-self::br")
		existBr := brPath.Exists(it.Node())

		// match text
		textPath := xmlpath.MustCompile("text()")
		text, _ := textPath.String(it.Node())

		// if match first br start match something
		if !firstBr && existBr {
			firstBr = true
			continue

			// if not first br and exit br, this is end of image or text
		} else if firstBr && existBr {
			preBr = true
			if text == "" {
				continue
			}

			// if start in tag img
		} else if !firstBr {
			continue
		}

		// match image get lazy load image src
		imgPath := xmlpath.MustCompile("img/@data-src")
		existImg := imgPath.Exists(it.Node())
		// if have img tag
		if existImg {
			// if text not empty, it mean not end
			if text == "" {
				// get image url
				if src, ok := imgPath.String(it.Node()); ok {

					// if not start image tag, delete image comment in jokeText and add to jokeImg
					if !inImg {
						var list []string
						list = append(list, jokeText[len(jokeText)-1])

						// delete last jokeText
						jokeText = jokeText[:len(jokeText)-1]

						jokeImg = append(jokeImg, list)
					}

					// parse url, let crawl url to lazy loaded url
					src = getImgRealUrl(src)

					// add image src to jokeImg image list
					jokeImg[len(jokeImg)-1] = append(jokeImg[len(jokeImg)-1], src)
					inImg = true
				}
			} else {
				// if end of joke list, break this cycle
				break
			}

			continue
		}

		// if tag p not have joke
		if text == "" {
			continue
		}

		// if first text trim left
		if !firstText {
			text = strings.TrimLeft(text, "冷兔槽，")
			firstText = true
		}

		// append text joke to jokeText
		if !preBr {
			jokeText[len(jokeText)-1] += "\n" + text
		} else {
			preBr = false
			jokeText = append(jokeText, strings.TrimSpace(text))
		}

		inImg = false
	}

	println("parse ok, start insert")

	print("text joke length: ", len(jokeText), "\n")

	print("image joke length: ", len(jokeImg), "\n")

	ch := make(chan int, len(jokeText))
	// insert text joke list to db
	for _, v := range jokeText {
		textJoke := &models.Text_joke{
			Source:    SOURCE,
			Create_at: time.Now().Unix(),
			C_id:      c_id,
			Content:   v,
		}
		textJoke.Save()

		// notice to phone
		go notice.Send(v, ch)
	}

	for i := 0; i < len(jokeText); i++ {
		<-ch
	}

	// insert image list to db
	for _, v := range jokeImg {

		comment := v[0]

		v = v[1:]

		var imgs []string
		for _, vv := range v {
			// save image to local
			src, _ := image.Save(vv)
			imgs = append(imgs, src)
		}

		localImgs, err := json.Marshal(imgs)
		if err != nil {
			panic(err)
		}

		realImgs, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}

		imageJoke := &models.Img_joke{
			C_id:        c_id,
			Comment:     comment,
			Img_list:    string(localImgs),
			Source_list: string(realImgs),
			Img_count:   int64(len(v)),
			Create_at:   time.Now().Unix(),
		}

		imageJoke.Save()
	}

}

func getImgRealUrl(url string) (src string) {
	endReg := regexp.MustCompile(`0\?wx_fmt=(.*)`)
	src = endReg.ReplaceAllString(url, "640?wx_fmt=$1&tp=webp&wxfrom=5&wx_lazy=1")

	src = strings.Replace(src, "/0", "640?tp=webp&wxfrom=5&wx_lazy=1", 1)

	return
}

func SaveHtml(html []byte, title string) {
	basePath := "public/html/" + time.Now().Format("20060102") + "/"

	os.MkdirAll(basePath, 755)

	fileName := basePath + title + ".html"

	// str, _ := ioutil.ReadAll(html)

	file, _ := os.Create(fileName)

	_, err := io.WriteString(file, string(html))
	if err != nil {
		panic(err)
	}
}

func unique(title string) (no string, ok bool) {

	// match【冷兔•槽】每日一冷NO.xxxx
	re := regexp.MustCompile(`[每日一冷NO\s]*\.(\d+)`)

	res := re.FindStringSubmatch(title)

	if len(res) > 1 {
		no = res[1]
		if models.FindByUnique(no, SOURCE) {
			return no, true
		}
	}

	return "", false
}
