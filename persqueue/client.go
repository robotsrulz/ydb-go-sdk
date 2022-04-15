package persqueue

import (
	"context"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/scheme"
)

type Client interface {
	DescribeStream(context.Context, scheme.Path) (StreamInfo, error)
	DropStream(context.Context, scheme.Path) error
	CreateStream(context.Context, scheme.Path, StreamSettings, ...StreamOption) error
	AlterStream(context.Context, scheme.Path, StreamSettings, ...StreamOption) error
	AddReadRule(context.Context, scheme.Path, ReadRule) error
	RemoveReadRule(context.Context, scheme.Path, Consumer) error
}

type Consumer string

type StreamSettings struct {
	// How many partitions in topic. Must less than database limit. Default limit - 10.
	PartitionsCount int
	// How long data in partition should be stored. Must be greater than 0 and less than limit for this database.
	// Default limit - 36 hours.
	RetentionPeriod time.Duration
	// How long last written seqno for message group should be stored. Must be greater then retention_period_ms
	// and less then limit for this database.  Default limit - 16 days.
	MessageGroupSeqnoRetentionPeriod time.Duration
	// How many last written seqno for various message groups should be stored per partition. Must be less than limit
	// for this database.  Default limit - 6*10^6 values.
	MaxPartitionMessageGroupsSeqnoStored int
	// List of allowed codecs for stream writes.
	// Writes with codec not from this list are forbidden.
	SupportedCodecs []Codec
	// Max storage usage for each topic's partition. Must be less than database limit. Default limit - 130 GB.
	MaxPartitionStorageSize int
	// Partition write speed in bytes per second. Must be less than database limit. Default limit - 1 MB/s.
	MaxPartitionWriteSpeed int
	// Burst size for write in partition, in bytes. Must be less than database limit. Default limit - 1 MB.
	MaxPartitionWriteBurst int

	// Max format version that is allowed for writers.
	// Writes with greater format version are forbidden.
	SupportedFormat Format
	// Disallows client writes. Used for mirrored topics in federation.
	ClientWriteDisabled bool

	// User and server attributes of topic. Server attributes starts from "_" and will be validated by server.
	Attributes map[string]string // TODO: что сюда можно написать и зачем

}

// Message for read rules description.
type ReadRule struct {
	// For what consumer this read rule is. Must be valid not empty consumer name.
	// Is key for read rules. There could be only one read rule with corresponding consumer name.
	Consumer Consumer
	// All messages with smaller timestamp of write will be skipped.
	StartingMessageTimestamp time.Time
	// Flag that this consumer is important.
	Important bool
	// Max format version that is supported by this consumer.
	// supported_format on topic must not be greater.
	SupportedFormat Format
	// List of supported codecs by this consumer.
	// supported_codecs on topic must be contained inside this list.
	Codecs []Codec

	// Read rule version. Any non-negative integer.
	Version int

	// Client service type.
	ServiceType string // TODO: что это и зачем
}

// Message for remote mirror rule description.
type RemoteMirrorRule struct {
	// Source cluster endpoint in format server:port.
	Endpoint string
	// Source topic that we want to mirror.
	SourceStream string
	// Source consumer for reading source topic.
	Consumer Consumer
	// Credentials for reading source topic by source consumer.
	Credentials RemoteMirrorCredentials
	// All messages with smaller timestamp of write will be skipped.
	StartingMessageTimestamp time.Time
	// Database
	Database string
}

type RemoteMirrorCredentials interface {
	isRemoteMirrorCredentials()
}

func (OAuthTokenCredentials) isRemoteMirrorCredentials()
func (JWTCredentials) isRemoteMirrorCredentials()
func (IAMCredentials) isRemoteMirrorCredentials()

type OAuthTokenCredentials string

type JWTCredentials string // TODO: что сюда писать?

type IAMCredentials struct {
	Endpoint          string
	ServiceAccountKey string
}

type StreamInfo struct {
	scheme.Entry
	StreamSettings

	// List of consumer read rules for this topic.
	ReadRules []ReadRule
	// remote mirror rule for this topic.
	RemoteMirrorRule RemoteMirrorRule // TODO: хотим ли выставлять это?
}
