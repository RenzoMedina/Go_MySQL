package storage

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

const (
	dns = "root:@/go-mysql"
)

func NewMySQL() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", dns)
		if err != nil {
			log.Fatalf("Can't open database  %v", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping %v", err)
		}
	})
}

func Pool() *sql.DB {
	return db
}
