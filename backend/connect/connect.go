package connect

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() *sql.DB {
	var err error
	db, err = sql.Open("mysql", "abinsms:password@tcp(localhost:3306)/mobiles")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
