package option

import (
	"fmt"
	"net/url"

	"github.com/sample/go-webservices/getyourguide"
)

const TimeFormat = "2006-01-02T15:04:05"

type Client struct {
	B getyourguide.Backend
}

func Get(params *getyourguide.GetOptionParams) (*getyourguide.OptionResp, error) {
	return getC().Get(params)
}

func ListAvailabilities(params *getyourguide.ListOptionAvailabilitiesParams) (*getyourguide.OptionAvailabilitiesResp, error) {
	return getC().ListAvailabilities(params)
}

func ListPricings(params *getyourguide.ListOptionPricingsParams) (*getyourguide.OptionPricingsResp, error) {
	return getC().ListPricings(params)
}

func (c Client) Get(params *getyourguide.GetOptionParams) (resp *getyourguide.OptionResp, err error) {
	v := url.Values{}

	v.Add("cnt_language", params.Language)
	v.Add("currency", params.Currency)

	path := "/options/" + params.OptionId

	err = c.B.Call("GET", path, &v, nil, &resp)
	return resp, err
}

func (c Client) ListAvailabilities(params *getyourguide.ListOptionAvailabilitiesParams) (resp *getyourguide.OptionAvailabilitiesResp, err error) {
	v := url.Values{}

	v.Add("cnt_language", params.Language)
	v.Add("currency", params.Currency)

	if params.Date != nil {
		startDateStr := params.Date.Format(TimeFormat)
		v.Add("date[]", startDateStr)
	}

	path := fmt.Sprintf("/options/%s/availabilities", params.OptionId)

	err = c.B.Call("GET", path, &v, nil, &resp)
	return resp, err
}

func (c Client) ListPricings(params *getyourguide.ListOptionPricingsParams) (resp *getyourguide.OptionPricingsResp, err error) {
	v := url.Values{}

	v.Add("cnt_language", params.Language)
	v.Add("currency", params.Currency)

	path := fmt.Sprintf("/options/%s/pricings", params.OptionId)

	err = c.B.Call("GET", path, &v, nil, &resp)
	return resp, err
}

func getC() Client {
	return Client{getyourguide.GetBackend(getyourguide.PublicBackend)}
}
