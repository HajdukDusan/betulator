package main

import (
	"betulator/internal/math"
	"betulator/internal/util"
	"betulator/pkg/model"
	"betulator/pkg/scrapers/mozzartbet"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shopspring/decimal"
)

type Event struct {
	Outcome   []string
	Odds      []decimal.Decimal
	House     []string
	StartTime time.Time
}

func CollectHouseSportEvents(getEvents func() ([]model.Event, error)) []model.Event {
	start := time.Now()

	events, err := getEvents()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nCollected: %d events\n", len(events))
	fmt.Printf("\nTook: %f secs\n", time.Since(start).Seconds())

	util.SortByOutcome(events)

	return events
}

func MergeEventsByBestOdds(events map[string]Event, newEvents []model.Event) {

	for _, event := range newEvents {

		outcomesKey := ""

		houses := []string{}

		for indx := range event.Outcome {
			outcomesKey += event.Outcome[indx] + ","
			houses = append(houses, "mozzartbet")
		}

		event := Event{
			Outcome:   event.Outcome,
			Odds:      event.Odds,
			House:     houses,
			StartTime: event.StartTime,
		}

		events[outcomesKey] = event
	}
}

func main() {

	mozzartbetFootballEvents := CollectHouseSportEvents(mozzartbet.GetFootballEvents)
	// meridianbetFootballEvents := CollectHouseSportEvents(meridianbet.GetFootballEvents)

	events := map[string]Event{}

	MergeEventsByBestOdds(events, mozzartbetFootballEvents)
	// MergeEventsByBestOdds(events, meridianbetFootballEvents)

	// // sort events array by time
	// sort.Slice(eventsArr, func(i, j int) bool {
	// 	return eventsArr[i].StartTime.Before(eventsArr[j].StartTime)
	// })

	// for _, event := range eventsArr {
	// 	printEvent(event)
	// }

}

func ShowEvent(event Event) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Outcome", "Odds", "House", "Bet Coefficient"})

	for indx := range event.Outcome {

		t.AppendRow(table.Row{
			event.Outcome[indx],
			event.Odds[indx],
			event.House[indx],
		})
		t.AppendSeparator()
	}

	totalOdds := math.CalculateOdds(event.Odds)

	t.AppendSeparator()
	t.AppendFooter(table.Row{"Total Odds", totalOdds})
	t.AppendFooter(table.Row{"Start Time", event.StartTime.Format(time.RFC822)})

	t.SetStyle(table.StyleLight)
	t.Render()
}
