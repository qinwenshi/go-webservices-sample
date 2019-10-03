package getyourguide

type GetTourParams struct {
	TourId   string
	Language string
	Currency string
}

type TourResp struct {
	Meta Meta         `json:"_metadata"`
	Data TourRespData `json:"data"`
}

type TourRespData struct {
	Tours []Tour `json:"tours"`
}

type Tour struct {
	TourID          int        `json:"tour_id"`
	TourCode        string     `json:"tour_code"`
	CondLanguage    []string   `json:"cond_language"`
	Title           string     `json:"title"`
	Abstract        string     `json:"abstract"`
	Bestseller      bool       `json:"bestseller"`
	Certified       bool       `json:"certified"`
	OverallRating   int        `json:"overall_rating"`
	NumberOfRatings int        `json:"number_of_ratings"`
	Pictures        []Picture  `json:"pictures"`
	Coordinates     Coordinate `json:"coordinates"`
	Price           Price      `json:"price"`
	Categories      []Category `json:"categories"`
	Locations       []Location `json:"locations"`
	URL             string     `json:"url"`
	Durations       []Duration `json:"durations"`
}

type Picture struct {
	ID       int    `json:"id"`
	URL      string `json:"url"`
	SslURL   string `json:"ssl_url"`
	Verified bool   `json:"verified"`
}

type Coordinate struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Viewport struct {
	SwLat  float64 `json:"sw_lat"`
	SwLong float64 `json:"sw_long"`
	NeLat  float64 `json:"ne_lat"`
	NeLong float64 `json:"ne_long"`
}

type Location struct {
	LocationID  int        `json:"location_id"`
	Type        string     `json:"type"`
	Name        string     `json:"name"`
	EnglishName string     `json:"english_name"`
	City        string     `json:"city"`
	Country     string     `json:"country"`
	Coordinates Coordinate `json:"coordinates"`
	Viewport    Viewport   `json:"viewport"`
	ParentId    int        `json:"parent_id"`
}

type Duration struct {
	Duration int    `json:"duration"`
	Unit     string `json:"unit"`
}

type Values struct {
	Amount float32 `json:"amount"`
}

type Price struct {
	Values      Values `json:"values"`
	Description string `json:"description"`
}
