package mozzartbet

import (
	"betulator/pkg/model"
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

func GetFootballEvents() ([]model.Event, error) {

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

	timeoutContext, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	var body string

	err = chromedp.Run(timeoutContext,
		chromedp.Navigate(`https://www.mozzartbet.com/en#/betting/?sid=1`),
		// chromedp.ScrollIntoView("div.paginator"),
		// chromedp.Click(`#example-After`, chromedp.NodeVisible),
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

	events := []model.Event{}

	doc.Find("selection.promo-matches, div.competition").Each(func(i int, query *goquery.Selection) {

		query.Find("article").Each(func(i int, query *goquery.Selection) {

			part1 := query.Find("div.part1")

			// skip bonus odds
			if part1.Find("span.bonus-group").Size() == 0 {

				outcomes := make([]string, 3)
				startTime := time.Now()

				part1.Find("a.pairs").Find("span").Each(func(i int, query *goquery.Selection) {

					if i == 0 {

						//startTime = nil

						outcomes[1] = "draw"
					} else if i == 1 {
						outcomes[0] = query.Text()
					} else if i == 2 {
						outcomes[2] = query.Text()
					}
				})

				part2 := query.Find("div.part2")

				odds := []decimal.Decimal{}

				part2.Find("span.odd-font").EachWithBreak(func(i int, query *goquery.Selection) bool {

					dec, err := decimal.NewFromString(query.Text())
					if err != nil {
						fmt.Println(err)
					}

					odds = append(odds, dec)

					return true
				})

				events = append(events, model.Event{
					Outcome:   outcomes,
					Odds:      odds,
					StartTime: startTime,
				})
			}
		})
	})

	return events, nil
}
