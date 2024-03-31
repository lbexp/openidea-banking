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
	GetAllByUserId(ctx context.Context, userId string, filters transaction_model.GetAllByUserIdFilters) ([]transaction_model.Transaction, int, error)
}

type TransactionServiceImpl struct {
	DBPool                *pgxpool.Pool
	TransactionRepository repository.TransactionRepository
	BalanceRepository     repository.BalanceRepository
}

func NewTransactionService(
	dbPool *pgxpool.Pool,
	transactionRepository repository.TransactionRepository,
	balanceRepository repository.BalanceRepository,
) TransactionService {
	return &TransactionServiceImpl{
		DBPool:                dbPool,
		TransactionRepository: transactionRepository,
		BalanceRepository:     balanceRepository,
	}
}

func (service *TransactionServiceImpl) Create(ctx context.Context, transaction transaction_model.Transaction) error {
	balance, err := service.BalanceRepository.GetByUserIdAndCurrency(ctx, service.DBPool, transaction.UserId, transaction.Currency)
	if err != nil {
		return err
	}

	if transaction.Balance > balance.Balance {
		return utils.ErrorBadRequest
	}

	balance.Balance = -transaction.Balance
	transaction.Balance = -transaction.Balance

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

	err = service.BalanceRepository.Upsert(ctx, tx, balance)
	if err != nil {
		return err
	}

	tx.Commit(ctx)

	return nil
}

func (service *TransactionServiceImpl) GetAllByUserId(ctx context.Context, userId string, filters transaction_model.GetAllByUserIdFilters) ([]transaction_model.Transaction, int, error) {
	transactions, totalItems, err := service.TransactionRepository.GetAllByUserId(ctx, service.DBPool, userId, filters)
	if err != nil {
		return nil, 0, err
	}

	return transactions, totalItems, nil
}
