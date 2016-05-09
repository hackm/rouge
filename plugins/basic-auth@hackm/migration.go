package plugin

import (
	// "database/sql"
  "../../engine"
)

var Migration map[string]engine.MigratorSet = make(map[string]engine.MigratorSet)