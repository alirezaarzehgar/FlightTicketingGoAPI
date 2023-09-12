package routes

import (
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/BaseMax/FlightTicketingGoAPI/api/handlers"
	"github.com/BaseMax/FlightTicketingGoAPI/api/middlewares"
	"github.com/BaseMax/FlightTicketingGoAPI/config"
)

func todo(c echo.Context) error {
	return c.String(http.StatusOK, "version:"+c.Param("version"))
}

func groupedByVersion(g *echo.Group) {
	g.POST("/register", handlers.Register)
	g.POST("/login", handlers.Login)

	g = g.Group("", echojwt.WithConfig(echojwt.Config{SigningKey: config.GetJwtSecret()}))
	g.GET("/users/:id", handlers.FetchUser)
	g.GET("/users", handlers.FetchUsers)
	g.PUT("/users/:id", handlers.EditUser)
	g.DELETE("/users/:id", handlers.DeleteUser, middlewares.AdminOnly)

	g.POST("/flights", todo, middlewares.AdminOnly)
	g.GET("/flights", todo, middlewares.AdminOnly)
	g.GET("/flights/:id", todo, middlewares.AdminOnly)
	g.PUT("/flights/:id", todo, middlewares.AdminOnly)
	g.DELETE("/flights/:id", todo, middlewares.AdminOnly)

	g.POST("/flights/:flight_id/booking", todo)
	g.GET("/tickets", todo)
	g.GET("/tickets/:id", todo)
	g.GET("/tickets/:flight_id", todo)
	g.PUT("/tickets/:id", todo)
	g.DELETE("/tickets/:id/cancel", todo, middlewares.AdminOnly)

	// We haven't plan for that
	g.POST("/payments", todo)
}

func InitRoutes() *echo.Echo {
	e := echo.New()

	middlewareList := []echo.MiddlewareFunc{
		middlewares.ValidVersion,
		middlewares.ValidIdParams,
	}
	groupedByVersion(e.Group("/v:version", middlewareList...))

	return e
}
