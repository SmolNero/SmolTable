package  ricefields

import(
	"testing"
	"os"
	"path/filepath"
	"database/sql"
)

// TestInitDB validates that InitDB creates a DB and applies schema.sql
func TestInitDB(t *testing.t) {
	schemeDir := filepath.Join("internal", "ricefields")
	schemaPath := filepath.Join(schemeDir, "schemea.sql")
	if err := os.MkdirAll(schemeDir, 0o755); err != nil {
		t.Fatalf("failed to create schema dir: %v", err)
	}

	
}