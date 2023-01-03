package api

import (
	"log"
	m "mig-attendance/models"
	"mig-attendance/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handler) CheckInAttendance(c echo.Context) (err error) {
	var (
		user_id    int
		check_in   string
		attendance m.Attendance
	)

	user_id, err = strconv.Atoi(c.FormValue("user_id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, "id must be an integer and can't be empty")
		return c.JSON(http.StatusBadRequest, res)
	}

	check_in = utils.TimeNow
	attendance, err = h.repository.CheckInAttendance(c, check_in, user_id)
	if err != nil {
		log.Println("[Delivery][CheckInAttendance] can't check-in, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, "failed to check-in")
		return c.JSON(http.StatusInternalServerError, res)
	}

	var data []interface{}
	data = append(data, attendance)

	res := m.SetResponse(http.StatusOK, "check-in success", data)
	return c.JSON(http.StatusOK, res)
}
func (h *handler) CheckOutAttendance(c echo.Context) (err error) {
	var (
		user_id    int
		check_out  string
		attendance m.Attendance
	)

	user_id, err = strconv.Atoi(c.FormValue("user_id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, "id must be an integer and can't be empty")
		return c.JSON(http.StatusBadRequest, res)
	}

	check_out = utils.TimeNow
	attendance, err = h.repository.CheckInAttendance(c, check_out, user_id)
	if err != nil {
		log.Println("[Delivery][CheckInAttendance] can't check-out, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, "failed to check-out")
		return c.JSON(http.StatusInternalServerError, res)
	}

	var data []interface{}
	data = append(data, attendance)

	res := m.SetResponse(http.StatusOK, "check-out success", data)
	return c.JSON(http.StatusOK, res)
}
func (h *handler) GetHistoryAttendance(c echo.Context) (err error) {
	return
}
