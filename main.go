package main

import (
	"context"
	"fmt"
	"log"
	"github.com/chromedp/chromedp"
	"time"
)


func main(){
	// 1번 크롤러
	// linkedList  := getLinkList()
	// for _, val := range linkedList{
	// 	getDescription(val)
	// }

	//2번 크롤러
	runCrawler("https://zigzag.kr/catalog/products/124759639", "1", "1")
}

func runCrawler(URL string, lineNum string, stationNm string) {
    
	// settings for crawling
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false))

		// create chrome instance
	contextVar, cancelFunc := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelFunc()

	contextVar, cancelFunc = chromedp.NewContext(contextVar)
	defer cancelFunc()


	var htmlContent string

	err := chromedp.Run(contextVar,
		chromedp.Navigate(URL),
		chromedp.WaitVisible(".pdp__title > h1"),
		chromedp.Click("eewsflh1"),
		chromedp.OuterHTML("html", &htmlContent, chromedp.ByQuery),
	)
	fmt.Println("html", htmlContent)

	if err != nil{
		panic(err)
	}
}

func getDescription(url string){
	fmt.Println(url)
	contextVar, cancelFunc := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancelFunc()
	
	contextVar, cancelFunc = context.WithTimeout(contextVar, 30 * time.Second)//timeout 값을 설정
	defer cancelFunc()
	contextVar, cancelFunc = chromedp.NewContext(contextVar)
	defer cancelFunc()

	var strVar string
	err := chromedp.Run(contextVar,
		chromedp.Navigate("https://www.youtube.com" + url),
		chromedp.Click("#primary div#primary-inner div#below ytd-watch-metadata div#above-the-fold div#bottom-row div#description tp-yt-paper-button#expand-sizer", chromedp.ByID ),
		chromedp.Text("#primary div#primary-inner div#below ytd-watch-metadata div#above-the-fold div#bottom-row div#description", &strVar,chromedp.ByID),
	)

	if err != nil{
		panic(err)
	}

	fmt.Println(strVar)
}

func getLinkList() []string{
	contextVar, cancelFunc := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancelFunc()

	contextVar, cancelFunc = context.WithTimeout(contextVar, 300 * time.Second) //timeout설정
	defer cancelFunc()

	err := chromedp.Run(contextVar,
		chromedp.Navigate("https://www.youtube.com/@paik_jongwon/videos"),
	)
	if err != nil{
		panic(err)
	}

	var oldHeight int
	var newHeight int
	for {
		err = chromedp.Run(contextVar,
			chromedp.Evaluate(`window.scrollTo(0,document.querySelector("body ytd-app div#content").clientHeight); document.querySelector("body ytd-app div#content").clientHeight;`, &newHeight),
			chromedp.Sleep(700*time.Millisecond),
		)
		if err != nil{
			panic(err)
		}
		if(oldHeight == newHeight){
			break
		}
		oldHeight = newHeight
	}

	attr := make([]map[string]string, 0)
	err = chromedp.Run(contextVar, 
		chromedp.AttributesAll("#primary ytd-rich-grid-renderer div#contents ytd-rich-grid-row div#contents ytd-rich-item-renderer #video-title-link", &attr,chromedp.ByQueryAll),
	)

	if err != nil{
		panic(err)
	}

	var linkedList []string
	for _, val := range attr {
		linkedList = append(linkedList,val["href"])
	}
	fmt.Println(len(linkedList))
	return linkedList
}