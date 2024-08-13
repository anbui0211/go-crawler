package services

import (
	"fmt"

	"github.com/gocolly/colly"
)

func crawleVibloTrending() {
	type VibloTrendingData struct {
		Title string
		Link  string
	}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting: ", r.URL)
	})

	c.OnHTML(".post-feed .link", func(e *colly.HTMLElement) {
		data := VibloTrendingData{}
		data.Title = e.Text
		data.Link = "https://viblo.asia" + e.Attr("href")

		if data.Title == "" || data.Link == "https://viblo.asia" {
			return
		}
		fmt.Println("title: ", data.Title, " | ", "link: ", data.Link)
	})

	for i := 1; i < 4; i++ {
		fullURL := fmt.Sprintf("https://viblo.asia/trending?page=%d", i)
		c.Visit(fullURL)
	}
}
