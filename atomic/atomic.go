package frsAtomic

import (
	"context"
	"errors"
	"log"
)

var InvalidAtomicSessionProvider error = errors.New("invalid_atomic_session_provider")

type AtomicSessionProvider interface {
	BeginSession(ctx context.Context) (*AtomicSessionContext, error)
}

type AtomicSession interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type AtomicSessionContext struct {
	context.Context
	AtomicSession
}

func NewAtomicSessionContext(ctx context.Context, session AtomicSession) *AtomicSessionContext {
	return &AtomicSessionContext{
		Context:       ctx,
		AtomicSession: session,
	}
}

func Atomic(ctx context.Context, provider AtomicSessionProvider, fn func(ctx context.Context) error) error {
	sessionCtx, err := provider.BeginSession(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if v := recover(); v != nil {
			if rbErr := sessionCtx.Rollback(ctx); rbErr != nil {
				log.Println(err)
			}
			panic(v)
		}
	}()

	if err := fn(sessionCtx); err != nil {
		log.Println(err)
		if rbErr := sessionCtx.Rollback(ctx); rbErr != nil {
			log.Println(err)
		}
		return err
	}

	if cmErr := sessionCtx.Commit(ctx); cmErr != nil {
		log.Println(err)
		return err
	}

	return nil
}
