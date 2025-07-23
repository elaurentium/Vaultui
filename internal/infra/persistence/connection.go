package persistence

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

type SqliteConfig struct {
	Path string
}

func NewSqlite() (*sql.DB, error) {
	config := &SqliteConfig {
		Path: os.Getenv("SQL_PATH"),
	}

	db, err := sql.Open("sqlite", config.Path)
	
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping sqlite database: %w", err)
	}

	return db, nil
}