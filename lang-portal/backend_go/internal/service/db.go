package service

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(dbPath string) error {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	return db.Ping()
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
