package getyourguide

import "time"

type Meta struct {
		Descriptor         string        `json:"descriptor"`
		Method             string        `json:"method"`
		Date               time.Time     `json:"date"`
		Status             string        `json:"status"`
		Query              string        `json:"query"`
		AvailableLanguages []string      `json:"availableLanguages"`
		Exchange           Exchange      `json:"exchange"`
		TotalCount int `json:"totalCount"`
		Limit      int `json:"limit"`
		Offset     int `json:"offset"`
}

type Exchange           struct {
	Rate     float64 `json:"rate"`
	Currency string  `json:"currency"`
}