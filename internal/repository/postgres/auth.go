package postgres

import (
	"fmt"
	"log"
	database "mig-attendance/database/queries"
	m "mig-attendance/models"
	"mig-attendance/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (r *repository) Login(c echo.Context, email string, password string) (m.User, error) {
	var (
		user m.User
		err  error
	)

	err = r.db.QueryRow(database.GetUserInfo, email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Println("[Repository][Login] can't get user info, err:", err.Error())
		return m.User{}, utils.WrongUsernamePass
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println("[Repository][Login] can't login, err:", err.Error())
		return m.User{}, utils.WrongUsernamePass
	}

	utils.SetCookie(c, "auth", "randomstring", time.Now().Add(24*time.Hour))

	return user, nil
}
func (r *repository) Logout(c echo.Context) error {
	fmt.Println(c.Cookies())
	cookie, err := c.Cookie("auth")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Cookie not found")
	}

	utils.InvalidateCookie(cookie)

	c.SetCookie(&http.Cookie{
		Name:    cookie.Name,
		Value:   "",
		Expires: time.Now().Add(-24 * time.Hour),
	})

	return nil
}
