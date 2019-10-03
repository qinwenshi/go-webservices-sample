package getyourguide

type ListToursParams struct {
	Language string
	Currency string

	// IATA airport code
	// https://en.wikipedia.org/wiki/IATA_airport_code
	IATA string

	// offsets represent offset of paging
	// default 0
	Offset int

	// how many results per response
	// default 10
	Limit int
}

type ToursResp struct {
	Meta Meta          `json:"_metadata"`
	Data ToursRespData `json:"data"`
}

type ToursRespData struct {
	Tours []Tour `json:"tours"`
}
