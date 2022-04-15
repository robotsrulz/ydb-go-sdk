package persqueue

import (
	"io"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/scheme"
)

type ReadStream interface {
	Send(ReadRequest) error
	Recv() (ReadStreamMessage, error)
	CloseSend() error
	Close() error
}

type WriteStream interface {
	Send(WriteRequest) error
	Recv() (WriteStreamMessage, error)
	CloseSend() error
	Close() error
}

// Common types for streaming

type Assign struct {
	Stream   string
	Cluster  string
	AssignID uint
}

type CommitCookie struct {
	AssignID        uint
	PartitionCookie uint
}

type ReadOffsetRange struct {
	AssignID    uint
	StartOffset uint
	EndOffset   uint
}

type ReadSelector struct {
	Stream             scheme.Path
	PartitionGroups    []int
	SkipMessagesBefore time.Time
}

type SizeReader interface {
	io.Reader
	Len() int
}

type ReadPartitionData struct {
	Stream    string
	Cluster   string
	Partition string

	CommitCookie CommitCookie

	Batches []ReadBatch
}

type ReadBatch struct {
	MessageGroupID string
	ExtraFields    map[string]string
	WrittenAt      time.Time
	IP             string

	MessagesData []ReadMessageData
}

type ReadMessageData struct {
	Offset    uint
	SeqNo     uint
	CreatedAt time.Time

	Data         SizeReader
	PartitionKey string
	ExplicitHash string
}

type WriteRecord struct {
	SeqNo            int
	CreatedAt        time.Time
	SentAt           time.Time
	Codec            Codec
	UncompressedSize int
	CompressedData   SizeReader
}

type WriteResult struct {
	SeqNo          int
	Offset         uint
	AlreadyWritten bool
}

type WriteStatistics struct {
	Persist              time.Duration
	QueuedInPartition    time.Duration
	ThrottledOnPartition time.Duration
	ThrottledOnStream    time.Duration
}

// ReadRequest is group of types sutable for send to ReadStream.Send
type ReadRequest interface {
	isReadRequest()
}

type ReadInit struct {
	readRequestMark

	ReadSettings []ReadSelector
	Consumer     string

	OnlyOriginal       bool
	MaxLag             time.Duration
	SkipMessagesBefore time.Time

	MaxBlockFormatVersion int
	MaxMetaCacheSize      int
}

type ReadStart struct {
	readRequestMark
	Assign

	ReadOffset   uint
	CommitOffset uint

	VerifyReadOffset bool
}

type ReadReleased struct {
	readRequestMark
	Assign
}

type ReadCommit struct {
	readRequestMark
	Cookies []CommitCookie
	Offsets []ReadOffsetRange
}

// ReadStreamMessage is group of message types which can be received from ReadStream.Recv
type ReadStreamMessage interface {
	isReadStreamMessage()
}

type ReadInitMessage struct {
	readStreamMessageMark
	SessionID                  string
	StreamsBlockFormatVersions map[string]int

	MaxMetaCacheSize int
}

type ReadAssigned struct {
	readStreamMessageMark
	Assign
	Partition int

	ReadOffset uint
	EndOffset  uint
}

type ReadRelease struct {
	readStreamMessageMark
	Assign
	Partition int

	Force             bool
	MaxCommitedOffset uint
}

type ReadCommitted struct {
	readStreamMessageMark
	Cookies []CommitCookie
	Offsets []ReadOffsetRange
}

type ReadData struct {
	readStreamMessageMark
	PartitionsData []ReadPartitionData
}

type ReadPartitionStatus struct {
	readStreamMessageMark
	Assign
	Partition int

	CommitedOffset uint
	EndOffset      uint
	WriteWatermark time.Time
}

// WriteRequest is group of types sutable for WriteStream.Send
type WriteRequest interface {
	isWriteRequest()
}

type WriteInit struct {
	writeRequestMark
	Stream            string
	MessageGroupID    string
	PartitionGroup    uint
	ExtraFields       map[string]string
	ConnectionAttempt int

	IdleTimeout time.Duration

	PreferredCluster string

	SessionID string // ???
}

type WriteMessages struct {
	writeRequestMark
	Records []WriteRecord
}

type WriteUpdateToken struct { // сам токен взять из dialer
	writeRequestMark
}

// WriteStreamMessage is group of message types which can be received from WriteStream.Recv
type WriteStreamMessage interface {
	isWriteStreamMessage()
}

type WriteInitMessage struct {
	writeStreamMessageMark
	LastSeqNo int
	SessionID string
	Stream    string
	Cluster   string
	Partition int
	Codecs    []Codec

	BlockFormatVersion int
	BlockSizeLimit     int

	FlushWindowLimit int
}

type WriteResponse struct {
	writeStreamMessageMark

	Results    []WriteResult
	Partition  int
	Statistics WriteStatistics // maybe should be lazy
}

type WriteTokenUpdated struct {
	writeStreamMessageMark
}

// Technical helper types

type readRequestMark struct{}

func (readRequestMark) isReadRequest() {}

type readStreamMessageMark struct{}

func (readStreamMessageMark) isReadStreamMessage() {}

type writeRequestMark struct{}

func (writeRequestMark) isWriteRequest() {}

type writeStreamMessageMark struct{}

func (writeStreamMessageMark) isWriteStreamMessage() {}
