package ports

import "student-crud/domain/models"

type StudentRepository interface {
	Save(student *models.Student) error
	FindAll() ([]models.Student, error)
	FindByID(id uint) (*models.Student, error)
	Update(student *models.Student) error
	Delete(student *models.Student) error
}

