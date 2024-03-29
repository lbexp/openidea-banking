package repository

import (
	"context"
	transaction_model "openidea-banking/src/model/transaction"
	"openidea-banking/src/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository interface {
	Create(ctx context.Context, tx pgx.Tx, transaction transaction_model.Transaction) error
	GetAllByUserId(ctx context.Context, conn *pgxpool.Pool, userId string, filters transaction_model.GetAllByUserIdFilters) ([]transaction_model.Transaction, int, error)
}

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (repository *TransactionRepositoryImpl) Create(ctx context.Context, tx pgx.Tx, transaction transaction_model.Transaction) error {
	QUERY := "INSERT INTO transactions(user_id, currency, balance, proof_image_url, bank_account_number, bank_name) values($1, $2, $3, $4, $5, $6) RETURNING transaction_id"

	res, err := tx.Exec(ctx, QUERY, transaction.UserId, transaction.Currency, transaction.Balance, transaction.ProofImageUrl, transaction.BankAccountNumber, transaction.BankName)
	if err != nil {
		if err == pgx.ErrNoRows {
			return utils.ErrorConflict
		} else {
			return utils.ErrorInternalServer
		}
	}

	if res.RowsAffected() == 0 {
		return utils.ErrorInternalServer
	}

	return nil
}

func (repository *TransactionRepositoryImpl) GetAllByUserId(ctx context.Context, conn *pgxpool.Pool, userId string, filters transaction_model.GetAllByUserIdFilters) ([]transaction_model.Transaction, int, error) {
	query := transaction_model.BuildGetAllByUserIdQuery(filters)

	rows, err := conn.Query(ctx, query, userId)
	if err != nil {
		return nil, 0, utils.ErrorInternalServer
	}
	defer rows.Close()

	var transactions []transaction_model.Transaction
	var totalItems int

	for rows.Next() {
		transaction := transaction_model.Transaction{}

		err = rows.Scan(
			&transaction.TransactionId,
			&transaction.Balance,
			&transaction.Currency,
			&transaction.ProofImageUrl,
			&transaction.CreatedAt,
			&transaction.BankAccountNumber,
			&transaction.BankName,
			&totalItems,
		)
		if err != nil {
			return nil, 0, utils.ErrorInternalServer
		}

		transactions = append(transactions, transaction)
	}

	return transactions, totalItems, nil
}
