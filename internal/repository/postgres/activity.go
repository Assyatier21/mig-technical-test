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
func (r *repository) AddEditActivity(c echo.Context, activity string, id int) (m.Attendance, error) {
	var (
		attendance m.Attendance
		err        error
	)

	rows, err := r.db.Exec(database.AddUpdateActivity, activity, id)
	if err != nil {
		log.Println("[UpdateArticle] can't update article, err:", err.Error())
		return m.Attendance{}, err
	}

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected > 0 {
		err = r.db.QueryRow(database.GetActivityByID, id).Scan(&attendance.Id, &attendance.UserID, &attendance.Activity, &attendance.CheckIn, &attendance.CheckOut)
		if err != nil {
			log.Println("[Repository][GetActivityByID] failed to scan attendance, err:", err.Error())
			return m.Attendance{}, err
		}

		attendance.CheckIn = utils.FormattedTime(attendance.CheckIn)
		attendance.CheckOut = utils.FormattedTime(attendance.CheckOut)

		return attendance, nil
	} else {
		log.Println("Repository][AddUpdateActivity] err:", utils.NoRowsAffected)
		return m.Attendance{}, utils.NoRowsAffected
	}
}
func (r *repository) DeleteActivity(c echo.Context, id int) error {
	var (
		err error
	)

	rows, err := r.db.Exec(database.DeleteActivity, id)
	if err != nil {
		log.Println("[UpdateArticle] can't update article, err:", err.Error())
		return err
	}

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected > 0 {
		return nil
	} else {
		log.Println("Repository][AddUpdateActivity] err:", utils.NoRowsAffected)
		return utils.NoRowsAffected
	}
}
