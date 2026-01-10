package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

// MakeArgsList takes a prefix (like domainID) and a slice of strings (like URLs or reports),
// and builds a [][]any that can be used with ExecBulkInsert.
func MakeArgsList(prefix any, values []string) [][]any {
	argsList := make([][]any, len(values))
	for i, val := range values {
		argsList[i] = []any{prefix, val}
	}
	return argsList
}

// ExecBulkInsert runs a prepared, transactional bulk insert.
// query: INSERT statement with placeholders.
// argsList: each inner slice represents one row of arguments.
func (db *DB) ExecBulkInsert(query string, argsList [][]any) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, args := range argsList {
		if _, err := stmt.Exec(args...); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// Exec executes an query with args without returning any rows.
func (db *DB) ExecInsert(query string, args ...any) error {
	_, err := db.Exec(query, args...)
	return err
}

// ExecStmt executes a statement without returning any rows.
func (db *DB) ExecStmt(stmt string) error {
	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}

// Connect opens a database connection to the specified path.
func (db *DB) Connect(dbPath string) error {
	var err error
	db.DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}
	if err := db.ExecStmt("PRAGMA foreign_keys = ON;"); err != nil {
		return fmt.Errorf("failed to enable foreign keys: %v", err)
	}
	return nil
}

// Close closes the database connection.
func (db *DB) Close() error {
	if err := db.DB.Close(); err != nil {
		return err
	}
	return nil
}
