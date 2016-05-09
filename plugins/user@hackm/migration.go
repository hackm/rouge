package plugin

import (
	"database/sql"
  "../../engine"
)

var migration map[string]engine.MigratorSet = make(map[string]engine.MigratorSet)

func GetMigration() map[string]engine.MigratorSet {
  migration["0.1.0"] = engine.MigratorSet {
    Up: func(db *sql.DB) {
      db.Exec(
      `SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
        SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
        SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

        -- -----------------------------------------------------
        -- Schema rouge
        -- -----------------------------------------------------
        CREATE SCHEMA IF NOT EXISTS ` + "`" + `rouge` + "`" + ` DEFAULT CHARACTER SET utf8 ;
        USE ` + "`" + `rouge` + "`" + ` ;

        -- -----------------------------------------------------
        -- Table ` + "`" + `rouge` + "`" + `.` + "`" + `users` + "`" + `
        -- -----------------------------------------------------
        CREATE TABLE IF NOT EXISTS ` + "`" + `rouge` + "`" + `.` + "`" + `users` + "`" + ` (
          ` + "`" + `id` + "`" + ` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
          ` + "`" + `name` + "`" + ` VARCHAR(32) NOT NULL,
          ` + "`" + `display_name` + "`" + ` VARCHAR(62) NOT NULL,
          ` + "`" + `email` + "`" + ` VARCHAR(255) NOT NULL,
          ` + "`" + `password` + "`" + ` VARCHAR(256) NOT NULL,
          ` + "`" + `create_at` + "`" + ` DATETIME NOT NULL,
          ` + "`" + `updated_at` + "`" + ` DATETIME NOT NULL,
          PRIMARY KEY (` + "`" + `id` + "`" + `),
          UNIQUE INDEX ` + "`" + `id_UNIQUE` + "`" + ` (` + "`" + `id` + "`" + ` ASC),
          UNIQUE INDEX ` + "`" + `email_UNIQUE` + "`" + ` (` + "`" + `email` + "`" + ` ASC),
          UNIQUE INDEX ` + "`" + `name_UNIQUE` + "`" + ` (` + "`" + `name` + "`" + ` ASC));
          
          SET SQL_MODE=@OLD_SQL_MODE;
          SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
          SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;`)
    }, 
    Down: func(db *sql.DB) {
      db.Exec(`DROP TABLE ` + "`" + `rouge` + "`" + `.` + "`" + `users` + "`" + `;`)
    }, 
  }
  return migration
}
