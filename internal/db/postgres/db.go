package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lord-server/panorama/internal/util/iterator"
	"github.com/lord-server/panorama/internal/util/result"
)

type Database struct {
	pool *pgxpool.Pool
}

func New() (*Database, error) {
	pool, err := pgxpool.New(context.Background(), "")
	if err != nil {
		return nil, err
	}

	return &Database{
		pool: pool,
	}, nil
}

func Execute(ctx context.Context, db *Database, sql string, args ...any) error {
	_, err := db.pool.Exec(ctx, sql, args...)

	return err
}

func QueryRows[T any](ctx context.Context, db *Database, sql string, args ...any) ([]T, error) {
	rows, err := db.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows[T](rows, pgx.RowToStructByNameLax)
}

func IterRows[T any](ctx context.Context, db *Database, sql string, args ...any) iterator.ResultSeq[T] {
	return func(yield func(result.Result[T]) bool) {
		rows, err := db.pool.Query(ctx, sql, args...)
		if err != nil {
			yield(result.Error[T](err))

			return
		}

		defer rows.Close()

		for rows.Next() {
			value, err := pgx.RowToStructByNameLax[T](rows)
			if err != nil {
				yield(result.Error[T](err))

				return
			}

			if !yield(result.Ok(value)) {
				return
			}
		}

		if err := rows.Err(); err != nil {
			yield(result.Error[T](err))
		}
	}
}
