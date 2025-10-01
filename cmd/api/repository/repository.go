package repository

import "gorm.io/gorm"

type GormRepository[T any] struct {
	db *gorm.DB
}

func NewGormRepository[T any](db *gorm.DB) Repository[T] {
	return &GormRepository[T]{db: db}
}

// Create inserts the entity and returns it
func (r *GormRepository[T]) Create(entity *T) (*T, error) {
	if err := r.db.Create(entity).Error; err != nil {
		return nil, err
	}
	// entity is updated with ID, CreatedAt, etc.
	return entity, nil
}

// GetAll returns all entities
func (r *GormRepository[T]) GetAll() ([]T, error) {
	var entities []T
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// GetByID returns a single entity by ID
func (r *GormRepository[T]) GetByID(id uint) (*T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Update saves the entity and returns it
func (r *GormRepository[T]) Update(id uint, updated *T) (*T, error) {
	var entity T

	// fetch the existing record first
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}

	// update only the fields from updated
	if err := r.db.Model(&entity).Updates(updated).Error; err != nil {
		return nil, err
	}

	// return the updated entity
	return &entity, nil
}

// Delete deletes the entity by ID and returns the deleted entity
func (r *GormRepository[T]) Delete(id uint) (*T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Delete(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Exec executes a raw SQL statement (INSERT, UPDATE, DELETE, DDL)
// Returns the number of rows affected
func (r *GormRepository[T]) Exec(query string, args ...interface{}) (int64, error) {
	result := r.db.Exec(query, args...)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
