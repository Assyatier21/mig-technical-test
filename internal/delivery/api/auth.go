package api

import (
	"log"
	m "mig-attendance/models"
	"mig-attendance/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) Login(ctx echo.Context) (err error) {
	var (
		email    string
		password string
	)

	if !utils.IsValidEmail(ctx.FormValue("email")) {
		res := m.SetError(http.StatusBadRequest, "wrong email format")
		return ctx.JSON(http.StatusBadRequest, res)
	} else {
		email = ctx.FormValue("email")
	}

	password = ctx.FormValue("password")

	user, err := h.repository.Login(ctx, email, password)
	if err != nil {
		log.Println("[Delivery][Login] can't get login details, err:", err.Error())
		if err == utils.WrongUsernamePass {
			res := m.SetError(http.StatusUnauthorized, err.Error())
			return ctx.JSON(http.StatusUnauthorized, res)
		} else {
			res := m.SetError(http.StatusInternalServerError, "failed to get login details")
			return ctx.JSON(http.StatusInternalServerError, res)
		}
	}
	var data []interface{}
	data = append(data, user)

	res := m.SetResponse(http.StatusOK, "logged in", data)
	return ctx.JSON(http.StatusBadGateway, res)
}
func (h *handler) Logout(ctx echo.Context) (err error) {
	err = h.repository.Logout(ctx)
	if err != nil {
		log.Println("[Delivery][Logout] can't logout, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "logged out"})
}
