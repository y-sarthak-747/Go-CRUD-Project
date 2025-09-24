package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"student-crud/config"
	"student-crud/domain/models"
	"student-crud/infrastructure/metrics"
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
			metrics.StudentCacheHits.WithLabelValues("findById").Inc()
			return &student, nil
		}
	}

	// Cache miss
	metrics.StudentCacheMisses.WithLabelValues("findById").Inc()
	return nil, err
}


func (r *StudentRedisRepository) Save(student *models.Student) error {
	key := fmt.Sprintf("student:%d", student.ID)

	data, err := json.Marshal(student)
	if err != nil {
		return err
	}

	err = config.Redis.Set(context.Background(), key, data, 10*time.Minute).Err()
	if err == nil {
		metrics.StudentCacheHits.WithLabelValues("save").Inc() // success means cache updated
	} else {
		metrics.StudentCacheMisses.WithLabelValues("save").Inc()
	}
	return err
}

