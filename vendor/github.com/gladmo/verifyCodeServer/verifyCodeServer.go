package verifyCodeServer

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sourcegraph/go-selenium"
)

var (
	Port         string = "2222"
	BrowserName  string = "firefox"
	RemoteUrl    string = "http://localhost:8910"
	ImgId        string = "#verify_img"
	ImgUrlSrc    string = "src"
	Log                 = log.New(os.Stderr, "[verifyCodeServer] ", log.Ltime|log.Lmicroseconds)
	templatePath string = "public/tpls/verificationCode.html"
	webDriver    selenium.WebDriver
	elem         selenium.WebElement
)

func HaveVerifyCode(url string) bool {

	var err error

	if Log == nil {
		selenium.Log = nil
	}

	caps := selenium.Capabilities(map[string]interface{}{"browserName": BrowserName})
	if webDriver, err = selenium.NewRemote(caps, RemoteUrl); err != nil {
		panic(err)
	}

	err = webDriver.Get(url)
	if err != nil {
		panic(err)
	}

	elem, err = webDriver.FindElement(selenium.ByCSSSelector, ImgId)
	if err != nil {
		if Log != nil {
			Log.Printf("Failed to get elem: %s\n", err)
		}
		return false
	}

	text, err := elem.GetAttribute(ImgUrlSrc)
	if err != nil {
		if Log != nil {
			Log.Printf("Failed to get text of element: %s\n", err)
		}
		return false
	}

	if text != "" {
		return true
	} else {
		return false
	}

}

func StartWebServer() {
	// close webDriver
	defer webDriver.Quit()

	// start verify code server, and wait for a verification code
	server := &WebServer{
		verifyCodeUrl: getImgBase64(),
	}

	server.webServer()
}

func getImgBase64() string {

	getImgBase64 := `
		if(typeof getBase64Image != "function"){
			function getBase64Image(img) {
				var canvas = document.createElement("canvas");
				canvas.width = img.width;
				canvas.height = img.height;

				var ctx = canvas.getContext("2d");
				ctx.drawImage(img, 0, 0, img.width, img.height);

				var dataURL = canvas.toDataURL("image/bmp");
				return dataURL
			}
		}

		return getBase64Image(document.getElementById('verify_img'));
	`

	data, err := webDriver.ExecuteScript(getImgBase64, nil)
	if err != nil {
		panic(err)
	}

	return data.(string)
}

type WebServer struct {
	verifyCodeUrl string
	ch            chan int
}

func (s *WebServer) webServer() {
	s.ch = make(chan int, 1)

	server := &http.Server{Addr: "localhost:" + Port, Handler: s}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			if Log != nil {
				Log.Printf("Httpserver: ListenAndServe() error: %s", err)
			}
		}
	}()

	<-s.ch

	// now close the server gracefully ("shutdown")
	// timeout could be given instead of nil as a https://golang.org/pkg/context/
	if err := server.Shutdown(nil); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}

	if Log != nil {
		Log.Println("End of web server")
	}
}

func (s *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method, r.RequestURI)

	switch r.Method {
	case "GET":
		switch r.RequestURI {
		case "/verify_change":
			// click tag img, change verify code
			elem.Click()
			time.Sleep(time.Second * 1)

			s.verifyCodeUrl = getImgBase64()

			fmt.Fprintf(w, s.verifyCodeUrl)

		default:
			data := map[string]interface{}{
				"img": s.verifyCodeUrl,
			}

			t, _ := template.ParseFiles(templatePath) //将一个文件读作模板
			t.Execute(w, data)
		}
	case "POST":
		// post verify code then submit form
		body, _ := ioutil.ReadAll(r.Body)

		result, err := webDriver.ExecuteScript("document.getElementById('input').value='"+string(body)+"';document.getElementById('bt').click();if(document.getElementById('verify_img') == null) return 1;else return 0;", nil)
		if err != nil {
			panic(err)
		}

		// not closed chan well fail
		r.Body.Close()
		fmt.Println("result:", result)

		// if verify code error
		if result.(float64) == 1 {
			time.Sleep(time.Second * 1)

			// get changed img
			s.verifyCodeUrl = getImgBase64()

			fmt.Fprintf(w, s.verifyCodeUrl)
		} else {
			fmt.Fprintf(w, "0")
			s.ch <- 1
		}
	}
}
