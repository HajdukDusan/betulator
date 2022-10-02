package mozzartbet

type FootballOddsResponse []struct {
	ID         int `json:"id"`
	GameCounts struct {
		FiveAe09C36Df923371Ecc9675A int `json:"5ae09c36df923371ecc9675a"`
		Total                       int `json:"total"`
		FiveAc7739Adf92330546247Efe int `json:"5ac7739adf92330546247efe"`
		FiveAda5917Df923343C2505133 int `json:"5ada5917df923343c2505133"`
		FiveAda5E04Df923343C2505137 int `json:"5ada5e04df923343c2505137"`
		FiveAe0989Adf923371Ecc96758 int `json:"5ae0989adf923371ecc96758"`
		FiveAda5C6Bdf923343C2505136 int `json:"5ada5c6bdf923343c2505136"`
		FiveAe09Cc0Df923371Ecc9675C int `json:"5ae09cc0df923371ecc9675c"`
		FiveAda5Fbedf923343C2505139 int `json:"5ada5fbedf923343c2505139"`
		FiveAe09C7Adf923371Ecc9675B int `json:"5ae09c7adf923371ecc9675b"`
		FiveB99492Ab8118876C583477F int `json:"5b99492ab8118876c583477f"`
		FiveAda5B7Edf923343C2505135 int `json:"5ada5b7edf923343c2505135"`
		FiveAda6262Df923343C250513B int `json:"5ada6262df923343c250513b"`
		FiveAda479Adf923343C250512D int `json:"5ada479adf923343c250512d"`
		FiveAda6394Df923343C250513C int `json:"5ada6394df923343c250513c"`
		FiveAe09Bb2Df923371Ecc96759 int `json:"5ae09bb2df923371ecc96759"`
		FiveAda61Addf923343C250513A int `json:"5ada61addf923343c250513a"`
		FiveAda49C3Df923343C250512F int `json:"5ada49c3df923343c250512f"`
		FiveAe08E02Df923371Ecc96751 int `json:"5ae08e02df923371ecc96751"`
		FiveAda5755Df923343C2505131 int `json:"5ada5755df923343c2505131"`
		FiveAda3Fc8Df923343C250512B int `json:"5ada3fc8df923343c250512b"`
	} `json:"gameCounts"`
	Kodds struct {
		Num1001001001 struct {
			ID              int    `json:"id"`
			SpecialOddValue string `json:"specialOddValue"`
			MatchID         int    `json:"matchId"`
			Value           string `json:"value"`
			WinStatus       string `json:"winStatus"`
			SubGame         struct {
				ID                  int    `json:"id"`
				SubGameID           int    `json:"subGameId"`
				GameID              int    `json:"gameId"`
				GameName            string `json:"gameName"`
				SubGameName         string `json:"subGameName"`
				GameShortName       string `json:"gameShortName"`
				SubGameDescription  string `json:"subGameDescription"`
				SpecialOddValueType string `json:"specialOddValueType"`
				Priority            bool   `json:"priority"`
			} `json:"subGame"`
		} `json:"1001001001"`
		Num1001001002 struct {
			ID              int    `json:"id"`
			SpecialOddValue string `json:"specialOddValue"`
			MatchID         int    `json:"matchId"`
			Value           string `json:"value"`
			WinStatus       string `json:"winStatus"`
			SubGame         struct {
				ID                  int    `json:"id"`
				SubGameID           int    `json:"subGameId"`
				GameID              int    `json:"gameId"`
				GameName            string `json:"gameName"`
				SubGameName         string `json:"subGameName"`
				GameShortName       string `json:"gameShortName"`
				SubGameDescription  string `json:"subGameDescription"`
				SpecialOddValueType string `json:"specialOddValueType"`
				Priority            bool   `json:"priority"`
			} `json:"subGame"`
		} `json:"1001001002"`
		Num1001001003 struct {
			ID              int    `json:"id"`
			SpecialOddValue string `json:"specialOddValue"`
			MatchID         int    `json:"matchId"`
			Value           string `json:"value"`
			WinStatus       string `json:"winStatus"`
			SubGame         struct {
				ID                  int    `json:"id"`
				SubGameID           int    `json:"subGameId"`
				GameID              int    `json:"gameId"`
				GameName            string `json:"gameName"`
				SubGameName         string `json:"subGameName"`
				GameShortName       string `json:"gameShortName"`
				SubGameDescription  string `json:"subGameDescription"`
				SpecialOddValueType string `json:"specialOddValueType"`
				Priority            bool   `json:"priority"`
			} `json:"subGame"`
		} `json:"1001001003"`
	} `json:"kodds"`
}
