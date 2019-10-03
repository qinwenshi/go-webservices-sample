package bemyguest2

import "time"

type ProductParams struct {
	// uuid of type of product
	Id string
	// language uuid, also language code may be provided
	Language string
	// product's prices start date, format YYYY-MM-DD
	DateStart time.Time
	// product's prices end date, format YYYY-MM-DD
	DateEnd time.Time
	// coma separated list of keys (fields) to be returned
	Fields string
}

func (p *ProductParams) SetId(id string) {
	p.Id = id
}

func (p *ProductParams) SetLanguage(language string) {
	p.Language = language
}

func (p *ProductParams) SetDateStart(dateStart time.Time) {
	p.DateStart = dateStart
}

func (p *ProductParams) SetDateEnd(dateEnd time.Time) {
	p.DateEnd = dateEnd
}

type ProductResp struct {
	Data Product `json:"data"`
}

type Product struct {
	UUID                     string                 `json:"uuid"`
	UpdatedAt                string                 `json:"updatedAt"`
	Title                    string                 `json:"title"`
	TitleTranslated          string                 `json:"titleTranslated"`
	Description              string                 `json:"description"`
	DescriptionTranslated    string                 `json:"descriptionTranslated"`
	Highlights               string                 `json:"highlights"`
	HighlightsTranslated     string                 `json:"highlightsTranslated"`
	AdditionalInfo           string                 `json:"additionalInfo"`
	AdditionalInfoTranslated string                 `json:"additionalInfoTranslated"`
	PriceIncludes            string                 `json:"priceIncludes"`
	PriceIncludesTranslated  string                 `json:"priceIncludesTranslated"`
	PriceExcludes            string                 `json:"priceExcludes"`
	PriceExcludesTranslated  string                 `json:"priceExcludesTranslated"`
	ValidFrom                *string            `json:"validFrom"`
	ValidThrough             *string            `json:"validThrough"`
	Itinerary                string                 `json:"itinerary"`
	ItineraryTranslated      string                 `json:"itineraryTranslated"`
	Warnings                 string                 `json:"warnings"`
	WarningsTranslated       string                 `json:"warningsTranslated"`
	Safety                   string                 `json:"safety"`
	SafetyTranslated         string                 `json:"safetyTranslated"`
	Latitude                 string                 `json:"latitude"`
	Longitude                string                 `json:"longitude"`
	MinPax                   int                    `json:"minPax"`
	MaxPax                   int                    `json:"maxPax"`
	BasePrice                float64                `json:"basePrice"`
	Currency                 Currency               `json:"currency"`
	IsFlatPaxPrice           bool                   `json:"isFlatPaxPrice"`
	ReviewCount              int                    `json:"reviewCount"`
	ReviewAverageScore       float64                `json:"reviewAverageScore"`
	TypeName                 string                 `json:"typeName"`
	TypeUUID                 string                 `json:"typeUuid"`
	PhotosURL                string                 `json:"photosUrl"`
	BusinessHoursFrom        string                 `json:"businessHoursFrom"`
	BusinessHoursTo          string                 `json:"businessHoursTo"`
	AverageDelivery          int                    `json:"averageDelivery"`
	HotelPickup              bool                   `json:"hotelPickup"`
	AirportPickup            bool                   `json:"airportPickup"`
	Photos                   []Photo                `json:"photos"`
	Categories               []ProductCategory      `json:"categories"`
	Locations                []ProductLocation      `json:"locations"`
	GuideLanguages           []GuideLanguage        `json:"guideLanguages"`
	AudioHeadsetLanguages    []AudioHeadsetLanguage `json:"audioHeadsetLanguages"`
	WrittenLanguages         []WrittenLanguage      `json:"writtenLanguages"`
	Links                    []Link                 `json:"links"`
}

type Photo struct {
	Caption interface{} `json:"caption"`
	UUID    string      `json:"uuid"`
	Paths   Paths       `json:"paths"`
}

type Paths struct {
	Original    string      `json:"original"`
	Seven5X50   string      `json:"75x50"`
	One75X112   string      `json:"175x112"`
	Six80X325   string      `json:"680x325"`
	One280X720  string      `json:"1280x720"`
	One920X1080 interface{} `json:"1920x1080"`
	Two048X1536 interface{} `json:"2048x1536"`
}

type ProductCategory struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type ProductLocation struct {
	City        string `json:"city"`
	CityUUID    string `json:"cityUuid"`
	State       string `json:"state"`
	StateUUID   string `json:"stateUuid"`
	Country     string `json:"country"`
	CountryUUID string `json:"countryUuid"`
}

type GuideLanguage struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type AudioHeadsetLanguage struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type WrittenLanguage struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}


type ProductBrief struct {
	UUID            string      `json:"uuid"`
	UpdatedAt       string      `json:"updatedAt"`
	Title           string      `json:"title"`
	TitleTranslated string      `json:"titleTranslated"`
	ValidFrom       *string `json:"validFrom"`
	ValidThrough    *string `json:"validThrough"`
	BasePrice       float64     `json:"basePrice"`
	TypeName        string      `json:"typeName"`
	TypeUUID        string      `json:"typeUuid"`
	Links           []Link  `json:"links"`
}