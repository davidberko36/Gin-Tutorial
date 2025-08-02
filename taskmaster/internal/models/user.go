package models


type User struct {
	ID int64 `db:"id"` // This helps to bind go fields to db columns using libraries like sqlx
	Username string `db:"username"`
	Email string `db:"email"`
	Password string `db:"password_hash"`
}