package models

import (
	"gorm.io/gorm"
)

func Create(model any) *gorm.DB {
	return db.Create(model)
}
