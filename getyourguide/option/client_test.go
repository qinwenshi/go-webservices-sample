package option

import (
	"os"
	"testing"
	"time"

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

func TestGetTourOption(t *testing.T) {
	defer gock.Off()

	language := "en"
	currency := "usd"
	optionId := "178119"

	gock.New("").
		Get("/options/"+optionId).
		MatchParam("cnt_language", language).
		MatchParam("currency", currency).
		Reply(200).
		File("./responses/option.json")

	params := &getyourguide.GetOptionParams{
		OptionId: optionId,
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

	assert.IsType(t, getyourguide.OptionRespData{}, resp.Data, "should have data in option")
	assert.NotEmpty(t, resp.Data.TourOptions, "should contains options")
	assert.Equal(t, 1, len(resp.Data.TourOptions), "number of option should be 1")
	assert.Equal(t, 178119, resp.Data.TourOptions[0].OptionID, "tour id should be 178119")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestGetTourOptionWithoutParams(t *testing.T) {
	defer gock.Off()

	optionId := "178119"

	gock.New("").
		Get("/options/" + optionId).
		Reply(400).
		File("./responses/err_missing_inputs.json")

	params := &getyourguide.GetOptionParams{
		OptionId: optionId,
	}
	resp, err := Get(params)

	assert.NotNil(t, err, "should be error")
	assert.IsType(t, getyourguide.ErrorResp{}, err, "should be a getyourguide error response")
	assert.Nil(t, resp, "response should be nil")

	e := err.(getyourguide.ErrorResp)
	assert.Equal(t, "GetYourGuide AG", e.Descriptor, "descriptor should be GetYourGuide AG")
	assert.Equal(t, "1", e.APIVersion, "apiVersion should be 1")
	assert.Equal(t, "getTourOptionByIdAction", e.Method, "method should be getTourOptionByIdAction")
	assert.Equal(t, "ERROR", e.Status, "status should be ERROR")
	assert.Equal(t, 2, len(e.Errors), "number of errors should be 2")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestListOptionAvailabilities(t *testing.T) {
	defer gock.Off()

	language := "en"
	currency := "usd"
	optionId := "178119"

	gock.New("").
		Get("/options/"+optionId+"/availabilities").
		MatchParam("cnt_language", language).
		MatchParam("currency", currency).
		Reply(200).
		File("./responses/option_availabilities.json")

	params := &getyourguide.ListOptionAvailabilitiesParams{
		GetOptionParams: getyourguide.GetOptionParams{
			OptionId: optionId,
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

	assert.IsType(t, getyourguide.OptionAvailabilitiesRespData{}, resp.Data, "should have data in option")
	assert.NotEmpty(t, resp.Data.OptionAvailabilities, "should contains option availabilities")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestListOptionAvailabilitiesByDate(t *testing.T) {
	defer gock.Off()

	language := "en"
	currency := "usd"
	optionId := "37361"

	dateStr := "2018-05-30T00:00:00"

	date, _ := time.Parse(TimeFormat, dateStr)

	gock.New("").
		Get("/options/"+optionId+"/availabilities").
		MatchParam("cnt_language", language).
		MatchParam("currency", currency).
		MatchParam("date[]", dateStr).
		Reply(200).
		File("./responses/option_availabilities_by_date.json")
		//EnableNetworking()

	params := &getyourguide.ListOptionAvailabilitiesParams{
		GetOptionParams: getyourguide.GetOptionParams{
			OptionId: optionId,
			Currency: currency,
			Language: language,
		},
		Date: &date,
	}
	resp, err := ListAvailabilities(params)

	assert.Nil(t, err, "should not have error")
	assert.IsType(t, getyourguide.Meta{}, resp.Meta, "should have meta in response")
	assert.Equal(t, "OK", resp.Meta.Status, "status should be OK")
	assert.IsType(t, getyourguide.Exchange{}, resp.Meta.Exchange, "should have exhange in meta")
	assert.Equal(t, 0, resp.Meta.Offset, "offset should be 0")
	assert.Equal(t, 10, resp.Meta.Limit, "limit should be 10")

	assert.IsType(t, getyourguide.OptionAvailabilitiesRespData{}, resp.Data, "should have data in option")
	assert.NotEmpty(t, resp.Data.OptionAvailabilities, "should contains option availabilities")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestListOptionPricings(t *testing.T) {
	defer gock.Off()

	language := "en"
	currency := "usd"
	optionId := "178119"

	gock.New("").
		Get("/options/"+optionId+"/pricings").
		MatchParam("cnt_language", language).
		MatchParam("currency", currency).
		Reply(200).
		File("./responses/option_pricings.json")

	params := &getyourguide.ListOptionPricingsParams{
		GetOptionParams: getyourguide.GetOptionParams{
			OptionId: optionId,
			Currency: currency,
			Language: language,
		},
	}
	resp, err := ListPricings(params)

	assert.Nil(t, err, "should not have error")
	assert.IsType(t, getyourguide.Meta{}, resp.Meta, "should have meta in response")
	assert.Equal(t, "OK", resp.Meta.Status, "status should be OK")
	assert.IsType(t, getyourguide.Exchange{}, resp.Meta.Exchange, "should have exhange in meta")
	assert.Equal(t, 0, resp.Meta.Offset, "offset should be 0")
	assert.Equal(t, 10, resp.Meta.Limit, "limit should be 10")
	assert.IsType(t, getyourguide.OptionPricingsRespData{}, resp.Data, "should have data in pricings")
	assert.NotEmpty(t, resp.Data.OptionPricings, "should contains tour options")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}
