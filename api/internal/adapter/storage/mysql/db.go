package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"leinadium.dev/wca-ranking/internal/adapter/config"
)

const (
	maxLifetime  = time.Minute * 3
	maxOpenConns = 10
	maxIdleConns = 10
)

type DB struct {
	db *sql.DB
}

func New(config *config.DB) (*DB, error) {
	db, err := sql.Open("mysql", createDNS(config))
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(maxLifetime)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	return &DB{db: db}, nil
}

func createDNS(config *config.DB) string {
	c := mysql.NewConfig()

	// basic parameters
	c.User = config.User
	c.Passwd = config.Password
	c.Addr = fmt.Sprintf("%s:%d", config.Host, config.Port)

	// advanced parameters
	c.ParseTime = true

	return c.FormatDSN()
}
