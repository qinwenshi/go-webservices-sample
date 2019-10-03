package bookings

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

func TestCreateCart(t *testing.T) {
	defer gock.Off()

	gock.New("").
		Post("/carts").
		Reply(200).
		File("./responses/create_cart.json")

	language := "en"
	currency := "gbp"
	cartId := "45942258"
	firstName := "first_name"
	lastName := "last_name"
	email := "testing@example.com"
	phoneNumber := "1234567890"
	countryCode := "HK"

	params := getyourguide.NewCreateCartParams(language, currency, cartId, firstName, lastName, email, phoneNumber, countryCode)
	resp, err := Create(params)

	assert.Nil(t, err, "should not have error")
	assert.IsType(t, getyourguide.Meta{}, resp.Meta, "should have meta in response")
	assert.Equal(t, "OK", resp.Meta.Status, "status should be OK")
	assert.Equal(t, "confirmShoppingCartAction", resp.Meta.Method, "method should be createBookingAction")
	assert.IsType(t, getyourguide.Exchange{}, resp.Meta.Exchange, "should have exhange in meta")
	assert.Equal(t, 0, resp.Meta.Offset, "offset should be 0")
	assert.Equal(t, 10, resp.Meta.Limit, "limit should be 10")

	assert.IsType(t, getyourguide.CreateCartRespData{}, resp.Data, "should have data in cart")
	assert.Equal(t, 45942258, resp.Data.ShoppingCartID, "cart id should be 45942258")
	assert.Equal(t, "completed", resp.Data.Status, "status should be completed")

	// asserts billing struct
	billing := resp.Data.Billing
	assert.Equal(t, firstName, billing.FirstName, "first name should be first_name")
	assert.Equal(t, lastName, billing.LastName, "last name should be last_name")
	assert.Equal(t, email, billing.Email, "email should be testing@example.com")
	assert.Equal(t, phoneNumber, billing.PhoneNumber, "phone number should be 1234567890")
	assert.Equal(t, countryCode, billing.CountryCode, "country code should be HK")

	// asserts bookings
	bookings := resp.Data.Bookings
	assert.Equal(t, 1, len(bookings), "number of booking should be 1")

	booking := bookings[0]
	assert.Equal(t, "confirmed", booking.BookingStatus, "booking status should be confirmed")
	assert.Equal(t, 45942258, booking.ShoppingCartID, "booking cart id should be 45942258")

	// asserts ticket
	assert.NotEmpty(t, booking.Ticket.TicketURL, "booking ticket url should bit be empty")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestCreateCartAlreadyProcessedError(t *testing.T) {
	defer gock.Off()

	gock.New("").
		Post("/carts").
		Reply(400).
		File("./responses/cart_is_already_processed.json")

	language := "en"
	currency := "gbp"
	cartId := "45942258"
	firstName := "first_name"
	lastName := "last_name"
	email := "testing@example.com"
	phoneNumber := "1234567890"
	countryCode := "HK"

	params := getyourguide.NewCreateCartParams(language, currency, cartId, firstName, lastName, email, phoneNumber, countryCode)
	resp, err := Create(params)

	assert.NotNil(t, err, "should be error")
	assert.IsType(t, getyourguide.ErrorResp{}, err, "should be a getyourguide error response")
	assert.Nil(t, resp, "response should be nil")

	e := err.(getyourguide.ErrorResp)
	assert.Equal(t, "GetYourGuide AG", e.Descriptor, "descriptor should be GetYourGuide AG")
	assert.Equal(t, "1", e.APIVersion, "apiVersion should be 1")
	assert.Equal(t, "confirmShoppingCartAction", e.Method, "method should be confirmShoppingCartAction")
	assert.Equal(t, "ERROR", e.Status, "status should be ERROR")
	assert.Equal(t, 1, len(e.Errors), "number of errors should be 1")
	assert.Equal(t, "Shopping cart is already processed", e.Errors[0].ErrorMessage, "error message should be Shopping cart is already processed")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}
