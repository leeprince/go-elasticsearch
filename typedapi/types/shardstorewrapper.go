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

// ShardStoreWrapper type.
type ShardStoreWrapper struct {
	Stores []ShardStore `json:"stores"`
}

// ShardStoreWrapperBuilder holds ShardStoreWrapper struct and provides a builder API.
type ShardStoreWrapperBuilder struct {
	v *ShardStoreWrapper
}

// NewShardStoreWrapper provides a builder for the ShardStoreWrapper struct.
func NewShardStoreWrapper() *ShardStoreWrapperBuilder {
	r := ShardStoreWrapperBuilder{
		&ShardStoreWrapper{},
	}

	return &r
}

// Build finalize the chain and returns the ShardStoreWrapper struct
func (rb *ShardStoreWrapperBuilder) Build() ShardStoreWrapper {
	return *rb.v
}

// Stores set the Stores property for ShardStoreWrapperBuilder.
func (rb *ShardStoreWrapperBuilder) Stores(stores ...ShardStore) *ShardStoreWrapperBuilder {
	rb.v.Stores = append(rb.v.Stores, stores...)
	return rb
}
