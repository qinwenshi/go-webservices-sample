package config

import "github.com/sample/go-webservices/bemyguest2"

type Client struct {
	B bemyguest2.Backend
}

func Get() (*bemyguest2.Config, error) {
	return getC().Get()
}

func (c Client) Get() (resp *bemyguest2.Config, err error) {
	err = c.B.Call("GET", "/v2/config", nil, nil, &resp)
	return resp, err
}

func getC() Client {
	return Client{bemyguest2.GetBackend(bemyguest2.PublicBackend)}
}
