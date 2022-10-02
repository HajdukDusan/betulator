package meridianbet

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

	for i := 0; i < 20; i++ {

		currentTime := time.Now().Format(time.RFC3339)
		jsonResult, err := httprequest.Get(
			"https://meridianbet.rs/sails/sport/58/date/" + currentTime +
				"/filter/all/offset/" + strconv.Itoa(i*20) + "?filterPositions=0,0,0")
		if err != nil {
			return nil, err
		}

		resultBytes := []byte(jsonResult)

		parsedData := Football{}

		err = json.Unmarshal(resultBytes, &parsedData)

		if err != nil {
			return nil, err
		}

		loc, _ := time.LoadLocation("Europe/Belgrade")

		for _, event := range parsedData.Events {

			if len(event.StandardShortMarketsMobile[0].Selection) == 0 {
				continue
			}

			oddsA, err := decimal.NewFromString(event.StandardShortMarketsMobile[0].Selection[0].Price)
			if err != nil {
				return nil, err
			}
			oddsDraw, err := decimal.NewFromString(event.StandardShortMarketsMobile[0].Selection[1].Price)
			if err != nil {
				return nil, err
			}
			oddsB, err := decimal.NewFromString(event.StandardShortMarketsMobile[0].Selection[2].Price)
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
				StartTime: event.StartTime.In(loc),
				House:     house,
			}

			events = append(events, newEvent)
		}
	}

	return events, nil
}
