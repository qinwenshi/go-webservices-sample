package productcatalog

import "fmt"

type ErrorResp struct {
	ErrorMessage string `json:"error"`
}

func (e ErrorResp) Error() string {
	return fmt.Sprintf("error: %s", e.ErrorMessage)
}
