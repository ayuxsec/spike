package shell

import "spike/internal/scanner/db"

// Context holds the current state of the shell.
type Context struct {
	DB     *db.DB
	Domain *db.Domain
	Data   []string // current working dataset (STDIN equivalent)
}
