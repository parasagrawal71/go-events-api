package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"` // auto-increment by default
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"-"`                                   // ignored in JSON output
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // custom column name, auto-set when inserting
	UpdatedAt time.Time `json:"updated_at"`                          // auto-set when inserting/updating
}

func (User) TableName() string {
	return "users" // custom table name
}
