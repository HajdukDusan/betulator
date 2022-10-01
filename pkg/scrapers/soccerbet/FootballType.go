package soccerbet

type Football []struct {
	ID                 int         `json:"Id"`
	HomeCompetitorName string      `json:"HomeCompetitorName"`
	AwayCompetitorName string      `json:"AwayCompetitorName"`
	Code               int         `json:"Code"`
	ExternalID         int         `json:"ExternalId"`
	StreamID           interface{} `json:"StreamId"`
	StartDate          string      `json:"StartDate"`
	Status             int         `json:"Status"`
	CompetitionID      int         `json:"CompetitionId"`
	SportID            int         `json:"SportId"`
	FavouriteBets      []struct {
		ID                   int         `json:"Id"`
		Odds                 float64     `json:"Odds"`
		IsEnabled            bool        `json:"IsEnabled"`
		HandicapOrTotalParam interface{} `json:"HandicapOrTotalParam"`
		ScoreParam           interface{} `json:"ScoreParam"`
		PeriodParam          interface{} `json:"PeriodParam"`
		AdditionalParam      interface{} `json:"AdditionalParam"`
		BetGameOutcomeID     int         `json:"BetGameOutcomeId"`
	} `json:"FavouriteBets"`
	AllBets interface{} `json:"AllBets"`
}
