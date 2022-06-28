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

// ClusterStateQueue type.
type ClusterStateQueue struct {
	Committed *int64 `json:"committed,omitempty"`
	Pending   *int64 `json:"pending,omitempty"`
	Total     *int64 `json:"total,omitempty"`
}

// ClusterStateQueueBuilder holds ClusterStateQueue struct and provides a builder API.
type ClusterStateQueueBuilder struct {
	v *ClusterStateQueue
}

// NewClusterStateQueue provides a builder for the ClusterStateQueue struct.
func NewClusterStateQueue() *ClusterStateQueueBuilder {
	r := ClusterStateQueueBuilder{
		&ClusterStateQueue{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterStateQueue struct
func (rb *ClusterStateQueueBuilder) Build() ClusterStateQueue {
	return *rb.v
}

// Committed set the Committed property for ClusterStateQueueBuilder.
func (rb *ClusterStateQueueBuilder) Committed(committed int64) *ClusterStateQueueBuilder {
	rb.v.Committed = &committed
	return rb
}

// Pending set the Pending property for ClusterStateQueueBuilder.
func (rb *ClusterStateQueueBuilder) Pending(pending int64) *ClusterStateQueueBuilder {
	rb.v.Pending = &pending
	return rb
}

// Total set the Total property for ClusterStateQueueBuilder.
func (rb *ClusterStateQueueBuilder) Total(total int64) *ClusterStateQueueBuilder {
	rb.v.Total = &total
	return rb
}
