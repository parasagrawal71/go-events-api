package dto

type Event struct {
	ID          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required,min=10"`
	Date        string `json:"date" binding:"required"`
	Location    string `json:"location" binding:"required"`
}
