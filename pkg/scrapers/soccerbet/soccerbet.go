package soccerbet

import (
	"betulator/pkg/httprequest"
	"betulator/pkg/model"
	"encoding/json"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

var houseLink string = "https://meridianbet.rs/"

func GetFootballEvents() ([]model.Event, error) {

	jsonResult, err := httprequest.Get("https://soccerbet.rs/api/Prematch/GetPagedSportMatches?sportId=1&pageNumber=1")
	if err != nil {
		return nil, err
	}

	resultBytes := []byte(jsonResult)

	parsedData := Football{}

	err = json.Unmarshal(resultBytes, &parsedData)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	events := []model.Event{}

	for _, event := range parsedData {

		odds := []decimal.Decimal{
			decimal.NewFromFloat(event.FavouriteBets[0].Odds),
			decimal.NewFromFloat(event.FavouriteBets[1].Odds),
			decimal.NewFromFloat(event.FavouriteBets[2].Odds),
		}

		outcomes := []string{
			event.HomeCompetitorName,
			"draw",
			event.AwayCompetitorName,
		}

		startTime, err := time.Parse(time.RFC3339, event.StartDate[:]+"Z")
		if err != nil {
			return nil, err
		}

		events = append(events, model.Event{
			Outcome:   outcomes,
			Odds:      odds,
			StartTime: startTime,
		})
	}

	return events, nil
}
