package src

import (
	"encoding/json"
	"log"

	"openidea-banking/src/config"
	"openidea-banking/src/middleware"
	"openidea-banking/src/utils"

	"github.com/gofiber/fiber/v2"
<<<<<<< HEAD
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
=======
	"github.com/gofiber/fiber/v2/middleware/logger"
>>>>>>> origin/main
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

	healthChecker := middleware.HealthChecker{
		DB:  db,
		App: app,
	}

	app.Use(healthChecker.HealthCheckMiddleware)
	app.Use(middleware.PrometheusMiddleware)
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	RegisterRoute(app, dbPool)
	app.Use(utils.HandleErrorNotFound)

	err := app.Listen(":" + port)
	log.Fatal(err)
}
