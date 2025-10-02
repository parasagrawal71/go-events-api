package models

import "time"

type Event struct {
	ID          uint      `gorm:"primaryKey" json:"id"` // auto-increment by default
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required,min=10"`
	Date        string    `json:"date" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	OwnerID     uint      `json:"owner_id"`                                      // FK field
	Owner       User      `gorm:"foreignKey:OwnerID;references:ID" json:"owner"` // relation
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Event) TableName() string {
	return "events" // custom table name
}
