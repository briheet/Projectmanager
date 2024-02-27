package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MYSQLSTORAGE struct {
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) *MYSQLSTORAGE {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MySql")

	return &MYSQLSTORAGE{
		db: db,
	}
}

func (s *MYSQLSTORAGE) Init() (*sql.DB, error) {
	// Tables
	return s.db, nil
}
