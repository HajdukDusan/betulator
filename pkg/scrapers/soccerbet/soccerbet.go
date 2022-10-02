package soccerbet

import (
	"betulator/pkg/httprequest"
	"betulator/pkg/model"
	"encoding/json"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

var houseLink string = "https://meridianbet.rs/"

func GetFootballEvents() ([]model.Event, error) {

	events := []model.Event{}

	for i := 1; i <= 20; i++ {

		jsonResult, err := httprequest.Get("https://soccerbet.rs/api/Prematch/GetPagedSportMatches?sportId=1&pageNumber=" + strconv.Itoa(i))
		if err != nil {
			return nil, err
		}

		resultBytes := []byte(jsonResult)

		parsedData := Football{}

		err = json.Unmarshal(resultBytes, &parsedData)

		if err != nil {
			return nil, err
		}

		for _, event := range parsedData {

			odds := []decimal.Decimal{
				decimal.NewFromFloat(event.FavouriteBets[0].Odds),
				decimal.NewFromFloat(event.FavouriteBets[1].Odds),
				decimal.NewFromFloat(event.FavouriteBets[2].Odds),
			}

			outcomes := []string{
				event.HomeCompetitorName,
				"Draw",
				event.AwayCompetitorName,
			}

			startTime, err := time.Parse(time.RFC3339, event.StartDate[:]+"Z")
			if err != nil {
				return nil, err
			}

			house := []string{
				"soccerbet",
				"soccerbet",
				"soccerbet",
			}

			events = append(events, model.Event{
				Outcome:   outcomes,
				Odds:      odds,
				StartTime: startTime,
				House:     house,
			})
		}
	}

	return events, nil
}
