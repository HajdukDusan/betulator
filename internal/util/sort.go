package util

import "betulator/pkg/model"

func SortIndividualEventsByOutcome(events []model.Event) {

	for _, event := range events {

		n := len(event.Outcome)
		for i := 0; i < n; i++ {

			if event.Outcome[i] == "Draw" {
				continue
			}

			var minIdx = i
			for j := i; j < n; j++ {

				if event.Outcome[j] == "Draw" {
					continue
				}

				if event.Outcome[j] < event.Outcome[minIdx] {
					minIdx = j
				}
			}
			event.Outcome[i], event.Outcome[minIdx] = event.Outcome[minIdx], event.Outcome[i]
			event.Odds[i], event.Odds[minIdx] = event.Odds[minIdx], event.Odds[i]
		}

	}
}
