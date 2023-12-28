package frsAtomic

import (
	"context"
	"log"

	frsAtomic "github.com/Risuii/frs-lib/atomic"
	"github.com/jmoiron/sqlx"

	"go.opentelemetry.io/otel/trace"
)

type SqlxAtomicSessionProvider struct {
	db    *sqlx.DB
	trace trace.Tracer
}

func NewSqlxAtomicSessionProvider(db *sqlx.DB, tr trace.Tracer) *SqlxAtomicSessionProvider {
	return &SqlxAtomicSessionProvider{
		db:    db,
		trace: tr,
	}
}

func (r *SqlxAtomicSessionProvider) BeginSession(ctx context.Context) (*frsAtomic.AtomicSessionContext, error) {
	ctx, span := r.trace.Start(ctx, "SqlxAtomicSessionProvider/BeginSession")
	defer span.End()

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	atomicSession := NewAtomicSession(tx, r.trace)
	return frsAtomic.NewAtomicSessionContext(ctx, atomicSession), nil
}
