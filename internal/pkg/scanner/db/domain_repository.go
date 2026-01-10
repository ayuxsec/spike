package db

import "time"

type Domain struct {
	Id        int
	Name      string
	IsScanned bool
	CreatedAt time.Time
}

type DomainRepository struct {
	db *DB
}

func NewDomainRepository(db *DB) *DomainRepository {
	return &DomainRepository{db: db}
}

func (r *DomainRepository) CreateTable() error {
	return r.db.ExecStmt(`
        CREATE TABLE IF NOT EXISTS domains (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            domain TEXT NOT NULL UNIQUE,
            is_scanned BOOLEAN DEFAULT FALSE,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
}

func (r *DomainRepository) Insert(domain string) error {
	return r.db.ExecInsert(
		`INSERT OR IGNORE INTO domains (domain) VALUES (?)`,
		domain,
	)
}

func (r *DomainRepository) BulkInsert(domains []string) error {
	argsList := make([][]any, len(domains))
	for i, d := range domains {
		argsList[i] = []any{d}
	}
	return r.db.ExecBulkInsert(
		`INSERT OR IGNORE INTO domains (domain) VALUES (?)`,
		argsList,
	)
}

func (r *DomainRepository) SelectByName(name string) (*Domain, error) {
	row := r.db.QueryRow(`
		SELECT id, domain, is_scanned, created_at
		FROM domains
		WHERE domain=?`, name)

	var d Domain
	if err := row.Scan(
		&d.Id,
		&d.Name,
		&d.IsScanned,
		&d.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &d, nil
}

func (r *DomainRepository) GetAll() ([]Domain, error) {
	rows, err := r.db.Query(`
		SELECT id, domain, is_scanned, created_at
		FROM domains`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Domain
	for rows.Next() {
		var d Domain
		if err := rows.Scan(
			&d.Id,
			&d.Name,
			&d.IsScanned,
			&d.CreatedAt,
		); err != nil {
			return nil, err
		}
		out = append(out, d)
	}
	return out, nil
}

func (r *DomainRepository) GetUnscanned() ([]Domain, error) {
	rows, err := r.db.Query(`
		SELECT id, domain, is_scanned, created_at
		FROM domains
		WHERE is_scanned=FALSE`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Domain
	for rows.Next() {
		var d Domain
		if err := rows.Scan(
			&d.Id,
			&d.Name,
			&d.IsScanned,
			&d.CreatedAt,
		); err != nil {
			return nil, err
		}
		out = append(out, d)
	}
	return out, nil
}

func (r *DomainRepository) GetScanned() ([]Domain, error) {
	rows, err := r.db.Query(`
		SELECT id, domain, is_scanned, created_at
		FROM domains
		WHERE is_scanned=TRUE`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Domain
	for rows.Next() {
		var d Domain
		if err := rows.Scan(
			&d.Id,
			&d.Name,
			&d.IsScanned,
			&d.CreatedAt,
		); err != nil {
			return nil, err
		}
		out = append(out, d)
	}
	return out, nil
}

func (r *DomainRepository) MarkAsScanned(name string) error {
	return r.db.ExecInsert(
		`UPDATE domains SET is_scanned=TRUE WHERE domain=?`,
		name,
	)
}
