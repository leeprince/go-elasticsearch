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

// IndexingPressureMemory type.
type IndexingPressureMemory struct {
	Current      *PressureMemory `json:"current,omitempty"`
	LimitInBytes *int64          `json:"limit_in_bytes,omitempty"`
	Total        *PressureMemory `json:"total,omitempty"`
}

// IndexingPressureMemoryBuilder holds IndexingPressureMemory struct and provides a builder API.
type IndexingPressureMemoryBuilder struct {
	v *IndexingPressureMemory
}

// NewIndexingPressureMemory provides a builder for the IndexingPressureMemory struct.
func NewIndexingPressureMemory() *IndexingPressureMemoryBuilder {
	r := IndexingPressureMemoryBuilder{
		&IndexingPressureMemory{},
	}

	return &r
}

// Build finalize the chain and returns the IndexingPressureMemory struct
func (rb *IndexingPressureMemoryBuilder) Build() IndexingPressureMemory {
	return *rb.v
}

// Current set the Current property for IndexingPressureMemoryBuilder.
func (rb *IndexingPressureMemoryBuilder) Current(current PressureMemory) *IndexingPressureMemoryBuilder {
	rb.v.Current = &current
	return rb
}

// LimitInBytes set the LimitInBytes property for IndexingPressureMemoryBuilder.
func (rb *IndexingPressureMemoryBuilder) LimitInBytes(limitinbytes int64) *IndexingPressureMemoryBuilder {
	rb.v.LimitInBytes = &limitinbytes
	return rb
}

// Total set the Total property for IndexingPressureMemoryBuilder.
func (rb *IndexingPressureMemoryBuilder) Total(total PressureMemory) *IndexingPressureMemoryBuilder {
	rb.v.Total = &total
	return rb
}
