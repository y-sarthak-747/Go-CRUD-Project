package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"student-crud/config"
	"student-crud/domain/models"
)

type StudentRedisRepository struct{}

func NewStudentRedisRepository() *StudentRedisRepository {
	return &StudentRedisRepository{}
}

func (r *StudentRedisRepository) FindByID(id uint) (*models.Student, error) {
	key := fmt.Sprintf("student:%d", id)

	// Try to get from Redis
	val, err := config.Redis.Get(config.Ctx, key).Result()
	if err == nil {
		// Cache hit
		var student models.Student
		if err := json.Unmarshal([]byte(val), &student); err == nil {
			return &student, nil
		}
	}

	// Cache miss (not in Redis)
	return nil, err
}

func (r *StudentRedisRepository) Save(student *models.Student) error {
	key := fmt.Sprintf("student:%d", student.ID)

	data, err := json.Marshal(student)
	if err != nil {
		return err
	}

	// Set with TTL of 10 minutes
	return config.Redis.Set(context.Background(), key, data, 10*time.Minute).Err()
}
