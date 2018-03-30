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

func Fetch(url string) ([]byte, error) {

	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request featch failed %v",err)
	}
	//增加header选项
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36")
	reqest.Header.Add("Cookie", "huxiu_analyzer_wcy_id=353onua5rf9pynpg3uj; gr_user_id=33f075cb-73c1-4b0b-a9cb-bac944b28fdc; b6a739d69e7ea5b6_gr_last_sent_cs1=0; _ga=GA1.2.994315101.1522250468; screen=%7B%22w%22%3A1920%2C%22h%22%3A1080%2C%22d%22%3A1%7D; aliyungf_tc=AQAAANez0HTZ3gsA3UKttKiV9xIVsNZi; _alicdn_sec__=5abe827d29d82040be4c3147827972db5f229485; b6a739d69e7ea5b6_gr_session_id=c60b48f1-b30d-47ea-a7d0-f1c9859fdd19; b6a739d69e7ea5b6_gr_last_sent_sid_with_cs1=c60b48f1-b30d-47ea-a7d0-f1c9859fdd19; _gid=GA1.2.427699895.1522434691; Hm_lvt_324368ef52596457d064ca5db8c6618e=1522250468,1522336543,1522434691; b6a739d69e7ea5b6_gr_cs1=0; _gat=1; Hm_lpvt_324368ef52596457d064ca5db8c6618e=1522435123; SERVERID=03a07aad3597ca2bb83bc3f5ca3decf7|1522435120|1522434687")

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
