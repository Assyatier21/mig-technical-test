package main

import (
	"fmt"
	db "mig-attendance/database"
	"mig-attendance/internal/delivery/api"
	"mig-attendance/internal/repository/postgres"
	"mig-attendance/routes"
)

func main() {
	db := db.Init()

	repository := postgres.New(db)
	handler := api.New(repository)
	echo := routes.GetRoutes(handler)

	host := fmt.Sprintf("%s:%s", "127.0.0.1", "8800")
	_ = echo.Start(host)
}
