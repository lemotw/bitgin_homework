package database

import "database/sql"

var dbconn *sql.DB

func init() {
	dbconn = nil
}

func SetDB(db *sql.DB) {
	dbconn = db
}

func GetDB() *sql.DB {
	return dbconn
}
