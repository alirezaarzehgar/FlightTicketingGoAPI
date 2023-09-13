package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/BaseMax/FlightTicketingGoAPI/models"
	"github.com/BaseMax/FlightTicketingGoAPI/utils"
)

func SearchAirline(c echo.Context) error {
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

	var airline models.Airline
	r := db.First(&airline, conditions...)
	// GORM divers haven't ErrInvalidField ErrorTranslator feature.
	// I should hack there.
	if r.Error == nil && r.RowsAffected == 0 || errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return echo.ErrNotFound
	} else if r.Error != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, airline)

}

func FetchAllAirlines(c echo.Context) error {
	return FetchAllModels[models.Airline](c, "")
}

func airlineActivity(c echo.Context, isActive bool) error {
	id, _ := strconv.Atoi(c.Param("id"))
	r := db.Model(&models.Airline{}).Where(id).Update("active", isActive)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func ActiveAirline(c echo.Context) error {
	return airlineActivity(c, true)
}

func DeactiveAirline(c echo.Context) error {
	return airlineActivity(c, false)
}
