package productcatalog

import (
	"time"
)

type Price struct {
	Id                string            `json:"id"`
	ProductId         string            `json:"product_id"`
	PackageId         string            `json:"package_id"`
	ItemId            string            `json:"item_id"`
	IsAvailable       bool              `json:"is_available"`
	Date              string            `json:"date"`
	StartAt           *string           `json:"start_at"`
	EndAt             *string           `json:"end_at"`
	Currency          string            `json:"currency"`
	PriceByHeadcounts PriceByHeadcounts `json:"price_by_headcounts"`
	IsDeleted         bool              `json:"is_deleted"`
	DeletedAt         *time.Time        `json:"deleted_at"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
}

type PriceByHeadcounts map[string]float64

type ListItemPricesParams struct {
	//ProductId string
	PackageId string
	ItemId    string
	Date      time.Time
	StartAt   string
	PaginationParams
}

type ItemPricesResp struct {
	PaginatedResp
	Prices []Price `json:"data"`
}
