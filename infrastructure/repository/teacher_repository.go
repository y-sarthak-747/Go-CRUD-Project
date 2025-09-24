package repository

import (
	"student-crud/config"
	"student-crud/domain/models"
)

type PostgresTeacherRepository struct{}

func NewPostgresTeacherRepository() *PostgresTeacherRepository {
	return &PostgresTeacherRepository{}
}

func (r *PostgresTeacherRepository) Save(teacher *models.Teacher) error {
	return config.DB.Create(teacher).Error
}

func (r *PostgresTeacherRepository) FindAll() ([]models.Teacher, error) {
	var teachers []models.Teacher
	err := config.DB.Find(&teachers).Error
	return teachers, err
}

func (r *PostgresTeacherRepository) FindByID(id uint) (*models.Teacher, error) {
	var teacher models.Teacher
	err := config.DB.First(&teacher, id).Error
	return &teacher, err
}

func (r *PostgresTeacherRepository) Update(teacher *models.Teacher) error {
	return config.DB.Save(teacher).Error
}

func (r *PostgresTeacherRepository) Delete(teacher *models.Teacher) error {
	return config.DB.Delete(teacher).Error
}
