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

// ShardStats type.
type ShardStats struct {
	Bulk            *BulkStats            `json:"bulk,omitempty"`
	Commit          *ShardCommit          `json:"commit,omitempty"`
	Completion      *CompletionStats      `json:"completion,omitempty"`
	Docs            *DocStats             `json:"docs,omitempty"`
	Fielddata       *FielddataStats       `json:"fielddata,omitempty"`
	Flush           *FlushStats           `json:"flush,omitempty"`
	Get             *GetStats             `json:"get,omitempty"`
	Indexing        *IndexingStats        `json:"indexing,omitempty"`
	Indices         *IndicesStats         `json:"indices,omitempty"`
	Merges          *MergesStats          `json:"merges,omitempty"`
	QueryCache      *ShardQueryCache      `json:"query_cache,omitempty"`
	Recovery        *RecoveryStats        `json:"recovery,omitempty"`
	Refresh         *RefreshStats         `json:"refresh,omitempty"`
	RequestCache    *RequestCacheStats    `json:"request_cache,omitempty"`
	RetentionLeases *ShardRetentionLeases `json:"retention_leases,omitempty"`
	Routing         *ShardRouting         `json:"routing,omitempty"`
	Search          *SearchStats          `json:"search,omitempty"`
	Segments        *SegmentsStats        `json:"segments,omitempty"`
	SeqNo           *ShardSequenceNumber  `json:"seq_no,omitempty"`
	ShardPath       *ShardPath            `json:"shard_path,omitempty"`
	ShardStats      *ShardsTotalStats     `json:"shard_stats,omitempty"`
	Shards          *ShardsTotalStats     `json:"shards,omitempty"`
	Store           *StoreStats           `json:"store,omitempty"`
	Translog        *TranslogStats        `json:"translog,omitempty"`
	Warmer          *WarmerStats          `json:"warmer,omitempty"`
}

// ShardStatsBuilder holds ShardStats struct and provides a builder API.
type ShardStatsBuilder struct {
	v *ShardStats
}

// NewShardStats provides a builder for the ShardStats struct.
func NewShardStats() *ShardStatsBuilder {
	r := ShardStatsBuilder{
		&ShardStats{},
	}

	return &r
}

// Build finalize the chain and returns the ShardStats struct
func (rb *ShardStatsBuilder) Build() ShardStats {
	return *rb.v
}

// Bulk set the Bulk property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Bulk(bulk BulkStats) *ShardStatsBuilder {
	rb.v.Bulk = &bulk
	return rb
}

// Commit set the Commit property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Commit(commit ShardCommit) *ShardStatsBuilder {
	rb.v.Commit = &commit
	return rb
}

// Completion set the Completion property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Completion(completion CompletionStats) *ShardStatsBuilder {
	rb.v.Completion = &completion
	return rb
}

// Docs set the Docs property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Docs(docs DocStats) *ShardStatsBuilder {
	rb.v.Docs = &docs
	return rb
}

// Fielddata set the Fielddata property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Fielddata(fielddata FielddataStats) *ShardStatsBuilder {
	rb.v.Fielddata = &fielddata
	return rb
}

// Flush set the Flush property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Flush(flush FlushStats) *ShardStatsBuilder {
	rb.v.Flush = &flush
	return rb
}

// Get set the Get property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Get(get GetStats) *ShardStatsBuilder {
	rb.v.Get = &get
	return rb
}

// Indexing set the Indexing property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Indexing(indexing IndexingStats) *ShardStatsBuilder {
	rb.v.Indexing = &indexing
	return rb
}

// Indices set the Indices property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Indices(indices IndicesStats) *ShardStatsBuilder {
	rb.v.Indices = &indices
	return rb
}

// Merges set the Merges property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Merges(merges MergesStats) *ShardStatsBuilder {
	rb.v.Merges = &merges
	return rb
}

// QueryCache set the QueryCache property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) QueryCache(querycache ShardQueryCache) *ShardStatsBuilder {
	rb.v.QueryCache = &querycache
	return rb
}

// Recovery set the Recovery property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Recovery(recovery RecoveryStats) *ShardStatsBuilder {
	rb.v.Recovery = &recovery
	return rb
}

// Refresh set the Refresh property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Refresh(refresh RefreshStats) *ShardStatsBuilder {
	rb.v.Refresh = &refresh
	return rb
}

// RequestCache set the RequestCache property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) RequestCache(requestcache RequestCacheStats) *ShardStatsBuilder {
	rb.v.RequestCache = &requestcache
	return rb
}

// RetentionLeases set the RetentionLeases property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) RetentionLeases(retentionleases ShardRetentionLeases) *ShardStatsBuilder {
	rb.v.RetentionLeases = &retentionleases
	return rb
}

// Routing set the Routing property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Routing(routing ShardRouting) *ShardStatsBuilder {
	rb.v.Routing = &routing
	return rb
}

// Search set the Search property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Search(search SearchStats) *ShardStatsBuilder {
	rb.v.Search = &search
	return rb
}

// Segments set the Segments property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Segments(segments SegmentsStats) *ShardStatsBuilder {
	rb.v.Segments = &segments
	return rb
}

// SeqNo set the SeqNo property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) SeqNo(seqno ShardSequenceNumber) *ShardStatsBuilder {
	rb.v.SeqNo = &seqno
	return rb
}

// ShardPath set the ShardPath property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) ShardPath(shardpath ShardPath) *ShardStatsBuilder {
	rb.v.ShardPath = &shardpath
	return rb
}

// ShardStats set the ShardStats property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) ShardStats(shardstats ShardsTotalStats) *ShardStatsBuilder {
	rb.v.ShardStats = &shardstats
	return rb
}

// Shards set the Shards property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Shards(shards ShardsTotalStats) *ShardStatsBuilder {
	rb.v.Shards = &shards
	return rb
}

// Store set the Store property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Store(store StoreStats) *ShardStatsBuilder {
	rb.v.Store = &store
	return rb
}

// Translog set the Translog property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Translog(translog TranslogStats) *ShardStatsBuilder {
	rb.v.Translog = &translog
	return rb
}

// Warmer set the Warmer property for ShardStatsBuilder.
func (rb *ShardStatsBuilder) Warmer(warmer WarmerStats) *ShardStatsBuilder {
	rb.v.Warmer = &warmer
	return rb
}
