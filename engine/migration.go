package engine

import (
	"database/sql"
)

type Migrator func(db *sql.DB)
type MigratorSet struct {
  Up Migrator
  Down Migrator
}
type Migration map[string]MigratorSet

func NewMigration() Migration {
  return Migration{}
}