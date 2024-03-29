package controller

import (
	common_model "openidea-banking/src/model/common"
	transaction_model "openidea-banking/src/model/transaction"
	"openidea-banking/src/service"
	"openidea-banking/src/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TransactionController interface {
	Create(ctx *fiber.Ctx) error
	GetAllByUserId(ctx *fiber.Ctx) error
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

func validateQueryParams(req map[string]string) (transaction_model.GetAllByUserIdFilters, error) {
	var filters transaction_model.GetAllByUserIdFilters
	limitVal, isLimitExists := req["limit"]

	if !isLimitExists && limitVal == "" {
		filters.Limit = 5
	} else {
		resultLimit, err := strconv.Atoi(limitVal)
		if err != nil {
			return transaction_model.GetAllByUserIdFilters{}, utils.ErrorBadRequest
		}

		if resultLimit < 0 {
			return transaction_model.GetAllByUserIdFilters{}, utils.ErrorBadRequest
		}

		filters.Limit = resultLimit
	}

	if isLimitExists && limitVal == "" {
		return transaction_model.GetAllByUserIdFilters{}, utils.ErrorBadRequest
	}

	offsetVal, isOffsetExists := req["offset"]
	if !isOffsetExists && offsetVal == "" {
		filters.Offset = 0
	} else {
		resultOffset, err := strconv.Atoi(offsetVal)
		if err != nil {
			return transaction_model.GetAllByUserIdFilters{}, utils.ErrorBadRequest
		}

		filters.Offset = resultOffset
	}

	if isOffsetExists && offsetVal == "" {
		return transaction_model.GetAllByUserIdFilters{}, utils.ErrorBadRequest
	}

	return filters, nil
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

func (controller *TransactionControllerImpl) GetAllByUserId(ctx *fiber.Ctx) error {
	filters, err := validateQueryParams(ctx.Queries())
	if err != nil {
		return err
	}

	userId, err := controller.AuthService.GetValidUser(ctx)
	if err != nil {
		return err
	}

	transactions, totalItems, err := controller.TransactionService.GetAllByUserId(ctx.UserContext(), userId, filters)
	if err != nil {
		return err
	}

	var data []transaction_model.TransactionGetAllData
	for _, transaction := range transactions {
		rowData := transaction_model.TransactionGetAllData{
			TransactionId:    transaction.TransactionId,
			Balance:          transaction.Balance,
			Currency:         transaction.Currency,
			TransferProofImg: transaction.ProofImageUrl,
			CreatedAt:        transaction.CreatedAt,
			Source: transaction_model.TransactionSource{
				BankAccountNumber: transaction.BankAccountNumber,
				BankName:          transaction.BankName,
			},
		}

		data = append(data, rowData)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		transaction_model.TransactionGetAllResponse{
			Message: "success",
			Data:    data,
			Meta: common_model.MetaResponse{
				Limit:  filters.Limit,
				Offset: filters.Offset,
				Total:  totalItems,
			},
		},
	)
}
