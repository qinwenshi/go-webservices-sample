package products

import (
	"fmt"
	"os"
	"testing"

	"github.com/sample/go-webservices/bemyguest2"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	apiUrl := ""
	apiKey := ""

	bemyguest2.Setup(apiUrl, apiKey)
	bemyguest2.SetDebug(true)
}

func TestGetProduct(t *testing.T) {
	defer gock.Off()

	productId := "987081b3-2965-5300-bf35-8baa682ed969"
	path := "/v2/products/" + productId

	gock.New("").
		Get(path).
		Reply(200).
		File("./responses/product.json")

	params := &bemyguest2.ProductParams{}
	params.SetId(productId)

	resp, err := Get(params)
	assert.Nil(t, err, "should not have error")
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, productId, resp.Data.UUID, "product id should be equal")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestGetNonExistProduct(t *testing.T) {
	defer gock.Off()

	productId := "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"

	path := "/v2/products/" + productId

	gock.New("").
		Get(path).
		Reply(404).
		File("./responses/error_not_found.json")

	params := &bemyguest2.ProductParams{}
	params.SetId(productId)

	resp, err := Get(params)
	assert.NotNil(t, err, "should be error")
	assert.Nil(t, resp, "response should be nil")
	assert.IsType(t, bemyguest2.ErrorResp{}, err, "should receive bemyguest2 error response")
	content := err.(bemyguest2.ErrorResp).Content
	assert.Equal(t, 404, content.HTTPCode)
	assert.Equal(t, "Listing was not found or expired, provided UUID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", content.Message)
	assert.Equal(t, "not_found", content.Code)
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestGetProducts(t *testing.T) {
	defer gock.Off()

	page := 2
	perPage := 10

	gock.New("").
		Get("/v2/products").
		MatchParam("page", fmt.Sprintf("%v", page)).
		MatchParam("per_page", fmt.Sprintf("%v", perPage)).
		Reply(200).
		File("./responses/products.json")

	params := &bemyguest2.ProductsParams{}
	params.SetPage(page)
	params.SetPerPage(perPage)
	resp, err := List(params)

	assert.Nil(t, err, "should not have error")
	assert.Equal(t, len(resp.Data), perPage, "should have number of product as required")
	assert.IsType(t, resp.Meta, bemyguest2.Meta{}, "should have meta in response")
	assert.IsType(t, resp.Meta.Pagination, bemyguest2.Pagination{}, "should have pagination in response")
	assert.Equal(t, page, resp.Meta.Pagination.CurrentPage, "should got current page")
	assert.Equal(t, perPage, resp.Meta.Pagination.PerPage, "should have same per page setting")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestListProductTypes(t *testing.T) {
	defer gock.Off()

	productId := "43096939-a213-560d-ab63-24825f7822ae"
	path := "/v2/products/" + productId + "/product-types"

	gock.New("").
		Get(path).
		Reply(200).
		File("./responses/product_types.json")

	params := &bemyguest2.ProductTypesParams{}
	params.SetId(productId)

	resp, err := ListProductTypes(params)
	assert.Nil(t, err, "should not have error")
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}