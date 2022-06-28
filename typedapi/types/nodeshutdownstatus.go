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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/shutdownstatus"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/shutdowntype"
)

// NodeShutdownStatus type.
type NodeShutdownStatus struct {
	NodeId                NodeId                        `json:"node_id"`
	PersistentTasks       PersistentTaskStatus          `json:"persistent_tasks"`
	Plugins               PluginsStatus                 `json:"plugins"`
	Reason                string                        `json:"reason"`
	ShardMigration        ShardMigrationStatus          `json:"shard_migration"`
	ShutdownStartedmillis EpochMillis                   `json:"shutdown_startedmillis"`
	Status                shutdownstatus.ShutdownStatus `json:"status"`
	Type                  shutdowntype.ShutdownType     `json:"type"`
}

// NodeShutdownStatusBuilder holds NodeShutdownStatus struct and provides a builder API.
type NodeShutdownStatusBuilder struct {
	v *NodeShutdownStatus
}

// NewNodeShutdownStatus provides a builder for the NodeShutdownStatus struct.
func NewNodeShutdownStatus() *NodeShutdownStatusBuilder {
	r := NodeShutdownStatusBuilder{
		&NodeShutdownStatus{},
	}

	return &r
}

// Build finalize the chain and returns the NodeShutdownStatus struct
func (rb *NodeShutdownStatusBuilder) Build() NodeShutdownStatus {
	return *rb.v
}

// NodeId set the NodeId property for NodeShutdownStatusBuilder.
func (rb *NodeShutdownStatusBuilder) NodeId(nodeid NodeId) *NodeShutdownStatusBuilder {
	rb.v.NodeId = nodeid
	return rb
}

// PersistentTasks set the PersistentTasks property for NodeShutdownStatusBuilder.
func (rb *NodeShutdownStatusBuilder) PersistentTasks(persistenttasks PersistentTaskStatus) *NodeShutdownStatusBuilder {
	rb.v.PersistentTasks = persistenttasks
	return rb
}

// Plugins set the Plugins property for NodeShutdownStatusBuilder.
func (rb *NodeShutdownStatusBuilder) Plugins(plugins PluginsStatus) *NodeShutdownStatusBuilder {
	rb.v.Plugins = plugins
	return rb
}

// Reason set the Reason property for NodeShutdownStatusBuilder.
func (rb *NodeShutdownStatusBuilder) Reason(reason string) *NodeShutdownStatusBuilder {
	rb.v.Reason = reason
	return rb
}

// ShardMigration set the ShardMigration property for NodeShutdownStatusBuilder.
func (rb *NodeShutdownStatusBuilder) ShardMigration(shardmigration ShardMigrationStatus) *NodeShutdownStatusBuilder {
	rb.v.ShardMigration = shardmigration
	return rb
}

// ShutdownStartedmillis set the ShutdownStartedmillis property for NodeShutdownStatusBuilder.
func (rb *NodeShutdownStatusBuilder) ShutdownStartedmillis(shutdownstartedmillis EpochMillis) *NodeShutdownStatusBuilder {
	rb.v.ShutdownStartedmillis = shutdownstartedmillis
	return rb
}

// Status set the Status property for NodeShutdownStatusBuilder.
func (rb *NodeShutdownStatusBuilder) Status(status shutdownstatus.ShutdownStatus) *NodeShutdownStatusBuilder {
	rb.v.Status = status
	return rb
}

// Type set the Type property for NodeShutdownStatusBuilder.
func (rb *NodeShutdownStatusBuilder) Type_(type_ shutdowntype.ShutdownType) *NodeShutdownStatusBuilder {
	rb.v.Type = type_
	return rb
}
