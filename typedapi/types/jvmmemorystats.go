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

// JvmMemoryStats type.
type JvmMemoryStats struct {
	HeapCommittedInBytes    *int64          `json:"heap_committed_in_bytes,omitempty"`
	HeapMaxInBytes          *int64          `json:"heap_max_in_bytes,omitempty"`
	HeapUsedInBytes         *int64          `json:"heap_used_in_bytes,omitempty"`
	HeapUsedPercent         *int64          `json:"heap_used_percent,omitempty"`
	NonHeapCommittedInBytes *int64          `json:"non_heap_committed_in_bytes,omitempty"`
	NonHeapUsedInBytes      *int64          `json:"non_heap_used_in_bytes,omitempty"`
	Pools                   map[string]Pool `json:"pools,omitempty"`
}

// JvmMemoryStatsBuilder holds JvmMemoryStats struct and provides a builder API.
type JvmMemoryStatsBuilder struct {
	v *JvmMemoryStats
}

// NewJvmMemoryStats provides a builder for the JvmMemoryStats struct.
func NewJvmMemoryStats() *JvmMemoryStatsBuilder {
	r := JvmMemoryStatsBuilder{
		&JvmMemoryStats{
			Pools: make(map[string]Pool, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the JvmMemoryStats struct
func (rb *JvmMemoryStatsBuilder) Build() JvmMemoryStats {
	return *rb.v
}

// HeapCommittedInBytes set the HeapCommittedInBytes property for JvmMemoryStatsBuilder.
func (rb *JvmMemoryStatsBuilder) HeapCommittedInBytes(heapcommittedinbytes int64) *JvmMemoryStatsBuilder {
	rb.v.HeapCommittedInBytes = &heapcommittedinbytes
	return rb
}

// HeapMaxInBytes set the HeapMaxInBytes property for JvmMemoryStatsBuilder.
func (rb *JvmMemoryStatsBuilder) HeapMaxInBytes(heapmaxinbytes int64) *JvmMemoryStatsBuilder {
	rb.v.HeapMaxInBytes = &heapmaxinbytes
	return rb
}

// HeapUsedInBytes set the HeapUsedInBytes property for JvmMemoryStatsBuilder.
func (rb *JvmMemoryStatsBuilder) HeapUsedInBytes(heapusedinbytes int64) *JvmMemoryStatsBuilder {
	rb.v.HeapUsedInBytes = &heapusedinbytes
	return rb
}

// HeapUsedPercent set the HeapUsedPercent property for JvmMemoryStatsBuilder.
func (rb *JvmMemoryStatsBuilder) HeapUsedPercent(heapusedpercent int64) *JvmMemoryStatsBuilder {
	rb.v.HeapUsedPercent = &heapusedpercent
	return rb
}

// NonHeapCommittedInBytes set the NonHeapCommittedInBytes property for JvmMemoryStatsBuilder.
func (rb *JvmMemoryStatsBuilder) NonHeapCommittedInBytes(nonheapcommittedinbytes int64) *JvmMemoryStatsBuilder {
	rb.v.NonHeapCommittedInBytes = &nonheapcommittedinbytes
	return rb
}

// NonHeapUsedInBytes set the NonHeapUsedInBytes property for JvmMemoryStatsBuilder.
func (rb *JvmMemoryStatsBuilder) NonHeapUsedInBytes(nonheapusedinbytes int64) *JvmMemoryStatsBuilder {
	rb.v.NonHeapUsedInBytes = &nonheapusedinbytes
	return rb
}

// Pools set the Pools property for JvmMemoryStatsBuilder.
func (rb *JvmMemoryStatsBuilder) Pools(value map[string]Pool) *JvmMemoryStatsBuilder {
	rb.v.Pools = value
	return rb
}
