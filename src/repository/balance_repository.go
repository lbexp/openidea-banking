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
	GetAll(ctx context.Context, conn *pgxpool.Pool, userId string) ([]balance_model.Balance, error)
}

type BalanceRepositoryImpl struct {
}

func NewBalanceRepository() BalanceRepository {
	return &BalanceRepositoryImpl{}
}

func (repository *BalanceRepositoryImpl) Upsert(ctx context.Context, tx pgx.Tx, balance balance_model.Balance) error {
	QUERY := "INSERT INTO balances(balance_id, user_id, currency, balance) values(gen_random_uuid(), $1, $2, $3) ON CONFLICT(user_id,currency) DO UPDATE SET balance = balance + $3 RETURNING balance_id"

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

func (repository *BalanceRepositoryImpl) GetAll(ctx context.Context, conn *pgxpool.Pool, userId string) ([]balance_model.Balance, error) {
	QUERY := "SELECT balance, currency FROM balances WHERE user_id = $1"

	rows, err := conn.Query(ctx, QUERY, userId)
	if err != nil {
		return nil, utils.ErrorInternalServer
	}
	defer rows.Close()

	var balances []balance_model.Balance

	for rows.Next() {
		balance := balance_model.Balance{}

		err = rows.Scan(
			&balance.BalanceId,
			&balance.Currency,
		)
		if err != nil {
			return nil, utils.ErrorInternalServer
		}

		balances = append(balances, balance)
	}

	return balances, nil
}
