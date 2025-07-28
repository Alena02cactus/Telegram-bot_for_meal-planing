package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"  // Драйвер MySQL
)

func NewMySQLDB(url string) (*sql.DB, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}