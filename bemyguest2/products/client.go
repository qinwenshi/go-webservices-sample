package products

import (
	"fmt"
	"net/url"

	"github.com/sample/go-webservices/bemyguest2"
)

type Client struct {
	B bemyguest2.Backend
}

func Get(params *bemyguest2.ProductParams) (*bemyguest2.ProductResp, error) {
	return getC().Get(params)
}

func List(params *bemyguest2.ProductsParams) (*bemyguest2.ProductsResp, error) {
	return getC().List(params)
}

func ListProductTypes(params *bemyguest2.ProductTypesParams) (*bemyguest2.ProductTypesResp, error) {
	return getC().ListProductTypes(params)
}

func (c Client) Get(params *bemyguest2.ProductParams) (resp *bemyguest2.ProductResp, err error) {
	path := "/v2/products/" + params.Id
	err = c.B.Call("GET", path, nil, nil, &resp)
	return resp, err
}

func (c Client) List(params *bemyguest2.ProductsParams) (resp *bemyguest2.ProductsResp, err error) {
	v := url.Values{}

	v.Add("page", fmt.Sprintf("%d", params.Page))
	v.Add("per_page", fmt.Sprintf("%d", params.PerPage))

	if params.Type != nil {
		v.Add("type", *params.Type)
	}
	if params.Country != nil {
		v.Add("country", *params.Country)
	}
	if params.City != nil {
		v.Add("city", *params.City)
	}
	if params.PriceMin != nil {
		v.Add("price_min", fmt.Sprintf("%v", *params.PriceMin))
	}
	if params.PriceMax != nil {
		v.Add("price_max", fmt.Sprintf("%v", *params.PriceMax))
	}
	if params.Category != nil {
		v.Add("category", *params.Category)
	}
	if params.Pax != nil {
		v.Add("pax", fmt.Sprintf("%v", *params.Pax))
	}
	if params.DurationDaysMin != nil {
		v.Add("duration_days_min", fmt.Sprintf("%v", *params.DurationDaysMin))
	}
	if params.DurationDaysMax != nil {
		v.Add("duration_days_max", fmt.Sprintf("%v", *params.DurationDaysMax))
	}

	err = c.B.Call("GET", "/v2/products", &v, nil, &resp)
	return resp, err
}

func (c Client) ListProductTypes(params *bemyguest2.ProductTypesParams) (resp *bemyguest2.ProductTypesResp, err error) {
	path := fmt.Sprintf("/v2/products/%s/product-types", params.Id)
	err = c.B.Call("GET", path, nil, nil, &resp)
	return resp, err
}

func getC() Client {
	return Client{bemyguest2.GetBackend(bemyguest2.PublicBackend)}
}
