package main

import (
	"student-crud/config"
	"student-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
