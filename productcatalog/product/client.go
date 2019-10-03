package product

import (
	"fmt"
	"net/url"

	"github.com/sample/go-webservices/productcatalog"
)

type Client struct {
	B productcatalog.Backend
}

func Get(id string) (*productcatalog.Product, error) {
	return getC().Get(id)
}

func List(params *productcatalog.PaginationParams) (*productcatalog.ProductsResp, error) {
	return getC().List(params)
}

func ListPackages(params *productcatalog.ListProductPackagesParams) (*productcatalog.PackagesResp, error) {
	return getC().ListPackages(params)
}

func ListItems(params *productcatalog.ListProductItemParams) (*productcatalog.ItemsResp, error) {
	return getC().ListItems(params)
}

func (c Client) Get(id string) (resp *productcatalog.Product, err error) {
	err = c.B.Call("GET", "/products/"+id, nil, nil, nil, &resp)
	return resp, err
}

func (c Client) List(params *productcatalog.PaginationParams) (resp *productcatalog.ProductsResp, err error) {
	v := url.Values{}

	v.Add("skip", fmt.Sprintf("%d", params.Skip))
	v.Add("limit", fmt.Sprintf("%d", params.Limit))

	err = c.B.Call("GET", "/products", &v, nil, nil, &resp)
	return resp, err
}

func (c Client) ListPackages(params *productcatalog.ListProductPackagesParams) (resp *productcatalog.PackagesResp, err error) {
	v := url.Values{}

	v.Add("skip", fmt.Sprintf("%d", params.Skip))
	v.Add("limit", fmt.Sprintf("%d", params.Limit))

	path := fmt.Sprintf("/products/%s/packages", params.ProductId)

	err = c.B.Call("GET", path, &v, nil, nil, &resp)
	return resp, err
}

func (c Client) ListItems(params *productcatalog.ListProductItemParams) (resp *productcatalog.ItemsResp, err error) {
	v := url.Values{}

	v.Add("skip", fmt.Sprintf("%d", params.Skip))
	v.Add("limit", fmt.Sprintf("%d", params.Limit))

	path := fmt.Sprintf("/products/%s/items", params.ProductId)

	err = c.B.Call("GET", path, &v, nil, nil, &resp)
	return resp, err
}

func getC() Client {
	return Client{productcatalog.GetBackend(productcatalog.PublicBackend)}
}
