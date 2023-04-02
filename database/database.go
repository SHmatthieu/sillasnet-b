// database package implements functions and struct
// to interact with the mysql db of the project

package database

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

type DbTable interface {
	AddInDB() error
	DeleteFromDB() error
}

var db *sql.DB

func DBConnection() (*sql.DB, error) {

	// UTILSER DES VARIABLES D'ENVIRONNEMENT
	cfg := mysql.Config{
		User:                 "usersec1",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "projecttest",
		AllowNativePasswords: true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
