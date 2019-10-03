package currency

type ConvertResp struct {
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	Result float64 `json:"result"`
	To     string  `json:"to"`
}
