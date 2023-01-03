package postgres

import (
	"log"
	database "mig-attendance/database/queries"
	m "mig-attendance/models"
	"mig-attendance/utils"

	"github.com/labstack/echo/v4"
)

func (r *repository) GetActivityByDate(c echo.Context, date_start string, date_end string, user_id int) ([]m.ResAttendanceByDate, error) {
	var (
		attendances []m.ResAttendanceByDate
		err         error
	)

	rows, err := r.db.Query(database.GetActivityByDate, date_start, date_end, user_id)
	if err != nil {
		log.Println("[Repository][GetActivityByDate] can't get attendances, err:", err.Error())
		return nil, err
	}

	for rows.Next() {
		var temp = m.ResAttendanceByDate{}
		if err := rows.Scan(&temp.Id, &temp.UserID, &temp.Activity, &temp.Date); err != nil {
			log.Println("[GetArticles] failed to scan article, err :", err.Error())
			return nil, err
		}
		temp.Date = utils.FormattedTimeActivity(temp.Date)
		attendances = append(attendances, temp)
	}

	if len(attendances) > 0 {
		return attendances, nil
	} else {
		return []m.ResAttendanceByDate{}, nil
	}
}
func (r *repository) AddEditActivity(c echo.Context, id int) (m.Attendance, error) {
	return m.Attendance{}, nil
}
func (r *repository) DeleteActivity(c echo.Context, id int) error {
	return nil
}
