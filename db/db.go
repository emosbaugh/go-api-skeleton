package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Init(dataSourceName string) (*sqlx.DB, error) {
	return sqlx.Open("postgres", dataSourceName)
}
