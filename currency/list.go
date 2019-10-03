package currency

import "time"

type ListResp struct {
	PaginatedResp
	Currency []Currency `json:"data"`
}

type Currency struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}
