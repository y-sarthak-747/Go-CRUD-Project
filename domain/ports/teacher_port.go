package ports

import "student-crud/domain/models"

type TeacherRepository interface {
	Save(teacher *models.Teacher) error
	FindAll() ([]models.Teacher, error)
	FindByID(id uint) (*models.Teacher, error)
	Update(teacher *models.Teacher) error
	Delete(teacher *models.Teacher) error
}
