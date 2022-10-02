package main

import (
	"betulator/internal/math"
	"betulator/internal/util"
	"betulator/pkg/model"
	"betulator/pkg/scrapers/meridianbet"
	"betulator/pkg/scrapers/mozzartbet"
	"betulator/pkg/scrapers/soccerbet"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

func CollectAndSortHouseSportEvents(getEvents func() ([]model.Event, error)) ([]model.Event, error) {

	events, err := getEvents()
	if err != nil {
		return nil, err
	}

	util.SortIndividualEventsByOutcome(events)

	return events, nil
}

func sortAllByTotalOdds(events []model.Event) {
	sort.Slice(events, func(i, j int) bool {
		return math.CalculateTotalOdds(events[i].Odds).Cmp(math.CalculateTotalOdds(events[j].Odds)) == -1
	})
	// sort events array by time
	// sort.Slice(eventsArr, func(i, j int) bool {
	// 	return eventsArr[i].StartTime.Before(eventsArr[j].StartTime)
	// })
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

	currentEventWordsMatrix := createOutcomeWordMatrix(currentEvent.Outcome)
	newEventWordsMatrix := createOutcomeWordMatrix(newEvent.Outcome)

	currentEventCombos := make([][]string, 0)
	newEventCombos := make([][]string, 0)

	createPossibleCombinations(&currentEventWordsMatrix, &currentEventCombos, 0, []string{})
	createPossibleCombinations(&newEventWordsMatrix, &newEventCombos, 0, []string{})

	for _, currEventCombo := range currentEventCombos {

		for _, newEventCombo := range newEventCombos {

			found := true

			for indx := range currEventCombo {

				// skip words with less than 4 letters or women
				if len(currEventCombo[indx]) < 4 ||
					currEventCombo[indx] == "women" ||
					currEventCombo[indx] == "city" {
					found = false
					break
				}

				if currEventCombo[indx] != newEventCombo[indx] {
					found = false
					break
				}
			}

			if found {
				return true
			}
		}
	}

	return false
}

func MergeEventsByBestOdds(events *[]model.Event, newEvents []model.Event) {

	for _, newEvent := range newEvents {

		found := false

		// try to find existing event
		for _, event := range *events {

			//[Maccabi Haifa FC Draw Maccabi Tel Aviv FC] [Maccabi Herzliya Draw Maccabi Ironi Ashdod]!!!

			if CheckIfEventsMatch(event, newEvent) {

				fmt.Println("Merged: ", event.Outcome, newEvent.Outcome)

				// merge best odds
				for indx := range event.Odds {
					if newEvent.Odds[indx].Cmp(event.Odds[indx]) == 1 {
						event.Outcome[indx] = newEvent.Outcome[indx]
						event.Odds[indx] = newEvent.Odds[indx]
						event.House[indx] = newEvent.House[indx]
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

func worker(getEvents func() ([]model.Event, error), house string, results chan<- *[]model.Event) {

	start := time.Now()

	fmt.Println("Started scraping", house, "..")
	events, err := CollectAndSortHouseSportEvents(getEvents)
	if err != nil {
		fmt.Println("Scraping", house, "failed with error:", err)
		results <- nil
	}
	fmt.Println("Finished scraping", house, "in", time.Since(start).Seconds(), "secs, took:", len(events), "events")

	results <- &events
}

func main() {

	const numJobs = 3
	results := make(chan *[]model.Event, numJobs)

	go worker(mozzartbet.GetFootballEvents, "mozzartbet", results)
	go worker(meridianbet.GetFootballEvents, "meridianbet", results)
	go worker(soccerbet.GetFootballEvents, "soccerbet", results)

	fetchedEvents := make([][]model.Event, numJobs)

	for i := 1; i <= numJobs; i++ {
		fetchedEvents = append(fetchedEvents, *<-results)
	}

	fmt.Println("Scraping Completed!")

	fmt.Println("\nMerging Events..")
	events := make([]model.Event, 0)

	for _, houseEvents := range fetchedEvents {
		MergeEventsByBestOdds(&events, houseEvents)
	}

	fmt.Println("Events Sum: ", len(events))

	sortAllByTotalOdds(events)

	// for _, event := range events {
	// 	ShowEvent(event)
	// }

	for i := 0; i < 10; i++ {
		ShowEvent(events[i])
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

	totalOdds := math.CalculateTotalOdds(event.Odds)

	t.AppendSeparator()
	t.AppendFooter(table.Row{"Total Odds", totalOdds})
	t.AppendFooter(table.Row{"Start Time", event.StartTime.Format(time.RFC822)})

	t.SetStyle(table.StyleLight)
	t.Render()
}
