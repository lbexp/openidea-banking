package src

import (
	"openidea-banking/src/controller"
	"openidea-banking/src/repository"
	"openidea-banking/src/service"
	"openidea-banking/src/validation"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoute(app *fiber.App, dbPool *pgxpool.Pool) {
	validator := validator.New()
	validation.RegisterValidation(validator)

	authService := service.NewAuthService()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(dbPool, userRepository, authService)
	userController := controller.NewUserController(validator, userService)

	transactionRepository := repository.NewTransactionRepository()

	balanceRepository := repository.NewBalanceRepository()
	balanceService := service.NewBalanceService(dbPool, balanceRepository, transactionRepository)
	balanceController := controller.NewBalanceController(validator, balanceService, authService)

	app.Post("/v1/user/register", userController.Register)
	app.Post("/v1/user/login", userController.Login)

	app.Post("/v1/balance", balanceController.Upsert)
	app.Get("/v1/balance", balanceController.GetAll)
}
