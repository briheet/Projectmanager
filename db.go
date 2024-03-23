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

	if err := s.createUserTable(); err != nil {
		return nil, err
	}

	if err := s.createTasksTable(); err != nil {
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
	if err != nil {
		return err
	}

	return nil
}

func (s *MYSQLSTORAGE) createTasksTable() error {
	_, err := s.db.Exec(`
    CREATE TABLE IF NOT EXISTS tasks (
     id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    status ENUM( 'TODO', 'IN_PROGRESS', 'IN_TESTING', 'DONE') NOT NULL DEFAULT 'TODO',
    projectID INT UNSIGNED NOT NULL,
    assignedToID INT UNSIGNED NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id),
    FOREIGN KEY (assignedToID) REFERENCES users(id),
    FOREIGN KEY (projectID) REFERENCES projects(id)

    )
    ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `)
	if err != nil {
		return err
	}

	return nil
}

func (s *MYSQLSTORAGE) createUserTable() error {
	_, err := s.db.Exec(`
    CREATE TABLE IF NOT EXISTS users (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL,
    firstname VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id),
    UNIQUE key(email)

    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `)
	if err != nil {
		return err
	}

	return nil
}
