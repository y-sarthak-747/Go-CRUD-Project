package models

type Teacher struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name   string `json:"name"`
	Subject string `json:"subject"`
}
