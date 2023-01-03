package postgres

import (
	"database/sql"
	m "mig-attendance/models"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	// Login Handler
	Login(ctx echo.Context, username string, password string) (m.User, error)
	Logout(ctx echo.Context) error

	// Attendance Handler
	CheckInAttendance(ctx echo.Context, check_in_date string, id int) (m.Attendance, error)
	CheckOutAttendance(ctx echo.Context, check_out_date string, id int) (m.Attendance, error)
	GetHistoryAttendance(ctx echo.Context, user_id int) ([]m.Attendance, error)

	// Activity Handler
	GetActivityByDate(ctx echo.Context, date string) (m.Attendance, error)
	AddEditActivity(ctx echo.Context, id int) (m.Attendance, error)
	DeleteActivity(ctx echo.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
