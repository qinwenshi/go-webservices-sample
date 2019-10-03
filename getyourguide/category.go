package getyourguide

type GetCategoryParams struct {
	CategoryId string
	Language   string
	Currency   string
}

type CategoryResp struct {
	Meta Meta             `json:"_metadata"`
	Data CategoryRespData `json:"data"`
}

type CategoryRespData struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	CategoryID    int    `json:"category_id"`
	Name          string `json:"name"`
	NumberOfTours int    `json:"number_of_tours"`
	ParentId      int    `json:"parent_id"`
}
