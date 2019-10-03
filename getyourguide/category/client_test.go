package category

import (
	"os"
	"testing"

	"github.com/sample/go-webservices/getyourguide"

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
	apiVersion := "1"

	getyourguide.Setup(apiUrl, apiKey, apiVersion)
	getyourguide.SetDebug(true)
}

func TestGetCategory(t *testing.T) {
	defer gock.Off()

	language := "en"
	currency := "usd"
	categoryId := "2"

	gock.New("").
		Get("/categories/2").
		MatchParam("cnt_language", language).
		MatchParam("currency", currency).
		Reply(200).
		File("./responses/category.json")

	params := &getyourguide.GetCategoryParams{
		CategoryId: categoryId,
		Currency:   currency,
		Language:   language,
	}
	resp, err := Get(params)

	assert.Nil(t, err, "should not have error")
	assert.IsType(t, getyourguide.Meta{}, resp.Meta, "should have meta in response")
	assert.Equal(t, "OK", resp.Meta.Status, "status should be OK")
	assert.IsType(t, getyourguide.Exchange{}, resp.Meta.Exchange, "should have exhange in meta")
	assert.Equal(t, 0, resp.Meta.Offset, "offset should be 0")
	assert.Equal(t, 10, resp.Meta.Limit, "limit should be 10")

	assert.IsType(t, getyourguide.CategoryRespData{}, resp.Data, "should have data in category")
	assert.NotEmpty(t, resp.Data.Categories, "should contains category")
	assert.Equal(t, 2, resp.Data.Categories[0].CategoryID, "first category id should be 2")
	assert.Equal(t, "Sightseeing Tours", resp.Data.Categories[0].Name, "first category name should be Sightseeing Tours")
	assert.Equal(t, 12176, resp.Data.Categories[0].NumberOfTours, "first category id should be 12176")
	assert.Equal(t, 1, resp.Data.Categories[0].ParentId, "first category parent id should be 1")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestGetCategoryWithoutParams(t *testing.T) {
	defer gock.Off()

	categoryId := "2"

	gock.New("").
		Get("/categories/2").
		Reply(400).
		File("./responses/err_missing_inputs.json")

	params := &getyourguide.GetCategoryParams{
		CategoryId: categoryId,
	}
	resp, err := Get(params)

	assert.NotNil(t, err, "should be error")
	assert.IsType(t, getyourguide.ErrorResp{}, err, "should be a getyourguide error response")
	assert.Nil(t, resp, "response should be nil")

	e := err.(getyourguide.ErrorResp)
	assert.Equal(t, "GetYourGuide AG", e.Descriptor, "descriptor should be GetYourGuide AG")
	assert.Equal(t, "1", e.APIVersion, "apiVersion should be 1")
	assert.Equal(t, "getByIdAction", e.Method, "method should be getByIdAction")
	assert.Equal(t, "ERROR", e.Status, "status should be ERROR")
	assert.Equal(t, 2, len(e.Errors), "number of errors should be 2")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}
