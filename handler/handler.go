package handler

import (
	"database/sql"

	"github.com/Andesson/arkscript/config"
)

var (
	db *sql.DB
)

func InitializeHandler() {
	db = config.GetSQLite()
}
