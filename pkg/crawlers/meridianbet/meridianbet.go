package meridianbet

import (
	"betulator/pkg/httprequest"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

var houseLink string = "https://meridianbet.rs/"

type Event struct {
	Outcome []string
	Odds    []decimal.Decimal
}

func GetFootballEventsSortedByOutcome() ([]Event, error) {

	events, err := GetFootballEvents()
	if err != nil {
		return nil, err
	}

	for _, event := range events {
		if event.Outcome[0] < event.Outcome[2] {
			event.Outcome[0], event.Outcome[2] = event.Outcome[2], event.Outcome[0]
			event.Odds[0], event.Odds[2] = event.Odds[2], event.Odds[0]
		}
	}

	return events, nil
}

func GetFootballEvents() ([]Event, error) {

	currentTime := time.Now().Format(time.RFC3339)

	jsonResult, err := httprequest.Get("https://meridianbet.rs/sails/sport-with-leagues/58/date/" + currentTime + "/filter/all/filterPosition/0,0,0")
	if err != nil {
		return nil, err
	}

	resultBytes := []byte(jsonResult)

	// // events[0, n]/events[0, n]/marketShort[0]/selection[0,1,2]/price
	// // events[0]/events[0]/team[0]/name		team[1]/name

	parsedData := Football{}

	err = json.Unmarshal(resultBytes, &parsedData)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	events := []Event{}

	for _, eventGroup := range parsedData.Events {

		for _, event := range eventGroup.Events {

			oddsA, err := decimal.NewFromString(event.MarketShort[0].Selection[0].Price)
			if err != nil {
				return nil, err
			}
			oddsDraw, err := decimal.NewFromString(event.MarketShort[0].Selection[1].Price)
			if err != nil {
				return nil, err
			}
			oddsB, err := decimal.NewFromString(event.MarketShort[0].Selection[2].Price)
			if err != nil {
				return nil, err
			}

			newEvent := Event{
				Outcome: []string{
					strings.ToLower(event.Team[0].Name),
					"draw",
					strings.ToLower(event.Team[1].Name),
				},
				Odds: []decimal.Decimal{
					oddsA,
					oddsDraw,
					oddsB,
				},
			}

			events = append(events, newEvent)
		}
	}

	return events, nil
}
