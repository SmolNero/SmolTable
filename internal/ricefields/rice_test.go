// internal/ricefields/rice_test.go
package ricefields

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// TestInitDB validates that InitDB creates a DB and applies schema.sql
func TestInitDB(t *testing.T) {
	// 1) Ensure schema.sql exists where InitDB expects it
	schemaDir := filepath.Join("internalSqlTest", "ricefieldsSqlTest")
	schemaPath := filepath.Join(schemaDir, "schema.sql")
	if err := os.MkdirAll(schemaDir, 0o755); err != nil {
		t.Fatalf("failed to create schema dir: %v", err)
	}

	// Minimal schema just for testing
	const schema = `
	CREATE TABLE IF NOT EXISTS ping (
		id INTEGER PRIMARY KEY
	);
	`
	if err := os.WriteFile(schemaPath, []byte(schema), 0o644); err != nil {
		t.Fatalf("failed to write schema.sql: %v", err)
	}

	// 2) Use a temp DB path so we don’t touch real data
	tmp := t.TempDir()
	dbPath := filepath.Join(tmp, "ricefields_test.db")

	// 3) Call InitDB
	db, err := InitDB(dbPath)
	if err != nil {
		t.Fatalf("InitDB failed: %v", err)
	}
	defer db.Close()

	// 4) Verify the test table exists
	if err := assertTableExists(db, "ping"); err != nil {
		t.Fatalf("expected table not found: %v", err)
	}
}

// assertTableExists checks if a table exists in the SQLite DB
func assertTableExists(db *sql.DB, table string) error {
	row := db.QueryRow(
		`SELECT name FROM sqlite_master WHERE type='table' AND name=?`,
		table,
	)
	var name string
	if err := row.Scan(&name); err != nil {
		return err
	}
	if name != table {
		return fmt.Errorf("expected %q but got %q", table, name)
	}
	return nil
}
