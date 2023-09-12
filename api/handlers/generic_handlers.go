package handlers

import (
	"net/http"
	"strconv"

	"github.com/BaseMax/FlightTicketingGoAPI/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Paginate(c echo.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func FetchModelById[T any](c echo.Context, idParam, omit string) error {
	var model T
	id, _ := strconv.Atoi(c.Param(idParam))
	r := db.Omit(omit).First(&model, id)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, model)
}

func FetchAllModels[T any](c echo.Context, omit string) error {
	var models []T
	r := db.Scopes(Paginate(c)).Omit(omit).Find(&models)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, models)
}

func DeleteById(c echo.Context, model any, idParam string) error {
	id, _ := strconv.Atoi(c.Param(idParam))
	r := db.Delete(model, id)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
