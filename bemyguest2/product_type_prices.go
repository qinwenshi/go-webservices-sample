package bemyguest2

import (
	"encoding/json"
	"time"
)

type ListPricesParams struct {
	// uuid of type of product
	Id string

	DateStart *time.Time
	DateEnd   *time.Time
}

func (p *ListPricesParams) SetId(id string) {
	p.Id = id
}

type GetPriceByDateParams struct {
	// uuid of type of product
	Id string

	Date time.Time
}

func (p *GetPriceByDateParams) SetId(id string) {
	p.Id = id
}

func (p *GetPriceByDateParams) SetDate(date time.Time) {
	p.Date = date
}

type ProductTypePricesResp struct {
	Data []PriceBrief `json:"data"`
}

type PriceBrief struct {
	Date      string `json:"date"`
	Weekday   string `json:"weekday"`
	Available bool   `json:"available"`
	Currency  string `json:"currency"`
	Prices    struct {
		Adults             PriceByQuantity `json:"adults"`
		Children           PriceByQuantity `json:"children"`
		Seniors            PriceByQuantity `json:"seniors"`
		CancellationPolicy []interface{}   `json:"cancellationPolicy"`
	} `json:"prices"`
	Timeslots []Timeslot `json:"timeslots"`
	Options   []struct {
		UUID     string `json:"uuid"`
		Required bool   `json:"required"`
		Price    int    `json:"price"`
	} `json:"options"`
	Links           []Link `json:"links"`
	VoucherValidity struct {
		Adults struct {
			Date string `json:"date"`
		} `json:"adults"`
		Children struct {
			Date string `json:"date"`
		} `json:"children"`
		Seniors struct {
			Date string `json:"date"`
		} `json:"seniors"`
	} `json:"voucherValidity"`
}

type PriceByQuantity map[string]float64

// UnmarshalJSON handles multi types value e.g. adults / children / seniors
func (p *PriceByQuantity) UnmarshalJSON(b []byte) error {
	var value map[string]float64
	err := json.Unmarshal(b, &value)
	if err != nil {
		return nil
	}

	*p = PriceByQuantity(value)
	return nil
}

type ProductTypePricesByDateResp struct {
	Data PriceBrief `json:"data"`
}
