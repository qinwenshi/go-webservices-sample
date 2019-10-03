package bemyguest2

import "time"

type ProductsParams struct {
	// uuid of type of product
	Type *string
	// uuid of country
	Country *string
	// uuid of city, it will always overwrite country parameter if provided
	City *string
	// minimal price in decimal format 000.00 - it's compared to base price
	PriceMin *float32
	// max price in decimal format 000.00
	PriceMax *float32
	// product duration minimum days (default 0)
	DurationDaysMin *int
	// product duration maximum days (default NULL)
	DurationDaysMax *int
	// page number for results
	Page int
	// how many results per page - if not provided default value from user account will be used
	PerPage int
	// uuid of litsing category
	Category *string
	// number of people
	Pax *int
	// sorting field, example: &sort=date,-price or &sort=price
	Sort *string
	// language uuid, also language code may be provided
	Language *string
	// product's prices start date, format YYYY-MM-DD
	DateStart time.Time
	// product's prices end date, format YYYY-MM-DD
	DateEnd time.Time
	// coma separated list of keys (fields) to be returned
	Fields *string
}

func (p *ProductsParams) SetType(tye string) {
	p.Type = &tye
}

func (p *ProductsParams) SetCountry(country string) {
	p.Country = &country
}

func (p *ProductsParams) SetCity(city string) {
	p.City = &city
}

func (p *ProductsParams) SetDurationDaysMin(min int) {
	p.DurationDaysMin = &min
}

func (p *ProductsParams) SetDurationDaysMax(max int) {
	p.DurationDaysMax = &max
}

func (p *ProductsParams) SetPage(page int) {
	p.Page = page
}

func (p *ProductsParams) SetPerPage(perPage int) {
	p.PerPage = perPage
}

func (p *ProductsParams) SetLanguage(language string) {
	p.Language = &language
}

func (p *ProductsParams) SetDateStart(dateStart time.Time) {
	p.DateStart = dateStart
}

func (p *ProductsParams) SetDateEnd(dateEnd time.Time) {
	p.DateEnd = dateEnd
}

type ProductsResp struct {
	Data []ProductBrief `json:"data"`
	Meta Meta      `json:"meta"`
}

func (r *ProductsResp) HasNext() bool {
	return r.Meta.Pagination.HasNext()
}

func (r *ProductsResp) GetNextPage() int {
	return r.Meta.Pagination.GetNextPage()
}

type Links struct {
	Next string `json:"next"`
	Previous string `json:"previous"`
}

type Link struct {
	Method string `json:"method"`
	Rel    string `json:"rel"`
	Href   string `json:"href"`
}
