package main

import (
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	//"log"

	"github.com/gocolly/colly"
)

type NewInfo struct {
	data  string
	title string
}

func main() {
	Url := "https://acc.cv.ua/"
	scrapPage(Url)

}

func scrapPage(url string) {

	c := colly.NewCollector()
	c.OnHTML("div.tab-latest-articles ul.tab-latest-articles-perdate li", func(e *colly.HTMLElement) {
		selection := e.DOM
		data := selection.Find("div.left_block").Text()
		title := selection.Find("div.right_block").Text()

		fmt.Println(data, "\t", title)

	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)
}
