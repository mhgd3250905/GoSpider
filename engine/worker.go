package engine

import (
	"GoSpider/fetcher"
	"log"
)

//Fetcher+Parse
func worker(r Request) (result ParseResult, err error) {
	//调用Fetcher来获取body
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher:error fetching url %s:%v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body,r.Url), nil
}