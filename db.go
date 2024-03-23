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

	if err := s.createProjectTable(); err != nil {
		return nil, err
	}

	if err := createUserTable(); err != nil {
		return nil, err
	}

	if err := createTasksTable(); err != nil {
		return nil, err
	}

	return s.db, nil
}

func (s *MYSQLSTORAGE) createProjectTable() error {
	_, err := s.db.Exec(`
    CREATE TABLE IF NOT EXISTS projects (
     id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id)

    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

    `)

	return err
}
