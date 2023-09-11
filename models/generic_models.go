package models

import (
	"errors"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ErrGormToHttp(r *gorm.DB) *echo.HTTPError {
	err := r.Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
	case errors.Is(err, gorm.ErrForeignKeyViolated):
		return echo.ErrNotFound
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return echo.ErrConflict
	}
	if err != nil {
		return echo.ErrInternalServerError
	}
	if r.RowsAffected == 0 {
		return echo.ErrNotFound
	}
	return nil
}

func Create[T any](model *T) error {
	return ErrGormToHttp(db.Create(&model))
}
