package models

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Attendance struct {
	Id       int    `json:"id" form:"id"`
	UserID   int    `json:"user_id" form:"user_id"`
	Activity string `json:"activity" form:"activity"`
	CheckIn  string `json:"check_in" form:"check_in"`
	CheckOut string `json:"check_out" form:"check_out"`
}

type ResAttendance struct {
	Id       int    `json:"id" form:"id"`
	UserID   int    `json:"user_id" form:"user_id"`
	CheckIn  string `json:"check_in" form:"check_in"`
	CheckOut string `json:"check_out" form:"check_out"`
}
