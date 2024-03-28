package repository

import (
	"context"
	user_model "openidea-banking/src/model/user"
	"openidea-banking/src/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Register(ctx context.Context, tx pgx.Tx, user user_model.User) (user_model.User, error)
	Login(ctx context.Context, conn *pgxpool.Pool, user user_model.User) (user_model.User, error)
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, tx pgx.Tx, user user_model.User) (user_model.User, error) {
	QUERY := "INSERT INTO users(user_id, email, password, name) values (gen_random_uuid(), $1, $2, $3) ON CONFLICT(email) DO NOTHING RETURNING user_id"

	var userId string

	err := tx.QueryRow(ctx, QUERY, user.Email, user.Password, user.Name).Scan(&userId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return user_model.User{}, utils.ErrorConflict
		} else {
			return user_model.User{}, utils.ErrorInternalServer
		}
	}

	user.UserId = userId

	return user, nil
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, conn *pgxpool.Pool, user user_model.User) (user_model.User, error) {
	QUERY := "SELECT user_id, email, password, name FROM users WHERE email = $1"

	userResult := user_model.User{}

	err := conn.QueryRow(ctx, QUERY, user.Email).Scan(
		&userResult.UserId,
		&userResult.Email,
		&userResult.Password,
		&userResult.Password,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return user_model.User{}, utils.ErrorConflict
		} else {
			return user_model.User{}, utils.ErrorInternalServer
		}
	}

	return userResult, nil
}
