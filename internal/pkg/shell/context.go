package shell

import "github.com/ayuxsec/spike/internal/pkg/scanner/db"

// Context holds the database connection and the selected domain for the current REPL shell session
type Context struct {
	DB     *db.DB
	Domain *db.Domain
}
