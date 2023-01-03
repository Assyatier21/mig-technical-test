package database

const (
	GetActivityByDate = `SELECT activity, check_in FROM mig_attendance
                            WHERE date(check_in) = $1 AND user_id = $2`

	AddUpdateActivity = `UPDATE mig_attendance SET activity = $1
                            WHERE id = $2`

	DeleteActivity = `UPDATE mig_attendance SET activity = ''
                            WHERE id = $1`
)
