package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetActivity(ctx echo.Context) (err error) {
	return
}
func (h *handler) GetActivityByDate(ctx echo.Context) (err error) {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Success Accessed!"})
}
func (h *handler) AddEditActivity(ctx echo.Context) (err error) {
	return
}
func (h *handler) DeleteActivity(ctx echo.Context) (err error) {
	return
}
