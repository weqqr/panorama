package worldrepo

import "github.com/lord-server/panorama/internal/db/postgres"

type WorldRepo struct {
	db *postgres.Database
}

func New(db *postgres.Database) *WorldRepo {
	return &WorldRepo{
		db: db,
	}
}
