package postgres

import (
	m "mig-attendance/models"

	"github.com/labstack/echo/v4"
)

func (r *repository) GetActivityByDate(c echo.Context, date string) (m.Attendance, error) {
	return m.Attendance{
		Id:       1,
		UserID:   1,
		Activity: "New Activity",
	}, nil
}
func (r *repository) AddEditActivity(c echo.Context, id int) (m.Attendance, error) {
	return m.Attendance{}, nil
}
func (r *repository) DeleteActivity(c echo.Context, id int) error {
	return nil
}
