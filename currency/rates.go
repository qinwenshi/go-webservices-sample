package currency

import "time"

type RatesResp struct {
	PaginatedResp
	Rates []Rate `json:"data"`
}

type Rate struct {
	Id        string                 `json:"id"`
	Base      string                 `json:"base"`
	Date      string                 `json:"date"`
	Rates     map[string]interface{} `json:"rates"`
	CreatedAt *time.Time             `json:"created_at"`
	UpdatedAt *time.Time             `json:"updated_at"`
}
