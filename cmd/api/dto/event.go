package dto

import "time"

type Event struct {
	ID          uint      `json:"id" `
	Name        string    `json:"name"`
	Description string    `json:"description" binding:"min=10"`
	Date        string    `json:"date"`
	Location    string    `json:"location"`
	OwnerID     uint      `json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
