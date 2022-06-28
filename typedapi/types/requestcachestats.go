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

// RequestCacheStats type.
type RequestCacheStats struct {
	Evictions         int64   `json:"evictions"`
	HitCount          int64   `json:"hit_count"`
	MemorySize        *string `json:"memory_size,omitempty"`
	MemorySizeInBytes int64   `json:"memory_size_in_bytes"`
	MissCount         int64   `json:"miss_count"`
}

// RequestCacheStatsBuilder holds RequestCacheStats struct and provides a builder API.
type RequestCacheStatsBuilder struct {
	v *RequestCacheStats
}

// NewRequestCacheStats provides a builder for the RequestCacheStats struct.
func NewRequestCacheStats() *RequestCacheStatsBuilder {
	r := RequestCacheStatsBuilder{
		&RequestCacheStats{},
	}

	return &r
}

// Build finalize the chain and returns the RequestCacheStats struct
func (rb *RequestCacheStatsBuilder) Build() RequestCacheStats {
	return *rb.v
}

// Evictions set the Evictions property for RequestCacheStatsBuilder.
func (rb *RequestCacheStatsBuilder) Evictions(evictions int64) *RequestCacheStatsBuilder {
	rb.v.Evictions = evictions
	return rb
}

// HitCount set the HitCount property for RequestCacheStatsBuilder.
func (rb *RequestCacheStatsBuilder) HitCount(hitcount int64) *RequestCacheStatsBuilder {
	rb.v.HitCount = hitcount
	return rb
}

// MemorySize set the MemorySize property for RequestCacheStatsBuilder.
func (rb *RequestCacheStatsBuilder) MemorySize(memorysize string) *RequestCacheStatsBuilder {
	rb.v.MemorySize = &memorysize
	return rb
}

// MemorySizeInBytes set the MemorySizeInBytes property for RequestCacheStatsBuilder.
func (rb *RequestCacheStatsBuilder) MemorySizeInBytes(memorysizeinbytes int64) *RequestCacheStatsBuilder {
	rb.v.MemorySizeInBytes = memorysizeinbytes
	return rb
}

// MissCount set the MissCount property for RequestCacheStatsBuilder.
func (rb *RequestCacheStatsBuilder) MissCount(misscount int64) *RequestCacheStatsBuilder {
	rb.v.MissCount = misscount
	return rb
}
