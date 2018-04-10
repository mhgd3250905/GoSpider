package persist

import (
	"os"

	"io"
	"net/http"
	"fmt"
	"strings"
)

func DownloadFile(url string, name string,id string) {

	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	header := map[string]string{
		"User-Agent": `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36`,
		"Referer":`http://www.mzitu.com`,
	}

	//增加header选项
	for key, value := range header {
		reqest.Header.Add(key, value)
	}

	//处理返回结果
	res, err := client.Do(reqest)
	if err != nil {
		return
	}

	name=strings.Replace(name,":","_",-1);

	dirPath:=fmt.Sprintf("downloads/%s",name)
	if exist, _ := PathExists(dirPath);!exist {
		os.Mkdir(dirPath,os.ModePerm)
	}

	f, err := os.Create(fmt.Sprintf("downloads/%s/%s.jpg", name,id))
	if err != nil {
		return
	}
	io.Copy(f, res.Body)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
