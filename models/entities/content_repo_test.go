package entity

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	con, err := sql.Open("mysql", "root:!Qaz2wsx@/rouge?interpolateParams=true&parseTime=true")
	if err != nil {
		panic(err)
	}
	db = con
}

func TestFindByID(t *testing.T) {
	fmt.Println(FindByID(db, 0))
}

func TestFindByPaperID(t *testing.T) {
	content, err := FindByPaperID(db, 0)
	fmt.Println(content[0], err)
}
