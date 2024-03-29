package service

import (
	"context"
	transaction_model "openidea-banking/src/model/transaction"
	"openidea-banking/src/repository"
	"openidea-banking/src/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionService interface {
	Create(ctx context.Context, transaction transaction_model.Transaction) error
}

type TransactionServiceImpl struct {
	DBPool                *pgxpool.Pool
	TransactionRepository repository.TransactionRepository
}

func NewTransactionService(
	dbPool *pgxpool.Pool,
	transactionRepository repository.TransactionRepository,
) TransactionService {
	return &TransactionServiceImpl{
		DBPool:                dbPool,
		TransactionRepository: transactionRepository,
	}
}

func (service *TransactionServiceImpl) Create(ctx context.Context, transaction transaction_model.Transaction) error {
	tx, err := service.DBPool.Begin(ctx)
	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback(ctx)

			if rollbackErr != nil {
				err = rollbackErr
			}
		}
	}()

	if err != nil {
		return utils.ErrorInternalServer
	}

	err = service.TransactionRepository.Create(ctx, tx, transaction)
	if err != nil {
		return err
	}

	tx.Commit(ctx)

	return nil
}
