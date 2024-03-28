package balance_model

type BalanceRequest struct {
	SenderBankAccountNumber string `json:"senderBankAccountNumber" validate:"required,min=5,max=30"`
	SenderBankName          string `json:"senderBankName" validate:"required,min=5,max=30"`
	AddedBalance            int    `json:"addedBalance" validate:"required,gt=0"`
	Currency                string `json:"currency" validate:"required,iso4217"`
	TransferProofImg        string `json:"transferProofImg" validate:"required,imageurl"`
}
