package controller

import (
	user_model "openidea-banking/src/model/user"
	"openidea-banking/src/service"
	"openidea-banking/src/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}

type UserControllerImpl struct {
	Validator   *validator.Validate
	UserService service.UserService
}

func NewUserController(validator *validator.Validate, userService service.UserService) UserController {
	return &UserControllerImpl{
		Validator:   validator,
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(ctx *fiber.Ctx) error {
	request := new(user_model.UserRegisterRequest)

	err := ctx.BodyParser(request)
	if err != nil {
		return utils.ErrorBadRequest
	}

	err = controller.Validator.Struct(request)
	if err != nil {
		return utils.ErrorBadRequest
	}

	user, err := controller.UserService.Register(ctx.UserContext(), user_model.User{
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
	})
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(user_model.UserResponse{
		Message: "User registered successfully",
		Data: user_model.UserData{
			Email:       user.Email,
			Name:        user.Name,
			AccessToken: user.AccessToken,
		},
	})
}

func (controller *UserControllerImpl) Login(ctx *fiber.Ctx) error {
	request := new(user_model.UserLoginRequest)

	err := ctx.BodyParser(request)
	if err != nil {
		return utils.ErrorBadRequest
	}

	err = controller.Validator.Struct(request)
	if err != nil {
		return utils.ErrorBadRequest
	}

	user, err := controller.UserService.Login(ctx.UserContext(), user_model.User{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(user_model.UserResponse{
		Message: "User logged successfully",
		Data: user_model.UserData{
			Email:       user.Email,
			Name:        user.Name,
			AccessToken: user.AccessToken,
		},
	})
}
