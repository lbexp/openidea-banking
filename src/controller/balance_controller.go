package controller

import (
	balance_model "openidea-banking/src/model/balance"
	transaction_model "openidea-banking/src/model/transaction"
	"openidea-banking/src/service"
	"openidea-banking/src/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type BalanceController interface {
	Upsert(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
}

type BalanceControllerImpl struct {
	Validator      *validator.Validate
	BalanceService service.BalanceService
	AuthService    service.AuthService
}

func NewBalanceController(
	validator *validator.Validate,
	balanceService service.BalanceService,
	authService service.AuthService,
) BalanceController {
	return &BalanceControllerImpl{
		Validator:      validator,
		BalanceService: balanceService,
		AuthService:    authService,
	}
}

func (controller *BalanceControllerImpl) Upsert(ctx *fiber.Ctx) error {
	request := new(balance_model.BalanceRequest)

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

	err = controller.BalanceService.Upsert(ctx.UserContext(), transaction_model.Transaction{
		UserId:            userId,
		Currency:          request.Currency,
		Balance:           request.AddedBalance,
		ProofImageUrl:     request.TransferProofImg,
		BankAccountNumber: request.SenderBankAccountNumber,
		BankName:          request.SenderBankName,
	})
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON("Success")
}

func (controller *BalanceControllerImpl) GetAll(ctx *fiber.Ctx) error {
	userId, err := controller.AuthService.GetValidUser(ctx)
	if err != nil {
		return err
	}

	balances, err := controller.BalanceService.GetAll(ctx.UserContext(), userId)
	if err != nil {
		return err
	}

	var data []balance_model.BalanceGetData
	for _, balance := range balances {
		rowData := balance_model.BalanceGetData{
			Balance: balance.Balance,
			Currecy: balance.Currency,
		}

		data = append(data, rowData)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		balance_model.BalanceGetResponse{
			Message: "success",
			Data:    data,
		},
	)
}
