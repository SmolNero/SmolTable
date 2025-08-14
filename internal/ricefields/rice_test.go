package  ricefields

import(
	"testing"
	"os"
	"path/filepath"
	"database/sql"
)

// TestInitDB validates that InitDB creates a DB and applies schema.sql
func TestInitDB