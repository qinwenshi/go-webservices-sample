package categories

import (
	"fmt"
	"net/url"

	"github.com/sample/go-webservices/getyourguide"
)

type Client struct {
	B getyourguide.Backend
}

func List(params *getyourguide.ListCategoriesParams) (*getyourguide.CategoriesResp, error) {
	return getC().List(params)
}

func (c Client) List(params *getyourguide.ListCategoriesParams) (resp *getyourguide.CategoriesResp, err error) {
	v := url.Values{}

	v.Add("offset", fmt.Sprintf("%d", params.Offset))
	v.Add("limit", fmt.Sprintf("%d", params.Limit))
	v.Add("cnt_language", params.Language)
	v.Add("currency", params.Currency)

	err = c.B.Call("GET", "/categories", &v, nil, &resp)
	return resp, err
}

func getC() Client {
	return Client{getyourguide.GetBackend(getyourguide.PublicBackend)}
}
