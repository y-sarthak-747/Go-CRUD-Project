package main

import (
	"student-crud/bootstrap"
	"student-crud/config"
	"student-crud/routes"

	"student-crud/infrastructure/metrics"
	"student-crud/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.Use(middleware.PrometheusMiddleware())

    // 1. Infrastructure
    config.ConnectDatabase()
    config.ConnectRedis()
    config.InitKafkaProducer()

    // 2. Routes (HTTP)
    routes.RegisterRoutes(r)

    // 3. Metrics
    metrics.RegisterMetricsEndpoint()

    // 4. Background consumers
    bootstrap.InitKafkaConsumers()

    // 5. Start API
    r.Run(":8080")
}
