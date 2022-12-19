package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"
)

var exitMp map[string]bool
var queue chan string

func init() {
	exitMp = make(map[string]bool)
	queue = make(chan string, 10000)
}

func main() {
	go parseUrl()
	queue <- "https://learnku.com/go"

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
}

func parseUrl() {
	for v := range queue {
		res, err := http.Get(v)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		go func() {
			// Load the HTML document
			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			// Find the review items
			doc.Find("a[href]").Each(func(i int, selection *goquery.Selection) {
				href, ok := selection.Attr("href")
				//name := selection.Text()
				//// 去除空格
				//name = strings.Replace(name, " ", "", -1)
				//// 去除换行符
				//name = strings.Replace(name, "\n", "", -1)
				if !ok {
					return
				}

				if exitMp[href] {
					return
				}

				if u, err := url.Parse(href); err == nil && u.Host == "learnku.com" {
					fmt.Println(href)
					exitMp[href] = true
					queue <- href
				}
			})
		}()
		time.Sleep(10 * time.Second)
	}
}
