package models

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"user" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"pass" db:"password"`
	Salt     string `db:"salt"`
}
