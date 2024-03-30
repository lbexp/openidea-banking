package src

import (
	"encoding/json"
	"log"

	"openidea-banking/src/config"
	"openidea-banking/src/middleware"
	"openidea-banking/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	app.Use(logger.New())

	db := GetConnectionDB()
	defer db.Close()

	app.Use(middleware.PrometheusMiddleware)

	RegisterRoute(app, dbPool)
	app.Use(utils.HandleErrorNotFound)

	err := app.Listen(":" + port)
	log.Fatal(err)
}
