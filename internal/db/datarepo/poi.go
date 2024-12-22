package datarepo

import (
	"context"

	"github.com/lord-server/panorama/internal/db/postgres"
	"github.com/lord-server/panorama/internal/domain"
	"github.com/lord-server/panorama/internal/util/layers/into"
)

type POI struct {
	ID   string `db:"id"`
	Name string `db:"name"`
	Type string `db:"type"`
}

func (p *POI) Into() domain.POI {
	return domain.POI{
		ID:   p.ID,
		Name: p.Name,
		Type: p.Type,
	}
}

func (r *DataRepo) ListPOIs(ctx context.Context) ([]domain.POI, error) {
	return into.SliceOf[domain.POI](postgres.QueryRows[POI](ctx, r.db, `
		SELECT id,
		       name,
		       type
		FROM pois`,
	))
}

func (r *DataRepo) CreatePOI(ctx context.Context, poi POI) error {
	return postgres.Execute(ctx, r.db, `
		INSERT INTO pois (id, name, type)
		VALUES ($1, $2, $3)`,

		poi.ID,
		poi.Name,
		poi.Type,
	)
}

func (r *DataRepo) DeletePOI(ctx context.Context, id string) error {
	return postgres.Execute(ctx, r.db, `
		DELETE FROM pois WHERE id = $1`,

		id,
	)
}
