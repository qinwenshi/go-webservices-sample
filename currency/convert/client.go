package convert

import (
	"fmt"
	"net/url"

	"github.com/sample/go-webservices/currency"
)

type Client struct {
	B currency.Backend
}

func Convert(from, to string, amount float64) (*currency.ConvertResp, error) {
	return getC().Convert(from, to, amount)
}

func (c Client) Convert(from, to string, amount float64) (resp *currency.ConvertResp, err error) {
	v := url.Values{}

	v.Add("from", from)
	v.Add("to", to)
	v.Add("amount", fmt.Sprintf("%+v", amount))

	path := "/currency/convert"
	err = c.B.Call("GET", path, &v, nil, &resp)
	return resp, err
}

func getC() Client {
	return Client{currency.GetBackend(currency.PublicBackend)}
}
