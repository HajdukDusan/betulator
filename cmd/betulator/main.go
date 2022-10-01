package main

import (
	"betulator/internal/math"
	"betulator/internal/util"
	"betulator/pkg/model"
	"betulator/pkg/scrapers/meridianbet"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

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

func createPossibleCombinations(wordsMatrix *[][]string, combinations *[][]string, rowIndx int, currComb []string) {

	if rowIndx >= len(*wordsMatrix) {
		(*combinations) = append(*combinations, currComb)
		return
	}

	for j := 0; j < len((*wordsMatrix)[rowIndx]); j++ {

		cpyComb := make([]string, len(currComb))
		copy(cpyComb, currComb)

		cpyComb = append(cpyComb, (*wordsMatrix)[rowIndx][j])

		createPossibleCombinations(wordsMatrix, combinations, rowIndx+1, cpyComb)
	}
}

func CheckIfEventsMatch(currentEvent model.Event, newEvent model.Event) bool {

	wordsMatrix := createOutcomeWordMatrix(currentEvent.Outcome)

	combinations := make([][]string, 0)

	createPossibleCombinations(&wordsMatrix, &combinations, 0, []string{})

	for _, combination := range combinations {

		found := true

		for outcomeIndx, word := range combination {

			wordFound := false

			outcomeWords := strings.Fields(strings.ToLower(newEvent.Outcome[outcomeIndx]))

			for _, outcomeWord := range outcomeWords {

				if outcomeWord == word {
					wordFound = true
					break
				}
			}

			if wordFound == false {
				found = false
				break
			}
		}

		if found {
			return true
		}
	}

	return false
}

func MergeEventsByBestOdds(events *[]model.Event, newEvents []model.Event, house string) {

	for _, newEvent := range newEvents {

		found := false

		// try to find existing event
		for _, event := range *events {
			if CheckIfEventsMatch(event, newEvent) {

				fmt.Println("isti su", event.Outcome, newEvent.Outcome)

				// merge best odds
				for indx := range event.Odds {
					if newEvent.Odds[indx].Cmp(event.Odds[indx]) == 1 {
						event.Outcome[indx] = newEvent.Outcome[indx]
						event.Odds[indx] = newEvent.Odds[indx]
						event.House[indx] = house
					}
				}
				found = true
				break
			}
		}

		if !found {

			(*events) = append(*events, model.Event{
				Outcome:   newEvent.Outcome,
				Odds:      newEvent.Odds,
				House:     newEvent.House,
				StartTime: newEvent.StartTime,
			})
		}
	}
}

func main() {

	// mozzartbetFootballEvents := CollectHouseSportEvents(mozzartbet.GetFootballEvents)
	meridianbetFootballEvents := CollectHouseSportEvents(meridianbet.GetFootballEvents)

	// soccerbetFootballEvents := CollectHouseSportEvents(soccerbet.GetFootballEvents)

	events := make([]model.Event, 0)

	// MergeEventsByBestOdds(events, mozzartbetFootballEvents)
	MergeEventsByBestOdds(&events, meridianbetFootballEvents, "meridianbet")
	// MergeEventsByBestOdds(events, soccerbetFootballEvents, "soccerbet")

	// // sort events array by time
	// sort.Slice(eventsArr, func(i, j int) bool {
	// 	return eventsArr[i].StartTime.Before(eventsArr[j].StartTime)
	// })

	fmt.Println(len(events))

	for _, event := range events {
		ShowEvent(event)
	}

}

func ShowEvent(event model.Event) {
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
