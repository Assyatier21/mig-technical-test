package database

const (
	GetActivityByDate = `SELECT id, user_id, activity, date(check_in) FROM mig_attendance
                            WHERE date(check_in) BETWEEN $1 AND $2 AND user_id = $3`

	GetActivityByID = `SELECT * FROM mig_attendance
                            WHERE id = $1`

	AddUpdateActivity = `UPDATE mig_attendance
                            SET activity = $1 WHERE id = $2 
                            AND check_in IS NOT NULL`

	DeleteActivity = `UPDATE mig_attendance SET activity = ''
                            WHERE id = $1`
)
