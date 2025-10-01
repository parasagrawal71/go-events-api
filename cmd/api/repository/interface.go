package repository

type Repository[T any] interface {
	// Create inserts the entity and returns the inserted entity with auto-generated fields
	Create(entity *T) (*T, error)

	// GetAll fetches all entities
	GetAll() ([]T, error)

	// GetByID fetches a single entity by ID
	GetByID(id uint) (*T, error)

	// Update updates the entity with the given ID using the provided body
	// and returns the updated entity
	Update(id uint, entity *T) (*T, error)

	// Delete deletes the entity by ID and returns the deleted entity
	Delete(id uint) (*T, error)

	// Exec executes raw SQL statements and returns rows affected
	Exec(query string, args ...interface{}) (int64, error)
}
