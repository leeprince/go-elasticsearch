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

// ShardQueryCache type.
type ShardQueryCache struct {
	CacheCount        int64 `json:"cache_count"`
	CacheSize         int64 `json:"cache_size"`
	Evictions         int64 `json:"evictions"`
	HitCount          int64 `json:"hit_count"`
	MemorySizeInBytes int64 `json:"memory_size_in_bytes"`
	MissCount         int64 `json:"miss_count"`
	TotalCount        int64 `json:"total_count"`
}

// ShardQueryCacheBuilder holds ShardQueryCache struct and provides a builder API.
type ShardQueryCacheBuilder struct {
	v *ShardQueryCache
}

// NewShardQueryCache provides a builder for the ShardQueryCache struct.
func NewShardQueryCache() *ShardQueryCacheBuilder {
	r := ShardQueryCacheBuilder{
		&ShardQueryCache{},
	}

	return &r
}

// Build finalize the chain and returns the ShardQueryCache struct
func (rb *ShardQueryCacheBuilder) Build() ShardQueryCache {
	return *rb.v
}

// CacheCount set the CacheCount property for ShardQueryCacheBuilder.
func (rb *ShardQueryCacheBuilder) CacheCount(cachecount int64) *ShardQueryCacheBuilder {
	rb.v.CacheCount = cachecount
	return rb
}

// CacheSize set the CacheSize property for ShardQueryCacheBuilder.
func (rb *ShardQueryCacheBuilder) CacheSize(cachesize int64) *ShardQueryCacheBuilder {
	rb.v.CacheSize = cachesize
	return rb
}

// Evictions set the Evictions property for ShardQueryCacheBuilder.
func (rb *ShardQueryCacheBuilder) Evictions(evictions int64) *ShardQueryCacheBuilder {
	rb.v.Evictions = evictions
	return rb
}

// HitCount set the HitCount property for ShardQueryCacheBuilder.
func (rb *ShardQueryCacheBuilder) HitCount(hitcount int64) *ShardQueryCacheBuilder {
	rb.v.HitCount = hitcount
	return rb
}

// MemorySizeInBytes set the MemorySizeInBytes property for ShardQueryCacheBuilder.
func (rb *ShardQueryCacheBuilder) MemorySizeInBytes(memorysizeinbytes int64) *ShardQueryCacheBuilder {
	rb.v.MemorySizeInBytes = memorysizeinbytes
	return rb
}

// MissCount set the MissCount property for ShardQueryCacheBuilder.
func (rb *ShardQueryCacheBuilder) MissCount(misscount int64) *ShardQueryCacheBuilder {
	rb.v.MissCount = misscount
	return rb
}

// TotalCount set the TotalCount property for ShardQueryCacheBuilder.
func (rb *ShardQueryCacheBuilder) TotalCount(totalcount int64) *ShardQueryCacheBuilder {
	rb.v.TotalCount = totalcount
	return rb
}
