package src

import (
	"encoding/json"
	"log"
	"openidea-banking/src/config"
	"openidea-banking/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func StartApplication(port string, prefork bool) {
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		IdleTimeout:  config.IdleTimeout,
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
		Prefork:      prefork,
	})

	app.Use(middleware.PrometheusMiddleware)

	err := app.Listen(":" + port)
	log.Fatal(err)
}
