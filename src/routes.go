package src

import (
	"openidea-banking/src/controller"
	"openidea-banking/src/middleware"
	"openidea-banking/src/repository"
	"openidea-banking/src/security"
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
	transactionRepository := repository.NewTransactionRepository()
	balanceRepository := repository.NewBalanceRepository()

	userService := service.NewUserService(dbPool, userRepository, authService)
	transactionService := service.NewTransactionService(dbPool, transactionRepository, balanceRepository)
	balanceService := service.NewBalanceService(dbPool, balanceRepository, transactionRepository)
	imageService := service.NewImageService(security.GetAwsS3Session())

	userController := controller.NewUserController(validator, userService)
	transactionController := controller.NewTransactionController(validator, transactionService, authService)
	balanceController := controller.NewBalanceController(validator, balanceService, authService)
	imageController := controller.NewImageController(authService, imageService)

	app.Post("/v1/user/register", userController.Register)
	app.Post("/v1/user/login", userController.Login)

	app.Use(middleware.JWTHeaderMiddleware)
	app.Use(middleware.JWTTokenMiddleware())

	app.Post("/v1/transaction", transactionController.Create)

	app.Post("/v1/balance", balanceController.Upsert)
	app.Get("/v1/balance", balanceController.GetAll)
	app.Get("/v1/balance/history", transactionController.GetAllByUserId)

	app.Post("/v1/image", imageController.UploadImage)
}
