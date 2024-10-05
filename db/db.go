package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewSqlStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Error in database connection", err)
	}
	return db, nil
}

func InitSqlStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Error in initalizing database", err)
	}
	log.Println("Database connected")
}
