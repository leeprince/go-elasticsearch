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

// ClusterIndicesShardsIndex type.
type ClusterIndicesShardsIndex struct {
	// Primaries Contains statistics about the number of primary shards assigned to selected
	// nodes.
	Primaries ClusterShardMetrics `json:"primaries"`
	// Replication Contains statistics about the number of replication shards assigned to
	// selected nodes.
	Replication ClusterShardMetrics `json:"replication"`
	// Shards Contains statistics about the number of shards assigned to selected nodes.
	Shards ClusterShardMetrics `json:"shards"`
}

// ClusterIndicesShardsIndexBuilder holds ClusterIndicesShardsIndex struct and provides a builder API.
type ClusterIndicesShardsIndexBuilder struct {
	v *ClusterIndicesShardsIndex
}

// NewClusterIndicesShardsIndex provides a builder for the ClusterIndicesShardsIndex struct.
func NewClusterIndicesShardsIndex() *ClusterIndicesShardsIndexBuilder {
	r := ClusterIndicesShardsIndexBuilder{
		&ClusterIndicesShardsIndex{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterIndicesShardsIndex struct
func (rb *ClusterIndicesShardsIndexBuilder) Build() ClusterIndicesShardsIndex {
	return *rb.v
}

// Primaries Contains statistics about the number of primary shards assigned to selected
// nodes.
func (rb *ClusterIndicesShardsIndexBuilder) Primaries(primaries ClusterShardMetrics) *ClusterIndicesShardsIndexBuilder {
	rb.v.Primaries = primaries
	return rb
}

// Replication Contains statistics about the number of replication shards assigned to
// selected nodes.
func (rb *ClusterIndicesShardsIndexBuilder) Replication(replication ClusterShardMetrics) *ClusterIndicesShardsIndexBuilder {
	rb.v.Replication = replication
	return rb
}

// Shards Contains statistics about the number of shards assigned to selected nodes.
func (rb *ClusterIndicesShardsIndexBuilder) Shards(shards ClusterShardMetrics) *ClusterIndicesShardsIndexBuilder {
	rb.v.Shards = shards
	return rb
}
