package Databases

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	Path string `yaml:"path"`
	Db   *sql.DB
}

func (sqlc *SQLite) Initial() (err error) {
	sqlc.Db, err = sql.Open("sqlite3", "./storage/calendly.db")
	if err != nil {
		panic(err)
	}
	return err
}

func (sqlc *SQLite) Close() {
	sqlc.Close()
}

func (sqlc *SQLite) GetConnection() *sql.DB {
	return sqlc.Db
}
