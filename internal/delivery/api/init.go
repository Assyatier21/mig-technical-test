package api

import (
	"mig-attendance/internal/repository/postgres"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	// Login Handler
	Login(c echo.Context) (err error)
	Logout(c echo.Context) (err error)

	// Attendance Handler
	CheckInAttendance(c echo.Context) (err error)
	CheckOutAttendance(c echo.Context) (err error)
	GetHistoryAttendance(c echo.Context) (err error)

	// Activity Handler
	GetActivity(c echo.Context) (err error)
	GetActivityByDate(c echo.Context) (err error)
	AddEditActivity(c echo.Context) (err error)
	DeleteActivity(c echo.Context) (err error)
}

type handler struct {
	repository postgres.Repository
}

func New(repository postgres.Repository) Handler {
	return &handler{
		repository: repository,
	}
}
