package dto

type UpdateEvent struct {
	Name        string `json:"name"`
	Description string `json:"description" binding:"min=10"`
	Date        string `json:"date"`
	Location    string `json:"location"`
}
