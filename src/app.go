package app

import (
	"log"
	"openidea-banking/src/config"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func Start(port string, prefork bool) {
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		IdleTimeout:  config.IdleTimeout,
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
		Prefork:      prefork,
	})

	dbPool := InitDbPool()
	defer dbPool.Close()

	RegisterRoutes(app, dbPool)

	err := app.Listen(":" + port)
	log.Fatal("Failed to start app - ", err)
}
