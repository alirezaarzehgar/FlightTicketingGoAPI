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
	g.POST("/register", handlers.RegisterPassenger)
	g.POST("/login", handlers.Login)

	g = g.Group("", echojwt.WithConfig(echojwt.Config{SigningKey: config.GetJwtSecret()}))
	g.POST("/register/employee", handlers.RegisterEmployee, middlewares.AdminOnly)
	g.GET("/users/:id", handlers.FetchUser)
	g.GET("/users", handlers.FetchUsers)
	g.PUT("/users/:id", handlers.EditUser)
	g.DELETE("/users/:id", handlers.DeleteUser, middlewares.AdminOnly)

	g.GET("/airlines/search", handlers.SearchAirline, middlewares.EmployeePrivilege)
	g.GET("/airlines", handlers.FetchAllAirlines, middlewares.EmployeePrivilege)
	g.POST("/airlines/:id/active", handlers.ActiveAirline, middlewares.EmployeePrivilege)
	g.POST("/airlines/:id/deactive", handlers.DeactiveAirline, middlewares.EmployeePrivilege)

	g.POST("/flights", handlers.NewFlight, middlewares.EmployeePrivilege)
	g.GET("/flights/search", handlers.SearchFlight, middlewares.EmployeePrivilege)
	g.GET("/flights", handlers.FetchAllFlights, middlewares.EmployeePrivilege)
	g.PUT("/flights/:id", handlers.EditFlight, middlewares.EmployeePrivilege)
	g.DELETE("/flights/:id", handlers.DeleteFlight, middlewares.EmployeePrivilege)

	g.POST("/flights/:flight_id/booking", todo)
	g.GET("/tickets", todo)
	g.GET("/tickets/:id", todo)
	g.GET("/tickets/:flight_id", todo)
	g.PUT("/tickets/:id", todo, middlewares.EmployeePrivilege)
	g.DELETE("/tickets/:id/cancel", todo, middlewares.EmployeePrivilege)

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
