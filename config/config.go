package config

import (
	"database/sql"
	"fmt"
)

var (
	db *sql.DB
)

func Init() error {
	var err error
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("error initializing SQLite: %v", err)
	}
	return nil
}

func GetSQLite() *sql.DB {
	return db
}
