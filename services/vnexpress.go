package services

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

type VnExpressFootballData struct {
	Title       string
	Description string
	PageNumber  int
}

func CrawlVnExpressFootball() {
	var dataList []VnExpressFootballData
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visitting:: ", r.URL)
	})

	c.OnHTML(".list-news-subfolder .item-news:not(.close_not_qc)", func(e *colly.HTMLElement) {
		data := VnExpressFootballData{}

		// Title
		data.Title = e.ChildText("h2 a")

		// Description
		data.Description = e.ChildText("p a")

		// Page number
		pageUrl := e.Request.URL.String()
		data.PageNumber, _ = strconv.Atoi(pageUrl[len(pageUrl)-1:])
		if data.PageNumber == 0 {
			data.PageNumber = 1
		}

		dataList = append(dataList, data)
	})

	numPageCrawl := 1
	for i := 0; i < numPageCrawl; i++ {
		fullUrl := fmt.Sprintf("https://vnexpress.net/bong-da-p%d", i)
		c.Visit(fullUrl)
	}

	fileName := "./files/json/vn_express_football.json"
	if err := WriteToJSONFile(fileName, dataList); err != nil {
		log.Println("Crawl error: ", err.Error())
	}
}
