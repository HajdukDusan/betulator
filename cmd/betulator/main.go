package main

import (
	"betulator/internal/math"
	"betulator/internal/util"
	"betulator/pkg/model"
	"betulator/pkg/scrapers/meridianbet"
	"betulator/pkg/scrapers/soccerbet"
	"fmt"
	"os"
	"strings"
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

func createOutcomeWordMatrix(outcomes []string) [][]string {
	result := make([][]string, len(outcomes))

	for indx := range outcomes {

		words := strings.Fields(outcomes[indx])

		result[indx] = make([]string, len(words))

		for i := range words {
			result[indx][i] = strings.ToLower(words[i])
		}
	}
	return result
}

func recursion(matrix *[][]string, combinations *[]string, rowIndx int, curr []string) {

	if row >= len(*matrix) {
		return
	}

	for x := row; x < len((*matrix)[row]); x++ {
		recursion(matrix, combinations, row+1)
	}
}

func createPossibleKeys(outcomes []string) {
	result := []string{}

	outcomesWords := createOutcomeWordMatrix(outcomes)

	for i := 0; i < len(outcomesWords); i++ {

		for j := 0; j < len(outcomesWords[j]); j++ {

		}
	}

	for _, words := range outcomesWords {

		for _, word := range words {

		}
	}

	for indx := range outcomes {
		result += strings.ToLower(event.Outcome[indx]) + ","
	}
	return result
}

func MergeEventsByBestOdds(events map[string]Event, newEvents []model.Event, house string) {

	for _, event := range newEvents {

		found := false

		for {
			_, ok := events[outcomesKey]
			if ok == true {

				for indx := range events[outcomesKey].Odds {

					if events[outcomesKey].Odds[indx].Cmp(event.Odds[indx]) == 1 {
						events[outcomesKey].Odds[indx] = event.Odds[indx]
						events[outcomesKey].House[indx] = house
					}
				}
			}
		}

		if found == false {
			houses := []string{}

			for i := 0; i < len(event.Outcome); i++ {
				houses = append(houses, house)
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
}

func main() {

	// mozzartbetFootballEvents := CollectHouseSportEvents(mozzartbet.GetFootballEvents)
	meridianbetFootballEvents := CollectHouseSportEvents(meridianbet.GetFootballEvents)
	soccerbetFootballEvents := CollectHouseSportEvents(soccerbet.GetFootballEvents)

	events := map[string]Event{}

	// MergeEventsByBestOdds(events, mozzartbetFootballEvents)
	MergeEventsByBestOdds(events, meridianbetFootballEvents, "meridianbet")
	MergeEventsByBestOdds(events, soccerbetFootballEvents, "soccerbet")

	// // sort events array by time
	// sort.Slice(eventsArr, func(i, j int) bool {
	// 	return eventsArr[i].StartTime.Before(eventsArr[j].StartTime)
	// })

	fmt.Println(len(events))

	for _, event := range events {
		ShowEvent(event)
	}

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
