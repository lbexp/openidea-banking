package transaction_model

type Transaction struct {
	TransactionId     string
	UserId            string
	Balance           int
	Currency          string
	ProofImageUrl     string
	BankAccountNumber string
	BankName          string
}
