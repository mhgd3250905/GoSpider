package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"io"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"log"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("wrong status code %d",resp.StatusCode)
	}

	//获取编码格式
	e := determineEncoding(resp.Body)

	//切换到UTF-8格式
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

//获取指定输入流的编码格式
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error %v",err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}