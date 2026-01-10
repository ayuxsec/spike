package db

import (
	"fmt"
)

// STRING-ONLY REPOSITORIES
type StringRepository struct {
	*StringRepositoryBase
	valueColumn string
}

// NewStringRepository creates a new string-only repo.
func NewStringRepository(db *DB, tableName, valueColumn, createSQL, insertSQL string) *StringRepository {
	return &StringRepository{
		StringRepositoryBase: NewStringRepositoryBase(db, tableName, createSQL, insertSQL),
		valueColumn:          valueColumn,
	}
}

// Fetch data by domain ID
func (r *StringRepository) Fetch(domainID int) ([]string, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE domain_id = ?", r.valueColumn, r.tableName)

	rows, err := r.db.Query(query, domainID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			return nil, err
		}
		results = append(results, value)
	}

	return results, nil
}

// INDIVIDUAL TOOL REPOS
func NewSubfinderRepository(db *DB) *StringRepository {
	return NewStringRepository(db,
		"subfinder",
		"subdomain",
		`CREATE TABLE IF NOT EXISTS subfinder (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			domain_id INTEGER NOT NULL,
			subdomain TEXT NOT NULL,
			FOREIGN KEY (domain_id) REFERENCES domains(id) ON DELETE CASCADE
		)`,
		`INSERT INTO subfinder (domain_id, subdomain) VALUES (?, ?)`,
	)
}

func NewHttpxRepository(db *DB) *StringRepository {
	return NewStringRepository(db,
		"httpx",
		"hosts",
		`CREATE TABLE IF NOT EXISTS httpx (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			domain_id INTEGER NOT NULL,
			hosts TEXT NOT NULL,
			FOREIGN KEY (domain_id) REFERENCES domains(id) ON DELETE CASCADE
		)`,
		`INSERT INTO httpx (domain_id, hosts) VALUES (?, ?)`,
	)
}

func NewGauRepository(db *DB) *StringRepository {
	return NewStringRepository(db,
		"gau",
		"url",
		`CREATE TABLE IF NOT EXISTS gau (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			domain_id INTEGER NOT NULL,
			url TEXT NOT NULL,
			FOREIGN KEY (domain_id) REFERENCES domains(id) ON DELETE CASCADE
		)`,
		`INSERT INTO gau (domain_id, url) VALUES (?, ?)`,
	)
}

func NewKatanaRepository(db *DB) *StringRepository {
	return NewStringRepository(db,
		"katana",
		"url",
		`CREATE TABLE IF NOT EXISTS katana (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			domain_id INTEGER NOT NULL,
			url TEXT NOT NULL,
			FOREIGN KEY (domain_id) REFERENCES domains(id) ON DELETE CASCADE
		)`,
		`INSERT INTO katana (domain_id, url) VALUES (?, ?)`,
	)
}

func NewNucleiRepository(db *DB) *StringRepository {
	return NewStringRepository(db,
		"nuclei",
		"report",
		`CREATE TABLE IF NOT EXISTS nuclei (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			domain_id INTEGER NOT NULL,
			report TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (domain_id) REFERENCES domains(id) ON DELETE CASCADE
		)`,
		`INSERT INTO nuclei (domain_id, report) VALUES (?, ?)`,
	)
}

func NewCachexRepository(db *DB) *StringRepository {
	return NewStringRepository(db,
		"cachex",
		"result",
		`CREATE TABLE IF NOT EXISTS cachex (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			domain_id INTEGER NOT NULL,
			result TEXT NOT NULL,
			FOREIGN KEY (domain_id) REFERENCES domains(id) ON DELETE CASCADE
		)`,
		`INSERT INTO cachex (domain_id, result) VALUES (?, ?)`,
	)
}

func NewUroRepository(db *DB) *StringRepository {
	return NewStringRepository(db,
		"uro",
		"url",
		`CREATE TABLE IF NOT EXISTS uro (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            domain_id INTEGER NOT NULL,
            url TEXT NOT NULL,
            FOREIGN KEY (domain_id) REFERENCES domains(id) ON DELETE CASCADE
        )`,
		`INSERT INTO uro (domain_id, url) VALUES (?, ?)`,
	)
}

// ---------------------------
//  AGGREGATED TOOLS REPOSITORY
// ---------------------------

type ToolsRepository struct {
	db        *DB
	Subfinder *StringRepository
	Httpx     *StringRepository
	Gau       *StringRepository
	Katana    *StringRepository
	Nuclei    *StringRepository
	Cachex    *StringRepository
	Uro       *StringRepository
}

func NewToolsRepository(db *DB) *ToolsRepository {
	return &ToolsRepository{
		db:        db,
		Subfinder: NewSubfinderRepository(db),
		Httpx:     NewHttpxRepository(db),
		Gau:       NewGauRepository(db),
		Katana:    NewKatanaRepository(db),
		Nuclei:    NewNucleiRepository(db),
		Cachex:    NewCachexRepository(db),
		Uro:       NewUroRepository(db),
	}
}

func (tr *ToolsRepository) CreateTables() error {
	repos := []*StringRepository{
		tr.Subfinder,
		tr.Httpx,
		tr.Gau,
		tr.Katana,
		tr.Nuclei,
		tr.Cachex,
		tr.Uro,
	}

	for _, repo := range repos {
		if err := repo.CreateTable(); err != nil {
			return err
		}
	}
	return nil
}
