package tours

import (
	"fmt"
	"net/url"

	"github.com/sample/go-webservices/getyourguide"
)

type Client struct {
	B getyourguide.Backend
}

func List(params *getyourguide.ListToursParams) (*getyourguide.ToursResp, error) {
	return getC().List(params)
}

func (c Client) List(params *getyourguide.ListToursParams) (resp *getyourguide.ToursResp, err error) {
	v := url.Values{}

	v.Add("offset", fmt.Sprintf("%d", params.Offset))
	v.Add("limit", fmt.Sprintf("%d", params.Limit))
	v.Add("cnt_language", params.Language)
	v.Add("currency", params.Currency)

	err = c.B.Call("GET", "/tours", &v, nil, &resp)
	return resp, err
}

func getC() Client {
	return Client{getyourguide.GetBackend(getyourguide.PublicBackend)}
}
