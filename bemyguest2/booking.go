package bemyguest2

import "time"

type BookingResp struct {
	Data Booking `json:"data"`
}

type BookingsResp struct {
	Data []Booking `json:"data"`
	Meta Meta      `json:"meta"`
}

func (r *BookingsResp) HasNext() bool {
	return r.Meta.Pagination.HasNext()
}

func (r *BookingsResp) GetNextPage() int {
	return r.Meta.Pagination.GetNextPage()
}

type ListBookingsParams struct {
	DateStart time.Time
	DateEnd   time.Time
	FirstName *string
	LastName  *string
	Email     *string
	Phone     *string
	// represents the given unique partnerReference ID of this booking
	PartnerReference *string
	Page             int
	PerPage          int
	// represents free phrase for text search for example &query=John
	Query *string
	// represents booking status, available values are
	// reserved
	// waiting
	// cancellation_requested
	// cancelled
	// approved
	// expired
	// rejected
	// refunded
	// refund_declined
	Status *string
}

type Customer struct {
	// salutation represents available options "Mr.", "Ms.", "Mrs."
	Salutation string `json:"salutation"`
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

type AmountBreakdown struct {
	Name                 string  `json:"name"`
	Quantity             int     `json:"quantity"`
	Price                float64 `json:"price"`
	PriceRequestCurrency float64 `json:"priceRequestCurrency"`
}

type CreateBookingParams struct {
	Message         string   `json:"message"`
	ProductTypeUUID string   `json:"productTypeUuid"`
	Customer        Customer `json:"customer"`
	Adults          int      `json:"adults"`
	Children        int      `json:"children"`
	Seniors         int      `json:"seniors"`

	// optional
	TimeSlotUUID     *string `json:"timeSlotUuid"`
	ArrivalDate      string  `json:"arrivalDate"`
	PartnerReference string  `json:"partnerReference"`
	Options          Options `json:"options"`
}

type Options struct {
	PerBooking []Option   `json:"perBooking"`
	PerPax     [][]Option `json:"perPax"`
}

type Option struct {
	UUID  string `json:"uuid"`
	Label string `json:"label"`
	Value string `json:"value"`
	// type inconsistent in POST /bookings and GET /bookings/:id
	Price interface{} `json:"price"`
}

type CreateBookingResp struct {
	Data Booking `json:"data"`
}

type Booking struct {
	UUID                       string            `json:"uuid"`
	Code                       string            `json:"code"`
	PartnerReference           *string           `json:"partnerReference"`
	Status                     string            `json:"status"`
	ProductTypeTitle           string            `json:"productTypeTitle"`
	ProductTypeTitleTranslated string            `json:"productTypeTitleTranslated"`
	ProductTypeUUID            string            `json:"productTypeUuid"`
	CurrencyCode               string            `json:"currencyCode"`
	CurrencyUUID               string            `json:"currencyUuid"`
	TotalAmount                float64           `json:"totalAmount"`
	AmountBreakdown            []AmountBreakdown `json:"amountBreakdown"`
	TotalAmountRequestCurrency float64           `json:"totalAmountRequestCurrency"`
	RequestCurrencyCode        string            `json:"requestCurrencyCode"`
	RequestCurrencyUUID        string            `json:"requestCurrencyUuid"`
	ArrivalDate                string            `json:"arrivalDate"`
	CreatedAt                  string            `json:"createdAt"`
	UpdatedAt                  string            `json:"updatedAt"`
	Salutation                 string            `json:"salutation"`
	FirstName                  string            `json:"firstName"`
	LastName                   string            `json:"lastName"`
	Email                      string            `json:"email"`
	Phone                      string            `json:"phone"`
	Adults                     int               `json:"adults"`
	Children                   int               `json:"children"`
	Seniors                    int               `json:"seniors"`
	Options                    []Option          `json:"options"`
	CompletedAt                *string           `json:"completedAt"`
	CancellationRequestAt      *string           `json:"cancellationRequestAt"`
	CancellationRequestStatus  *string           `json:"cancellationRequestStatus"`
	CancellationStatus         *string           `json:"cancellationStatus"`
	RefundDate                 *string           `json:"refundDate"`
	RefundAmount               *float64          `json:"refundAmount"`
	RefundTransaction          interface{}       `json:"refundTransaction"`
	Links                      []Link            `json:"links"`
}
