package repository

import (
	"student-crud/config"
	"student-crud/domain/models"
)

type StudentRepository struct{}

func NewPostgresStudentRepository() *StudentRepository {
	return &StudentRepository{}
}

func (r *StudentRepository) Save(student *models.Student) error {
	return config.DB.Create(student).Error
}

func (r *StudentRepository) FindAll() ([]models.Student, error) {
	var students []models.Student
	err := config.DB.Find(&students).Error
	return students, err
}

func (r *StudentRepository) FindByID(id uint) (*models.Student, error) {
	var student models.Student
	err := config.DB.First(&student, id).Error
	return &student, err
}

func (r *StudentRepository) Update(student *models.Student) error {
	return config.DB.Save(student).Error
}

func (r *StudentRepository) Delete(student *models.Student) error {
	return config.DB.Delete(student).Error
}
