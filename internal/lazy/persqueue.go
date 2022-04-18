package lazy

import (
	"context"
	"sync"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/database"
	intpq "github.com/ydb-platform/ydb-go-sdk/v3/internal/persqueue"
	"github.com/ydb-platform/ydb-go-sdk/v3/persqueue"
	"github.com/ydb-platform/ydb-go-sdk/v3/persqueue/config"
)

type LazyPersqueue struct {
	db      database.Connection
	options []config.Option
	once    sync.Once
	c       persqueue.Client
}

func Persqueue(db database.Connection, options []config.Option) *LazyPersqueue {
	return &LazyPersqueue{
		db:      db,
		options: options,
	}
}

func (l *LazyPersqueue) Client() persqueue.Client {
	l.once.Do(func() {
		l.c = intpq.New(l.db, l.options)
	})

	return l.c
}

func (l *LazyPersqueue) Close(ctx context.Context) error {
	l.once.Do(func() {})
	if l.c == nil {
		return nil
	}
	return l.c.Close(ctx)
}
