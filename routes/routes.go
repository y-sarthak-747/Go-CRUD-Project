package routes

import (
	"student-crud/application/services"
	"student-crud/infrastructure/controllers"
	"student-crud/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    // === Student Setup ===
    dbRepo := repository.NewPostgresStudentRepository()
	redisRepo := repository.NewStudentRedisRepository()
	hybridRepo := repository.NewStudentHybridRepository(dbRepo, redisRepo)

	studentService := services.NewStudentService(hybridRepo)
	studentController := controllers.NewStudentController(studentService)

    studentRoutes := r.Group("/students")
    {
        studentRoutes.POST("", studentController.CreateStudent)
        studentRoutes.GET("", studentController.GetStudents)
        studentRoutes.PUT("/:id", studentController.UpdateStudent)
        studentRoutes.DELETE("/:id", studentController.DeleteStudent)
    }

    // === Teacher Setup ===
    teacherRepo := repository.NewPostgresTeacherRepository()
    teacherService := services.NewTeacherService(teacherRepo)
    teacherController := controllers.NewTeacherController(teacherService)

    teacherRoutes := r.Group("/teachers")
    {
        teacherRoutes.POST("", teacherController.CreateTeacher)
        teacherRoutes.GET("", teacherController.GetTeachers)
        teacherRoutes.PUT("/:id", teacherController.UpdateTeacher)
        teacherRoutes.DELETE("/:id", teacherController.DeleteTeacher)
    }
}

/*
Student APIs

POST /students
GET /students
PUT /students/:id
DELETE /students/:id



Teacher APIs

POST /teachers
GET /teachers
PUT /teachers/:id
DELETE /teachers/:id
*/
