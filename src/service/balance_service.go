package service

import (
	"context"
	balance_model "openidea-banking/src/model/balance"
	transaction_model "openidea-banking/src/model/transaction"
	"openidea-banking/src/repository"
	"openidea-banking/src/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BalanceService interface {
	Upsert(ctx context.Context, transaction transaction_model.Transaction) error
	GetAll(ctx context.Context, userId string) ([]balance_model.Balance, error)
}

type BalanceServiceImpl struct {
	DBPool                *pgxpool.Pool
	BalanceRepository     repository.BalanceRepository
	TransactionRepository repository.TransactionRepository
}

func NewBalanceService(
	dbPool *pgxpool.Pool,
	balanceRepository repository.BalanceRepository,
	transactionRepository repository.TransactionRepository,
) BalanceService {
	return &BalanceServiceImpl{
		DBPool:                dbPool,
		BalanceRepository:     balanceRepository,
		TransactionRepository: transactionRepository,
	}
}

func (service *BalanceServiceImpl) Upsert(ctx context.Context, transaction transaction_model.Transaction) error {
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

	err = service.BalanceRepository.Upsert(ctx, tx, balance_model.Balance{
		UserId:   transaction.UserId,
		Currency: transaction.Currency,
		Balance:  transaction.Balance,
	})
	if err != nil {
		return err
	}

	err = service.TransactionRepository.Create(ctx, tx, transaction)
	if err != nil {
		return err
	}

	return nil
}

func (service *BalanceServiceImpl) GetAll(ctx context.Context, userId string) ([]balance_model.Balance, error) {
	balances, err := service.BalanceRepository.GetAll(ctx, service.DBPool, userId)
	if err != nil {
		return nil, err
	}

	return balances, nil
}
