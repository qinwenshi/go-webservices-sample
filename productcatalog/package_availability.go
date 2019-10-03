package productcatalog

import (
	"time"
)

type Availability struct {
	Id          string     `json:"id"`
	ProductId   string     `json:"product_id"`
	PackageId   string     `json:"package_id"`
	Date        string     `json:"date"`
	IsAvailable bool       `json:"is_available"`
	Timeslots   []Timeslot `json:"timeslots"`
	IsDeleted   bool       `json:"is_deleted"`
	DeletedAt   *time.Time `json:"deleted_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Timeslot struct {
	StartAt     string  `json:"start_at"`
	EndAt       *string `json:"end_at"`
	IsAvailable bool    `json:"is_available"`
	ReferenceId string  `json:"reference_id"`
}
