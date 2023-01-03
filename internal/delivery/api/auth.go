package api

import (
	"log"
	m "mig-attendance/models"
	"mig-attendance/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) Login(c echo.Context) (err error) {
	var (
		email    string
		password string
	)

	if !utils.IsValidEmail(c.FormValue("email")) {
		res := m.SetError(http.StatusBadRequest, "wrong email format")
		return c.JSON(http.StatusBadRequest, res)
	} else {
		email = c.FormValue("email")
	}

	password = c.FormValue("password")

	user, err := h.repository.Login(c, email, password)
	if err != nil {
		log.Println("[Delivery][Login] can't get login details, err:", err.Error())
		if err == utils.WrongUsernamePass {
			res := m.SetError(http.StatusUnauthorized, err.Error())
			return c.JSON(http.StatusUnauthorized, res)
		} else {
			res := m.SetError(http.StatusInternalServerError, "failed to get login details")
			return c.JSON(http.StatusInternalServerError, res)
		}
	}
	var data []interface{}
	data = append(data, user)

	res := m.SetResponse(http.StatusOK, "logged in", data)
	return c.JSON(http.StatusOK, res)
}
func (h *handler) Logout(c echo.Context) (err error) {
	err = h.repository.Logout(c)
	if err != nil {
		log.Println("[Delivery][Logout] can't logout, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, err.Error())
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "logged out"})
}
