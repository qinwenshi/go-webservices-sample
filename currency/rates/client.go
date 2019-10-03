package rates

import (
	"fmt"
	"net/url"

	"github.com/sample/go-webservices/currency"
)

type Client struct {
	B currency.Backend
}

func ListRates(id string, params *currency.PaginationParams) (*currency.RatesResp, error) {
	return getC().ListRates(id, params)
}

func (c Client) ListRates(id string, params *currency.PaginationParams) (resp *currency.RatesResp, err error) {
	v := url.Values{}

	v.Add("skip", fmt.Sprintf("%d", params.Skip))
	v.Add("limit", fmt.Sprintf("%d", params.Limit))
	v.Add("order_by", fmt.Sprintf("%s", params.OrderBy))

	err = c.B.Call("GET", "/currencies/"+id+"/rates", &v, nil, &resp)
	return resp, err
}

func getC() Client {
	return Client{currency.GetBackend(currency.PublicBackend)}
}
