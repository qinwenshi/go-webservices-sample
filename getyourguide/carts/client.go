package bookings

import (
	"github.com/sample/go-webservices/getyourguide"
)

type Client struct {
	B getyourguide.Backend
}

func Create(params *getyourguide.CreateCartParams) (*getyourguide.CreateCartResp, error) {
	return getC().Create(params)
}

func (c Client) Create(params *getyourguide.CreateCartParams) (resp *getyourguide.CreateCartResp, err error) {
	err = c.B.Call("POST", "/carts", nil, params, &resp)
	return resp, err
}

func getC() Client {
	return Client{getyourguide.GetBackend(getyourguide.PublicBackend)}
}
