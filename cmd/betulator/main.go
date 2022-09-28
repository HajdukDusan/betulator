package main

import (
	"betulator/internal/math"
	"betulator/pkg/crawlers/meridianbet"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shopspring/decimal"
)

type Event struct {
	Outcome []string
	Odds    []decimal.Decimal
	House   []string
}

func main() {

	meridianbetFootballEvents, err := meridianbet.GetFootballEventsSortedByOutcome()
	if err != nil {
		fmt.Println(err)
	}

	footballEvents := map[string]Event{}

	for _, event := range meridianbetFootballEvents {

		outcomesKey := ""

		houses := []string{}

		for indx := range event.Outcome {
			outcomesKey += event.Outcome[indx] + ","
			houses = append(houses, "meridianbet")
		}

		footballEvents[outcomesKey] = Event{
			Outcome: event.Outcome,
			Odds:    event.Odds,
			House:   houses,
		}
	}

	for _, event := range footballEvents {
		printEvent(event)
	}
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
	t.AppendFooter(table.Row{"TOTAL ODDS", totalOdds})

	t.SetStyle(table.StyleLight)
	t.Render()
}
