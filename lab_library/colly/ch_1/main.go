package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
	"time"
)

func main() {
	scrap("https://tieba.baidu.com")
}

func scrap(url string) {
	urls := make(map[string]struct{})
	threatUrl := make(map[string]struct{})

	u, err := url.Parse(url)

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("tieba.baidu.com"),
		colly.MaxDepth(3),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  `tieba\.baidu\.com`,
		RandomDelay: 5 * time.Second,
		Parallelism: 8,
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Visit link found on page
		// Only those links are visited which are matched by  any of the URLFilter regexps
		_, ok := urls[link]
		if !ok {
			// Print link
			name := strings.Replace(e.Text, " ", "", -1)
			//// 去除换行符
			name = strings.Replace(name, "\n", "", -1)
			fmt.Printf("Link found: %q -> %s\n", name, link)
			urls[link] = struct{}{}
			if u, err := url.Parse(link); err == nil && u.Host == "tieba.baidu.com" {
				c.Visit(e.Request.AbsoluteURL(link))
			}
		}

		//c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("OnError", err.Error())
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://tieba.baidu.com/f?kw=golang&ie=utf-8")

	c.Wait()
}
