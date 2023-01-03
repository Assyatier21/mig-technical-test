package api

import (
	"mig-attendance/internal/repository/postgres"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	// Login Handler
	Login(ctx echo.Context) (err error)
	Logout(ctx echo.Context) (err error)

	// Attendance Handler
	CheckInAttendance(ctx echo.Context) (err error)
	CheckOutAttendance(ctx echo.Context) (err error)
	GetHistoryAttendance(ctx echo.Context) (err error)

	// Activity Handler
	GetActivity(ctx echo.Context) (err error)
	GetActivityByDate(ctx echo.Context) (err error)
	AddEditActivity(ctx echo.Context) (err error)
	DeleteActivity(ctx echo.Context) (err error)
}

type handler struct {
	repository postgres.Repository
}

func New(repository postgres.Repository) Handler {
	return &handler{
		repository: repository,
	}
}
