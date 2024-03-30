package transaction_model

import "time"

type Transaction struct {
	TransactionId     string
	UserId            string
	Balance           int
	Currency          string
	ProofImageUrl     string
	BankAccountNumber string
	BankName          string
	CreatedAt         *time.Time
}
