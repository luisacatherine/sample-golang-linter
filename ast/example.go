package main

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type SqlConfig struct {
	Connection        *sqlx.DB
	DSN               string
	MaxOpenConnection int
}

func (pg *SqlConfig) Connect(maxconn int) {
	dbobj, err := sqlx.Open("postgres", pg.DSN)
	if err != nil {
		log.Fatal("failed to open postgres connection with error ", err)
	}

	pg.Connection = dbobj
	pg.Connection.SetMaxOpenConns(maxconn)
	pg.Connection.SetConnMaxIdleTime(time.Duration(maxconn))
	pg.Connection.SetConnMaxLifetime(time.Duration(maxconn))

	err = pg.Connection.Ping()
	if err != nil {
		log.Fatal("failed to ping postgres with error ", err)
	}
}
