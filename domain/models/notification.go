package models

import "time"

type Notification struct {
    ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
    StudentID uint      `json:"student_id"`
    TeacherID uint      `json:"teacher_id"`
    Message   string    `json:"message"`
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (Notification) TableName() string {
    return "student_teacher_notifications"
}