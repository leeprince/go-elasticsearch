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

// CacheStats type.
type CacheStats struct {
	Count     int `json:"count"`
	Evictions int `json:"evictions"`
	Hits      int `json:"hits"`
	Misses    int `json:"misses"`
	NodeId    Id  `json:"node_id"`
}

// CacheStatsBuilder holds CacheStats struct and provides a builder API.
type CacheStatsBuilder struct {
	v *CacheStats
}

// NewCacheStats provides a builder for the CacheStats struct.
func NewCacheStats() *CacheStatsBuilder {
	r := CacheStatsBuilder{
		&CacheStats{},
	}

	return &r
}

// Build finalize the chain and returns the CacheStats struct
func (rb *CacheStatsBuilder) Build() CacheStats {
	return *rb.v
}

// Count set the Count property for CacheStatsBuilder.
func (rb *CacheStatsBuilder) Count(count int) *CacheStatsBuilder {
	rb.v.Count = count
	return rb
}

// Evictions set the Evictions property for CacheStatsBuilder.
func (rb *CacheStatsBuilder) Evictions(evictions int) *CacheStatsBuilder {
	rb.v.Evictions = evictions
	return rb
}

// Hits set the Hits property for CacheStatsBuilder.
func (rb *CacheStatsBuilder) Hits(hits int) *CacheStatsBuilder {
	rb.v.Hits = hits
	return rb
}

// Misses set the Misses property for CacheStatsBuilder.
func (rb *CacheStatsBuilder) Misses(misses int) *CacheStatsBuilder {
	rb.v.Misses = misses
	return rb
}

// NodeId set the NodeId property for CacheStatsBuilder.
func (rb *CacheStatsBuilder) NodeId(nodeid Id) *CacheStatsBuilder {
	rb.v.NodeId = nodeid
	return rb
}
