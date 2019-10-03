package bookings

import (
	"fmt"
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

func TestGetBooking(t *testing.T) {
	defer gock.Off()

	language := "en"
	currency := "usd"
	bookingId := "27896440"

	gock.New("").
		Get("/bookings").
		MatchParam("cnt_language", language).
		MatchParam("currency", currency).
		Reply(200).
		File("./responses/booking.json")

	params := &getyourguide.GetBookingParams{
		Currency: currency,
		Language: language,
		Id:       bookingId,
	}
	resp, err := Get(params)

	assert.Nil(t, err, "should not have error")
	assert.IsType(t, getyourguide.Meta{}, resp.Meta, "should have meta in response")
	assert.Equal(t, "OK", resp.Meta.Status, "status should be OK")
	assert.IsType(t, getyourguide.Exchange{}, resp.Meta.Exchange, "should have exhange in meta")
	assert.Equal(t, 0, resp.Meta.Offset, "offset should be 0")
	assert.Equal(t, 10, resp.Meta.Limit, "limit should be 10")

	assert.IsType(t, getyourguide.BookingRespData{}, resp.Data, "should have data in booking")
	assert.NotEmpty(t, resp.Data.Bookings, "should contains booking")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestGetBookings(t *testing.T) {
	defer gock.Off()

	language := "en"
	currency := "usd"
	offset := 0
	limit := 10

	gock.New("").
		Get("/bookings").
		MatchParam("cnt_language", language).
		MatchParam("currency", currency).
		MatchParam("offset", fmt.Sprintf("%v", offset)).
		MatchParam("limit", fmt.Sprintf("%v", limit)).
		Reply(200).
		File("./responses/bookings.json")

	params := &getyourguide.ListBookingsParams{
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

	assert.IsType(t, getyourguide.BookingsRespData{}, resp.Data, "should have data in tours")
	assert.NotEmpty(t, resp.Data.Bookings, "should contains tours")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestGetBookingsWithoutParams(t *testing.T) {
	defer gock.Off()

	offset := 0
	limit := 10

	gock.New("").
		Get("/bookings").
		MatchParam("offset", fmt.Sprintf("%v", offset)).
		MatchParam("limit", fmt.Sprintf("%v", limit)).
		Reply(400).
		File("./responses/err_missing_inputs.json")

	params := &getyourguide.ListBookingsParams{
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
	assert.Equal(t, "getBookingsByQueryAction", e.Method, "method should be getBookingsByQueryAction")
	assert.Equal(t, "ERROR", e.Status, "status should be ERROR")
	assert.Equal(t, 2, len(e.Errors), "number of errors should be 2")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestCreateBooking(t *testing.T) {
	defer gock.Off()

	gock.New("").
		Post("/bookings").
		Reply(200).
		File("./responses/create_booking.json")

	language := "en"
	currency := "usd"
	optionId := "178119"
	date := time.Now().UTC()

	categoryId := getyourguide.CateogryId(1)
	numberOfParticipant := getyourguide.NumberOfParticipant(1)
	inputs := map[getyourguide.CateogryId]getyourguide.NumberOfParticipant{
		categoryId: numberOfParticipant,
	}

	price := float32(123)
	referenceId := "testing"

	params := getyourguide.NewCreateBookingParams(language, currency, optionId, date, inputs, price, referenceId)
	resp, err := Create(params)

	assert.Nil(t, err, "should not have error")
	assert.IsType(t, getyourguide.Meta{}, resp.Meta, "should have meta in response")
	assert.Equal(t, "OK", resp.Meta.Status, "status should be OK")
	assert.Equal(t, "createBookingAction", resp.Meta.Method, "method should be createBookingAction")
	assert.IsType(t, getyourguide.Exchange{}, resp.Meta.Exchange, "should have exhange in meta")
	assert.Equal(t, 0, resp.Meta.Offset, "offset should be 0")
	assert.Equal(t, 10, resp.Meta.Limit, "limit should be 10")

	assert.IsType(t, getyourguide.CreateBookingRespData{}, resp.Data, "should have data in booking")
	assert.NotEmpty(t, resp.Data.Bookings, "should contains booking object")
	assert.Equal(t, "temp", resp.Data.Bookings.Status, "status should be temp")
	assert.Equal(t, 0, resp.Data.Bookings.ReturnCode, "return code should be 0")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}
