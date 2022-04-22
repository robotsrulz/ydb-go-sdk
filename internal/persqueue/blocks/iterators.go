package blocks

import (
	"github.com/ydb-platform/ydb-go-sdk/v3/persqueue"
)

type MessageIterator interface {
	NextMessage() (data persqueue.EncodeReader, end bool)
}

type BlockIteratot interface {
	NextBlock() (block *Block, end bool) // Returned block can not be used after next call
}
