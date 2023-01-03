package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetActivity(c echo.Context) (err error) {
	return
}
func (h *handler) GetActivityByDate(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, map[string]string{"message": "Success Accessed!"})
}
func (h *handler) AddEditActivity(c echo.Context) (err error) {
	return
}
func (h *handler) DeleteActivity(c echo.Context) (err error) {
	return
}
