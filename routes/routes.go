package routes

import (
	"mig-attendance/internal/delivery/api"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetRoutes(handler api.Handler) *echo.Echo {
	e := echo.New()
	useMiddlewares(e)

	e.POST("/login", handler.Login)
	e.POST("/logout", handler.Logout)

	g := e.Group("/api/v1")
	g.Use(checkAuth)
	g.GET("/testing", handler.GetActivityByDate)

	return e
}
func useMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch},
	}))
}
func checkAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("auth")
		if err != nil {
			return echo.ErrUnauthorized
		}

		value := cookie.Value
		if value != "randomstring" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
