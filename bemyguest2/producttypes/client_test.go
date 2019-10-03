package producttypes

import (
	"fmt"
	"os"
	"testing"

	"github.com/sample/go-webservices/bemyguest2"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"time"
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

func TestGetProductType(t *testing.T) {
	defer gock.Off()

	productTypeId := "0c7491a2-355f-5b4d-82c9-c1b86236653f"
	path := "/v2/product-types/" + productTypeId

	gock.New("").
		Get(path).
		Reply(200).
		File("./responses/product_type.json")

	params := &bemyguest2.ProductTypesParams{}
	params.SetId(productTypeId)

	resp, err := Get(params)
	assert.Nil(t, err, "should not have error")
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, productTypeId, resp.Data.UUID, "product id should be equal")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestProductTypePrices(t *testing.T) {
	defer gock.Off()

	productTypeId := "0c7491a2-355f-5b4d-82c9-c1b86236653f"
	path := fmt.Sprintf("/v2/product-types/%s/price-lists", productTypeId)

	gock.New("").
		Get(path).
		Reply(200).
		File("./responses/product_type_prices.json")

	start := time.Now().UTC()
	end := start.AddDate(0, 0, 7)

	params := &bemyguest2.ListPricesParams{}
	params.SetId(productTypeId)
	params.DateStart = &start
	params.DateEnd = &end

	resp, err := ListPrices(params)

	assert.Nil(t, err, "should not have error")
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestProductTypePricesByDate(t *testing.T) {
	defer gock.Off()

	productTypeId := "0c7491a2-355f-5b4d-82c9-c1b86236653f"
	dateStr := "2006-01-02"
	date, _ := time.Parse("2006-01-02", dateStr)
	path := fmt.Sprintf("/v2/product-types/%s/price-lists/%s", productTypeId, dateStr)

	gock.New("").
		Get(path).
		Reply(200).
		File("./responses/product_type_prices_by_date.json")

	params := &bemyguest2.GetPriceByDateParams{}
	params.SetId(productTypeId)
	params.SetDate(date)
	resp, err := GetPriceByDate(params)

	assert.Nil(t, err, "should not have error")
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}
