package data

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DBTX interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryxContext(context.Context, string, ...any) (*sqlx.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
	QueryRowxContext(context.Context, string, ...any) *sqlx.Row
}

type DataStore interface {
	DB() DBTX
}

type dataStore struct {
	conn *sqlx.DB
	db   DBTX
}

func NewDataStore(db *sqlx.DB) DataStore {
	return &dataStore{
		conn: db,
		db:   db,
	}
}

func (s *dataStore) DB() DBTX {
	return s.db
}

func (s *dataStore) atomic(ctx context.Context, fn func(DataStore) error) error {
	tx, err := s.conn.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	err = fn(&dataStore{conn: s.conn, db: tx})
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return err
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func Atomic[T any](ctx context.Context, baseDs DataStore, wrap func(DataStore) T, handler func(T) error) error {
	ds, ok := baseDs.(*dataStore)
	if !ok {
		panic(fmt.Sprintf("baseDs is not a dataStore: %T", baseDs))
	}

	return ds.atomic(ctx, func(ds DataStore) error {
		domainDs := wrap(ds)
		if err := handler(domainDs); err != nil {
			return err
		}

		return nil
	})
}
