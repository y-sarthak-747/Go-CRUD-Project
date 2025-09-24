package services

import (
	"errors"
	"student-crud/domain/models"
	"student-crud/domain/ports"
)

type TeacherService struct {
	repo ports.TeacherRepository
}

func NewTeacherService(r ports.TeacherRepository) *TeacherService {
	return &TeacherService{repo: r}
}

func (s *TeacherService) Create(teacher *models.Teacher) error {
	return s.repo.Save(teacher)
}

func (s *TeacherService) GetAll() ([]models.Teacher, error) {
	return s.repo.FindAll()
}

func (s *TeacherService) Update(id uint, input *models.Teacher) (*models.Teacher, error) {
	teacher, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("teacher not found")
	}
	teacher.Name = input.Name
	teacher.Subject = input.Subject
	err = s.repo.Update(teacher)
	return teacher, err
}

func (s *TeacherService) Delete(id uint) error {
	teacher, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("teacher not found")
	}
	return s.repo.Delete(teacher)
}
