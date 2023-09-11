package models

import (
	"gorm.io/gorm"
)

func Create[T any](model *T) *gorm.DB {
	return db.Create(model)
}
