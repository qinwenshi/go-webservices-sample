package pkg

import (
	"fmt"
	"net/url"

	"github.com/sample/go-webservices/productcatalog"
)

type Client struct {
	B productcatalog.Backend
}

func Get(id string) (*productcatalog.Package, error) {
	return getC().Get(id)
}

func ListItems(params *productcatalog.ListPackageItemsParams) (*productcatalog.ItemsResp, error) {
	return getC().ListItems(params)
}

func ListAvailabilities(params *productcatalog.ListPackageAvailabilitiesParams) (*productcatalog.PackageAvailabilitiesResp, error) {
	return getC().ListAvailabilities(params)
}

func (c Client) Get(id string) (resp *productcatalog.Package, err error) {
	err = c.B.Call("GET", "/packages/"+id, nil, nil, nil, &resp)
	return resp, err
}

func (c Client) ListItems(params *productcatalog.ListPackageItemsParams) (resp *productcatalog.ItemsResp, err error) {
	v := url.Values{}

	v.Add("skip", fmt.Sprintf("%d", params.Skip))
	v.Add("limit", fmt.Sprintf("%d", params.Limit))

	path := fmt.Sprintf("/packages/%s/items", params.PackageId)

	err = c.B.Call("GET", path, &v, nil, nil, &resp)
	return resp, err
}

func (c Client) ListAvailabilities(params *productcatalog.ListPackageAvailabilitiesParams) (resp *productcatalog.PackageAvailabilitiesResp, err error) {
	v := url.Values{}

	v.Add("skip", fmt.Sprintf("%d", params.Skip))
	v.Add("limit", fmt.Sprintf("%d", params.Limit))

	if !params.StartDate.IsZero() {
		v.Add("start_date", params.StartDate.Format("2006-01-02"))
	}
	if !params.EndDate.IsZero() {
		v.Add("end_date", params.EndDate.Format("2006-01-02"))
	}

	path := fmt.Sprintf("/packages/%s/availabilities", params.PackageId)

	err = c.B.Call("GET", path, &v, nil, nil, &resp)
	return resp, err
}

func getC() Client {
	return Client{productcatalog.GetBackend(productcatalog.PublicBackend)}
}
