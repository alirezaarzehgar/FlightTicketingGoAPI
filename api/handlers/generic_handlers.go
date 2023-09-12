package handlers

import (
	"net/http"
	"strconv"

	"github.com/BaseMax/FlightTicketingGoAPI/utils"
	"github.com/labstack/echo/v4"
)

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
	r := db.Omit(omit).Find(&models)
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
