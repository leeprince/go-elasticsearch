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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/shardstoreallocation"
)

// ShardStore type.
type ShardStore struct {
	Allocation       shardstoreallocation.ShardStoreAllocation `json:"allocation"`
	AllocationId     Id                                        `json:"allocation_id"`
	Attributes       map[string]interface{}                    `json:"attributes"`
	Id               Id                                        `json:"id"`
	LegacyVersion    VersionNumber                             `json:"legacy_version"`
	Name             Name                                      `json:"name"`
	StoreException   ShardStoreException                       `json:"store_exception"`
	TransportAddress TransportAddress                          `json:"transport_address"`
}

// ShardStoreBuilder holds ShardStore struct and provides a builder API.
type ShardStoreBuilder struct {
	v *ShardStore
}

// NewShardStore provides a builder for the ShardStore struct.
func NewShardStore() *ShardStoreBuilder {
	r := ShardStoreBuilder{
		&ShardStore{
			Attributes: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ShardStore struct
func (rb *ShardStoreBuilder) Build() ShardStore {
	return *rb.v
}

// Allocation set the Allocation property for ShardStoreBuilder.
func (rb *ShardStoreBuilder) Allocation(allocation shardstoreallocation.ShardStoreAllocation) *ShardStoreBuilder {
	rb.v.Allocation = allocation
	return rb
}

// AllocationId set the AllocationId property for ShardStoreBuilder.
func (rb *ShardStoreBuilder) AllocationId(allocationid Id) *ShardStoreBuilder {
	rb.v.AllocationId = allocationid
	return rb
}

// Attributes set the Attributes property for ShardStoreBuilder.
func (rb *ShardStoreBuilder) Attributes(value map[string]interface{}) *ShardStoreBuilder {
	rb.v.Attributes = value
	return rb
}

// Id set the Id property for ShardStoreBuilder.
func (rb *ShardStoreBuilder) Id(id Id) *ShardStoreBuilder {
	rb.v.Id = id
	return rb
}

// LegacyVersion set the LegacyVersion property for ShardStoreBuilder.
func (rb *ShardStoreBuilder) LegacyVersion(legacyversion VersionNumber) *ShardStoreBuilder {
	rb.v.LegacyVersion = legacyversion
	return rb
}

// Name set the Name property for ShardStoreBuilder.
func (rb *ShardStoreBuilder) Name(name Name) *ShardStoreBuilder {
	rb.v.Name = name
	return rb
}

// StoreException set the StoreException property for ShardStoreBuilder.
func (rb *ShardStoreBuilder) StoreException(storeexception ShardStoreException) *ShardStoreBuilder {
	rb.v.StoreException = storeexception
	return rb
}

// TransportAddress set the TransportAddress property for ShardStoreBuilder.
func (rb *ShardStoreBuilder) TransportAddress(transportaddress TransportAddress) *ShardStoreBuilder {
	rb.v.TransportAddress = transportaddress
	return rb
}
