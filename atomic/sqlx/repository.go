package frsAtomic

import (
	"context"
	"log"

	frsAtomic "github.com/Risuii/frs-lib/atomic"
	"github.com/jmoiron/sqlx"
)

type SqlxAtomicSessionProvider struct {
	db *sqlx.DB
}

func NewSqlxAtomicSessionProvider(db *sqlx.DB) *SqlxAtomicSessionProvider {
	return &SqlxAtomicSessionProvider{
		db: db,
	}
}

func (r *SqlxAtomicSessionProvider) BeginSession(ctx context.Context) (*frsAtomic.AtomicSessionContext, error) {

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	atomicSession := NewAtomicSession(tx)
	return frsAtomic.NewAtomicSessionContext(ctx, atomicSession), nil
}
