package db

import (
	"database/sql"
	"fmt"
)

type ScanTracker struct {
	Id        int
	DomainID  int
	ToolName  string
	Status    string
	CreatedAt string
}

type ScanTrackerRepository struct {
	db *DB
}

func NewScanTrackerRepository(db *DB) *ScanTrackerRepository {
	return &ScanTrackerRepository{db: db}
}

func (r *ScanTrackerRepository) CreateTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS scan_tracker (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            domain_id INTEGER NOT NULL,
            tool_name TEXT NOT NULL,
            status TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            UNIQUE(domain_id, tool_name),
            FOREIGN KEY (domain_id) REFERENCES domains(id) ON DELETE CASCADE
        );
    `
	return r.db.ExecStmt(query)
}

// Mark scan completed using FULL UPSERT
func (r *ScanTrackerRepository) MarkScanCompleted(domainID int, toolName string) error {
	_, err := r.db.Exec(`
        INSERT INTO scan_tracker (domain_id, tool_name, status)
        VALUES (?, ?, 'completed')
        ON CONFLICT(domain_id, tool_name) DO UPDATE SET status='completed';
    `, domainID, toolName)
	return err
}

// Update status manually
func (r *ScanTrackerRepository) UpdateScanStatus(domainID int, toolName, status string) error {
	_, err := r.db.Exec(`
        INSERT INTO scan_tracker (domain_id, tool_name, status)
        VALUES (?, ?, ?)
        ON CONFLICT(domain_id, tool_name) DO UPDATE SET status=excluded.status;
    `, domainID, toolName, status)
	return err
}

// Check if scan is completed
func (r *ScanTrackerRepository) IsScanCompleted(domainID int, toolName string) (bool, error) {
	var status string

	err := r.db.QueryRow(`
        SELECT status FROM scan_tracker
        WHERE domain_id = ? AND tool_name = ?
        LIMIT 1;
    `, domainID, toolName).Scan(&status)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // Not completed
		}
		return false, fmt.Errorf("query failed: %w", err)
	}

	return status == "completed", nil
}
