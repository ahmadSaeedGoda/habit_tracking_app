package repositories

import (
	"gorm.io/gorm"
)

// GenericRepository interface defines common methods for all repositories
type GenericRepository interface {
	Create(model interface{}) error
	Update(id uint, model interface{}) error
	FindById(id uint, model interface{}) error
	List(models interface{}) error
}

type BaseRepository struct {
	db *gorm.DB
}
