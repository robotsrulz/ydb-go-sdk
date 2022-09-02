package ydb

import (
	"context"
	"sync/atomic"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xerrors"
)

var nextID = uint64(0)

func (c *connection) With(ctx context.Context, opts ...Option) (Connection, error) {
	id := atomic.AddUint64(&nextID, 1)

	child, err := open(
		ctx,
		append(
			append(
				c.opts,
				WithBalancer(
					c.config.Balancer(),
				),
				withOnClose(func(child *connection) {
					c.childrenMtx.Lock()
					defer c.childrenMtx.Unlock()

					delete(c.children, id)
				}),
				withConnPool(c.pool),
			),
			opts...,
		)...,
	)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	c.childrenMtx.Lock()
	defer c.childrenMtx.Unlock()

	c.children[id] = child

	return child, nil
}
