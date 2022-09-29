package mozzartbet

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/shopspring/decimal"
)

type Event struct {
	Outcome   []string
	Odds      []decimal.Decimal
	StartTime time.Time
}

func GetFootballEvents() ([]Event, error) {

	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		return nil, err
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

	// var res string
	var body string

	err = chromedp.Run(ctx,
		chromedp.Navigate(`https://www.mozzartbet.com/en#/betting/?sid=1`),
		chromedp.OuterHTML(".sportsoffer", &body, chromedp.ByQuery),
	)
	if err != nil {
		return nil, err
	}

	err = chromedp.Cancel(ctx)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	events := []Event{}

	doc.Find("selection.promo-matches, div.competition").Find("div.part1").Each(func(i int, query *goquery.Selection) {

		outcomes := make([]string, 3)
		startTime := time.Now()

		query.Find("a.pairs").Find("span").Each(func(i int, query *goquery.Selection) {

			if i == 0 {

				//startTime = nil

				outcomes[1] = "draw"
			} else if i == 1 {
				outcomes[0] = query.Text()
			} else if i == 2 {
				outcomes[2] = query.Text()
			}
		})

		events = append(events, Event{Outcome: outcomes, StartTime: startTime})
	})

	doc.Find("selection.promo-matches, div.competition").Find("div.part2").Each(func(i int, query *goquery.Selection) {

		odds := []decimal.Decimal{}

		query.Find("span.odd-font").EachWithBreak(func(i int, query *goquery.Selection) bool {

			dec, err := decimal.NewFromString(query.Text())
			if err != nil {
				fmt.Println(err)
			}

			odds = append(odds, dec)

			return true
		})

		events[i].Odds = odds
	})

	fmt.Printf("\nTook: %f secs\n", time.Since(start).Seconds())

	return events, nil
}
