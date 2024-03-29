package transaction_model

type TransactionRequest struct {
	RecipientBankAccountNumber string `json:"recipientBankAccountNumber" validate:"required,min=5,max=30"`
	RecipientBankName          string `json:"recipientBankName" validate:"required,min=5,max=30"`
	FromCurrency               string `json:"fromCurrency" validate:"required,iso4217"`
	Balances                   int    `json:"balances" validate:"required,number"`
}

type GetAllByUserIdFilters struct {
	Limit  int `json:"limit" validate:"number,gte=0"`
	Offset int `json:"offset" validate:"number,gte=0"`
}
