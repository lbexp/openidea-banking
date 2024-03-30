package middleware

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestProm = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_histogram",
		Help:    "Histogram of the http request duration.",
		Buckets: prometheus.LinearBuckets(1, 1, 10),
	}, []string{"path", "method", "status"})
)

func PrometheusMiddleware(ctx *fiber.Ctx) error {
	start := time.Now()
	err := ctx.Next()

	status := ctx.Response().StatusCode()
	elapsedTime := float64(time.Since(start).Milliseconds())
	httpRequestProm.WithLabelValues(ctx.Path(), ctx.Method(), http.StatusText(status)).Observe(elapsedTime)
	return err
}
