package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


func Connect(dbURL string) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", dbURL)
}