package productcatalog

type PaginatedResp struct {
	Limit int  `json:"limit"`
	Total int  `json:"total"`
	Count int  `json:"count"`
	Next  Next `json:"next"`
}

type Next struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

type PaginationParams struct {
	Skip    int
	Limit   int
	OrderBy string
}
