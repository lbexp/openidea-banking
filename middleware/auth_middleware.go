package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	auth := ctx.Get("Authorization")
	if auth == "" {
		return fiber.NewError(401, "token not found")
	}

	splitted := strings.Split(auth, " ")

	if splitted[0] != "Bearer" {
		return fiber.NewError(401, "token not found")
	}

	//TODO : verify token here

	return ctx.Next()
}
