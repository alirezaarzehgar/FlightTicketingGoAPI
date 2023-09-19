package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BaseMax/FlightTicketingGoAPI/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	r := db.Omit(omit).Preload(clause.Associations).First(&model, id)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, model)
}

func FetchAllModels[T any](c echo.Context, omit string) error {
	var models []T
	r := db.Scopes(Paginate(c)).Omit(omit).Preload(clause.Associations).Find(&models)
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

func SearchModel[T any](c echo.Context) error {
	if len(c.QueryParams()) == 0 {
		return echo.ErrBadRequest
	}

	conditionStr := ""
	conditions := []any{""}
	for qp := range c.QueryParams() {
		conditionStr += fmt.Sprintf("%s = ? AND", qp)
		conditions = append(conditions, c.QueryParam(qp))
	}
	conditionStr = conditionStr[:len(conditionStr)-3]
	conditions[0] = conditionStr

	var model T
	r := db.Preload(clause.Associations).First(&model, conditions...)
	// GORM divers haven't ErrInvalidField ErrorTranslator feature.
	// I should hack there.
	if r.Error == nil && r.RowsAffected == 0 || errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return echo.ErrNotFound
	} else if r.Error != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, model)
}
