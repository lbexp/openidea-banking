package service

import (
	"context"
	user_model "openidea-banking/src/model/user"
	"openidea-banking/src/repository"
	"openidea-banking/src/security"
	"openidea-banking/src/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService interface {
	Register(ctx context.Context, user user_model.User) (user_model.User, error)
	Login(ctx context.Context, user user_model.User) (user_model.User, error)
}

type UserServiceImpl struct {
	DBPool         *pgxpool.Pool
	UserRepository repository.UserRepository
	AuthService    AuthService
}

func NewUserService(dbPool *pgxpool.Pool, userRepository repository.UserRepository, authService AuthService) UserService {
	return &UserServiceImpl{
		DBPool:         dbPool,
		UserRepository: userRepository,
		AuthService:    authService,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, user user_model.User) (user_model.User, error) {
	tx, err := service.DBPool.Begin(ctx)
	if err != nil {
		return user_model.User{}, utils.ErrorInternalServer
	}

	hashedPass, err := security.GenerateHashedPassword(user.Password)
	if err != nil {
		tx.Rollback(ctx)
		return user_model.User{}, err
	}

	user.Password = hashedPass

	userResult, err := service.UserRepository.Register(ctx, tx, user)
	if err != nil {
		tx.Rollback(ctx)
		return user_model.User{}, err
	}

	accessToken, err := service.AuthService.GenerateToken(ctx, userResult.UserId)
	if err != nil {
		tx.Rollback(ctx)
		return user_model.User{}, err
	}

	userResult.Password = hashedPass
	userResult.AccessToken = accessToken

	tx.Commit(ctx)

	return userResult, nil
}

func (service *UserServiceImpl) Login(ctx context.Context, user user_model.User) (user_model.User, error) {
	userResult, err := service.UserRepository.Login(ctx, service.DBPool, user)
	if err != nil {
		return user_model.User{}, err
	}

	err = security.ComparePassword(userResult.Password, user.Password)
	if err != nil {
		return user_model.User{}, err
	}

	accessToken, err := service.AuthService.GenerateToken(ctx, userResult.UserId)
	if err != nil {
		return user_model.User{}, err
	}

	userResult.AccessToken = accessToken

	return userResult, nil
}
