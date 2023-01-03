package utils

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	WrongUsernamePass = errors.New("wrong username or password")
	ErrNotFound       = errors.New("data not found")
	NoRowsAffected    = errors.New("no rows affected")
	jakartaLoc, _     = time.LoadLocation("Asia/Jakarta")
	TimeNow           = time.Now().In(jakartaLoc).Format("2006-01-02T15:04:05Z")
)

func IsValidEmail(s string) bool {
	regex, _ := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regex.MatchString(s)
}

func IsValidAlphabet(s string) bool {
	regex, _ := regexp.Compile(`^[a-zA-Z]*$`)
	return regex.MatchString(s)
}

func IsValidNumeric(s string) bool {
	regex, _ := regexp.Compile(`([0-9])`)
	return regex.MatchString(s)
}

func IsValidAlphaNumeric(s string) bool {
	regex, _ := regexp.Compile(`(^[a-zA-Z0-9]*$)`)
	return regex.MatchString(s)
}

func IsValidSlug(s string) bool {
	regex, _ := regexp.Compile(`^[a-z0-9-]+$`)
	return regex.MatchString(s)
}

func FormattedTime(ts string) string {
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		log.Println(err)
		return ""
	}

	formattedTime := t.Format("2006-01-02 15:04:05")
	return formattedTime
}

func SetCookie(c echo.Context, name string, value string, expiration time.Time) {
	cookie := &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expiration,
	}

	c.SetCookie(cookie)
}
func InvalidateCookie(cookie *http.Cookie) {
	cookie.Expires = time.Now().Add(-24 * time.Hour)
}
