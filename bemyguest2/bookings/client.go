package bookings

import (
	"fmt"
	"net/url"

	"github.com/sample/go-webservices/bemyguest2"
)

type Status string

const (
	StatusConfirm Status = "confirm"
	StatusCancel  Status = "cancel"
)

type Client struct {
	B bemyguest2.Backend
}

func Get(id string) (*bemyguest2.BookingResp, error) {
	return getC().Get(id)
}

func Create(params *bemyguest2.CreateBookingParams) (*bemyguest2.CreateBookingResp, error) {
	return getC().Create(params)
}

func Update(id string, status Status) (*bemyguest2.CreateBookingResp, error) {
	return getC().Update(id, status)
}

func List(params *bemyguest2.ListBookingsParams) (*bemyguest2.BookingsResp, error) {
	return getC().List(params)
}

func ListVouchers(id string) (*bemyguest2.VouchersResp, error) {
	return getC().ListVouchers(id)
}

func (c Client) Get(id string) (resp *bemyguest2.BookingResp, err error) {
	path := "/v2/bookings/" + id
	err = c.B.Call("GET", path, nil, nil, &resp)
	return resp, err
}

func (c Client) Create(params *bemyguest2.CreateBookingParams) (resp *bemyguest2.CreateBookingResp, err error) {
	err = c.B.Call("POST", "/v2/bookings", nil, params, &resp)
	return resp, err
}

func (c Client) Update(id string, status Status) (resp *bemyguest2.CreateBookingResp, err error) {
	path := "/v2/bookings/" + id + "/" + string(status)
	err = c.B.Call("PUT", path, nil, nil, &resp)
	return resp, err
}

func (c Client) List(params *bemyguest2.ListBookingsParams) (resp *bemyguest2.BookingsResp, err error) {
	v := url.Values{}

	v.Add("page", fmt.Sprintf("%d", params.Page))
	v.Add("per_page", fmt.Sprintf("%d", params.PerPage))

	if params.FirstName != nil {
		v.Add("first_name", *params.FirstName)
	}
	if params.LastName != nil {
		v.Add("last_name", *params.LastName)
	}
	if params.Email != nil {
		v.Add("email", *params.Email)
	}
	if params.Phone != nil {
		v.Add("phone", *params.Phone)
	}
	if params.PartnerReference != nil {
		v.Add("numberpartner_reference", *params.PartnerReference)
	}
	if params.Query != nil {
		v.Add("query", *params.Query)
	}
	if params.Status != nil {
		v.Add("status", *params.Status)
	}

	err = c.B.Call("GET", "/v2/bookings", &v, nil, &resp)
	return resp, err
}

func (c Client) ListVouchers(id string) (resp *bemyguest2.VouchersResp, err error) {
	path := fmt.Sprintf("/v2/bookings/%s/vouchers", id)
	err = c.B.Call("GET", path, nil, nil, &resp)
	return resp, err
}

func getC() Client {
	return Client{bemyguest2.GetBackend(bemyguest2.PublicBackend)}
}
