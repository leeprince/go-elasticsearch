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

// ClusterIndices type.
type ClusterIndices struct {
	// Analysis Contains statistics about analyzers and analyzer components used in selected
	// nodes.
	Analysis CharFilterTypes `json:"analysis"`
	// Completion Contains statistics about memory used for completion in selected nodes.
	Completion CompletionStats `json:"completion"`
	// Count Total number of indices with shards assigned to selected nodes.
	Count int64 `json:"count"`
	// Docs Contains counts for documents in selected nodes.
	Docs DocStats `json:"docs"`
	// Fielddata Contains statistics about the field data cache of selected nodes.
	Fielddata FielddataStats `json:"fielddata"`
	// Mappings Contains statistics about field mappings in selected nodes.
	Mappings FieldTypesMappings `json:"mappings"`
	// QueryCache Contains statistics about the query cache of selected nodes.
	QueryCache QueryCacheStats `json:"query_cache"`
	// Segments Contains statistics about segments in selected nodes.
	Segments SegmentsStats `json:"segments"`
	// Shards Contains statistics about indices with shards assigned to selected nodes.
	Shards ClusterIndicesShards `json:"shards"`
	// Store Contains statistics about the size of shards assigned to selected nodes.
	Store    StoreStats        `json:"store"`
	Versions []IndicesVersions `json:"versions,omitempty"`
}

// ClusterIndicesBuilder holds ClusterIndices struct and provides a builder API.
type ClusterIndicesBuilder struct {
	v *ClusterIndices
}

// NewClusterIndices provides a builder for the ClusterIndices struct.
func NewClusterIndices() *ClusterIndicesBuilder {
	r := ClusterIndicesBuilder{
		&ClusterIndices{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterIndices struct
func (rb *ClusterIndicesBuilder) Build() ClusterIndices {
	return *rb.v
}

// Analysis Contains statistics about analyzers and analyzer components used in selected
// nodes.
func (rb *ClusterIndicesBuilder) Analysis(analysis CharFilterTypes) *ClusterIndicesBuilder {
	rb.v.Analysis = analysis
	return rb
}

// Completion Contains statistics about memory used for completion in selected nodes.
func (rb *ClusterIndicesBuilder) Completion(completion CompletionStats) *ClusterIndicesBuilder {
	rb.v.Completion = completion
	return rb
}

// Count Total number of indices with shards assigned to selected nodes.
func (rb *ClusterIndicesBuilder) Count(count int64) *ClusterIndicesBuilder {
	rb.v.Count = count
	return rb
}

// Docs Contains counts for documents in selected nodes.
func (rb *ClusterIndicesBuilder) Docs(docs DocStats) *ClusterIndicesBuilder {
	rb.v.Docs = docs
	return rb
}

// Fielddata Contains statistics about the field data cache of selected nodes.
func (rb *ClusterIndicesBuilder) Fielddata(fielddata FielddataStats) *ClusterIndicesBuilder {
	rb.v.Fielddata = fielddata
	return rb
}

// Mappings Contains statistics about field mappings in selected nodes.
func (rb *ClusterIndicesBuilder) Mappings(mappings FieldTypesMappings) *ClusterIndicesBuilder {
	rb.v.Mappings = mappings
	return rb
}

// QueryCache Contains statistics about the query cache of selected nodes.
func (rb *ClusterIndicesBuilder) QueryCache(querycache QueryCacheStats) *ClusterIndicesBuilder {
	rb.v.QueryCache = querycache
	return rb
}

// Segments Contains statistics about segments in selected nodes.
func (rb *ClusterIndicesBuilder) Segments(segments SegmentsStats) *ClusterIndicesBuilder {
	rb.v.Segments = segments
	return rb
}

// Shards Contains statistics about indices with shards assigned to selected nodes.
func (rb *ClusterIndicesBuilder) Shards(shards ClusterIndicesShards) *ClusterIndicesBuilder {
	rb.v.Shards = shards
	return rb
}

// Store Contains statistics about the size of shards assigned to selected nodes.
func (rb *ClusterIndicesBuilder) Store(store StoreStats) *ClusterIndicesBuilder {
	rb.v.Store = store
	return rb
}

// Versions set the Versions property for ClusterIndicesBuilder.
func (rb *ClusterIndicesBuilder) Versions(versions ...IndicesVersions) *ClusterIndicesBuilder {
	rb.v.Versions = append(rb.v.Versions, versions...)
	return rb
}
