package repository

import (
	"go-events-api/cmd/api/config"
	"go-events-api/cmd/api/models"
)

// Expose initialized repositories here
var (
	EventRepo Repository[models.Event]
	UserRepo  Repository[models.User]
)

// Init initializes all repositories once
func Init() {
	EventRepo = NewGormRepository[models.Event](config.DB)
	UserRepo = NewGormRepository[models.User](config.DB)
}
