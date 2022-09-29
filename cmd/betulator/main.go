package main

import (
	"betulator/internal/math"
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

func main() {

	mozzartbetFootballEvents, err := mozzartbet.GetFootballEvents()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\nCollected: %d events\n", len(mozzartbetFootballEvents))

	footballEvents := map[string]Event{}

	eventsArr := []Event{}

	for _, event := range mozzartbetFootballEvents {

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

		footballEvents[outcomesKey] = event
		eventsArr = append(eventsArr, event)
	}

	// // sort events array by time
	// sort.Slice(eventsArr, func(i, j int) bool {
	// 	return eventsArr[i].StartTime.Before(eventsArr[j].StartTime)
	// })

	for _, event := range eventsArr {
		printEvent(event)
	}

	return

	// meridianbetFootballEvents, err := meridianbet.GetFootballEventsSortedByOutcome()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// footballEvents := map[string]Event{}

	// eventsArr := []Event{}

	// for _, event := range meridianbetFootballEvents {

	// 	outcomesKey := ""

	// 	houses := []string{}

	// 	for indx := range event.Outcome {
	// 		outcomesKey += event.Outcome[indx] + ","
	// 		houses = append(houses, "meridianbet")
	// 	}

	// 	event := Event{
	// 		Outcome:   event.Outcome,
	// 		Odds:      event.Odds,
	// 		House:     houses,
	// 		StartTime: event.StartTime,
	// 	}

	// 	footballEvents[outcomesKey] = event
	// 	eventsArr = append(eventsArr, event)
	// }

	// // sort events array by time
	// sort.Slice(eventsArr, func(i, j int) bool {
	// 	return eventsArr[i].StartTime.Before(eventsArr[j].StartTime)
	// })

	// for _, event := range eventsArr {
	// 	printEvent(event)
	// }

}

func printEvent(event Event) {
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
