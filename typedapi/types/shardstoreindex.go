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

// ShardStoreIndex type.
type ShardStoreIndex struct {
	Aliases []Name          `json:"aliases,omitempty"`
	Filter  *QueryContainer `json:"filter,omitempty"`
}

// ShardStoreIndexBuilder holds ShardStoreIndex struct and provides a builder API.
type ShardStoreIndexBuilder struct {
	v *ShardStoreIndex
}

// NewShardStoreIndex provides a builder for the ShardStoreIndex struct.
func NewShardStoreIndex() *ShardStoreIndexBuilder {
	r := ShardStoreIndexBuilder{
		&ShardStoreIndex{},
	}

	return &r
}

// Build finalize the chain and returns the ShardStoreIndex struct
func (rb *ShardStoreIndexBuilder) Build() ShardStoreIndex {
	return *rb.v
}

// Aliases set the Aliases property for ShardStoreIndexBuilder.
func (rb *ShardStoreIndexBuilder) Aliases(aliases ...Name) *ShardStoreIndexBuilder {
	rb.v.Aliases = append(rb.v.Aliases, aliases...)
	return rb
}

// Filter set the Filter property for ShardStoreIndexBuilder.
func (rb *ShardStoreIndexBuilder) Filter(filter QueryContainer) *ShardStoreIndexBuilder {
	rb.v.Filter = &filter
	return rb
}
