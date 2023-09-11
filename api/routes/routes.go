package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/FlightTicketingGoAPI/api/middlewares"
)

func todo(c echo.Context) error {
	return c.String(http.StatusOK, "version:"+c.Param("version"))
}

func groupedByVersion(g *echo.Group) {
	g.POST("/register", todo)
	g.POST("/login", todo)
	g.GET("/users/:id", todo)
	g.GET("/users", todo)
	g.PUT("/users/:id", todo)
	g.DELETE("/users/:id", todo)

	g.POST("/flights", todo)
	g.GET("/flights", todo)
	g.GET("/flights/:id", todo)
	g.PUT("/flights/:id", todo)
	g.DELETE("/flights/:id", todo)

	g.POST("/flights/:flight_id/booking", todo)
	g.GET("/tickets", todo)
	g.GET("/tickets/:id", todo)
	g.GET("/tickets/:flight_id", todo)
	g.PUT("/tickets/:id", todo)
	g.DELETE("/tickets/:id", todo)

	// We haven't plan for that
	g.POST("/payments", todo)
}

func InitRoutes() *echo.Echo {
	e := echo.New()

	groupedByVersion(e.Group("/v:version", middlewares.ValidVersion))

	return e
}
