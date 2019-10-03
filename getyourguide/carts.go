package getyourguide

type CreateCartParams struct {
	BaseData BaseData         `json:"base_data"`
	Data     ShoppingCartData `json:"data"`
}

func NewCreateCartParams(language, currency, cartId, firstName, lastName, email, phoneNumber, countryCode string) *CreateCartParams {

	billing := Billing{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		CountryCode: countryCode,
		PhoneNumber: phoneNumber,
	}

	traveler := &Traveler{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
	}

	payment := Payment{
		MasterBill: true,
	}

	params := &CreateCartParams{
		BaseData: BaseData{
			CntLanguage: language,
			Currency:    currency,
		},
		Data: ShoppingCartData{
			ShoppingCart: ShoppingCart{
				ShoppingCartID: cartId,
				Billing:        billing,
				Traveler:       traveler,
				Payment:        payment,
			},
		},
	}

	return params
}

type ShoppingCartData struct {
	ShoppingCart ShoppingCart `json:"shopping_cart"`
}

type ShoppingCart struct {
	ShoppingCartID string    `json:"shopping_cart_id"`
	Billing        Billing   `json:"billing"`
	Traveler       *Traveler `json:"traveler,omitempty"`
	Payment        Payment   `json:"payment"`
}

type Billing struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`

	// optional
	SalutationCode string `json:"salutation_code,omitempty"`
	// optional
	IsCompany bool `json:"is_company,omitempty,omitempty"`
	// optional
	CompanyName string `json:"company_name,omitempty"`
	// optional
	Invoice bool `json:"invoice,omitempty"`
	// optional
	AddressLine1 string `json:"address_line_1,omitempty"`
	// optional
	AddressLine2 string `json:"address_line_2,omitempty"`
	// optional
	City string `json:"city,omitempty"`
	// optional
	PostalCode string `json:"postal_code,omitempty"`
	// optional
	State string `json:"state,omitempty"`
}

// optional
type Traveler struct {
	SalutationCode string `json:"salutation_code,omitempty"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
}

type Payment struct {
	MasterBill bool `json:"master_bill"`
}

type CreateCartResp struct {
	Meta Meta               `json:"_metadata"`
	Data CreateCartRespData `json:"data"`
}

type CreateCartRespData struct {
	ShoppingCartID   int         `json:"shopping_cart_id"`
	ShoppingCartHash string      `json:"shopping_cart_hash"`
	Billing          Billing     `json:"billing"`
	Traveler         Traveler    `json:"traveler"`
	Status           string      `json:"status"`
	Bookings         []Booking   `json:"bookings"`
	PaymentInfo      PaymentInfo `json:"payment_info"`
}

type PaymentInfo struct {
	PaymentCurrency  string  `json:"payment_currency"`
	TotalPrice       float64 `json:"total_price"`
	PrecouponPrice   float64 `json:"precoupon_price"`
	PaymentMethod    string  `json:"payment_method"`
	InvoiceReference string  `json:"invoice_reference"`
}
