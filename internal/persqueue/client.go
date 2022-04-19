package persqueue

import (
	"context"

	"github.com/ydb-platform/ydb-go-genproto/Ydb_PersQueue_V1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	pqproto "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_PersQueue_V1"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/operation"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xerrors"
	"github.com/ydb-platform/ydb-go-sdk/v3/persqueue"
	"github.com/ydb-platform/ydb-go-sdk/v3/persqueue/config"
	"github.com/ydb-platform/ydb-go-sdk/v3/retry"
	"github.com/ydb-platform/ydb-go-sdk/v3/scheme"
)

var _ persqueue.Client = &Client{}

type Client struct {
	config  config.Config
	service Ydb_PersQueue_V1.PersQueueServiceClient
}

func New(cc grpc.ClientConnInterface, options []config.Option) *Client {
	return &Client{
		config:  config.New(options...),
		service: Ydb_PersQueue_V1.NewPersQueueServiceClient(cc),
	}
}

func (c *Client) Close(_ context.Context) error {
	return nil
}

func (c *Client) DescribeStream(ctx context.Context, topic scheme.Path) (info persqueue.StreamInfo, err error) {
	err = retry.Retry(ctx, func(ctx context.Context) (err error) {
		info, err = c.describeStream(ctx, topic)
		return xerrors.WithStackTrace(err)
	}, retry.WithIdempotent(true))
	return info, err
}

func (c *Client) DropStream(ctx context.Context, topic scheme.Path) error {
	return retry.Retry(ctx, func(ctx context.Context) error {
		return c.dropStream(ctx, topic)
	})
}

func (c *Client) CreateStream(ctx context.Context, topic scheme.Path,
	settings persqueue.StreamSettings, opts ...persqueue.StreamOption,
) error {
	return retry.Retry(ctx, func(ctx context.Context) error {
		return c.createStream(ctx, topic, settings, opts...)
	})
}

func (c *Client) AlterStream(ctx context.Context, topic scheme.Path,
	settings persqueue.StreamSettings, opts ...persqueue.StreamOption,
) error {
	return retry.Retry(ctx, func(ctx context.Context) error {
		return c.alterStream(ctx, topic, settings, opts...)
	})
}

func (c *Client) AddReadRule(ctx context.Context, topic scheme.Path, rule persqueue.ReadRule) error {
	return retry.Retry(ctx, func(ctx context.Context) error {
		return c.addReadRule(ctx, topic, rule)
	})
}

func (c *Client) RemoveReadRule(ctx context.Context, topic scheme.Path, consumer persqueue.Consumer) error {
	return retry.Retry(ctx, func(ctx context.Context) error {
		return c.removeReadRule(ctx, topic, consumer)
	})
}

func (c *Client) describeStream(ctx context.Context, topic scheme.Path) (persqueue.StreamInfo, error) {
	var result persqueue.StreamInfo

	response, err := c.service.DescribeTopic(ctx,
		&pqproto.DescribeTopicRequest{
			Path:            string(topic),
			OperationParams: operation.Sync(ctx, c.config),
		},
	)
	if err != nil {
		return result, xerrors.WithStackTrace(err)
	}

	var opResult pqproto.DescribeTopicResult
	err = proto.Unmarshal(response.GetOperation().GetResult().GetValue(), &opResult)
	if err != nil {
		return result, xerrors.WithStackTrace(err)
	}
	result.Entry.From(opResult.Self)
	result.StreamSettings.From(opResult.Settings)

	result.ReadRules = make([]persqueue.ReadRule, len(opResult.Settings.ReadRules))
	for i := range result.ReadRules {
		result.ReadRules[i].From(opResult.Settings.ReadRules[i])
	}
	result.RemoteMirrorRule.From(opResult.Settings.RemoteMirrorRule)

	return result, nil
}

func (c *Client) dropStream(ctx context.Context, topic scheme.Path) error {
	_, err := c.service.DropTopic(ctx,
		&pqproto.DropTopicRequest{
			Path:            string(topic),
			OperationParams: operation.Sync(ctx, c.config),
		},
	)
	return xerrors.WithStackTrace(err)
}

func (c *Client) createStream(ctx context.Context, topic scheme.Path,
	settings persqueue.StreamSettings, opts ...persqueue.StreamOption,
) error {
	_, err := c.service.CreateTopic(ctx,
		&pqproto.CreateTopicRequest{
			Path:            string(topic),
			OperationParams: operation.Sync(ctx, c.config),
			Settings:        encodeTopicSettings(settings, opts...),
		},
	)
	return xerrors.WithStackTrace(err)
}

func (c *Client) alterStream(ctx context.Context, topic scheme.Path,
	settings persqueue.StreamSettings, opts ...persqueue.StreamOption,
) error {
	_, err := c.service.AlterTopic(ctx,
		&pqproto.AlterTopicRequest{
			Path:            string(topic),
			OperationParams: operation.Sync(ctx, c.config),
			Settings:        encodeTopicSettings(settings, opts...),
		},
	)
	return xerrors.WithStackTrace(err)
}

func (c *Client) addReadRule(ctx context.Context, topic scheme.Path, rule persqueue.ReadRule) error {
	_, err := c.service.AddReadRule(ctx,
		&pqproto.AddReadRuleRequest{
			Path:            string(topic),
			OperationParams: operation.Sync(ctx, c.config),
			ReadRule:        encodeReadRule(rule),
		},
	)
	return xerrors.WithStackTrace(err)
}

func (c *Client) removeReadRule(ctx context.Context, topic scheme.Path, consumer persqueue.Consumer) error {
	_, err := c.service.RemoveReadRule(ctx,
		&pqproto.RemoveReadRuleRequest{
			Path:            string(topic),
			OperationParams: operation.Sync(ctx, c.config),
			ConsumerName:    string(consumer),
		},
	)
	return xerrors.WithStackTrace(err)
}
