package bemyguest2

type ProductTypesParams struct {
	// uuid of type of product
	Id string
}

func (p *ProductTypesParams) SetId(id string) {
	p.Id = id
}

type ProductTypesResp struct {
	Data []struct {
		UUID  string `json:"uuid"`
		Title string `json:"title"`
		Links []Link `json:"links"`
	} `json:"data"`
}

type ProductTypeResp struct {
	Data ProductType `json:"data"`
}

type ProductType struct {
	UUID                               string        `json:"uuid"`
	Title                              string        `json:"title"`
	TitleTranslated                    string        `json:"titleTranslated"`
	Description                        string        `json:"description"`
	DescriptionTranslated              string        `json:"descriptionTranslated"`
	DurationDays                       int           `json:"durationDays"`
	DurationHours                      int           `json:"durationHours"`
	DurationMinutes                    int           `json:"durationMinutes"`
	DaysInAdvance                      *int          `json:"daysInAdvance"`
	CutOffTime                         interface{}   `json:"cutOffTime"`
	IsNonRefundable                    bool          `json:"isNonRefundable"`
	MinPax                             int           `json:"minPax"`
	MaxPax                             int           `json:"maxPax"`
	MinAdultAge                        int           `json:"minAdultAge"`
	MaxAdultAge                        int           `json:"maxAdultAge"`
	HasChildPrice                      bool          `json:"hasChildPrice"`
	AllowChildren                      bool          `json:"allowChildren"`
	MinChildren                        int           `json:"minChildren"`
	MaxChildren                        int           `json:"maxChildren"`
	MinChildAge                        int           `json:"minChildAge"`
	MaxChildAge                        int           `json:"maxChildAge"`
	AllowSeniors                       bool          `json:"allowSeniors"`
	MinSeniors                         *int          `json:"minSeniors"`
	MaxSeniors                         *int          `json:"maxSeniors"`
	MinSeniorAge                       *int          `json:"minSeniorAge"`
	MaxSeniorAge                       *int          `json:"maxSeniorAge"`
	AllowInfant                        bool          `json:"allowInfant"`
	MinInfantAge                       *int          `json:"minInfantAge"`
	MaxInfantAge                       *int          `json:"maxInfantAge"`
	MaxGroup                           *int          `json:"maxGroup"`
	MinGroup                           *int          `json:"minGroup"`
	InstantConfirmation                bool          `json:"instantConfirmation"`
	NonInstantVoucher                  bool          `json:"nonInstantVoucher"`
	DirectAdmission                    bool          `json:"directAdmission"`
	VoucherUse                         string        `json:"voucherUse"`
	VoucherUseTranslated               string        `json:"voucherUseTranslated"`
	VoucherRedemptionAddress           string        `json:"voucherRedemptionAddress"`
	VoucherRedemptionAddressTranslated string        `json:"voucherRedemptionAddressTranslated"`
	VoucherRequiresPrinting            bool          `json:"voucherRequiresPrinting"`
	MeetingTime                        string        `json:"meetingTime"`
	MeetingAddress                     string        `json:"meetingAddress"`
	MeetingLocation                    string        `json:"meetingLocation"`
	MeetingLocationTranslated          string        `json:"meetingLocationTranslated"`
	CancellationPolicies               []interface{} `json:"cancellationPolicies"`
	RecommendedMarkup                  *float64      `json:"recommendedMarkup"`
	ChildRecommendedMarkup             *float64      `json:"childRecommendedMarkup"`
	SeniorRecommendedMarkup            *float64      `json:"seniorRecommendedMarkup"`
	AdultParityPrice                   *float64      `json:"adultParityPrice"`
	ChildParityPrice                   *float64      `json:"childParityPrice"`
	SeniorParityPrice                  *float64      `json:"seniorParityPrice"`
	Validity                           Validity      `json:"validity"`
	Timeslots                          []Timeslot    `json:"timeslots"`
	Options                            struct {
		PerBooking []PerBooking `json:"perBooking"`
		PerPax     []PerPax     `json:"perPax"`
	} `json:"options"`
	Links []Link `json:"links"`
}
