package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"time"
)

type ScraperResponse struct {
	Hash           string `json:"hash"`
	Date           string `json:"date"`
	Assets         string `json:"assets"`
	AddressSpender string `json:"address_spender"`
	Allowance      string `json:"allowance"`
}

func main() {
	c := colly.NewCollector()
	c.UserAgent = "Go program"
	c.SetRequestTimeout(time.Second * 120)
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting ", request.URL)
		for k, v := range *request.Headers {
			fmt.Sprintf("Key: %v, Value: %v", k, v)
		}
		fmt.Println("Method ", request.Method)
	})

	//c.OnHTML("tr td a", func(element *colly.HTMLElement) {
	//	//fmt.Println(element.DOM.Nodes[0])
	//	if element.DOM.HasClass("hash-tag text-truncate") {
	//		fmt.Println(element.Text)
	//	}
	//})
	//
	//c.OnHTML("tr td", func(e *colly.HTMLElement) {
	//	if e.DOM.HasClass("showDate") {
	//		fmt.Println(e.Text)
	//	}
	//})
	res := make([]*ScraperResponse, 0)

	c.OnHTML("table#mytable > tbody", func(h *colly.HTMLElement) {
		response := &ScraperResponse{}
		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			//fmt.Println(el.ChildText("td:nth-child(1)"))
			//fmt.Println()
			//fmt.Println(el.ChildText("td:nth-child(2)"))
			//fmt.Println()
			//fmt.Println(el.ChildText("td:nth-child(3)"))
			//fmt.Println()
			//fmt.Println(el.ChildText("td:nth-child(4)"))
			//fmt.Println()
			//fmt.Println(el.ChildText("td:nth-child(5)"))
			//fmt.Println()
			//fmt.Println(el.ChildText("td:nth-child(6)"))
			//fmt.Println()
			response.Hash = el.ChildText("td:nth-child(1)")
			response.Date = el.ChildText("td:nth-child(2)")
			response.Assets = el.ChildText("td:nth-child(3)")
			response.AddressSpender = el.ChildText("td:nth-child(4)")
			response.Allowance = el.ChildText("td:nth-child(5)")
			res = append(res, response)
		})
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Response ", response.Request.URL)
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Error ", err)
	})

	c.Visit("https://etherscan.io/tokenapprovalchecker?type=0&search=0x893e2765c5c63e60f297c2ad3420f8c23c7b8ce5")
	fmt.Println(len(res))

	for _, r := range res {
		fmt.Println(r)
	}
}
