package controllers

import (
	"net/http"
	"strconv"
	"student-crud/application/services"
	"student-crud/domain/models"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	service *services.StudentService
}

func NewStudentController(s *services.StudentService) *StudentController {
	return &StudentController{service: s}
}

func (ctrl *StudentController) CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.service.Create(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save student"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (ctrl *StudentController) GetStudents(c *gin.Context) {
	students, err := ctrl.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch students"})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (ctrl *StudentController) UpdateStudent(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input models.Student
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student, err := ctrl.service.Update(uint(id), &input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (ctrl *StudentController) DeleteStudent(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	if err := ctrl.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
