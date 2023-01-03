package postgres

import (
	"log"
	database "mig-attendance/database/queries"
	m "mig-attendance/models"

	"github.com/labstack/echo/v4"
)

func (r *repository) CheckInAttendance(c echo.Context, check_in_date string, user_id int) (m.ResAttendance, error) {
	var (
		attendance m.ResAttendance
		err        error
	)

	err = r.db.QueryRow(database.CheckInAttendance, user_id, check_in_date).Scan(&attendance.Id, &attendance.UserID, &attendance.CheckIn, &attendance.CheckOut)
	if err != nil {
		log.Println("[Repository][CheckInAttendance] can't check-in, err:", err.Error())
		return m.ResAttendance{}, err
	}

	return attendance, nil
}
func (r *repository) CheckOutAttendance(c echo.Context, check_out_date string, user_id int) (m.Attendance, error) {
	var (
		attendance m.Attendance
		err        error
	)

	err = r.db.QueryRow(database.CheckOutAttendance, check_out_date, user_id).Scan(&attendance.Id, &attendance.UserID, &attendance.Activity, &attendance.CheckIn, &attendance.CheckOut)
	if err != nil {
		log.Println("[Repository][CheckOutAttendance] can't check-out, err:", err.Error())
		return m.Attendance{}, err
	}

	return attendance, nil
}
func (r *repository) GetHistoryAttendance(c echo.Context, user_id int) ([]m.Attendance, error) {
	return []m.Attendance{}, nil
}
