package src

import (
	"encoding/json"
	"log"

	"openidea-banking/src/config"
	"openidea-banking/src/middleware"
	"openidea-banking/src/utils"

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

	db := GetConnectionDB()
	defer db.Close()

	app.Use(middleware.PrometheusMiddleware)

	RegisterRoute(app, dbPool)
	app.Use(utils.HandleErrorNotFound)

	err := app.Listen(":" + port)
	log.Fatal(err)
}
