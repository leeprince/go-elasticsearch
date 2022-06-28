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

// AllocationStore type.
type AllocationStore struct {
	AllocationId        string `json:"allocation_id"`
	Found               bool   `json:"found"`
	InSync              bool   `json:"in_sync"`
	MatchingSizeInBytes int64  `json:"matching_size_in_bytes"`
	MatchingSyncId      bool   `json:"matching_sync_id"`
	StoreException      string `json:"store_exception"`
}

// AllocationStoreBuilder holds AllocationStore struct and provides a builder API.
type AllocationStoreBuilder struct {
	v *AllocationStore
}

// NewAllocationStore provides a builder for the AllocationStore struct.
func NewAllocationStore() *AllocationStoreBuilder {
	r := AllocationStoreBuilder{
		&AllocationStore{},
	}

	return &r
}

// Build finalize the chain and returns the AllocationStore struct
func (rb *AllocationStoreBuilder) Build() AllocationStore {
	return *rb.v
}

// AllocationId set the AllocationId property for AllocationStoreBuilder.
func (rb *AllocationStoreBuilder) AllocationId(allocationid string) *AllocationStoreBuilder {
	rb.v.AllocationId = allocationid
	return rb
}

// Found set the Found property for AllocationStoreBuilder.
func (rb *AllocationStoreBuilder) Found(found bool) *AllocationStoreBuilder {
	rb.v.Found = found
	return rb
}

// InSync set the InSync property for AllocationStoreBuilder.
func (rb *AllocationStoreBuilder) InSync(insync bool) *AllocationStoreBuilder {
	rb.v.InSync = insync
	return rb
}

// MatchingSizeInBytes set the MatchingSizeInBytes property for AllocationStoreBuilder.
func (rb *AllocationStoreBuilder) MatchingSizeInBytes(matchingsizeinbytes int64) *AllocationStoreBuilder {
	rb.v.MatchingSizeInBytes = matchingsizeinbytes
	return rb
}

// MatchingSyncId set the MatchingSyncId property for AllocationStoreBuilder.
func (rb *AllocationStoreBuilder) MatchingSyncId(matchingsyncid bool) *AllocationStoreBuilder {
	rb.v.MatchingSyncId = matchingsyncid
	return rb
}

// StoreException set the StoreException property for AllocationStoreBuilder.
func (rb *AllocationStoreBuilder) StoreException(storeexception string) *AllocationStoreBuilder {
	rb.v.StoreException = storeexception
	return rb
}
