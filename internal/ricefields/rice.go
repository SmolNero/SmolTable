// internal/ricefields/rice.go
package ricefields

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB opens (or creates) a SQLite database and runs migrations.
func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	if err := runMigrations(db); err != nil { // now defined below
		return nil, fmt.Errorf("failed to apply migrations: %w", err)
	}
	return db, nil
}

// runMigrations reads internal/ricefields/schema.sql and executes it.
func runMigrations(db *sql.DB) error {
	schemaPath := filepath.Join("internalSqlTest", "ricefieldsSqlTest", "schema.sql")
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("read schema: %w", err)
	}
	if _, err := db.Exec(string(schema)); err != nil {
		return fmt.Errorf("exec schema: %w", err)
	}
	return nil
}

