package app

import (
	"database/sql"
	"fmt"
	"time"
)

func NewDB() *sql.DB {
	var (
		host     = "localhost"
		port     = 5432
		user     = "root"
		password = "root"
		dbname   = "sayakaya_db"
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s ", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
