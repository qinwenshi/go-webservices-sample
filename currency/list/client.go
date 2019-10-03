package list

import (
	"fmt"
	"net/url"

	"github.com/sample/go-webservices/currency"
)

type Client struct {
	B currency.Backend
}

func Get(id string) (*currency.Currency, error) {
	return getC().Get(id)
}

func List(params *currency.PaginationParams) (*currency.ListResp, error) {
	return getC().List(params)
}

func (c Client) Get(id string) (resp *currency.Currency, err error) {
	err = c.B.Call("GET", "/currencies/"+id, nil, nil, &resp)
	return resp, err
}

func (c Client) List(params *currency.PaginationParams) (resp *currency.ListResp, err error) {
	v := url.Values{}

	v.Add("skip", fmt.Sprintf("%d", params.Skip))
	v.Add("limit", fmt.Sprintf("%d", params.Limit))
	v.Add("order_by", fmt.Sprintf("%s", params.OrderBy))

	path := "/currencies"
	err = c.B.Call("GET", path, &v, nil, &resp)
	return resp, err
}

func getC() Client {
	return Client{currency.GetBackend(currency.PublicBackend)}
}
