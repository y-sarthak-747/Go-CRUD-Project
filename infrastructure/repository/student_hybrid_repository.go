package repository

import (
	"student-crud/domain/models"
)

type StudentHybridRepository struct {
	dbRepo    *PostgresStudentRepository
	redisRepo *StudentRedisRepository
}

func NewStudentHybridRepository(dbRepo *PostgresStudentRepository, redisRepo *StudentRedisRepository) *StudentHybridRepository {
	return &StudentHybridRepository{dbRepo: dbRepo, redisRepo: redisRepo}
}

func (r *StudentHybridRepository) FindByID(id uint) (*models.Student, error) {
	// 1. Try cache
	student, err := r.redisRepo.FindByID(id)
	if err == nil && student != nil {
		return student, nil
	}

	// 2. Fallback to DB
	student, err = r.dbRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// 3. Update cache
	_ = r.redisRepo.Save(student)

	return student, nil
}

func (r *StudentHybridRepository) Save(student *models.Student) error {
	// Save in DB
	err := r.dbRepo.Save(student)
	if err != nil {
		return err
	}

	// Update cache
	return r.redisRepo.Save(student)
}

func (r *StudentHybridRepository) FindAll() ([]models.Student, error) {
	return r.dbRepo.FindAll()
}

func (r *StudentHybridRepository) Update(student *models.Student) error {
	err := r.dbRepo.Update(student)
	if err != nil {
		return err
	}
	return r.redisRepo.Save(student)
}

func (r *StudentHybridRepository) Delete(student *models.Student) error {
	return r.dbRepo.Delete(student)
}
