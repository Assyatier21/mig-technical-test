package database

const (
	CheckInAttendance = `INSERT INTO mig_attendance (user_id, check_in) 
                            VALUES ($1, $2) RETURNING id, user_id, check_in, check_out`

	CheckOutAttendance = `UPDATE mig_attendance SET check_out = $1, activity = COALESCE(activity, '')
                            WHERE user_id = $2 AND date(check_in) = date(NOW()) AND check_out IS NULL
                            RETURNING id, user_id, activity, check_in, check_out`

	GetHistoryAttendance = `SELECT * FROM mig_attendance 
                                WHERE user_id = $1 
                                ORDER BY id LIMIT $2 OFFSET $3`
)
