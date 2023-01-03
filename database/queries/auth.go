package database

const (
	GetUserInfo = `SELECT * FROM mig_users WHERE email = $1`
)
