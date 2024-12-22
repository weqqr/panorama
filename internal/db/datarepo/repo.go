package datarepo

import "github.com/lord-server/panorama/internal/db/postgres"

type DataRepo struct {
	db *postgres.Database
}

func New(db *postgres.Database) *DataRepo {
	return &DataRepo{
		db: db,
	}
}
