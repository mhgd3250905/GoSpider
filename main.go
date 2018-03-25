package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/net/html/charset"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("http err statusCode ", resp.StatusCode)
		return
	}

	//获取编码格式
	e := determineEncoding(resp.Body)

	//切换到UTF-8格式
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", all)
	//打印所有的城市信息
	printCityList(all)
}

//获取指定输入流的编码格式
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic("determine encoding failed ")
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

//正则表达式来匹配所有的城市列表
func printCityList(contents []byte)  {
	re:=regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matchs:=re.FindAllSubmatch(contents,-1)
	for _,m:=range matchs{
		fmt.Printf("City: %s, URL: %s\n",m[2],m[1])
	}
	fmt.Printf("Matchs found : %d\n",len(matchs))
}
