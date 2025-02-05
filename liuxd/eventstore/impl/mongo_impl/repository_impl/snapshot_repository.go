package repository_impl

import (
	"context"
	"github.com/dapr/components-contrib/liuxd/eventstore/domain/model"
	"github.com/dapr/components-contrib/liuxd/eventstore/domain/repository"
	"github.com/dapr/components-contrib/liuxd/eventstore/impl/mongo_impl/db"
	"go.mongodb.org/mongo-driver/bson"
)

type snapshotRepository struct {
	dao *dao[*model.Snapshot]
}

func NewSnapshotRepository(mongodb *db.MongoDbConfig, collName string) repository.SnapshotRepository {
	return &snapshotRepository{
		dao: NewDao[*model.Snapshot](mongodb, collName),
	}
}

func (r *snapshotRepository) Create(ctx context.Context, tenantId string, v *model.Snapshot) error {
	return r.dao.Insert(ctx, v)
}

func (r *snapshotRepository) DeleteById(ctx context.Context, tenantId string, id string) error {
	return r.dao.DeleteById(ctx, tenantId, id)
}

func (r *snapshotRepository) DeleteByAggregateId(ctx context.Context, tenantId, aggregateId string) error {
	filter := bson.M{
		TenantIdField:    tenantId,
		AggregateIdField: aggregateId,
	}
	return r.dao.deleteByFilter(ctx, tenantId, filter)
}

func (r *snapshotRepository) Update(ctx context.Context, tenantId string, v *model.Snapshot) error {
	return r.dao.Update(ctx, v)
}

func (r *snapshotRepository) FindById(ctx context.Context, tenantId string, id string) (*model.Snapshot, bool, error) {
	return r.dao.FindById(ctx, tenantId, id)
}

func (r *snapshotRepository) FindByAggregateId(ctx context.Context, tenantId string, aggregateId string) ([]*model.Snapshot, bool, error) {
	filter := bson.M{
		TenantIdField:    tenantId,
		AggregateIdField: aggregateId,
	}
	return r.dao.findList(ctx, tenantId, filter, nil)
}

func (r *snapshotRepository) DeleteByAggregateType(ctx context.Context, tenantId, aggregateType string) error {
	filter := bson.M{
		TenantIdField:      tenantId,
		AggregateTypeField: aggregateType,
	}
	return r.dao.deleteByFilter(ctx, tenantId, filter, msgOptions)
}

func (r *snapshotRepository) FindByMaxSequenceNumber(ctx context.Context, tenantId string, aggregateId string, aggregateType string) (*model.Snapshot, bool, error) {
	filter := bson.M{
		TenantIdField:      tenantId,
		AggregateIdField:   aggregateId,
		AggregateTypeField: aggregateType,
	}
	options := NewOptions().SetSort(bson.D{{SequenceNumberField, -1}})
	return r.dao.findOne(ctx, tenantId, filter, options)
}
