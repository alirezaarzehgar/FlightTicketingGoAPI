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
	case err != nil && r.RowsAffected == 0:
		return echo.ErrNotFound
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return echo.ErrConflict
	case err != nil:
		return echo.ErrInternalServerError
	}
	return nil
}

func Create[T any](model *T) *echo.HTTPError {
	return ErrGormToHttp(db.Create(&model))
}
