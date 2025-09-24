package main

import (
	"student-crud/config"
	"student-crud/routes"

	"student-crud/infrastructure/metrics"
	"student-crud/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.PrometheusMiddleware())

	config.ConnectDatabase()

	routes.RegisterRoutes(r)

	metrics.RegisterMetricsEndpoint()

	r.Run(":8080")
}
