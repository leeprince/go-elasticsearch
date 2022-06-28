// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/watcherstate"
)

// WatcherNodeStats type.
type WatcherNodeStats struct {
	CurrentWatches      []WatchRecordStats        `json:"current_watches,omitempty"`
	ExecutionThreadPool ExecutionThreadPool       `json:"execution_thread_pool"`
	NodeId              Id                        `json:"node_id"`
	QueuedWatches       []WatchRecordQueuedStats  `json:"queued_watches,omitempty"`
	WatchCount          int64                     `json:"watch_count"`
	WatcherState        watcherstate.WatcherState `json:"watcher_state"`
}

// WatcherNodeStatsBuilder holds WatcherNodeStats struct and provides a builder API.
type WatcherNodeStatsBuilder struct {
	v *WatcherNodeStats
}

// NewWatcherNodeStats provides a builder for the WatcherNodeStats struct.
func NewWatcherNodeStats() *WatcherNodeStatsBuilder {
	r := WatcherNodeStatsBuilder{
		&WatcherNodeStats{},
	}

	return &r
}

// Build finalize the chain and returns the WatcherNodeStats struct
func (rb *WatcherNodeStatsBuilder) Build() WatcherNodeStats {
	return *rb.v
}

// CurrentWatches set the CurrentWatches property for WatcherNodeStatsBuilder.
func (rb *WatcherNodeStatsBuilder) CurrentWatches(current_watches ...WatchRecordStats) *WatcherNodeStatsBuilder {
	rb.v.CurrentWatches = append(rb.v.CurrentWatches, current_watches...)
	return rb
}

// ExecutionThreadPool set the ExecutionThreadPool property for WatcherNodeStatsBuilder.
func (rb *WatcherNodeStatsBuilder) ExecutionThreadPool(executionthreadpool ExecutionThreadPool) *WatcherNodeStatsBuilder {
	rb.v.ExecutionThreadPool = executionthreadpool
	return rb
}

// NodeId set the NodeId property for WatcherNodeStatsBuilder.
func (rb *WatcherNodeStatsBuilder) NodeId(nodeid Id) *WatcherNodeStatsBuilder {
	rb.v.NodeId = nodeid
	return rb
}

// QueuedWatches set the QueuedWatches property for WatcherNodeStatsBuilder.
func (rb *WatcherNodeStatsBuilder) QueuedWatches(queued_watches ...WatchRecordQueuedStats) *WatcherNodeStatsBuilder {
	rb.v.QueuedWatches = append(rb.v.QueuedWatches, queued_watches...)
	return rb
}

// WatchCount set the WatchCount property for WatcherNodeStatsBuilder.
func (rb *WatcherNodeStatsBuilder) WatchCount(watchcount int64) *WatcherNodeStatsBuilder {
	rb.v.WatchCount = watchcount
	return rb
}

// WatcherState set the WatcherState property for WatcherNodeStatsBuilder.
func (rb *WatcherNodeStatsBuilder) WatcherState(watcherstate watcherstate.WatcherState) *WatcherNodeStatsBuilder {
	rb.v.WatcherState = watcherstate
	return rb
}
