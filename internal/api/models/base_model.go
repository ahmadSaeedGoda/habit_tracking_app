package models

import (
	"gorm.io/gorm"
)

// BaseModel defines the common fields for all models.
type BaseModel struct {
	gorm.Model
}
