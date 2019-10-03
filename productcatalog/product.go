package productcatalog

import (
	"time"
)

type Product struct {
	Id                    string     `json:"id"`
	Type                  string     `json:"type"`
	Name                  string     `json:"name"`
	Description           *string    `json:"description"`
	IsPublished           bool       `json:"-"`
	PublishedAt           *time.Time `json:"published_at,omitempty"`
	IsAvailable           bool       `json:"is_available"`
	IsRequireDate         *bool      `json:"is_require_date"`
	IsRequireTimeslot     *bool      `json:"is_require_timeslot"`
	IsRequirePaperVoucher *bool      `json:"is_require_paper_voucher"`
	HasHotelPickup        *bool      `json:"has_hotel_pickup"`
	HasAirportPickup      *bool      `json:"has_airport_pickup"`
	ConfirmationPeriod    *int       `json:"confirmation_period"`
	Duration              *int       `json:"duration"`
	Photos                []string   `json:"photos"`
	CategoryIds           []string   `json:"category_ids"`
	TagIds                []string   `json:"tag_ids"`
	Location              *Location  `json:"location,omitempty"`
	City                  *string    `json:"city"`
	EarliestAvailableDate *string    `json:"earliest_available_date,omitempty"`
	PurchasedCount        int        `json:"purchased_count"`
	Price                 *float64   `json:"price,omitempty"`
	Currency              *string    `json:"currency,omitempty"`
	//Question              *Question  `json:"question"`
	Provider    string     `json:"provider"`
	ReferenceId string     `json:"reference_id"`
	IsDeleted   bool       `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ProductsResp struct {
	PaginatedResp
	Products []Product `json:"data"`
}

type ListProductPackagesParams struct {
	ProductId string
	PaginationParams
}

type ListProductItemParams struct {
	ProductId string
	PaginationParams
}
