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

// IndexRouting type.
type IndexRouting struct {
	Allocation *IndexRoutingAllocation `json:"allocation,omitempty"`
	Rebalance  *IndexRoutingRebalance  `json:"rebalance,omitempty"`
}

// IndexRoutingBuilder holds IndexRouting struct and provides a builder API.
type IndexRoutingBuilder struct {
	v *IndexRouting
}

// NewIndexRouting provides a builder for the IndexRouting struct.
func NewIndexRouting() *IndexRoutingBuilder {
	r := IndexRoutingBuilder{
		&IndexRouting{},
	}

	return &r
}

// Build finalize the chain and returns the IndexRouting struct
func (rb *IndexRoutingBuilder) Build() IndexRouting {
	return *rb.v
}

// Allocation set the Allocation property for IndexRoutingBuilder.
func (rb *IndexRoutingBuilder) Allocation(allocation IndexRoutingAllocation) *IndexRoutingBuilder {
	rb.v.Allocation = &allocation
	return rb
}

// Rebalance set the Rebalance property for IndexRoutingBuilder.
func (rb *IndexRoutingBuilder) Rebalance(rebalance IndexRoutingRebalance) *IndexRoutingBuilder {
	rb.v.Rebalance = &rebalance
	return rb
}
