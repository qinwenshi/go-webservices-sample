package categories

import (
	"fmt"
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

func TestGetCategories(t *testing.T) {
	defer gock.Off()

	language := "en"
	currency := "usd"
	offset := 0
	limit := 10

	gock.New("").
		Get("/categories").
		MatchParam("cnt_language", language).
		MatchParam("currency", currency).
		MatchParam("offset", fmt.Sprintf("%v", offset)).
		MatchParam("limit", fmt.Sprintf("%v", limit)).
		Reply(200).
		File("./responses/categories.json")

	params := &getyourguide.ListCategoriesParams{
		Currency: currency,
		Language: language,
		Offset:   offset,
		Limit:    limit,
	}
	resp, err := List(params)

	assert.Nil(t, err, "should not have error")
	assert.IsType(t, getyourguide.Meta{}, resp.Meta, "should have meta in response")
	assert.Equal(t, "OK", resp.Meta.Status, "status should be OK")
	assert.IsType(t, getyourguide.Exchange{}, resp.Meta.Exchange, "should have exhange in meta")
	assert.Equal(t, 0, resp.Meta.Offset, "offset should be 0")
	assert.Equal(t, 10, resp.Meta.Limit, "limit should be 10")

	assert.IsType(t, getyourguide.CategoriesRespData{}, resp.Data, "should have data in categories")
	assert.NotEmpty(t, resp.Data.Categories, "should contains categories")
	assert.Equal(t, 1, resp.Data.Categories[0].CategoryID, "first category id should be 1")
	assert.Equal(t, "Tours", resp.Data.Categories[0].Name, "first category name should be Tours")
	assert.Equal(t, 17424, resp.Data.Categories[0].NumberOfTours, "first category id should be 17424")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestGetCategoriesWithoutParams(t *testing.T) {
	defer gock.Off()

	offset := 0
	limit := 10

	gock.New("").
		Get("/categories").
		MatchParam("offset", fmt.Sprintf("%v", offset)).
		MatchParam("limit", fmt.Sprintf("%v", limit)).
		Reply(400).
		File("./responses/err_missing_inputs.json")

	params := &getyourguide.ListCategoriesParams{
		Offset: offset,
		Limit:  limit,
	}
	resp, err := List(params)

	assert.NotNil(t, err, "should be error")
	assert.IsType(t, getyourguide.ErrorResp{}, err, "should be a getyourguide error response")
	assert.Nil(t, resp, "response should be nil")

	e := err.(getyourguide.ErrorResp)
	assert.Equal(t, "GetYourGuide AG", e.Descriptor, "descriptor should be GetYourGuide AG")
	assert.Equal(t, "1", e.APIVersion, "apiVersion should be 1")
	assert.Equal(t, "getAction", e.Method, "method should be getByIdAction")
	assert.Equal(t, "ERROR", e.Status, "status should be ERROR")
	assert.Equal(t, 2, len(e.Errors), "number of errors should be 2")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}
