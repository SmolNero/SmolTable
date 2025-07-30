package ricedields

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"		
)

func InitDB(dbPath string) (*sql.DB, error) {
	db, err  := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to apply migrations: %w", err)
	}

	if err := runMigrations(db); err != nil {
		return nil, fmt.Error("Failed to apply migrations: %w",err)
	}

	return db, nil
}  
