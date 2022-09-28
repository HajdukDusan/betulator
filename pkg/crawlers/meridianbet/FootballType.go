package meridianbet

import "time"

type Football struct {
	Events []struct {
		ID               int    `json:"id"`
		Name             string `json:"name"`
		NameTranslations []struct {
			Locale      string `json:"locale"`
			Translation string `json:"translation"`
		} `json:"nameTranslations"`
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
			MarketShort []struct {
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
			} `json:"marketShort"`
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
			TeamHomeTranslations []struct {
				Locale      string `json:"locale"`
				Translation string `json:"translation"`
			} `json:"teamHomeTranslations"`
			TeamAwayTranslations []struct {
				Locale      string `json:"locale"`
				Translation string `json:"translation"`
			} `json:"teamAwayTranslations"`
			RegionNameSlugs []struct {
				Locale string `json:"locale"`
				Slug   string `json:"slug"`
			} `json:"regionNameSlugs"`
			SportNameSlugs []struct {
				Locale string `json:"locale"`
				Slug   string `json:"slug"`
			} `json:"sportNameSlugs"`
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
			} `json:"betradarUnified"`
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
			MostPlayedPriority         int         `json:"mostPlayedPriority"`
			NumberOfActiveSelections   int         `json:"numberOfActiveSelections"`
			CreatedAt                  time.Time   `json:"createdAt"`
			UpdatedAt                  time.Time   `json:"updatedAt"`
			BetradarEventID            interface{} `json:"betradarEventId"`
			CurrentMinute              interface{} `json:"currentMinute"`
			Inc                        int         `json:"inc"`
			IsLowPriorityChange        bool        `json:"isLowPriorityChange"`
			IsNewEvent                 bool        `json:"isNewEvent"`
			OldState                   string      `json:"oldState"`
			SkipCreatingLeagueNameSlug bool        `json:"skipCreatingLeagueNameSlug"`
			SkipCreatingRegionNameSlug bool        `json:"skipCreatingRegionNameSlug"`
			SkipCreatingSportNameSlug  bool        `json:"skipCreatingSportNameSlug"`
			WasStandardActive          bool        `json:"wasStandardActive"`
			WasTopMatch                bool        `json:"wasTopMatch"`
		} `json:"events"`
		NumberOfEvents int  `json:"numberOfEvents"`
		IsOldDesign    bool `json:"isOldDesign,omitempty"`
	} `json:"events"`
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
