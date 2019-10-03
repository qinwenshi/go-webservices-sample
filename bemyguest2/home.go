package bemyguest2

type Home struct {
	Name         string    `json:"name"`
	BaseURL      string    `json:"baseUrl"`
	Encoding     string    `json:"encoding"`
	ContentTypes []string  `json:"contentTypes"`
	Resources    Resources `json:"resources"`
}

type Resources struct {
	Config             Endpoints `json:"config"`
	Products           Endpoints `json:"products"`
	ProductTypes       Endpoints `json:"product-types"`
	ProductTypeBundles Endpoints `json:"product-type-bundles"`
	Bookings           Endpoints `json:"bookings"`
	BulkProductTypes   Endpoints `json:"bulk-product-types"`
	BulkRequests       Endpoints `json:"bulk-requests"`
}

type Endpoints struct {
	Description string   `json:"description"`
	Actions     []Action `json:"actions"`
}

type Action struct {
	Method      string      `json:"method"`
	Endpoint    string      `json:"endpoint"`
	Description string      `json:"description"`
	Example     interface{} `json:"example,omitempty"`
}
