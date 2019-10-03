package productcatalog

import (
	"time"
)

type Package struct {
	Id                    string     `json:"id"`
	Name                  string     `json:"name"`
	Description           *string    `json:"description"`
	ProductId             string     `json:"product_id"`
	IsPublished           bool       `json:"is_published"`
	PublishedAt           *time.Time `json:"published_at"`
	IsAvailable           bool       `json:"is_available"`
	AvailableFrom         *time.Time `json:"available_from"`
	AvailableUntil        *time.Time `json:"available_until"`
	IsRequireDate         bool       `json:"is_require_date"`
	IsRequireTimeslot     bool       `json:"is_require_timeslot"`
	EarliestAvailableDate string     `json:"earliest_available_date"`
	PurchasedCount        int        `json:"purchased_count"`
	Price                 *float64   `json:"price"`
	Currency              *string    `json:"currency"`
	//Question              *Question  `json:"question"`
	Provider    string     `json:"provider"`
	ReferenceId string     `json:"reference_id"`
	IsDeleted   bool       `json:"is_deleted"`
	DeletedAt   *time.Time `json:"deleted_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type PackagesResp struct {
	PaginatedResp
	Packages []Package `json:"data"`
}

type ListPackageItemsParams struct {
	//ProductId string
	PackageId string
	PaginationParams
}

type ListPackageAvailabilitiesParams struct {
	//ProductId string
	PackageId string
	StartDate time.Time
	EndDate   time.Time
	PaginationParams
}

type PackageAvailabilitiesResp struct {
	PaginatedResp
	Availabilities []Availability `json:"data"`
}
