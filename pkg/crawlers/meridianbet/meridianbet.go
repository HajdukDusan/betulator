package meridianbet

import (
	"betulator/pkg/httprequest"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type outcome struct {
	winner string
	odds   decimal.Decimal
}

func GetBets() (map[[3]string][]decimal.Decimal, error) {

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

	result := make(map[[3]string][]decimal.Decimal)

	for _, eventGroup := range parsedData.Events {

		for _, event := range eventGroup.Events {

			teamA := strings.ToLower(event.Team[0].Name)
			teamB := strings.ToLower(event.Team[1].Name)

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

			if teamA >= teamB {
				result[[3]string{teamA, "draw", teamB}] = []decimal.Decimal{oddsA, oddsDraw, oddsB}
			} else {
				result[[3]string{teamB, "draw", teamA}] = []decimal.Decimal{oddsB, oddsDraw, oddsA}
			}
		}
	}

	return result, nil
}
