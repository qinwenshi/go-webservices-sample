package bemyguest2

type Config struct {
	Data Data `json:"data"`
}

type Data struct {
	Now Now `json:"now"`
	Version           int          `json:"version"`
	ServerURL         string       `json:"serverUrl"`
	PhotosURL         string       `json:"photosUrl"`
	ActivitiesSorting []string     `json:"activitiesSorting"`
	User              User         `json:"user"`
	LanguageData      LanguageData `json:"languages"`
	CurrencyData      CurrencyData `json:"currencies"`
	TypeData          TypeData     `json:"types"`
	CategoryData      CategoryData `json:"categories"`
	LocationData      LocationData `json:"locations"`
	OptionInputTypes      []OptionInputType `json:"optionInputTypes"`
}

type User struct {
	Name                   string      `json:"name"`
	Email                  string      `json:"email"`
	UUID                   string      `json:"uuid"`
	ContinueURL            string      `json:"continueUrl"`
	NotifyURL              string      `json:"notifyUrl"`
	DefaultPagination      int         `json:"defaultPagination"`
	DefaultSortBy          string      `json:"defaultSortBy"`
	DefaultCurrencyUUID    string      `json:"defaultCurrencyUuid"`
	DefaultCurrencyCode    string `json:"defaultCurrencyCode"`
	DefaultLanguageUUID    string `json:"defaultLanguageUuid"`
	DefaultLanguageCode    string `json:"defaultLanguageCode"`
	WalletBalance          float32         `json:"walletBalance"`
	WalletBlockedBalance   float32         `json:"walletBlockedBalance"`
	WalletAvailableBalance float32         `json:"walletAvailableBalance"`
	WalletAlertValue       float32         `json:"walletAlertValue"`
}

type LanguageData struct {
	Data []Language `json:"data"`
}

type Language struct {
	Name string `json:"name"`
	Code string `json:"code"`
	UUID string `json:"uuid"`
}

type CurrencyData struct {
	Data []Currency `json:"data"`
}

type Currency struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Symbol string `json:"symbol"`
	UUID   string `json:"uuid"`
}

type TypeData struct {
	Data []Type `json:"data"`
}

type Type struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type CategoryData struct {
	Data []Category `json:"data"`
}

type Category struct {
	Name     string     `json:"name"`
	UUID     string     `json:"uuid"`
	Children []Children `json:"children"`
}

type Children struct {
	Name        string        `json:"name"`
	UUID        string        `json:"uuid"`
	SubChildren []SubChildren `json:"children"`
}

type SubChildren struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type LocationData struct {
	Data []Location `json:"data"`
}

type Location struct {
	Continent   string      `json:"continent"`
	Code        string      `json:"code"`
	UUID        string      `json:"uuid"`
	CountryData CountryData `json:"countries"`
}

type CountryData struct {
	Data []Country `json:"data"`
}

type Country struct {
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	UUID      string    `json:"uuid"`
	StateData StateData `json:"states"`
}

type StateData struct {
	Data []State `json:"data"`
}

type State struct {
	Name     string   `json:"name"`
	UUID     string   `json:"uuid"`
	CityData CityData `json:"cities"`
}

type CityData struct {
	Data []City `json:"data"`
}

type City struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type Now struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}

type OptionInputType struct {
	Type1  string `json:"1"`
	Type2  string `json:"2"`
	Type3  string `json:"3"`
	Type4  string `json:"4"`
	Type5  string `json:"5"`
	Type6  string `json:"6"`
	Type7  string `json:"7"`
	Type8  string `json:"8"`
	Type9  string `json:"9"`
	Type10 string `json:"10"`
	Type11 string `json:"11"`
	Type12 string `json:"12"`
	Type13 string `json:"13"`
	Type14 string `json:"14"`
}