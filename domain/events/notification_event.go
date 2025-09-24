package events

// NotificationEvent is the payload produced/consumed via Kafka.
// Placing it in domain/events makes it available to consumer and service.
type NotificationEvent struct {
    StudentID uint   `json:"student_id"`
    TeacherID uint   `json:"teacher_id"`
    Message   string `json:"message"`
}
