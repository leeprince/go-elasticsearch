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

// SearchableSnapshots type.
type SearchableSnapshots struct {
	Available               bool `json:"available"`
	Enabled                 bool `json:"enabled"`
	FullCopyIndicesCount    *int `json:"full_copy_indices_count,omitempty"`
	IndicesCount            int  `json:"indices_count"`
	SharedCacheIndicesCount *int `json:"shared_cache_indices_count,omitempty"`
}

// SearchableSnapshotsBuilder holds SearchableSnapshots struct and provides a builder API.
type SearchableSnapshotsBuilder struct {
	v *SearchableSnapshots
}

// NewSearchableSnapshots provides a builder for the SearchableSnapshots struct.
func NewSearchableSnapshots() *SearchableSnapshotsBuilder {
	r := SearchableSnapshotsBuilder{
		&SearchableSnapshots{},
	}

	return &r
}

// Build finalize the chain and returns the SearchableSnapshots struct
func (rb *SearchableSnapshotsBuilder) Build() SearchableSnapshots {
	return *rb.v
}

// Available set the Available property for SearchableSnapshotsBuilder.
func (rb *SearchableSnapshotsBuilder) Available(available bool) *SearchableSnapshotsBuilder {
	rb.v.Available = available
	return rb
}

// Enabled set the Enabled property for SearchableSnapshotsBuilder.
func (rb *SearchableSnapshotsBuilder) Enabled(enabled bool) *SearchableSnapshotsBuilder {
	rb.v.Enabled = enabled
	return rb
}

// FullCopyIndicesCount set the FullCopyIndicesCount property for SearchableSnapshotsBuilder.
func (rb *SearchableSnapshotsBuilder) FullCopyIndicesCount(fullcopyindicescount int) *SearchableSnapshotsBuilder {
	rb.v.FullCopyIndicesCount = &fullcopyindicescount
	return rb
}

// IndicesCount set the IndicesCount property for SearchableSnapshotsBuilder.
func (rb *SearchableSnapshotsBuilder) IndicesCount(indicescount int) *SearchableSnapshotsBuilder {
	rb.v.IndicesCount = indicescount
	return rb
}

// SharedCacheIndicesCount set the SharedCacheIndicesCount property for SearchableSnapshotsBuilder.
func (rb *SearchableSnapshotsBuilder) SharedCacheIndicesCount(sharedcacheindicescount int) *SearchableSnapshotsBuilder {
	rb.v.SharedCacheIndicesCount = &sharedcacheindicescount
	return rb
}
