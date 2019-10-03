package bemyguest2

type Item struct {
	Label string   `json:"label"`
	Value string   `json:"value"`
	Price *float64 `json:"price"`
}

type PerPax struct {
	UUID        string   `json:"uuid"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Required    bool     `json:"required"`
	FormatRegex *string  `json:"formatRegex"`
	InputType   int      `json:"inputType"`
	ValidFrom   *string  `json:"validFrom"`
	ValidTo     *string  `json:"validTo"`
	Price       *float64 `json:"price"`
	Items       []Item   `json:"items,omitempty"`
}

type PerBooking struct {
	UUID        string   `json:"uuid"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Required    bool     `json:"required"`
	FormatRegex *string  `json:"formatRegex"`
	InputType   int      `json:"inputType"`
	ValidFrom   *string  `json:"validFrom"`
	ValidTo     *string  `json:"validTo"`
	Price       *float64 `json:"price"`
	Items       []Item   `json:"items,omitempty"`
}

type Validity struct {
	Type string  `json:"type"`
	Days *int    `json:"days"`
	Date *string `json:"date"`
}

type Timeslot struct {
	UUID      string `json:"uuid"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
