package balance_model

type BalanceGetData struct {
	Balance int    `json:"balance"`
	Currecy string `json:"currency"`
}

type BalanceGetResponse struct {
	Message string           `json:"message"`
	Data    []BalanceGetData `json:"data"`
}
