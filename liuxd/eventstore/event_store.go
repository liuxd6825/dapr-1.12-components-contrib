package eventstore

import (
	"context"
	"github.com/dapr/components-contrib/liuxd/common"
	"github.com/dapr/components-contrib/liuxd/eventstore/dto"
	pubsub_adapter "github.com/dapr/dapr/pkg/runtime/pubsub"
	"github.com/dapr/kit/logger"
)

type Session interface {
	UseTransaction(context.Context, SessionFunc) error
}

type SessionFunc func(ctx context.Context) error

type GetPubsubAdapter func() pubsub_adapter.Adapter

type Options struct {
	Metadata       common.Metadata
	PubsubAdapter  GetPubsubAdapter
	EventRepos     interface{}
	SnapshotRepos  interface{}
	AggregateRepos interface{}
	RelationRepos  interface{}
	MessageRepos   interface{}
	SnapshotCount  uint64
	Session        Session
}

// EventStore 领域事件存储接口
type EventStore interface {
	// Init 初始化
	Init(ctx context.Context, opts *Options) error

	GetLogger() logger.Logger

	// LoadEvent 加载事件
	LoadEvent(ctx context.Context, req *dto.LoadEventRequest) (*dto.LoadResponse, error)

	/*	// CreateEvent 创建聚合事件
		CreateEvent(ctx context.Context, req *dto.CreateEventRequest) (*dto.CreateEventResponse, error)

		// DeleteEvent 删除聚合事件
		DeleteEvent(ctx context.Context, req *dto.DeleteEventRequest) (*dto.DeleteEventResponse, error)
	*/

	// ApplyEvent 应用事件
	ApplyEvent(ctx context.Context, req *dto.ApplyEventsRequest) (*dto.ApplyEventsResponse, error)

	// DeleteAggregate 销毁聚合根
	DeleteAggregate(ctx context.Context, req *dto.DeleteAggregateRequest) (*dto.DeleteAggregateResponse, error)

	// FindEvents 查找事件
	FindEvents(ctx context.Context, req *dto.FindEventsRequest) (*dto.FindEventsResponse, error)

	// SaveSnapshot 保存镜像对象
	SaveSnapshot(ctx context.Context, req *dto.SaveSnapshotRequest) (*dto.SaveSnapshotResponse, error)

	// FindRelations 获取聚合根关系
	FindRelations(ctx context.Context, req *dto.FindRelationsRequest) (*dto.FindRelationsResponse, error)

	// RepublishMessage 重新发送未发出的消息
	RepublishMessage(ctx context.Context, req *dto.RepublishMessageRequest) (*dto.RepublishMessageResponse, error)

	Commit(ctx context.Context, req *dto.CommitRequest) (*dto.CommitResponse, error)

	Rollback(ctx context.Context, req *dto.RollbackRequest) (*dto.RollbackResponse, error)
}
