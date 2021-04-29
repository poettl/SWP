package DBHelper

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
)

var db *sql.DB = nil

func GetDB() *sql.DB {
	if db == nil {
		fmt.Printf("new object\n")

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		newdb, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}

		err = newdb.Ping()
		if err != nil {
			panic(err)
		}
		db = newdb
		return db
	}
	return db

}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
