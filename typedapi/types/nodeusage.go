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

// NodeUsage type.
type NodeUsage struct {
	Aggregations map[string]interface{} `json:"aggregations"`
	RestActions  map[string]int         `json:"rest_actions"`
	Since        EpochMillis            `json:"since"`
	Timestamp    EpochMillis            `json:"timestamp"`
}

// NodeUsageBuilder holds NodeUsage struct and provides a builder API.
type NodeUsageBuilder struct {
	v *NodeUsage
}

// NewNodeUsage provides a builder for the NodeUsage struct.
func NewNodeUsage() *NodeUsageBuilder {
	r := NodeUsageBuilder{
		&NodeUsage{
			Aggregations: make(map[string]interface{}, 0),
			RestActions:  make(map[string]int, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodeUsage struct
func (rb *NodeUsageBuilder) Build() NodeUsage {
	return *rb.v
}

// Aggregations set the Aggregations property for NodeUsageBuilder.
func (rb *NodeUsageBuilder) Aggregations(value map[string]interface{}) *NodeUsageBuilder {
	rb.v.Aggregations = value
	return rb
}

// RestActions set the RestActions property for NodeUsageBuilder.
func (rb *NodeUsageBuilder) RestActions(value map[string]int) *NodeUsageBuilder {
	rb.v.RestActions = value
	return rb
}

// Since set the Since property for NodeUsageBuilder.
func (rb *NodeUsageBuilder) Since(since EpochMillis) *NodeUsageBuilder {
	rb.v.Since = since
	return rb
}

// Timestamp set the Timestamp property for NodeUsageBuilder.
func (rb *NodeUsageBuilder) Timestamp(timestamp EpochMillis) *NodeUsageBuilder {
	rb.v.Timestamp = timestamp
	return rb
}
