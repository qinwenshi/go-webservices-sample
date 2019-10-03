package getyourguide

import (
	"fmt"
	"time"
)

type GetBookingParams struct {
	Language string
	Currency string
	Id       string
}

type ListBookingsParams struct {
	Language string
	Currency string

	// offsets represent offset of paging
	// default 0
	Offset int

	// how many results per response
	// default 10
	Limit int
}

type BookingsResp struct {
	Meta Meta             `json:"_metadata"`
	Data BookingsRespData `json:"data"`
}

type BookingsRespData struct {
	Bookings []Booking `json:"bookings"`
}

type Booking struct {
	BookingID        int      `json:"booking_id"`
	BookingHash      string   `json:"booking_hash"`
	BookingStatus    string   `json:"booking_status"`
	ShoppingCartID   int      `json:"shopping_cart_id"`
	ShoppingCartHash string   `json:"shopping_cart_hash"`
	Bookable         Bookable `json:"bookable"`
	Ticket           Ticket   `json:"ticket"`
}

type Bookable struct {
	TourID                 int                         `json:"tour_id"`
	OptionID               int                         `json:"option_id"`
	Datetime               string                      `json:"datetime"`
	DatetimeUtc            time.Time                   `json:"datetime_utc"`
	DatetimeType           string                      `json:"datetime_type"`
	Price                  float64                     `json:"price"`
	Categories             []BookableCategory          `json:"categories"`
	BookingParameters      []BookableBookingParameters `json:"booking_parameters"`
	ValidUntil             string                      `json:"valid_until"`
	CancellationPolicyText string                      `json:"cancellation_policy_text"`
}

type BookableCategory struct {
	CategoryID           int    `json:"category_id"`
	Name                 string `json:"name"`
	NumberOfParticipants int    `json:"number_of_participants"`
}

func NewBookableCategory(categoryId, numberOfParticipant int) *BookableCategory {
	return &BookableCategory{
		CategoryID:           categoryId,
		NumberOfParticipants: numberOfParticipant,
	}
}

type BookableBookingParameters struct {
	Name   string `json:"name"`
	Value1 string `json:"value_1"`
	Value2 string `json:"value_2"`
}

type Ticket struct {
	BookingReference     string                `json:"booking_reference"`
	ExternalReferenceID  string                `json:"external_reference_id"`
	SupplierBookingCodes []SupplierBookingCode `json:"supplier_booking_codes"`
	VoucherInformation   string                `json:"voucher_information"`
	EmergencyPhoneNumber string                `json:"emergency_phone_number"`
	EmergencyEmail       string                `json:"emergency_email"`
	TicketURL            string                `json:"ticket_url"`
	TicketHash           string                `json:"ticket_hash"`
}

type SupplierBookingCode struct {
	TicketHash string `json:"ticket_hash"`
	Label      string `json:"label"`
	Type       string `json:"type"`
	Code       string `json:"code"`
}

type CreateBookingParams struct {
	BaseData BaseData    `json:"base_data"`
	Data     BookingData `json:"data"`
}

type CateogryId int
type NumberOfParticipant int

func NewCreateBookingParams(language, currency, optionId string, date time.Time, inputs map[CateogryId]NumberOfParticipant, price float32, referenceId string) *CreateBookingParams {
	categories := make([]BookableCategory, 0, len(inputs))

	for categoryId, numberOfParticipant := range inputs {
		category := NewBookableCategory(int(categoryId), int(numberOfParticipant))
		categories = append(categories, *category)
	}

	bookable := CreateBookable{
		ExternalReferenceID: referenceId,
		OptionID:            optionId,
		Datetime:            date.Format("2006-01-02T15:04:05"),
		Price:               fmt.Sprintf("%v", price),
		Categories:          categories,
	}

	params := &CreateBookingParams{
		BaseData: BaseData{
			CntLanguage: language,
			Currency:    currency,
		},
		Data: BookingData{
			Booking: BookingReq{
				Bookable: bookable,
			},
		},
	}

	return params
}

type BaseData struct {
	CntLanguage string `json:"cnt_language"`
	Currency    string `json:"currency"`
}

type BookingData struct {
	Booking BookingReq `json:"booking"`
}

type BookingReq struct {
	Bookable CreateBookable `json:"bookable"`
}

type CreateBookable struct {
	ExternalReferenceID string             `json:"external_reference_id"`
	OptionID            string             `json:"option_id"`
	Datetime            string             `json:"datetime"`
	Price               string             `json:"price"`
	Categories          []BookableCategory `json:"categories"`
}

type CreateBookingResp struct {
	Meta Meta                  `json:"_metadata"`
	Data CreateBookingRespData `json:"data"`
}

type CreateBookingRespData struct {
	Bookings Booking2 `json:"bookings"`
}

type Booking2 struct {
	ShoppingCartID   int    `json:"shopping_cart_id"`
	ShoppingCartHash string `json:"shopping_cart_hash"`
	BookingID        int    `json:"booking_id"`
	BookingHash      string `json:"booking_hash"`
	Status           string `json:"status"`
	ReturnCode       int    `json:"return_code"`
}

type BookingResp struct {
	Meta Meta            `json:"_metadata"`
	Data BookingRespData `json:"data"`
}

type BookingRespData struct {
	Bookings Booking `json:"booking"`
}
