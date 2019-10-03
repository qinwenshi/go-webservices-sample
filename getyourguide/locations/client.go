package locations

import (
	"fmt"
	"net/url"

	"github.com/sample/go-webservices/getyourguide"
)

type Client struct {
	B getyourguide.Backend
}

func List(params *getyourguide.ListLocationsParams) (*getyourguide.LocationsResp, error) {
	return getC().List(params)
}

func (c Client) List(params *getyourguide.ListLocationsParams) (resp *getyourguide.LocationsResp, err error) {
	v := url.Values{}

	v.Add("offset", fmt.Sprintf("%d", params.Offset))
	v.Add("limit", fmt.Sprintf("%d", params.Limit))
	v.Add("cnt_language", params.Language)
	v.Add("currency", params.Currency)

	err = c.B.Call("GET", "/locations", &v, nil, &resp)
	return resp, err
}

func getC() Client {
	return Client{getyourguide.GetBackend(getyourguide.PublicBackend)}
}
