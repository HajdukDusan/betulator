package mozzartbet

type FootballEventsResponse struct {
	Matches []struct {
		CompetitionCommentSr string `json:"competitionComment_sr,omitempty"`
		ID                   int    `json:"id"`
		StartTime            int64  `json:"startTime"`
		SpecialType          int    `json:"specialType"`
		CompetitionNameSr    string `json:"competition_name_sr"`
		MatchNumber          int    `json:"matchNumber"`
		GameCounts           struct {
			Total                       int `json:"total"`
			FiveAe09C36Df923371Ecc9675A int `json:"5ae09c36df923371ecc9675a"`
			FiveAc7739Adf92330546247Efe int `json:"5ac7739adf92330546247efe"`
			FiveAda5917Df923343C2505133 int `json:"5ada5917df923343c2505133"`
			FiveAe0989Adf923371Ecc96758 int `json:"5ae0989adf923371ecc96758"`
			FiveAda5E04Df923343C2505137 int `json:"5ada5e04df923343c2505137"`
			FiveAda5C6Bdf923343C2505136 int `json:"5ada5c6bdf923343c2505136"`
			FiveAe09Cc0Df923371Ecc9675C int `json:"5ae09cc0df923371ecc9675c"`
			FiveAda5Fbedf923343C2505139 int `json:"5ada5fbedf923343c2505139"`
			FiveB99492Ab8118876C583477F int `json:"5b99492ab8118876c583477f"`
			FiveAe09C7Adf923371Ecc9675B int `json:"5ae09c7adf923371ecc9675b"`
			FiveAda5B7Edf923343C2505135 int `json:"5ada5b7edf923343c2505135"`
			FiveAda479Adf923343C250512D int `json:"5ada479adf923343c250512d"`
			FiveAda6262Df923343C250513B int `json:"5ada6262df923343c250513b"`
			FiveAda6394Df923343C250513C int `json:"5ada6394df923343c250513c"`
			FiveAe09Bb2Df923371Ecc96759 int `json:"5ae09bb2df923371ecc96759"`
			FiveAda61Addf923343C250513A int `json:"5ada61addf923343c250513a"`
			FiveAda49C3Df923343C250512F int `json:"5ada49c3df923343c250512f"`
			FiveAe08E02Df923371Ecc96751 int `json:"5ae08e02df923371ecc96751"`
			FiveAda5755Df923343C2505131 int `json:"5ada5755df923343c2505131"`
			FiveAda3Fc8Df923343C250512B int `json:"5ada3fc8df923343c250512b"`
		} `json:"gameCounts"`
		CountKodds   int `json:"countKodds"`
		OddsCount    int `json:"oddsCount"`
		Participants []struct {
			ID          int    `json:"id"`
			Description string `json:"description"`
			Name        string `json:"name"`
			ShortName   string `json:"shortName"`
		} `json:"participants"`
		MainMatch   string `json:"mainMatch"`
		Competition struct {
			ID                   int `json:"id"`
			Priority             int `json:"priority"`
			PriorityByGroupation struct {
				OrderItems []struct {
					GroupationID int `json:"groupationId"`
					Rank         int `json:"rank"`
				} `json:"orderItems"`
			} `json:"priorityByGroupation"`
			Name  string `json:"name"`
			Sport struct {
				ID                   int `json:"id"`
				PriorityByGroupation struct {
					OrderItems []struct {
						GroupationID int `json:"groupationId"`
						Rank         int `json:"rank"`
					} `json:"orderItems"`
				} `json:"priorityByGroupation"`
				Priority int    `json:"priority"`
				Name     string `json:"name"`
			} `json:"sport"`
			ShortName string `json:"shortName"`
			Country   struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"country"`
		} `json:"competition"`
	} `json:"matches"`
	Total int `json:"total"`
}
