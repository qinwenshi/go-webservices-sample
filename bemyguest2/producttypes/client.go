package producttypes

import (
	"fmt"
	"net/url"

	"github.com/sample/go-webservices/bemyguest2"
)

type Client struct {
	B bemyguest2.Backend
}

func Get(params *bemyguest2.ProductTypesParams) (*bemyguest2.ProductTypeResp, error) {
	return getC().Get(params)
}

func ListPrices(params *bemyguest2.ListPricesParams) (*bemyguest2.ProductTypePricesResp, error) {
	return getC().ListPrices(params)
}

func GetPriceByDate(params *bemyguest2.GetPriceByDateParams) (*bemyguest2.ProductTypePricesByDateResp, error) {
	return getC().GetPriceByDate(params)
}

func (c Client) Get(params *bemyguest2.ProductTypesParams) (resp *bemyguest2.ProductTypeResp, err error) {
	path := "/v2/product-types/" + params.Id
	err = c.B.Call("GET", path, nil, nil, &resp)
	return resp, err
}

func (c Client) ListPrices(params *bemyguest2.ListPricesParams) (resp *bemyguest2.ProductTypePricesResp, err error) {
	v := url.Values{}

	if params.DateStart != nil {
		v.Add("date_start", (*params.DateStart).Format("2006-01-02"))
	}
	if params.DateEnd != nil {
		v.Add("date_end", (*params.DateEnd).Format("2006-01-02"))
	}

	path := fmt.Sprintf("/v2/product-types/%s/price-lists", params.Id)
	err = c.B.Call("GET", path, &v, nil, &resp)
	return resp, err
}

func (c Client) GetPriceByDate(params *bemyguest2.GetPriceByDateParams) (resp *bemyguest2.ProductTypePricesByDateResp, err error) {
	path := fmt.Sprintf("/v2/product-types/%s/price-lists/%s", params.Id, params.Date.Format("2006-01-02"))
	err = c.B.Call("GET", path, nil, nil, &resp)
	return resp, err
}

func getC() Client {
	return Client{bemyguest2.GetBackend(bemyguest2.PublicBackend)}
}
