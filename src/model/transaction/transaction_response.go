package transaction_model

import common_model "openidea-banking/src/model/common"

type TransactionSource struct {
	BankAccountNumber string `json:"bankAccountNumber"`
	BankName          string `json:"bankName"`
}

type TransactionGetAllData struct {
	TransactionId    string            `json:"transactionId"`
	Balance          int               `json:"balance"`
	Currency         string            `json:"currency"`
	TransferProofImg string            `json:"transferProofImg"`
	CreatedAt        string            `json:"createdAt"`
	Source           TransactionSource `json:"source"`
}

type TransactionGetAllResponse struct {
	Message string                    `json:"message"`
	Data    []TransactionGetAllData   `json:"data"`
	Meta    common_model.MetaResponse `json:"meta"`
}
