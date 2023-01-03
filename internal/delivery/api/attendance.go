package api

import (
	"database/sql"
	"log"
	m "mig-attendance/models"
	"mig-attendance/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *handler) CheckInAttendance(c echo.Context) (err error) {
	var (
		user_id    int
		check_in   string
		attendance m.ResAttendance
	)

	user_id, err = strconv.Atoi(c.FormValue("user_id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, "id must be an integer and can't be empty")
		return c.JSON(http.StatusBadRequest, res)
	}

	check_in = time.Now().In(utils.JakartaLoc).Format("2006-01-02T15:04:05")
	attendance, err = h.repository.CheckInAttendance(c, check_in, user_id)
	if err != nil {
		log.Println("[Delivery][CheckInAttendance] can't check-in, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, "failed to check-in")
		return c.JSON(http.StatusInternalServerError, res)
	}

	attendance.CheckIn = utils.FormattedTime(attendance.CheckIn)

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

	check_out = time.Now().In(utils.JakartaLoc).Format("2006-01-02T15:04:05")
	attendance, err = h.repository.CheckOutAttendance(c, check_out, user_id)
	if err != nil {
		log.Println("[Delivery][CheckOutAttendance] can't check-out, err:", err.Error())
		if err == sql.ErrNoRows {
			res := m.SetError(http.StatusNoContent, "already checked-out")
			return c.JSON(http.StatusOK, res)
		}
		res := m.SetError(http.StatusInternalServerError, "failed to check-out")
		return c.JSON(http.StatusInternalServerError, res)
	}

	attendance.CheckIn = utils.FormattedTime(attendance.CheckIn)
	attendance.CheckOut = utils.FormattedTime(attendance.CheckOut)

	var data []interface{}
	data = append(data, attendance)

	res := m.SetResponse(http.StatusOK, "check-out success", data)
	return c.JSON(http.StatusOK, res)
}
func (h *handler) GetHistoryAttendance(c echo.Context) (err error) {
	return
}
