package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.cnblogs.com/pu369/p/10315988.html`),
		chromedp.OuterHTML(`document.querySelector("html")`, &res, chromedp.ByJSPath),
	)

	if err != nil {
		log.Fatal(err)
	}

	//log.Println(strings.TrimSpace(res))

	ioutil.WriteFile("./1.html", []byte(res), 0666)

	time.Sleep(2 * time.Second)
}

// https://blog.csdn.net/neosmith/article/details/106021625
