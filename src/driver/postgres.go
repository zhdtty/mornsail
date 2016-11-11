package driver

import (
	"config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

func newPostgres() *sql.DB {
	if len(*config.FLAG_POSTGRES_DRIVER) <= 0 {
		return nil
	}
	db, err := sql.Open(*config.FLAG_POSTGRES_DRIVER, strings.Replace(*config.FLAG_POSTGRES_SOURCE, ":", "=", -1))
	if err != nil {
		fmt.Println("Postgres open failed!")
		return nil
	}
	db.SetMaxIdleConns(90)
	db.SetMaxOpenConns(90)
	return db
}

var PGPool *sql.DB //= newPostgres()
