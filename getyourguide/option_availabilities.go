package getyourguide

import (
	"time"
)

type ListOptionAvailabilitiesParams struct {
	GetOptionParams

	// offsets represent offset of paging
	// default 0
	Offset int

	// how many results per response
	// default 10
	Limit int

	Date *time.Time
}

type OptionAvailabilitiesResp struct {
	Meta Meta                         `json:"_metadata"`
	Data OptionAvailabilitiesRespData `json:"data"`
}

type OptionAvailabilitiesRespData struct {
	OptionAvailabilities []OptionAvailability `json:"availabilities"`
}

type OptionAvailability struct {
	StartTime string `json:"start_time"`
	PricingID int    `json:"pricing_id"`
	Vacancies int    `json:"vacancies"`
}
