package tour

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
	getyourguide.SetDebug(false)
}

func TestGetTour(t *testing.T) {
	defer gock.Off()

	language := "en"
	currency := "usd"
	tourId := "121665"

	gock.New("").
		Get("/tours/"+tourId).
		MatchParam("cnt_language", language).
		MatchParam("currency", currency).
		Reply(200).
		File("./responses/tour.json")

	params := &getyourguide.GetTourParams{
		TourId:   tourId,
		Currency: currency,
		Language: language,
	}
	resp, err := Get(params)

	assert.Nil(t, err, "should not have error")
	assert.IsType(t, getyourguide.Meta{}, resp.Meta, "should have meta in response")
	assert.Equal(t, "OK", resp.Meta.Status, "status should be OK")
	assert.IsType(t, getyourguide.Exchange{}, resp.Meta.Exchange, "should have exhange in meta")
	assert.Equal(t, 0, resp.Meta.Offset, "offset should be 0")
	assert.Equal(t, 10, resp.Meta.Limit, "limit should be 10")

	assert.IsType(t, getyourguide.TourRespData{}, resp.Data, "should have data in tour")
	assert.NotEmpty(t, resp.Data.Tours, "should contains tour")
	assert.Equal(t, 1, len(resp.Data.Tours), "number of tour should be 1")
	assert.Equal(t, 121665, resp.Data.Tours[0].TourID, "tour id should be 121665")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestGetTourWithoutParams(t *testing.T) {
	defer gock.Off()

	tourId := "121665"

	gock.New("").
		Get("/tours/" + tourId).
		Reply(400).
		File("./responses/err_missing_inputs.json")

	params := &getyourguide.GetTourParams{
		TourId: tourId,
	}
	resp, err := Get(params)

	assert.NotNil(t, err, "should be error")
	assert.IsType(t, getyourguide.ErrorResp{}, err, "should be a getyourguide error response")
	assert.Nil(t, resp, "response should be nil")

	e := err.(getyourguide.ErrorResp)
	assert.Equal(t, "GetYourGuide AG", e.Descriptor, "descriptor should be GetYourGuide AG")
	assert.Equal(t, "1", e.APIVersion, "apiVersion should be 1")
	assert.Equal(t, "getTourByIdAction", e.Method, "method should be getTourByIdAction")
	assert.Equal(t, "ERROR", e.Status, "status should be ERROR")
	assert.Equal(t, 2, len(e.Errors), "number of errors should be 2")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestListTourAvailabilities(t *testing.T) {
	defer gock.Off()

	language := "en"
	currency := "usd"
	tourId := "121665"

	gock.New("").
		Get("/tours/"+tourId+"/availabilities").
		MatchParam("cnt_language", language).
		MatchParam("currency", currency).
		Reply(200).
		File("./responses/tour_availabilities.json")

	params := &getyourguide.ListTourAvailabilitiesParams{
		GetTourParams: getyourguide.GetTourParams{
			TourId:   tourId,
			Currency: currency,
			Language: language,
		},
	}
	resp, err := ListAvailabilities(params)

	assert.Nil(t, err, "should not have error")
	assert.IsType(t, getyourguide.Meta{}, resp.Meta, "should have meta in response")
	assert.Equal(t, "OK", resp.Meta.Status, "status should be OK")
	assert.IsType(t, getyourguide.Exchange{}, resp.Meta.Exchange, "should have exhange in meta")
	assert.Equal(t, 0, resp.Meta.Offset, "offset should be 0")
	assert.Equal(t, 10, resp.Meta.Limit, "limit should be 10")

	assert.IsType(t, getyourguide.TourAvailabilitiesRespData{}, resp.Data, "should have data in tour")
	assert.NotEmpty(t, resp.Data.TourAvailabilities, "should contains tour availabilities")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestListTourOptions(t *testing.T) {
	defer gock.Off()

	language := "en"
	currency := "usd"
	tourId := "121665"

	gock.New("").
		Get("/tours/"+tourId+"/options").
		MatchParam("cnt_language", language).
		MatchParam("currency", currency).
		Reply(200).
		File("./responses/tour_options.json")

	params := &getyourguide.ListTourOptionsParams{
		GetTourParams: getyourguide.GetTourParams{
			TourId:   tourId,
			Currency: currency,
			Language: language,
		},
	}
	resp, err := ListOptions(params)

	assert.Nil(t, err, "should not have error")
	assert.IsType(t, getyourguide.Meta{}, resp.Meta, "should have meta in response")
	assert.Equal(t, "OK", resp.Meta.Status, "status should be OK")
	assert.IsType(t, getyourguide.Exchange{}, resp.Meta.Exchange, "should have exhange in meta")
	assert.Equal(t, 0, resp.Meta.Offset, "offset should be 0")
	assert.Equal(t, 10, resp.Meta.Limit, "limit should be 10")
	assert.IsType(t, getyourguide.TourOptionsRespData{}, resp.Data, "should have data in tour")
	assert.NotEmpty(t, resp.Data.TourOptions, "should contains tour options")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}
