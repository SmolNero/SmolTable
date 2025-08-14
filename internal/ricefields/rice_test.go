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
}