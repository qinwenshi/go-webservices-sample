package getyourguide

type ListTourOptionsParams struct {
	GetTourParams

	// offsets represent offset of paging
	// default 0
	Offset int

	// how many results per response
	// default 10
	Limit int
}

type TourOptionsResp struct {
	Meta Meta                `json:"_metadata"`
	Data TourOptionsRespData `json:"data"`
}

type TourOptionsRespData struct {
	TourOptions []TourOption `json:"tour_options"`
}
