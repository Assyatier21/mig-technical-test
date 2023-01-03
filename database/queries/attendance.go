package database

const (
	CheckInAttendance = `INSERT INTO mig_attendance (user_id, check_in) 
                            VALUES ($1, $2) RETURNING id, user_id, check_in, check_out`

	CheckOutAttendance = `UPDATE mig_attendance SET check_out = $1
                            WHERE id = $2 RETURNING id, user_id, activity, check_in, check_out`

	GetHistoryAttendance = `SELECT * FROM mig_attendance 
                                WHERE user_id = $1`
)
