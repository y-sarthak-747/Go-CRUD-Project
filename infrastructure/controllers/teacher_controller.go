package controllers

import (
	"net/http"
	"strconv"
	"student-crud/application/services"
	"student-crud/domain/models"

	"github.com/gin-gonic/gin"
)

type TeacherController struct {
	service *services.TeacherService
}

func NewTeacherController(s *services.TeacherService) *TeacherController {
	return &TeacherController{service: s}
}

func (ctrl *TeacherController) CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.service.Create(&teacher); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save teacher"})
		return
	}
	c.JSON(http.StatusOK, teacher)
}

func (ctrl *TeacherController) GetTeachers(c *gin.Context) {
	teachers, err := ctrl.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch teachers"})
		return
	}
	c.JSON(http.StatusOK, teachers)
}

func (ctrl *TeacherController) UpdateTeacher(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input models.Teacher
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacher, err := ctrl.service.Update(uint(id), &input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "teacher not found"})
		return
	}
	c.JSON(http.StatusOK, teacher)
}

func (ctrl *TeacherController) DeleteTeacher(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	if err := ctrl.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "teacher not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
