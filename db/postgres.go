package db

import (
	"database/sql"
	"otter-calendar/config"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
	cfg := config.Get()
	var err error
	db, err = sql.Open("postgres", "user="+cfg.DBUser+" password="+cfg.DBPwd+" dbname="+cfg.DBName+" sslmode=disable")

	if err != nil {
		panic(err)
	}
}

func Prepare(query string) (*sql.Stmt, error) {
	return db.Prepare(query)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(query, args...)
}

func QueryRow(query string, args ...interface{}) *sql.Row {
	return db.QueryRow(query, args...)
}
