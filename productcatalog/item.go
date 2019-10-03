package productcatalog

import (
	"time"
)

type Item struct {
	Id             string     `json:"id"`
	Name           string     `json:"name"`
	IsPrimary      bool       `json:"is_primary"`
	Description    *string    `json:"description"`
	IsPublished    bool       `json:"is_published"`
	PublishedAt    *time.Time `json:"published_at"`
	IsAvailable    bool       `json:"is_available"`
	PurchasedCount int        `json:"purchased_count"`
	Provider       string     `json:"provider"`
	ReferenceId    string     `json:"reference_id"`
	ProductId      string     `json:"product_id"`
	PackageIds     []string   `json:"package_ids"`
	//Question       *Question  `json:"question"`
	IsDeleted bool       `json:"is_deleted"`
	DeletedAt *time.Time `json:"deleted_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type ItemsResp struct {
	PaginatedResp
	Items []Item `json:"data"`
}
