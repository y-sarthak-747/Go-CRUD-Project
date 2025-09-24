package ports

import "student-crud/domain/models"

type NotificationRepository interface {
    Save(notification *models.Notification) error
    // Add other methods if you need reads, listing, etc.
}
