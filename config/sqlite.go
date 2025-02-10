package config

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func InitializeSQLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "dados.db")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS itens (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			item TEXT UNIQUE NOT NULL,
			bpitem TEXT NOT NULL,
			coinsItem TEXT,
			coinsItemBP TEXT,
			coinsItemDesconto TEXT,
			coinsItemBPDesconto TEXT
		);
	
		CREATE TABLE IF NOT EXISTS dinos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			dino TEXT UNIQUE NOT NULL,
			bp TEXT NOT NULL,
			coins TEXT,
			desconto TEXT,
			tokens TEXT
		);
	
		CREATE TABLE IF NOT EXISTS chibis (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			chibi TEXT UNIQUE NOT NULL,
			bpchibi TEXT NOT NULL
		);
		`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
