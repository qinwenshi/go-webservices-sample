package getyourguide

type ListCategoriesParams struct {
	Language string
	Currency string

	// offsets represent offset of paging
	// default 0
	Offset int

	// how many results per response
	// default 10
	Limit int
}

type CategoriesResp struct {
	Meta Meta               `json:"_metadata"`
	Data CategoriesRespData `json:"data"`
}

type CategoriesRespData struct {
	Categories []Category `json:"categories"`
}
