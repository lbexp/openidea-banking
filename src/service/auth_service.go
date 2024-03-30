package service

import (
	"context"
	"openidea-banking/src/middleware"
	"openidea-banking/src/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var tokenExpDuration = time.Now().Add(time.Minute * 30).Unix()

type AuthService interface {
	GenerateToken(ctx context.Context, userId string) (string, error)
	GetValidUser(ctx *fiber.Ctx) (string, error)
}

type AuthServiceImpl struct {
}

func NewAuthService() AuthService {
	return &AuthServiceImpl{}
}

func (service *AuthServiceImpl) GenerateToken(ctx context.Context, userId string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     tokenExpDuration,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		return "", utils.ErrorInternalServer
	}

	return signedToken, nil
}

func (service *AuthServiceImpl) GetValidUser(ctx *fiber.Ctx) (string, error) {
	userInfo := ctx.Locals(middleware.JWT_CONTEXT_KEY).(*jwt.Token)
	claims := userInfo.Claims.((jwt.MapClaims))
	userId := claims["user_id"].(string)

	return userId, nil
}
