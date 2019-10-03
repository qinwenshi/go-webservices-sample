package bemyguest2

import "fmt"

type ErrorResp struct {
	Content Error `json:"error"`
}

type Error struct {
	Code     string `json:"code"`
	HTTPCode int    `json:"http_code"`
	Message  string `json:"message"`
}

func (e ErrorResp) Error() string {
	return fmt.Sprintf("code: %s, status_code: %d, message: %s", e.Content.Code, e.Content.HTTPCode, e.Content.Message)
}