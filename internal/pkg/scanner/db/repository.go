package db

type StringRepositoryBase struct {
	db        *DB
	tableName string
	createSQL string
	insertSQL string
}

// NewStringRepositoryBase creates a new string-only repo.
func NewStringRepositoryBase(db *DB, tableName, createSQL, insertSQL string) *StringRepositoryBase {
	return &StringRepositoryBase{
		db:        db,
		tableName: tableName,
		createSQL: createSQL,
		insertSQL: insertSQL,
	}
}

// CreateTable runs the CREATE TABLE query.
func (r *StringRepositoryBase) CreateTable() error {
	return r.db.ExecStmt(r.createSQL)
}

// BulkInsert inserts many string rows for a domain.
func (r *StringRepositoryBase) BulkInsert(domainID int, items []string) error {
	if len(items) == 0 {
		return nil
	}

	// Chunk into groups of 1000
	chunks := chunkSlice(items, 1000)

	for _, chunk := range chunks {
		argsList := MakeArgsList(domainID, chunk)
		if err := r.db.ExecBulkInsert(r.insertSQL, argsList); err != nil {
			return err
		}
	}

	return nil
}

// Insert inserts a single string row for a domain.
func (r *StringRepositoryBase) Insert(domainID int, value string) error {
	return r.db.ExecInsert(r.insertSQL, domainID, value)
}
