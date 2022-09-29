package mozzartbet

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/shopspring/decimal"
)

type Event struct {
	Outcome   []string
	Odds      []decimal.Decimal
	StartTime time.Time
}

func GetFootballEvents() ([]Event, error) {

	// response, err := http.Get("https://www.mozzartbet.com/en#/betting/?sid=1")
	// if err != nil {
	// 	return nil, err
	// }

	// ctx, cancel := chromedp.NewContext(context.Background())
	// defer cancel()

	// // run task list
	// var res string
	// err := chromedp.Run(ctx,
	// 	chromedp.Navigate(`https://pkg.go.dev/time`),
	// 	chromedp.Text(`.Documentation-overview`, &res, chromedp.NodeVisible),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(strings.TrimSpace(res))

	scrap()

	return nil, nil
}

func scrap() {

	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserDataDir(dir),
		chromedp.Flag("headless", false),
		chromedp.Flag("restore-on-startup", true),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, _ := chromedp.NewContext(allocCtx)

	start := time.Now()

	var res string
	err = chromedp.Run(ctx,
		// emulation.SetUserAgentOverride("WebScraper 1.0"),
		chromedp.Navigate(`https://www.mozzartbet.com/en#/betting/?sid=1`),

		// wait for footer element is visible (ie, page is loaded)

		// chromedp.WaitVisible(`competition`, chromedp.),
		// chromedp.Text(``, &res, chromedp.NodeVisible, chromedp.ByQuery),
		// chromedp.Text(`competition`, &res, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Sleep(time.Second*10),
		// chromedp.ScrollIntoView(`.widget-footer-v2`),
		// chromedp.WaitVisible(`competition`),
		chromedp.Text(`body`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("h1 contains: '%s'\n", res)
	fmt.Printf("\nTook: %f secs\n", time.Since(start).Seconds())

}
