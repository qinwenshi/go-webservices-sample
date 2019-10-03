package bookings

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

func TestCreateBooking(t *testing.T) {
	defer gock.Off()

	gock.New("").
		Post("/v2/bookings").
		Reply(200).
		File("./responses/create_booking.json")

	productTypeId := "8f6ef765-4ccd-5e30-ab29-7270569646f1"
	message := "testing"
	arrivalDate := "2018-06-29"

	customer := bemyguest2.Customer{
		Salutation: "Mr.",
		FirstName:  "John",
		LastName:   "Doe",
		Email:      "john@sample.com",
		Phone:      "222333111",
	}

	params := &bemyguest2.CreateBookingParams{}
	params.Message = message
	params.ProductTypeUUID = productTypeId
	params.Customer = customer
	params.Adults = 1
	params.Children = 1
	params.ArrivalDate = arrivalDate
	params.PartnerReference = "random-string"

	option := bemyguest2.Option{
		UUID:  "fe5f4edc-7d76-494d-a49b-e967afd7215a",
		Value: "red",
	}
	options := []bemyguest2.Option{
		option,
	}

	params.Options = bemyguest2.Options{
		PerPax: [][]bemyguest2.Option{
			options,
		},
	}

	resp, err := Create(params)
	assert.Nil(t, err)
	assert.NotNil(t, resp, "response should not be nil")
	assert.NotEqual(t, "", resp.Data.UUID, "booking id should not be empty")
	assert.NotEqual(t, "", resp.Data.Code, "booking code should not be empty")
	assert.Equal(t, "reserved", resp.Data.Status, "booking status should be reserved")
	assert.Equal(t, productTypeId, resp.Data.ProductTypeUUID, "product type id should be equal")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestGetBooking(t *testing.T) {
	defer gock.Off()

	bookingId := "5e9f3deb-d0e3-4c20-80c2-1668e6c919be"
	path := "/v2/bookings/" + bookingId

	gock.New("").
		Get(path).
		Reply(200).
		File("./responses/booking.json")

	resp, err := Get(bookingId)

	assert.Nil(t, err)
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, bookingId, resp.Data.UUID, "booking id should be equal")
	assert.NotEqual(t, "", resp.Data.Code, "booking code should not be empty")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestGetNonExistBooking(t *testing.T) {
	defer gock.Off()

	nonExistBookingId := "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	path := "/v2/bookings/" + nonExistBookingId

	gock.New("").
		Get(path).
		Reply(400).
		File("./responses/error_not_found.json")

	resp, err := Get(nonExistBookingId)

	assert.IsType(t, bemyguest2.ErrorResp{}, err, "should return a not found error")
	assert.Nil(t, resp, "resp should be nil")
	castedErr := err.(bemyguest2.ErrorResp)
	assert.Equal(t, "GEN-WRONG-ARGS", castedErr.Content.Code, "error code should be equal")
	assert.Equal(t, 400, castedErr.Content.HTTPCode, "http code should be equal")
	assert.Equal(t, "Booking not found, UUID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", castedErr.Content.Message, "message should be equal")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestConfirmBooking(t *testing.T) {
	defer gock.Off()

	bookingId := "6b2964d7-e5cd-41d4-bb12-5a892562f2c0"
	path := "/v2/bookings/" + bookingId + "/" + string(StatusConfirm)

	gock.New("").
		Put(path).
		Reply(200).
		File("./responses/confirm_booking.json")

	resp, err := Update(bookingId, StatusConfirm)

	assert.Nil(t, err)
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, "waiting", resp.Data.Status, "booking status should be confirmed")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestCancelBooking(t *testing.T) {
	defer gock.Off()

	bookingId := "ab0df6a1-e42f-4330-ba7d-0aff24963650"
	path := "/v2/bookings/" + bookingId + "/" + string(StatusCancel)

	gock.New("").
		Put(path).
		Reply(200).
		File("./responses/cancel_booking.json")

	resp, err := Update(bookingId, StatusCancel)

	assert.Nil(t, err)
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, "canceled", resp.Data.Status, "booking status should be cancelled")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestListBookings(t *testing.T) {
	defer gock.Off()

	page := 2
	perPage := 1

	gock.New("").
		Get("/v2/bookings").
		MatchParam("page", fmt.Sprintf("%v", page)).
		MatchParam("per_page", fmt.Sprintf("%v", perPage)).
		Reply(200).
		File("./responses/bookings.json")

	params := &bemyguest2.ListBookingsParams{
		Page:    page,
		PerPage: perPage,
	}
	resp, err := List(params)

	assert.Nil(t, err)
	assert.NotNil(t, resp, "response should not be nil")
	assert.NotNil(t, resp.Meta, "meta field should not be nil")
	assert.NotNil(t, resp.Meta.Pagination, "pagination field should not be nil")
	assert.NotNil(t, resp.Data, "data field should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestListBookingVouchers(t *testing.T) {
	defer gock.Off()

	bookingId := "f3359b1b-f718-4ec5-ab31-50f98373d6e7"
	path := fmt.Sprintf("/v2/bookings/%s/vouchers", bookingId)

	gock.New("").
		Get(path).
		Reply(200).
		File("./responses/booking_vouchers.json")

	resp, err := ListVouchers(bookingId)

	assert.Nil(t, err)
	assert.NotNil(t, resp, "response should not be nil")
	//assert.NotNil(t, resp.Meta, "meta field should not be nil")
	//assert.NotNil(t, resp.Meta.Pagination, "pagination field should not be nil")
	assert.NotNil(t, resp.Data, "data field should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}
