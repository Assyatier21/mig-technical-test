package postgres

import (
	m "mig-attendance/models"

	"github.com/labstack/echo/v4"
)

func (r *repository) CheckInAttendance(ctx echo.Context, check_in_date string, id int) (m.Attendance, error) {
	return m.Attendance{}, nil
}
func (r *repository) CheckOutAttendance(ctx echo.Context, check_out_date string, id int) (m.Attendance, error) {
	return m.Attendance{}, nil
}
func (r *repository) GetHistoryAttendance(ctx echo.Context, user_id int) ([]m.Attendance, error) {
	return []m.Attendance{}, nil
}
