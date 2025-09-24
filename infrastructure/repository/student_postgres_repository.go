package repository

import (
	"student-crud/config"
	"student-crud/domain/models"
)

type PostgresStudentRepository struct{}

func NewPostgresStudentRepository() *PostgresStudentRepository {
	return &PostgresStudentRepository{}
}

func (r *PostgresStudentRepository) Save(student *models.Student) error {
	return config.DB.Create(student).Error
}

func (r *PostgresStudentRepository) FindAll() ([]models.Student, error) {
	var students []models.Student
	err := config.DB.Find(&students).Error
	return students, err
}

func (r *PostgresStudentRepository) FindByID(id uint) (*models.Student, error) {
	var student models.Student
	err := config.DB.First(&student, id).Error
	return &student, err
}

func (r *PostgresStudentRepository) Update(student *models.Student) error {
	return config.DB.Save(student).Error
}

func (r *PostgresStudentRepository) Delete(student *models.Student) error {
	return config.DB.Delete(student).Error
}
