package repository

import (
	"context"
	balance_model "openidea-banking/src/model/balance"
	"openidea-banking/src/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BalanceRepository interface {
	Upsert(ctx context.Context, tx pgx.Tx, balance balance_model.Balance) error
	Update(ctx context.Context, tx pgx.Tx, balance balance_model.Balance) error
	GetAll(ctx context.Context, conn *pgxpool.Pool, userId string) ([]balance_model.Balance, error)
	GetByUserIdAndCurrency(ctx context.Context, conn *pgxpool.Pool, userId string, currency string) (balance_model.Balance, error)
}

type BalanceRepositoryImpl struct {
}

func NewBalanceRepository() BalanceRepository {
	return &BalanceRepositoryImpl{}
}

func (repository *BalanceRepositoryImpl) Upsert(ctx context.Context, tx pgx.Tx, balance balance_model.Balance) error {
	QUERY := "INSERT INTO balances(balance_id, user_id, currency, balance) VALUES(gen_random_uuid(), $1, $2, $3) ON CONFLICT(user_id,currency) DO UPDATE SET balance = balances.balance + $3"

	res, err := tx.Exec(ctx, QUERY, balance.UserId, balance.Currency, balance.Balance)
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

func (repository *BalanceRepositoryImpl) Update(ctx context.Context, tx pgx.Tx, balance balance_model.Balance) error {
	QUERY := "UPDATE balances SET balance = balance + $1 WHERE balance_id = $2"

	res, err := tx.Exec(ctx, QUERY, balance.Balance, balance.BalanceId)
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

func (repository *BalanceRepositoryImpl) GetAll(ctx context.Context, conn *pgxpool.Pool, userId string) ([]balance_model.Balance, error) {
	QUERY := "SELECT balance, currency FROM balances WHERE user_id = $1 ORDER BY balance DESC"

	rows, err := conn.Query(ctx, QUERY, userId)
	if err != nil {
		return nil, utils.ErrorInternalServer
	}
	defer rows.Close()

	var balances []balance_model.Balance

	for rows.Next() {
		balance := balance_model.Balance{}

		err = rows.Scan(
			&balance.Balance,
			&balance.Currency,
		)
		if err != nil {
			return nil, utils.ErrorInternalServer
		}

		balances = append(balances, balance)
	}

	return balances, nil
}

func (repository *BalanceRepositoryImpl) GetByUserIdAndCurrency(ctx context.Context, conn *pgxpool.Pool, userId string, currency string) (balance_model.Balance, error) {
	QUERY := "SELECT balance_id, balance, currency FROM balances WHERE user_id = $1 AND currency = $2"

	balance := balance_model.Balance{
		UserId: userId,
	}
	err := conn.QueryRow(ctx, QUERY, userId, currency).Scan(
		&balance.BalanceId,
		&balance.Balance,
		&balance.Currency,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return balance_model.Balance{}, utils.ErrorBadRequest
		} else {
			return balance_model.Balance{}, utils.ErrorInternalServer
		}
	}

	return balance, nil
}
