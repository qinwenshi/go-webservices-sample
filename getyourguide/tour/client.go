package tour

import (
	"fmt"
	"net/url"

	"github.com/sample/go-webservices/getyourguide"
)

type Client struct {
	B getyourguide.Backend
}

func Get(params *getyourguide.GetTourParams) (*getyourguide.TourResp, error) {
	return getC().Get(params)
}

func ListAvailabilities(params *getyourguide.ListTourAvailabilitiesParams) (*getyourguide.TourAvailabilitiesResp, error) {
	return getC().ListAvailabilities(params)
}

func ListOptions(params *getyourguide.ListTourOptionsParams) (*getyourguide.TourOptionsResp, error) {
	return getC().ListOptions(params)
}

func (c Client) Get(params *getyourguide.GetTourParams) (resp *getyourguide.TourResp, err error) {
	v := url.Values{}

	v.Add("cnt_language", params.Language)
	v.Add("currency", params.Currency)

	path := "/tours/" + params.TourId

	err = c.B.Call("GET", path, &v, nil, &resp)
	return resp, err
}

func (c Client) ListAvailabilities(params *getyourguide.ListTourAvailabilitiesParams) (resp *getyourguide.TourAvailabilitiesResp, err error) {
	v := url.Values{}

	v.Add("cnt_language", params.Language)
	v.Add("currency", params.Currency)

	path := fmt.Sprintf("/tours/%s/availabilities", params.TourId)

	err = c.B.Call("GET", path, &v, nil, &resp)
	return resp, err
}

func (c Client) ListOptions(params *getyourguide.ListTourOptionsParams) (resp *getyourguide.TourOptionsResp, err error) {
	v := url.Values{}

	v.Add("cnt_language", params.Language)
	v.Add("currency", params.Currency)

	path := fmt.Sprintf("/tours/%s/options", params.TourId)

	err = c.B.Call("GET", path, &v, nil, &resp)
	return resp, err
}

func getC() Client {
	return Client{getyourguide.GetBackend(getyourguide.PublicBackend)}
}
