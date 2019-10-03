package getyourguide

type GetOptionParams struct {
	OptionId string
	Language string
	Currency string
}

type OptionResp struct {
	Meta Meta           `json:"_metadata"`
	Data OptionRespData `json:"data"`
}

type OptionRespData struct {
	TourOptions []TourOption `json:"tour_options"`
}

type TourOption struct {
	OptionID          int                `json:"option_id"`
	TourID            int                `json:"tour_id"`
	Title             string             `json:"title"`
	MeetingPoint      string             `json:"meeting_point"`
	DropOff           string             `json:"drop_off"`
	Duration          int                `json:"duration"`
	DurationUnit      string             `json:"duration_unit"`
	CondLanguage      CondLanguage       `json:"cond_language"`
	BookingParameters []BookingParameter `json:"booking_parameter"`
	Service           Service            `json:"services"`
	CoordinateType    string             `json:"coordinate_type"`
	Coordinate        Coordinate         `json:"coordinates"`
	Price             Price              `json:"price"`
	FreeSale          bool               `json:"free_sale"`
}

type CondLanguage struct {
	LanguageAudio   []interface{} `json:"language_audio"`
	LanguageBooklet []interface{} `json:"language_booklet"`
	LanguageLive    []string      `json:"language_live"`
}

type BookingParameter struct {
	Name        string `json:"name"`
	Mandatory   bool   `json:"mandatory"`
	Description string `json:"description,omitempty"`
}

type Service struct {
	Prt bool `json:"prt"`
	Stl bool `json:"stl"`
	Wlc bool `json:"wlc"`
}
