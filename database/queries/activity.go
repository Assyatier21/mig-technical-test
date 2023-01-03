package database

const (
	GetActivityByDate = `SELECT activity, check_in FROM mig_attendance
                            WHERE date(check_in) BETWEEN $1 AND $2 AND user_id = $3`

	AddUpdateActivity = `UPDATE mig_attendance
                            SET activity = $1 WHERE id = $2 
                            AND check_in IS NOT NULL`

	DeleteActivity = `UPDATE mig_attendance SET activity = ''
                            WHERE id = $1`
)
