package routes

import (
	"mig-attendance/internal/delivery/api"
	m "mig-attendance/models"
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
	g.POST("/check-in", handler.CheckInAttendance)
	g.PATCH("/check-out", handler.CheckOutAttendance)

	g.GET("/activity", handler.GetActivityByDate)
	g.PATCH("/activity", handler.AddEditActivity)
	g.PATCH("/activity/delete", handler.DeleteActivity)

	g.GET("/attendance/history", handler.GetHistoryAttendance)

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
			res := m.SetError(http.StatusUnauthorized, "Unauthorized")
			return c.JSON(http.StatusUnauthorized, res)
		}

		value := cookie.Value
		if value != "randomstring" {
			res := m.SetError(http.StatusUnauthorized, "Unauthorized")
			return c.JSON(http.StatusUnauthorized, res)
		}
		return next(c)
	}
}
