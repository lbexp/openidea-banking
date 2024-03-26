package src

import (
	"log"
	"openidea-banking/configs"

	"github.com/gofiber/fiber/v2"
)

func StartApplication(port string, prefork bool) {
	app := fiber.New(fiber.Config{
		IdleTimeout:  configs.IdleTimeout,
		WriteTimeout: configs.WriteTimeout,
		ReadTimeout:  configs.ReadTimeout,
		Prefork:      prefork,
	})

	err := app.Listen(":" + port)
	log.Fatal(err)
}
