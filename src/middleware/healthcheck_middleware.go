package middleware

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthChecker struct {
	DB  *pgxpool.Pool
	App *fiber.App
}

func NewHealthChecker(
	db *pgxpool.Pool,
	app *fiber.App,
) HealthChecker {
	return HealthChecker{
		DB:  db,
		App: app,
	}
}

func (hc *HealthChecker) HealthCheckMiddleware(ctx *fiber.Ctx) error {
	if ctx.Path() != "/healthz" {
		return ctx.Next()
	}

	_, err := http.Get(ctx.BaseURL() + "/")
	if err != nil {
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(map[string]string{
			"service": err.Error(),
		})
	}

	err = hc.DB.Ping(context.Background())
	if err != nil {
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(map[string]string{
			"service": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]string{
		"service": "oke",
	})
}
