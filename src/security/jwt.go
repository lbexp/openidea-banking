package security

import (
	"openidea-banking/src/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var tokenExpDuration = time.Now().Add(time.Minute * 2).Unix()

func GenerateToken(userId, name string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"name":    name,
		"exp":     tokenExpDuration,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return signedToken, nil
}

func GetValidUser(ctx *fiber.Ctx) UserPayload {
	userInfo := ctx.Locals(middleware.JWT_CONTEXT_KEY).(*jwt.Token)
	claims := userInfo.Claims.((jwt.MapClaims))
	userId := claims["user_id"].(string)
	name := claims["name"].(string)

	return UserPayload{
		UserId: userId,
		Name:   name,
	}
}

type UserPayload struct {
	UserId string
	Name   string
}
