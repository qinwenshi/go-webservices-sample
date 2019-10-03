package getyourguide

import (
	"fmt"
	"time"
)

type ErrorResp struct {
	Descriptor string    `json:"descriptor"`
	APIVersion string    `json:"apiVersion"`
	Method     string    `json:"method"`
	Date       time.Time `json:"date"`
	Status     string    `json:"status"`
	Errors     []Error   `json:"errors"`
	HelpURL    string    `json:"helpURL"`
}

type Error struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func (e ErrorResp) Error() string {
	message := fmt.Sprintf("descriptor: %s, api_version: %s, method: %s, date: %v, status: %s, help_url: %s", e.Descriptor, e.APIVersion, e.Method, e.Date, e.Status, e.HelpURL)

	for index, err := range e.Errors {
		message += fmt.Sprintf(" err_%d_code: %d, err_%d_message: %s", index, err.ErrorCode, index, err.ErrorMessage)
	}

	return message
}
