package controller

import (
	transaction_model "openidea-banking/src/model/transaction"
	"openidea-banking/src/service"
	"openidea-banking/src/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TransactionController interface {
	Create(ctx *fiber.Ctx) error
}

type TransactionControllerImpl struct {
	Validator          *validator.Validate
	TransactionService service.TransactionService
	AuthService        service.AuthService
}

func NewTransactionController(
	validator *validator.Validate,
	transactionService service.TransactionService,
	authService service.AuthService,
) TransactionController {
	return &TransactionControllerImpl{
		Validator:          validator,
		TransactionService: transactionService,
		AuthService:        authService,
	}
}

func (controller *TransactionControllerImpl) Create(ctx *fiber.Ctx) error {
	request := new(transaction_model.TransactionRequest)

	err := ctx.BodyParser(request)
	if err != nil {
		return utils.ErrorBadRequest
	}

	err = controller.Validator.Struct(request)
	if err != nil {
		return utils.ErrorBadRequest
	}

	userId, err := controller.AuthService.GetValidUser(ctx)
	if err != nil {
		return err
	}

	err = controller.TransactionService.Create(ctx.UserContext(), transaction_model.Transaction{
		UserId:            userId,
		Balance:           request.Balances,
		Currency:          request.FromCurrency,
		BankAccountNumber: request.RecipientBankAccountNumber,
		BankName:          request.RecipientBankName,
	})
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON("Success")
}
