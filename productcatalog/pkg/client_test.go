package pkg

import (
	"fmt"
	"os"
	"testing"

	"github.com/sample/go-webservices/productcatalog"

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
	apiKey := "testing"

	productcatalog.Setup(apiUrl, apiKey)
	productcatalog.SetDebug(true)
}

func TestGetPackage(t *testing.T) {
	defer gock.Off()

	packageId := "5b6aac7e869ff10005bd468f"
	path := "/packages/" + packageId

	gock.New("").
		Get(path).
		Reply(200).
		File("./responses/package.json")
	//EnableNetworking()

	resp, err := Get(packageId)
	assert.Nil(t, err, "should not have error")
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestListPackageItems(t *testing.T) {
	defer gock.Off()

	skip := 0
	limit := 10

	packageId := "5b6aac7e869ff10005bd468f"
	path := fmt.Sprintf("/packages/%s/items", packageId)

	gock.New("").
		Get(path).
		MatchParam("skip", fmt.Sprintf("%v", skip)).
		MatchParam("limit", fmt.Sprintf("%v", limit)).
		Reply(200).
		File("./responses/package_items.json")
	//EnableNetworking()

	params := &productcatalog.ListPackageItemsParams{}
	params.PackageId = packageId
	params.Skip = skip
	params.Limit = limit

	resp, err := ListItems(params)

	assert.Nil(t, err, "should not have error")
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, limit, resp.Limit, "limit should be equal")
	assert.NotNil(t, resp.Items, "should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestListPackageAvailabilities(t *testing.T) {
	defer gock.Off()

	skip := 0
	limit := 10

	packageId := "5b6aac7e869ff10005bd468f"
	path := fmt.Sprintf("/packages/%s/availabilities", packageId)

	gock.New("").
		Get(path).
		MatchParam("skip", fmt.Sprintf("%v", skip)).
		MatchParam("limit", fmt.Sprintf("%v", limit)).
		Reply(200).
		File("./responses/package_availabilities.json")
	//EnableNetworking()

	params := &productcatalog.ListPackageAvailabilitiesParams{}
	params.PackageId = packageId
	params.StartDate, _ = time.Parse("2016-01-02", "2018-08-17")
	params.EndDate, _ = time.Parse("2016-01-02", "2018-08-17")
	params.Skip = skip
	params.Limit = limit

	resp, err := ListAvailabilities(params)

	assert.Nil(t, err, "should not have error")
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, limit, resp.Limit, "limit should be equal")
	assert.NotNil(t, resp.Availabilities, "should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}
