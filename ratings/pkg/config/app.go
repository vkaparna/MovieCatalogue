package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sqlx.DB
)

func Connect() {
	d, err := sqlx.Open("sqlite3", "../database/ratingsData.db")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *sqlx.DB {
	return db
}
