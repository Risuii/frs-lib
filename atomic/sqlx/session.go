package frsAtomic

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type SqlxAtomicSession struct {
	tx *sqlx.Tx
}

func NewAtomicSession(tx *sqlx.Tx) *SqlxAtomicSession {
	return &SqlxAtomicSession{
		tx: tx,
	}
}

func (s SqlxAtomicSession) Commit(ctx context.Context) error {

	err := s.tx.Commit()
	if err != nil {
		log.Println(err)
	}
	return err
}

func (s SqlxAtomicSession) Rollback(ctx context.Context) error {

	err := s.tx.Rollback()
	if err != nil {
		log.Println(err)
	}
	return err
}

func (s SqlxAtomicSession) Tx() *sqlx.Tx {
	return s.tx
}
