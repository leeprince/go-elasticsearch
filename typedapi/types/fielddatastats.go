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

// FielddataStats type.
type FielddataStats struct {
	Evictions         *int64                     `json:"evictions,omitempty"`
	Fields            map[Field]FieldMemoryUsage `json:"fields,omitempty"`
	MemorySize        *ByteSize                  `json:"memory_size,omitempty"`
	MemorySizeInBytes int64                      `json:"memory_size_in_bytes"`
}

// FielddataStatsBuilder holds FielddataStats struct and provides a builder API.
type FielddataStatsBuilder struct {
	v *FielddataStats
}

// NewFielddataStats provides a builder for the FielddataStats struct.
func NewFielddataStats() *FielddataStatsBuilder {
	r := FielddataStatsBuilder{
		&FielddataStats{
			Fields: make(map[Field]FieldMemoryUsage, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the FielddataStats struct
func (rb *FielddataStatsBuilder) Build() FielddataStats {
	return *rb.v
}

// Evictions set the Evictions property for FielddataStatsBuilder.
func (rb *FielddataStatsBuilder) Evictions(evictions int64) *FielddataStatsBuilder {
	rb.v.Evictions = &evictions
	return rb
}

// Fields set the Fields property for FielddataStatsBuilder.
func (rb *FielddataStatsBuilder) Fields(value map[Field]FieldMemoryUsage) *FielddataStatsBuilder {
	rb.v.Fields = value
	return rb
}

// MemorySize set the MemorySize property for FielddataStatsBuilder.
func (rb *FielddataStatsBuilder) MemorySize(memorysize ByteSize) *FielddataStatsBuilder {
	rb.v.MemorySize = &memorysize
	return rb
}

// MemorySizeInBytes set the MemorySizeInBytes property for FielddataStatsBuilder.
func (rb *FielddataStatsBuilder) MemorySizeInBytes(memorysizeinbytes int64) *FielddataStatsBuilder {
	rb.v.MemorySizeInBytes = memorysizeinbytes
	return rb
}
