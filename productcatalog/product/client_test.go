package product

import (
	"fmt"
	"os"
	"testing"

	"github.com/sample/go-webservices/productcatalog"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	apiUrl := ""
	apiKey := "testing"

	productcatalog.Setup(apiUrl, apiKey)
	productcatalog.SetDebug(true)
}

func TestListProducts(t *testing.T) {
	defer gock.Off()

	path := "/products"
	skip := 1000
	limit := 10

	gock.New("").
		Get(path).
		MatchParam("skip", fmt.Sprintf("%v", skip)).
		MatchParam("limit", fmt.Sprintf("%v", limit)).
		Reply(200).
		File("./responses/products.json")
	//EnableNetworking()

	params := &productcatalog.PaginationParams{}
	params.Skip = skip
	params.Limit = limit

	resp, err := List(params)
	assert.Nil(t, err, "should not have error")
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, limit, resp.Limit, "limit should be equal")
	assert.NotNil(t, resp.Products, "should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestGetProduct(t *testing.T) {
	defer gock.Off()

	productId := "5b6aac55869ff10005bd4633"
	path := "/products/" + productId

	gock.New("").
		Get(path).
		Reply(200).
		File("./responses/product.json")
	//EnableNetworking()

	resp, err := Get(productId)
	assert.Nil(t, err, "should not have error")
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestListProductPackages(t *testing.T) {
	defer gock.Off()

	productId := "5b6aac55869ff10005bd4633"
	path := fmt.Sprintf("/products/%s/packages", productId)
	skip := 0
	limit := 10

	gock.New("").
		Get(path).
		MatchParam("skip", fmt.Sprintf("%v", skip)).
		MatchParam("limit", fmt.Sprintf("%v", limit)).
		Reply(200).
		File("./responses/product_packages.json")
	//EnableNetworking()

	params := &productcatalog.ListProductPackagesParams{}
	params.ProductId = productId
	params.Skip = skip
	params.Limit = limit

	resp, err := ListPackages(params)

	assert.Nil(t, err, "should not have error")
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, limit, resp.Limit, "limit should be equal")
	assert.NotNil(t, resp.Packages, "should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestListProductItems(t *testing.T) {
	defer gock.Off()

	productId := "5b6aac55869ff10005bd4633"
	path := fmt.Sprintf("/products/%s/items", productId)
	skip := 0
	limit := 10

	gock.New("").
		Get(path).
		MatchParam("skip", fmt.Sprintf("%v", skip)).
		MatchParam("limit", fmt.Sprintf("%v", limit)).
		Reply(200).
		File("./responses/product_items.json")
	//EnableNetworking()

	params := &productcatalog.ListProductItemParams{}
	params.ProductId = productId
	params.Skip = skip
	params.Limit = limit

	resp, err := ListItems(params)

	assert.Nil(t, err, "should not have error")
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, limit, resp.Limit, "limit should be equal")
	assert.NotNil(t, resp.Items, "should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}
