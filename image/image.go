package image

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Save(url string) (fileName string, err error) {

	basePath := "public/store/"

	dir := basePath + time.Now().Format("200601") + "/"
	return GetImage(dir, url)
}

func GetImage(dir string, url string) (fileName string, err error) {
	//判断目录是否存在
	TouchDir(dir)

	//获取图片文件扩展名
	var name string

	ext := GetImageExt(url)
	name = GetImageName(dir, ext)

	fileName = dir + name

	out, err := os.Create(fileName)

	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	_, err = io.Copy(out, bytes.NewReader(pix))

	return fileName, err
}

//创建目录
func TouchDir(dir string) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
}

func GetImageExt(url string) string {
	// default webp
	return "webp"

	res := strings.Split(url, ".")
	return res[len(res)-1]
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))

	return hex.EncodeToString(h.Sum(nil))
}

func GetImageName(path string, ext string) string {
	str := strconv.FormatInt(time.Now().UnixNano(), 10) + strconv.FormatFloat(rand.Float64(), 'g', 30, 32)
	var name string
	for {
		name = Md5(str) + "." + ext
		_, err := os.Stat(path + name)
		if os.IsNotExist(err) {
			break
		}
	}
	return name
}
