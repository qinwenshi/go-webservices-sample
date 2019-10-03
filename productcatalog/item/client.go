package item

import (
	"fmt"
	"net/url"

	"github.com/sample/go-webservices/productcatalog"
)

type Client struct {
	B productcatalog.Backend
}

func Get(id string) (*productcatalog.Item, error) {
	return getC().Get(id)
}

func ListItemPrices(params *productcatalog.ListItemPricesParams) (*productcatalog.ItemPricesResp, error) {
	return getC().ListItemPrices(params)
}

func (c Client) Get(id string) (resp *productcatalog.Item, err error) {
	err = c.B.Call("GET", "/items/"+id, nil, nil, nil, &resp)
	return resp, err
}

func (c Client) ListItemPrices(params *productcatalog.ListItemPricesParams) (resp *productcatalog.ItemPricesResp, err error) {
	v := url.Values{}

	v.Add("skip", fmt.Sprintf("%d", params.Skip))
	v.Add("limit", fmt.Sprintf("%d", params.Limit))

	if params.PackageId != "" {
		v.Add("package_id", params.PackageId)
	}
	if !params.Date.IsZero() {
		v.Add("date", params.Date.Format("2006-01-02"))
	}
	if params.StartAt != "" {
		v.Add("start_at", params.StartAt)
	}

	path := fmt.Sprintf("/items/%s/prices", params.ItemId)

	err = c.B.Call("GET", path, &v, nil, nil, &resp)
	return resp, err
}

func getC() Client {
	return Client{productcatalog.GetBackend(productcatalog.PublicBackend)}
}
