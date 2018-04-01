package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"log"
)

func Fetch(url string,header map[string]string) ([]byte, error) {

	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request featch failed %v",err)
	}
	//增加header选项
	for key,value:=range header{
		reqest.Header.Add(key, value)
	}

	//处理返回结果
	resp, err := client.Do(reqest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code %d", resp.StatusCode)
	}

	//获取编码格式
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)

	//切换到UTF-8格式
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	//content,err:=ioutil.ReadAll(utf8Reader)
	//fmt.Printf("%s",content)


	return ioutil.ReadAll(utf8Reader)
}

//获取指定输入流的编码格式
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
