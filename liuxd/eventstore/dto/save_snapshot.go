package dto

import "github.com/dapr/components-contrib/liuxd/eventstore/domain/model"

type SaveSnapshotRequest struct {
	TenantId         string                 `json:"tenantId"`
	AggregateId      string                 `json:"aggregateId"`
	AggregateType    string                 `json:"aggregateType"`
	AggregateData    map[string]interface{} `json:"aggregateData"`
	AggregateVersion string                 `json:"aggregateVersion"`
	SequenceNumber   uint64                 `json:"sequenceNumber"`
	Metadata         model.Metadata         `json:"metadata"`
}

type SaveSnapshotResponse struct {
	Headers *ResponseHeaders `json:"headers"`
}
