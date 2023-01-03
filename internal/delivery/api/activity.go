package api

import (
	"log"
	m "mig-attendance/models"
	"mig-attendance/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetActivityByDate(c echo.Context) (err error) {
	var (
		date_start string
		date_end   string
		user_id    int
	)

	if !utils.IsValidDate(c.FormValue("date_start")) {
		res := m.SetError(http.StatusBadRequest, "wrong date_start format")
		return c.JSON(http.StatusBadRequest, res)
	} else {
		date_start = c.FormValue("date_start")

	}

	if !utils.IsValidDate(c.FormValue("date_end")) {
		res := m.SetError(http.StatusBadRequest, "wrong date_end format")
		return c.JSON(http.StatusBadRequest, res)
	} else {
		date_end = c.FormValue("date_end")
	}

	user_id, err = strconv.Atoi(c.FormValue("user_id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, "user_id must be an integer")
		return c.JSON(http.StatusBadRequest, res)
	}

	attendances, err := h.repository.GetActivityByDate(c, date_start, date_end, user_id)
	if err != nil {
		log.Println("[Delivery][GetActivityByDate] can't get attendances, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, "failed to get attendances")
		return c.JSON(http.StatusInternalServerError, res)
	}

	var data []interface{}
	for _, v := range attendances {
		data = append(data, v)
	}

	res := m.SetResponse(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, res)
}
func (h *handler) AddEditActivity(c echo.Context) (err error) {
	var (
		activity string
		id       int
	)

	activity = c.FormValue("activity")
	id, err = strconv.Atoi(c.FormValue("id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, "id must be an integer")
		return c.JSON(http.StatusBadRequest, res)
	}

	attendance, err := h.repository.AddEditActivity(c, activity, id)
	if err != nil {
		log.Println("[Delivery][AddEditActivity] can't add or update activity, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, "failed to add or update activity")
		return c.JSON(http.StatusInternalServerError, res)
	}

	var data []interface{}
	data = append(data, attendance)

	res := m.SetResponse(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, res)
}
func (h *handler) DeleteActivity(c echo.Context) (err error) {
	var (
		id int
	)

	id, err = strconv.Atoi(c.FormValue("id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, "id must be an integer")
		return c.JSON(http.StatusBadRequest, res)
	}

	err = h.repository.DeleteActivity(c, id)
	if err != nil {
		log.Println("[Delivery][AddEditActivity] can't add or update activity, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, "failed to add or update activity")
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK,
		map[string]string{
			"status":  "success",
			"message": "activity deleted",
		})
}
