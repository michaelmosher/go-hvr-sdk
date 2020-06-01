package postgresql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type service struct {
	db *sqlx.DB
}

func New(connectionString string) (service, error) {
	db, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		return service{}, fmt.Errorf("error connecting to postgres database: %s", err)
	}

	return service{db: db}, nil
}
