package getyourguide

type ListOptionPricingsParams struct {
	GetOptionParams

	// offsets represent offset of paging
	// default 0
	Offset int

	// how many results per response
	// default 10
	Limit int
}

type OptionPricingsResp struct {
	Meta Meta                   `json:"_metadata"`
	Data OptionPricingsRespData `json:"data"`
}

type OptionPricingsRespData struct {
	OptionPricings []OptionPricing `json:"pricing"`
}

type OptionPricing struct {
	PricingID                int             `json:"pricing_id"`
	TotalMinimumParticipants int             `json:"total_minimum_participants"`
	Categories               []PriceCategory `json:"categories"`
}

type PriceCategory struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	MinAge     int     `json:"min_age"`
	MaxAge     int     `json:"max_age"`
	StandAlone bool    `json:"stand_alone"`
	Addon      bool    `json:"addon"`
	Scales     []Scale `json:"scale"`
}

type Scale struct {
	MinParticipants int     `json:"min_participants"`
	MaxParticipants int     `json:"max_participants"`
	RetailPrice     float64 `json:"retail_price"`
	NetPrice        float64 `json:"net_price"`
	Type            string  `json:"type"`
}
