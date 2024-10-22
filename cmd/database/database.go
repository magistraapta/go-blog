package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() error {
	var err error
	DB, err = sqlx.Connect("postgres", "user=postgres dbname=go-blog sslmode=disable host=localhost port=5432")
	if err != nil {
		return err
	}
	return DB.Ping()
}
