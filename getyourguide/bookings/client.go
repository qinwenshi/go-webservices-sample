package bookings

import (
	"fmt"
	"net/url"

	"github.com/sample/go-webservices/getyourguide"
)

type Client struct {
	B getyourguide.Backend
}

func Get(params *getyourguide.GetBookingParams) (*getyourguide.BookingResp, error) {
	return getC().Get(params)
}

func List(params *getyourguide.ListBookingsParams) (*getyourguide.BookingsResp, error) {
	return getC().List(params)
}

func Create(params *getyourguide.CreateBookingParams) (*getyourguide.CreateBookingResp, error) {
	return getC().Create(params)
}

func (c Client) Get(params *getyourguide.GetBookingParams) (resp *getyourguide.BookingResp, err error) {
	v := url.Values{}

	v.Add("cnt_language", params.Language)
	v.Add("currency", params.Currency)

	path := "/bookings/" + params.Id

	err = c.B.Call("GET", path, &v, nil, &resp)
	return resp, err
}

func (c Client) List(params *getyourguide.ListBookingsParams) (resp *getyourguide.BookingsResp, err error) {
	v := url.Values{}

	v.Add("offset", fmt.Sprintf("%d", params.Offset))
	v.Add("limit", fmt.Sprintf("%d", params.Limit))
	v.Add("cnt_language", params.Language)
	v.Add("currency", params.Currency)

	err = c.B.Call("GET", "/bookings", &v, nil, &resp)
	return resp, err
}

func (c Client) Create(params *getyourguide.CreateBookingParams) (resp *getyourguide.CreateBookingResp, err error) {
	err = c.B.Call("POST", "/bookings", nil, params, &resp)
	return resp, err
}

func getC() Client {
	return Client{getyourguide.GetBackend(getyourguide.PublicBackend)}
}
