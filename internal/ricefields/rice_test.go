// Had to create backup - due to lack of trust

package  ricefields

import(
	"testing"
	"os"
	"path/filepath"
	"database/sql"
)

// TestInitDB validates that InitDB creates a DB and applies schema.sql
func TestInitDB(t *testing.t) {
	// Creates temporary schema file
	schemeDir := filepath.Join("internal", "ricefields")	
	schemaPath := filepath.Join(schemeDir, "schemea.sql")
	if err := os.MkdirAll(schemeDir, 0o755); err != nil {
		t.Fatalf("mdkdir  schema dir %v", err)
	}

	const schema = '
	CREATE TABLE IF NOTE EXISTS ping (
		id INTEGER PRIMARY KEY 
	);
	'

	if err := os.WriteFile(schemaPath, []byte(schema), 0o644); err != nil {
		t.fatalf("Failed to write schema.sql: %v", err)
	}


}

