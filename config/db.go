package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func dbCon() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "db_gocrud"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}
