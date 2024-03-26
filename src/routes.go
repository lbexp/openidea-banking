<<<<<<< HEAD
package src

func RegisterRoute() {

=======
package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(app *fiber.App, dbPool *pgxpool.Pool) {
	// TODO: Pass this to controller
	validator := validator.New()
>>>>>>> origin/main
}
