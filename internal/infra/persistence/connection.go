package persistence

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteConfig struct {
	Path string
}

func NewSqlite() (*sql.DB, error) {
	config := &SqliteConfig {
		Path: getDBPath(),
	}

	db, err := sql.Open("sqlite3", config.Path)
	
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping sqlite database: %w", err)
	}

	return db, nil
}

func getDBPath() string {
	path := os.Getenv("SQL_PATH")

	if path == " " {
		return "sqlite not found"
	}

	return path
}