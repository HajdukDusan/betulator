package meridianbet

import (
	"betulator/pkg/httprequest"
	"encoding/json"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

// func GetBetOdds() (map[] event, error) {

// }

type FootballData struct {
	selectedFilter string `json:"selectedFilter"`
	events         []struct {
		events []struct {
			team []struct {
				name string `json:name`
			} `json:"events"`
		} `json:"events"`
	} `json:"events"`
}

func GetBetOdds() (map[interface{}][]decimal.Decimal, error) {

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

	for _, eventGroup := range parsedData.Events {

		for _, event := range eventGroup.Events {

			fmt.Println(event.Team[0])
			fmt.Println(event.Team[1])

			fmt.Println(event.MarketShort[0].Selection[0].Price)
			fmt.Println(event.MarketShort[0].Selection[1].Price)
			fmt.Println(event.MarketShort[0].Selection[2].Price)
		}
	}

	return nil, nil
}
