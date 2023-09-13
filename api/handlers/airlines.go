package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/BaseMax/FlightTicketingGoAPI/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

func ActiveAirline(c echo.Context) error {
	return nil
}

func DeactiveAirline(c echo.Context) error {
	return nil
}
