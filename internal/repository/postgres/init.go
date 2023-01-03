package postgres

import (
	"database/sql"
	m "mig-attendance/models"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	// Login Handler
	Login(c echo.Context, username string, password string) (m.User, error)
	Logout(c echo.Context) error

	// Attendance Handler
	CheckInAttendance(c echo.Context, check_in_date string, id int) (m.ResAttendance, error)
	CheckOutAttendance(c echo.Context, check_out_date string, id int) (m.Attendance, error)
	GetHistoryAttendance(c echo.Context, user_id int) ([]m.Attendance, error)

	// Activity Handler
	GetActivityByDate(c echo.Context, date string) (m.Attendance, error)
	AddEditActivity(c echo.Context, id int) (m.Attendance, error)
	DeleteActivity(c echo.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
