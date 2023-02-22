package database

import (
	"github.com/jmoiron/sqlx"
	"taalhach/go-http-server/configs"

	_ "github.com/lib/pq"
)

func ConnectDatabase(configs configs.DBConfigs) (*sqlx.DB, error) {
	// Open database connection
	db, err := sqlx.Connect("postgres", configs.PGConnectionString())
	if err != nil {
		return nil, err
	}

	// Test database connection
	err = db.Ping()
	return db, err
}
