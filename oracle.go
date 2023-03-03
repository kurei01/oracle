// connect oracle
package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/godror/godror"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("godror", "SYSTEM/ntanh01699909202@192.168.12.39:1521/xe")
	if err != nil {
		panic(err)
	}

	Ping(db);

	db.SetConnMaxLifetime(7 * time.Minute)
	db.SetConnMaxIdleTime(7 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("Connected to database")
	return db, nil
}

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func Ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
}
