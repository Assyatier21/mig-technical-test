package database

const (
	CheckInAttendance = `UPDATE mig_attendance SET check_in = $1
                            WHERE id = $2`

	CheckOutAttendance = `UPDATE mig_attendance SET check_out = $1
                            WHERE id = $2`

	GetHistoryAttendance = `SELECT * FROM mig_attendance 
                                WHERE user_id = $1`
)
