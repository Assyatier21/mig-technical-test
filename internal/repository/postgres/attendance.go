package postgres

import (
	"log"
	database "mig-attendance/database/queries"
	m "mig-attendance/models"

	"github.com/labstack/echo/v4"
)

func (r *repository) CheckInAttendance(c echo.Context, check_in_date string, id int) (m.ResAttendance, error) {
	var (
		attendance m.ResAttendance
		err        error
	)

	err = r.db.QueryRow(database.CheckInAttendance, check_in_date, id).Scan(&attendance.Id, &attendance.UserID, &attendance.CheckIn, &attendance.CheckOut)
	if err != nil {
		log.Println("[Repository][CheckInAttendance] can't check-in, err:", err.Error())
		return m.ResAttendance{}, err
	}

	return attendance, nil
}
func (r *repository) CheckOutAttendance(c echo.Context, check_out_date string, id int) (m.Attendance, error) {
	return m.Attendance{}, nil
}
func (r *repository) GetHistoryAttendance(c echo.Context, user_id int) ([]m.Attendance, error) {
	return []m.Attendance{}, nil
}
