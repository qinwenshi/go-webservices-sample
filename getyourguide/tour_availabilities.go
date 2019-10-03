package getyourguide


type ListTourAvailabilitiesParams struct {
	GetTourParams

	// offsets represent offset of paging
	// default 0
	Offset int

	// how many results per response
	// default 10
	Limit int
}

type TourAvailabilitiesResp struct {
	Meta Meta          `json:"_metadata"`
	Data TourAvailabilitiesRespData `json:"data"`
}

type TourAvailabilitiesRespData struct {
	TourAvailabilities []TourAvailability `json:"availabilities"`
}

type TourAvailability struct {
	StartTime string `json:"start_time"`
	PricingID int    `json:"pricing_id"`
	Vacancies int    `json:"vacancies"`
}