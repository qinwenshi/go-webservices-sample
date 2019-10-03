package getyourguide

type ListLocationsParams struct {
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

type LocationsResp struct {
	Meta Meta              `json:"_metadata"`
	Data LocationsRespData `json:"data"`
}

type LocationsRespData struct {
	Locations []Location `json:"locations"`
}
