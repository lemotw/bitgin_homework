package database

import "github.com/jmoiron/sqlx"

var dbconn *sqlx.DB

func init() {
	dbconn = nil
}

func SetDB(db *sqlx.DB) {
	dbconn = db
}

func GetDB() *sqlx.DB {
	return dbconn
}
