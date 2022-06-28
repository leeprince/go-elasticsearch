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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/shardroutingstate"
)

// NodeShard type.
type NodeShard struct {
	AllocationId   map[string]Id                       `json:"allocation_id,omitempty"`
	Index          IndexName                           `json:"index"`
	Node           *NodeName                           `json:"node,omitempty"`
	Primary        bool                                `json:"primary"`
	RecoverySource map[string]Id                       `json:"recovery_source,omitempty"`
	RelocatingNode NodeId                              `json:"relocating_node,omitempty"`
	Shard          int                                 `json:"shard"`
	State          shardroutingstate.ShardRoutingState `json:"state"`
	UnassignedInfo *UnassignedInformation              `json:"unassigned_info,omitempty"`
}

// NodeShardBuilder holds NodeShard struct and provides a builder API.
type NodeShardBuilder struct {
	v *NodeShard
}

// NewNodeShard provides a builder for the NodeShard struct.
func NewNodeShard() *NodeShardBuilder {
	r := NodeShardBuilder{
		&NodeShard{
			AllocationId:   make(map[string]Id, 0),
			RecoverySource: make(map[string]Id, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodeShard struct
func (rb *NodeShardBuilder) Build() NodeShard {
	return *rb.v
}

// AllocationId set the AllocationId property for NodeShardBuilder.
func (rb *NodeShardBuilder) AllocationId(value map[string]Id) *NodeShardBuilder {
	rb.v.AllocationId = value
	return rb
}

// Index set the Index property for NodeShardBuilder.
func (rb *NodeShardBuilder) Index(index IndexName) *NodeShardBuilder {
	rb.v.Index = index
	return rb
}

// Node set the Node property for NodeShardBuilder.
func (rb *NodeShardBuilder) Node(node NodeName) *NodeShardBuilder {
	rb.v.Node = &node
	return rb
}

// Primary set the Primary property for NodeShardBuilder.
func (rb *NodeShardBuilder) Primary(primary bool) *NodeShardBuilder {
	rb.v.Primary = primary
	return rb
}

// RecoverySource set the RecoverySource property for NodeShardBuilder.
func (rb *NodeShardBuilder) RecoverySource(value map[string]Id) *NodeShardBuilder {
	rb.v.RecoverySource = value
	return rb
}

// RelocatingNode set the RelocatingNode property for NodeShardBuilder.
func (rb *NodeShardBuilder) RelocatingNode(relocatingnode NodeId) *NodeShardBuilder {
	rb.v.RelocatingNode = relocatingnode
	return rb
}

// Shard set the Shard property for NodeShardBuilder.
func (rb *NodeShardBuilder) Shard(shard int) *NodeShardBuilder {
	rb.v.Shard = shard
	return rb
}

// State set the State property for NodeShardBuilder.
func (rb *NodeShardBuilder) State(state shardroutingstate.ShardRoutingState) *NodeShardBuilder {
	rb.v.State = state
	return rb
}

// UnassignedInfo set the UnassignedInfo property for NodeShardBuilder.
func (rb *NodeShardBuilder) UnassignedInfo(unassignedinfo UnassignedInformation) *NodeShardBuilder {
	rb.v.UnassignedInfo = &unassignedinfo
	return rb
}
