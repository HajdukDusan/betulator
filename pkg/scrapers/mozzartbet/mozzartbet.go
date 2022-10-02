package mozzartbet

import (
	"betulator/pkg/httprequest"
	"betulator/pkg/model"
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

type requestEvents struct {
	Date                string      `json:"date"`
	SportIds            []int       `json:"sportIds"`
	CompetitionIds      []int       `json:"competitionIds"`
	Sort                string      `json:"sort"`
	Specials            interface{} `json:"specials"`
	Subgames            []int       `json:"subgames"`
	Size                int         `json:"size"`
	MostPlayed          bool        `json:"mostPlayed"`
	Type                string      `json:"type"`
	NumberOfGames       int         `json:"numberOfGames"`
	ActiveCompleteOffer bool        `json:"activeCompleteOffer"`
	Lang                string      `json:"lang"`
	Offset              int         `json:"offset"`
}

type requestOdds struct {
	MatchIds []int `json:"matchIds"`
	Subgames []int `json:"subgames"`
}

func GetFootballEvents() ([]model.Event, error) {

	eventsData, err := getFootballEvents(20)
	if err != nil {
		return nil, err
	}

	eventsIds := []int{}

	eventsMap := map[int]model.Event{}

	for _, match := range eventsData.Matches {

		// skip 4+1 combos etc
		if match.SpecialType == 0 {

			if len(match.Participants) >= 2 {

				eventsIds = append(eventsIds, match.ID)
				eventsMap[match.ID] = model.Event{
					Outcome: []string{
						match.Participants[0].Name,
						"Draw",
						match.Participants[1].Name,
					},
					House: []string{
						"mozzartbet",
						"mozzartbet",
						"mozzartbet",
					},
					StartTime: time.Unix(match.StartTime, 0),
				}
			}
		}
	}

	oddsData, err := getFootballOdds(eventsIds)
	if err != nil {
		return nil, err
	}

	for _, match := range oddsData {

		if event, ok := eventsMap[match.ID]; ok {

			teamA, _ := decimal.NewFromString(match.Kodds.Num1001001001.Value)
			draw, _ := decimal.NewFromString(match.Kodds.Num1001001002.Value)
			teamB, _ := decimal.NewFromString(match.Kodds.Num1001001003.Value)

			event.Odds = []decimal.Decimal{
				teamA, draw, teamB,
			}

			eventsMap[match.ID] = event
		}
	}

	events := []model.Event{}
	for _, event := range eventsMap {
		events = append(events, event)
	}

	return events, nil
}

func getFootballEvents(numOfEvents int) (FootballEventsResponse, error) {
	requestData := requestEvents{
		Date:                string(time.Now().Format("2022-10-02")),
		SportIds:            []int{1},
		CompetitionIds:      []int{},
		Sort:                "bytime",
		Specials:            nil,
		Subgames:            []int{},
		Size:                numOfEvents,
		MostPlayed:          false,
		Type:                "betting",
		NumberOfGames:       1000,
		ActiveCompleteOffer: false,
		Lang:                "en",
		Offset:              0,
	}

	json_data, err := json.Marshal(requestData)
	if err != nil {
		return FootballEventsResponse{}, err
	}

	jsonResult, err := httprequest.PostJson("https://www.mozzartbet.com/betOffer2", json_data)
	if err != nil {
		return FootballEventsResponse{}, err
	}

	resultBytes := []byte(jsonResult)

	parsedData := FootballEventsResponse{}

	err = json.Unmarshal(resultBytes, &parsedData)
	if err != nil {
		return FootballEventsResponse{}, err
	}

	return parsedData, nil
}

func getFootballOdds(eventIds []int) (FootballOddsResponse, error) {
	requestData := requestOdds{
		MatchIds: eventIds,
		// event outcome selection
		Subgames: []int{
			1001001001,
			1001001002,
			1001001003,
			1001002001,
			1001002002,
		},
	}

	json_data, err := json.Marshal(requestData)
	if err != nil {
		return FootballOddsResponse{}, err
	}

	jsonResult, err := httprequest.PostJson("https://www.mozzartbet.com/getBettingOdds", json_data)
	if err != nil {
		return FootballOddsResponse{}, err
	}

	resultBytes := []byte(jsonResult)

	parsedData := FootballOddsResponse{}

	err = json.Unmarshal(resultBytes, &parsedData)
	if err != nil {
		return FootballOddsResponse{}, err
	}

	return parsedData, nil
}
