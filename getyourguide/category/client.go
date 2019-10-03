package category

import (
	"net/url"

	"github.com/sample/go-webservices/getyourguide"
)

type Client struct {
	B getyourguide.Backend
}

func Get(params *getyourguide.GetCategoryParams) (*getyourguide.CategoryResp, error) {
	return getC().Get(params)
}

func (c Client) Get(params *getyourguide.GetCategoryParams) (resp *getyourguide.CategoryResp, err error) {
	v := url.Values{}

	v.Add("cnt_language", params.Language)
	v.Add("currency", params.Currency)

	path := "/categories/" + params.CategoryId

	err = c.B.Call("GET", path, &v, nil, &resp)
	return resp, err
}

func getC() Client {
	return Client{getyourguide.GetBackend(getyourguide.PublicBackend)}
}
