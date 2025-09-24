package routes

import (
	"student-crud/application/services"
	"student-crud/infrastructure/controllers"
	"student-crud/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Setup repository + service + controller
	repo := repository.NewStudentRepository()
	service := services.NewStudentService(repo)
	controller := controllers.NewStudentController(service)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/students", controller.CreateStudent)
	r.GET("/students", controller.GetStudents)
	r.PUT("/students/:id", controller.UpdateStudent)
	r.DELETE("/students/:id", controller.DeleteStudent)
}
