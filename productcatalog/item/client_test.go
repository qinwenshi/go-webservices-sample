package item

import (
	"fmt"
	"os"
	"testing"
	"time"

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

func TestGetItem(t *testing.T) {
	defer gock.Off()

	itemId := "5b6aac59869ff10005bd4635"
	path := "/items/" + itemId

	gock.New("").
		Get(path).
		Reply(200).
		File("./responses/item.json")
	//EnableNetworking()

	resp, err := Get(itemId)
	assert.Nil(t, err, "should not have error")
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestListItemPrices(t *testing.T) {
	defer gock.Off()

	skip := 0
	limit := 10

	//productId := "5b6aac55869ff10005bd4633"
	packageId := "5b6aac7e869ff10005bd468f"
	itemId := "5b6aac59869ff10005bd4635"

	path := fmt.Sprintf("/items/%s/prices", itemId)

	gock.New("").
		Get(path).
		MatchParam("skip", fmt.Sprintf("%v", skip)).
		MatchParam("limit", fmt.Sprintf("%v", limit)).
		Reply(200).
		File("./responses/item_prices.json")
		//EnableNetworking()

	params := &productcatalog.ListItemPricesParams{}
	params.PackageId = packageId
	params.ItemId = itemId
	params.Date, _ = time.Parse("2016-01-02", "2018-08-15")
	params.StartAt = "15:00:00"
	params.Skip = skip
	params.Limit = limit

	resp, err := ListItemPrices(params)

	assert.Nil(t, err, "should not have error")
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, limit, resp.Limit, "limit should be equal")
	assert.NotNil(t, resp.Prices, "should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}
