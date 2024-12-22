package worldrepo

import (
	"context"

	"github.com/lord-server/panorama/internal/db/postgres"
	"github.com/lord-server/panorama/internal/util/iterator"
)

type RawBlock struct {
	X    int    `db:"posx"`
	Y    int    `db:"posy"`
	Z    int    `db:"posz"`
	Data []byte `db:"data"`
}

func (w *WorldRepo) GetBlocksAlongY(ctx context.Context, x, z int) iterator.ResultSeq[RawBlock] {
	return postgres.IterRows[RawBlock](ctx, w.db, `
		SELECT posx,
			   posy,
			   posz,
			   data
		FROM blocks
		WHERE posx = $1 AND posy = $2
		ORDER BY posy`,
	)
}
