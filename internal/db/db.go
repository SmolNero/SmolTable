package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"

	// embed schema.sql
 _ "embed"
)

var schema string

func Init(dbPath string) (*sql.DB, error) {
	d, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open db: %W", err)
	}
	if err := migrate(d); err != nil {
		d.Close()
		return nil, err
	}
	return d, nil
}

func migrate(d *sql.DB) error {
	if _, err := d.Exec(schema); err != nil{
		return fmt.Errorf("exec schema: %W", err)
	}
	return nil
}
