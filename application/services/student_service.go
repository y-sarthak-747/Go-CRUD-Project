package services

import (
	"errors"
	"student-crud/domain/models"
	"student-crud/domain/ports"
)

type StudentService struct {
	repo ports.StudentRepository
}

func NewStudentService(r ports.StudentRepository) *StudentService {
	return &StudentService{repo: r}
}

func (s *StudentService) Create(student *models.Student) error {
	return s.repo.Save(student)
}

func (s *StudentService) GetAll() ([]models.Student, error) {
	return s.repo.FindAll()
}

func (s *StudentService) Update(id uint, input *models.Student) (*models.Student, error) {
	student, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("student not found")
	}
	student.Name = input.Name
	student.Number = input.Number
	err = s.repo.Update(student)
	return student, err
}

func (s *StudentService) Delete(id uint) error {
	student, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("student not found")
	}
	return s.repo.Delete(student)
}
