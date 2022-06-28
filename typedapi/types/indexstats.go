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

// IndexStats type.
type IndexStats struct {
	Bulk *BulkStats `json:"bulk,omitempty"`
	// Completion Contains statistics about completions across all shards assigned to the node.
	Completion *CompletionStats `json:"completion,omitempty"`
	// Docs Contains statistics about documents across all primary shards assigned to the
	// node.
	Docs *DocStats `json:"docs,omitempty"`
	// Fielddata Contains statistics about the field data cache across all shards assigned to
	// the node.
	Fielddata *FielddataStats `json:"fielddata,omitempty"`
	// Flush Contains statistics about flush operations for the node.
	Flush *FlushStats `json:"flush,omitempty"`
	// Get Contains statistics about get operations for the node.
	Get *GetStats `json:"get,omitempty"`
	// Indexing Contains statistics about indexing operations for the node.
	Indexing *IndexingStats `json:"indexing,omitempty"`
	// Indices Contains statistics about indices operations for the node.
	Indices *IndicesStats `json:"indices,omitempty"`
	// Merges Contains statistics about merge operations for the node.
	Merges *MergesStats `json:"merges,omitempty"`
	// QueryCache Contains statistics about the query cache across all shards assigned to the
	// node.
	QueryCache *QueryCacheStats `json:"query_cache,omitempty"`
	// Recovery Contains statistics about recovery operations for the node.
	Recovery *RecoveryStats `json:"recovery,omitempty"`
	// Refresh Contains statistics about refresh operations for the node.
	Refresh *RefreshStats `json:"refresh,omitempty"`
	// RequestCache Contains statistics about the request cache across all shards assigned to the
	// node.
	RequestCache *RequestCacheStats `json:"request_cache,omitempty"`
	// Search Contains statistics about search operations for the node.
	Search *SearchStats `json:"search,omitempty"`
	// Segments Contains statistics about segments across all shards assigned to the node.
	Segments   *SegmentsStats    `json:"segments,omitempty"`
	ShardStats *ShardsTotalStats `json:"shard_stats,omitempty"`
	// Store Contains statistics about the size of shards assigned to the node.
	Store *StoreStats `json:"store,omitempty"`
	// Translog Contains statistics about transaction log operations for the node.
	Translog *TranslogStats `json:"translog,omitempty"`
	// Warmer Contains statistics about index warming operations for the node.
	Warmer *WarmerStats `json:"warmer,omitempty"`
}

// IndexStatsBuilder holds IndexStats struct and provides a builder API.
type IndexStatsBuilder struct {
	v *IndexStats
}

// NewIndexStats provides a builder for the IndexStats struct.
func NewIndexStats() *IndexStatsBuilder {
	r := IndexStatsBuilder{
		&IndexStats{},
	}

	return &r
}

// Build finalize the chain and returns the IndexStats struct
func (rb *IndexStatsBuilder) Build() IndexStats {
	return *rb.v
}

// Bulk set the Bulk property for IndexStatsBuilder.
func (rb *IndexStatsBuilder) Bulk(bulk BulkStats) *IndexStatsBuilder {
	rb.v.Bulk = &bulk
	return rb
}

// Completion Contains statistics about completions across all shards assigned to the node.
func (rb *IndexStatsBuilder) Completion(completion CompletionStats) *IndexStatsBuilder {
	rb.v.Completion = &completion
	return rb
}

// Docs Contains statistics about documents across all primary shards assigned to the
// node.
func (rb *IndexStatsBuilder) Docs(docs DocStats) *IndexStatsBuilder {
	rb.v.Docs = &docs
	return rb
}

// Fielddata Contains statistics about the field data cache across all shards assigned to
// the node.
func (rb *IndexStatsBuilder) Fielddata(fielddata FielddataStats) *IndexStatsBuilder {
	rb.v.Fielddata = &fielddata
	return rb
}

// Flush Contains statistics about flush operations for the node.
func (rb *IndexStatsBuilder) Flush(flush FlushStats) *IndexStatsBuilder {
	rb.v.Flush = &flush
	return rb
}

// Get Contains statistics about get operations for the node.
func (rb *IndexStatsBuilder) Get(get GetStats) *IndexStatsBuilder {
	rb.v.Get = &get
	return rb
}

// Indexing Contains statistics about indexing operations for the node.
func (rb *IndexStatsBuilder) Indexing(indexing IndexingStats) *IndexStatsBuilder {
	rb.v.Indexing = &indexing
	return rb
}

// Indices Contains statistics about indices operations for the node.
func (rb *IndexStatsBuilder) Indices(indices IndicesStats) *IndexStatsBuilder {
	rb.v.Indices = &indices
	return rb
}

// Merges Contains statistics about merge operations for the node.
func (rb *IndexStatsBuilder) Merges(merges MergesStats) *IndexStatsBuilder {
	rb.v.Merges = &merges
	return rb
}

// QueryCache Contains statistics about the query cache across all shards assigned to the
// node.
func (rb *IndexStatsBuilder) QueryCache(querycache QueryCacheStats) *IndexStatsBuilder {
	rb.v.QueryCache = &querycache
	return rb
}

// Recovery Contains statistics about recovery operations for the node.
func (rb *IndexStatsBuilder) Recovery(recovery RecoveryStats) *IndexStatsBuilder {
	rb.v.Recovery = &recovery
	return rb
}

// Refresh Contains statistics about refresh operations for the node.
func (rb *IndexStatsBuilder) Refresh(refresh RefreshStats) *IndexStatsBuilder {
	rb.v.Refresh = &refresh
	return rb
}

// RequestCache Contains statistics about the request cache across all shards assigned to the
// node.
func (rb *IndexStatsBuilder) RequestCache(requestcache RequestCacheStats) *IndexStatsBuilder {
	rb.v.RequestCache = &requestcache
	return rb
}

// Search Contains statistics about search operations for the node.
func (rb *IndexStatsBuilder) Search(search SearchStats) *IndexStatsBuilder {
	rb.v.Search = &search
	return rb
}

// Segments Contains statistics about segments across all shards assigned to the node.
func (rb *IndexStatsBuilder) Segments(segments SegmentsStats) *IndexStatsBuilder {
	rb.v.Segments = &segments
	return rb
}

// ShardStats set the ShardStats property for IndexStatsBuilder.
func (rb *IndexStatsBuilder) ShardStats(shardstats ShardsTotalStats) *IndexStatsBuilder {
	rb.v.ShardStats = &shardstats
	return rb
}

// Store Contains statistics about the size of shards assigned to the node.
func (rb *IndexStatsBuilder) Store(store StoreStats) *IndexStatsBuilder {
	rb.v.Store = &store
	return rb
}

// Translog Contains statistics about transaction log operations for the node.
func (rb *IndexStatsBuilder) Translog(translog TranslogStats) *IndexStatsBuilder {
	rb.v.Translog = &translog
	return rb
}

// Warmer Contains statistics about index warming operations for the node.
func (rb *IndexStatsBuilder) Warmer(warmer WarmerStats) *IndexStatsBuilder {
	rb.v.Warmer = &warmer
	return rb
}
