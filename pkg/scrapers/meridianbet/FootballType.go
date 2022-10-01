package meridianbet

import "time"

type Football struct {
	Events []struct {
		LeagueNameTranslations []struct {
			Locale      string `json:"locale"`
			Translation string `json:"translation"`
		} `json:"leagueNameTranslations"`
		SportNameTranslations []struct {
			Locale      string `json:"locale"`
			Translation string `json:"translation"`
		} `json:"sportNameTranslations"`
		RegionNameTranslations []struct {
			Locale      string `json:"locale"`
			Translation string `json:"translation"`
		} `json:"regionNameTranslations"`
		MarketOrder []interface{} `json:"marketOrder"`
		Result      []struct {
			ExtraData []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"extraData"`
		} `json:"result"`
		Team []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"team"`
		TeamHomeTranslations []interface{} `json:"teamHomeTranslations"`
		TeamAwayTranslations []interface{} `json:"teamAwayTranslations"`
		RegionNameSlugs      []struct {
			Locale string `json:"locale"`
			Slug   string `json:"slug"`
		} `json:"regionNameSlugs"`
		SportNameSlugs []struct {
			Locale string `json:"locale"`
			Slug   string `json:"slug"`
		} `json:"sportNameSlugs"`
		LiveShortMarkets     []interface{} `json:"liveShortMarkets"`
		StandardShortMarkets []struct {
			Selection []struct {
				Name             string `json:"name"`
				Price            string `json:"price"`
				State            string `json:"state"`
				ActivationPrice  string `json:"activationPrice"`
				NameTranslations []struct {
					Locale      string `json:"locale"`
					Translation string `json:"translation"`
				} `json:"nameTranslations"`
			} `json:"selection"`
			ID               string    `json:"id"`
			Name             string    `json:"name"`
			TemplateID       string    `json:"templateId"`
			TemplateName     string    `json:"templateName"`
			State            string    `json:"state"`
			CreationTime     time.Time `json:"creationTime"`
			ActivationTime   time.Time `json:"activationTime"`
			WinnerPrice      string    `json:"winnerPrice"`
			NameTranslations []struct {
				Locale      string `json:"locale"`
				Translation string `json:"translation"`
			} `json:"nameTranslations"`
			Inc                 int         `json:"inc"`
			OverUnder           interface{} `json:"overUnder"`
			Handicap            interface{} `json:"handicap"`
			ActivationOverUnder string      `json:"activationOverUnder,omitempty"`
		} `json:"standardShortMarkets"`
		LiveShortMarketsMobile     []interface{} `json:"liveShortMarketsMobile"`
		StandardShortMarketsMobile []struct {
			Selection []struct {
				Name             string `json:"name"`
				Price            string `json:"price"`
				State            string `json:"state"`
				ActivationPrice  string `json:"activationPrice"`
				NameTranslations []struct {
					Locale      string `json:"locale"`
					Translation string `json:"translation"`
				} `json:"nameTranslations"`
			} `json:"selection,omitempty"`
			ID               string    `json:"id,omitempty"`
			Name             string    `json:"name,omitempty"`
			TemplateID       string    `json:"templateId,omitempty"`
			TemplateName     string    `json:"templateName,omitempty"`
			State            string    `json:"state,omitempty"`
			CreationTime     time.Time `json:"creationTime,omitempty"`
			ActivationTime   time.Time `json:"activationTime,omitempty"`
			WinnerPrice      string    `json:"winnerPrice,omitempty"`
			NameTranslations []struct {
				Locale      string `json:"locale"`
				Translation string `json:"translation"`
			} `json:"nameTranslations,omitempty"`
			Inc                 int         `json:"inc,omitempty"`
			OverUnder           interface{} `json:"overUnder,omitempty"`
			Handicap            interface{} `json:"handicap,omitempty"`
			ActivationOverUnder string      `json:"activationOverUnder,omitempty"`
			Selections          []struct {
				Selection []struct {
					Name             string `json:"name"`
					Price            string `json:"price"`
					State            string `json:"state"`
					ActivationPrice  string `json:"activationPrice"`
					NameTranslations []struct {
						Locale      string `json:"locale"`
						Translation string `json:"translation"`
					} `json:"nameTranslations"`
				} `json:"selection"`
				ID               string    `json:"id"`
				Name             string    `json:"name"`
				TemplateID       string    `json:"templateId"`
				TemplateName     string    `json:"templateName"`
				State            string    `json:"state"`
				CreationTime     time.Time `json:"creationTime"`
				ActivationTime   time.Time `json:"activationTime"`
				WinnerPrice      string    `json:"winnerPrice"`
				NameTranslations []struct {
					Locale      string `json:"locale"`
					Translation string `json:"translation"`
				} `json:"nameTranslations"`
				Inc                    int         `json:"inc"`
				OverUnder              interface{} `json:"overUnder"`
				Handicap               interface{} `json:"handicap"`
				SelectionPositionIndex int         `json:"selectionPositionIndex"`
			} `json:"selections,omitempty"`
		} `json:"standardShortMarketsMobile"`
		ID                        string    `json:"id"`
		Code                      string    `json:"code"`
		TemplateID                string    `json:"templateId"`
		TemplateName              string    `json:"templateName"`
		State                     string    `json:"state"`
		CreationTime              time.Time `json:"creationTime"`
		StartTime                 time.Time `json:"startTime"`
		OfferType                 string    `json:"offerType"`
		LeagueID                  int       `json:"leagueId"`
		LeagueName                string    `json:"leagueName"`
		SportID                   int       `json:"sportId"`
		SportName                 string    `json:"sportName"`
		IconName                  string    `json:"iconName"`
		FeedBetradarunified       bool      `json:"feedBetradarunified"`
		IsBetradarUnifiedOutright bool      `json:"isBetradarUnifiedOutright"`
		BetradarUnified           struct {
			Locale string `json:"locale"`
			Type   string `json:"type"`
			ID     string `json:"id"`
		} `json:"betradarUnified,omitempty"`
		SetEventOrder              string      `json:"setEventOrder"`
		WasSetEventPriorityEvent   bool        `json:"wasSetEventPriorityEvent"`
		TopMatch                   bool        `json:"topMatch"`
		PeriodDuration             string      `json:"periodDuration"`
		Name                       string      `json:"name"`
		RegionID                   string      `json:"regionId"`
		RegionName                 string      `json:"regionName"`
		LeagueNameSlug             string      `json:"leagueNameSlug"`
		PowerLeaguesPriority       int         `json:"powerLeaguesPriority"`
		HasMarket                  bool        `json:"hasMarket"`
		FirstMarketPrice           string      `json:"firstMarketPrice"`
		IsPowerLeague              bool        `json:"isPowerLeague"`
		IsVisibleStandard          bool        `json:"isVisibleStandard"`
		PreviousVisibleLiveProp    bool        `json:"previousVisibleLiveProp"`
		IsVisibleLive              bool        `json:"isVisibleLive"`
		IsVisibleRacing            bool        `json:"isVisibleRacing"`
		MostPlayed                 bool        `json:"mostPlayed"`
		NumberOfActiveSelections   int         `json:"numberOfActiveSelections"`
		CreatedAt                  time.Time   `json:"createdAt"`
		UpdatedAt                  time.Time   `json:"updatedAt"`
		BetradarEventID            interface{} `json:"betradarEventId,omitempty"`
		CurrentMinute              interface{} `json:"currentMinute,omitempty"`
		Inc                        int         `json:"inc,omitempty"`
		IsLowPriorityChange        bool        `json:"isLowPriorityChange,omitempty"`
		IsNewEvent                 bool        `json:"isNewEvent,omitempty"`
		OldState                   string      `json:"oldState,omitempty"`
		SkipCreatingLeagueNameSlug bool        `json:"skipCreatingLeagueNameSlug,omitempty"`
		SkipCreatingRegionNameSlug bool        `json:"skipCreatingRegionNameSlug,omitempty"`
		SkipCreatingSportNameSlug  bool        `json:"skipCreatingSportNameSlug,omitempty"`
		WasStandardActive          bool        `json:"wasStandardActive,omitempty"`
		WasTopMatch                bool        `json:"wasTopMatch,omitempty"`
	} `json:"events"`
	HasMore     bool `json:"hasMore"`
	SportObject struct {
		ID               int    `json:"id"`
		Name             string `json:"name"`
		NameTranslations []struct {
			Locale      string `json:"locale"`
			Translation string `json:"translation"`
		} `json:"nameTranslations"`
		SportNameSlugs []struct {
			Locale string `json:"locale"`
			Slug   string `json:"slug"`
		} `json:"sportNameSlugs"`
		IconName string `json:"iconName"`
	} `json:"sportObject"`
	FirstEventsStartTime time.Time `json:"firstEventsStartTime"`
	SelectedFilter       string    `json:"selectedFilter"`
	IsOldDesign          bool      `json:"isOldDesign"`
}
