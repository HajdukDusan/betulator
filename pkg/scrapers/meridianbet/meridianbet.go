package meridianbet

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

	currentTime := time.Now().Format(time.RFC3339)

	jsonResult, err := httprequest.Get("https://meridianbet.rs/sails/sport/58/date/" + currentTime + "/filter/all/offset/0?filterPositions=0,0,0")
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

	for _, event := range parsedData.Events {

		oddsA, err := decimal.NewFromString(event.StandardShortMarkets[0].Selection[0].Price)
		if err != nil {
			return nil, err
		}
		oddsDraw, err := decimal.NewFromString(event.StandardShortMarkets[0].Selection[1].Price)
		if err != nil {
			return nil, err
		}
		oddsB, err := decimal.NewFromString(event.StandardShortMarkets[0].Selection[2].Price)
		if err != nil {
			return nil, err
		}

		house := []string{
			"meridianbet",
			"meridianbet",
			"meridianbet",
		}

		newEvent := model.Event{
			Outcome: []string{
				event.Team[0].Name,
				"Draw",
				event.Team[1].Name,
			},
			Odds: []decimal.Decimal{
				oddsA,
				oddsDraw,
				oddsB,
			},
			StartTime: event.StartTime,
			House:     house,
		}

		events = append(events, newEvent)
	}

	return events, nil
}
