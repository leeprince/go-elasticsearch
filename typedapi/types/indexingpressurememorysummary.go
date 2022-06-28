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

// IndexingPressureMemorySummary type.
type IndexingPressureMemorySummary struct {
	AllInBytes                            int64  `json:"all_in_bytes"`
	CombinedCoordinatingAndPrimaryInBytes int64  `json:"combined_coordinating_and_primary_in_bytes"`
	CoordinatingInBytes                   int64  `json:"coordinating_in_bytes"`
	CoordinatingRejections                *int64 `json:"coordinating_rejections,omitempty"`
	PrimaryInBytes                        int64  `json:"primary_in_bytes"`
	PrimaryRejections                     *int64 `json:"primary_rejections,omitempty"`
	ReplicaInBytes                        int64  `json:"replica_in_bytes"`
	ReplicaRejections                     *int64 `json:"replica_rejections,omitempty"`
}

// IndexingPressureMemorySummaryBuilder holds IndexingPressureMemorySummary struct and provides a builder API.
type IndexingPressureMemorySummaryBuilder struct {
	v *IndexingPressureMemorySummary
}

// NewIndexingPressureMemorySummary provides a builder for the IndexingPressureMemorySummary struct.
func NewIndexingPressureMemorySummary() *IndexingPressureMemorySummaryBuilder {
	r := IndexingPressureMemorySummaryBuilder{
		&IndexingPressureMemorySummary{},
	}

	return &r
}

// Build finalize the chain and returns the IndexingPressureMemorySummary struct
func (rb *IndexingPressureMemorySummaryBuilder) Build() IndexingPressureMemorySummary {
	return *rb.v
}

// AllInBytes set the AllInBytes property for IndexingPressureMemorySummaryBuilder.
func (rb *IndexingPressureMemorySummaryBuilder) AllInBytes(allinbytes int64) *IndexingPressureMemorySummaryBuilder {
	rb.v.AllInBytes = allinbytes
	return rb
}

// CombinedCoordinatingAndPrimaryInBytes set the CombinedCoordinatingAndPrimaryInBytes property for IndexingPressureMemorySummaryBuilder.
func (rb *IndexingPressureMemorySummaryBuilder) CombinedCoordinatingAndPrimaryInBytes(combinedcoordinatingandprimaryinbytes int64) *IndexingPressureMemorySummaryBuilder {
	rb.v.CombinedCoordinatingAndPrimaryInBytes = combinedcoordinatingandprimaryinbytes
	return rb
}

// CoordinatingInBytes set the CoordinatingInBytes property for IndexingPressureMemorySummaryBuilder.
func (rb *IndexingPressureMemorySummaryBuilder) CoordinatingInBytes(coordinatinginbytes int64) *IndexingPressureMemorySummaryBuilder {
	rb.v.CoordinatingInBytes = coordinatinginbytes
	return rb
}

// CoordinatingRejections set the CoordinatingRejections property for IndexingPressureMemorySummaryBuilder.
func (rb *IndexingPressureMemorySummaryBuilder) CoordinatingRejections(coordinatingrejections int64) *IndexingPressureMemorySummaryBuilder {
	rb.v.CoordinatingRejections = &coordinatingrejections
	return rb
}

// PrimaryInBytes set the PrimaryInBytes property for IndexingPressureMemorySummaryBuilder.
func (rb *IndexingPressureMemorySummaryBuilder) PrimaryInBytes(primaryinbytes int64) *IndexingPressureMemorySummaryBuilder {
	rb.v.PrimaryInBytes = primaryinbytes
	return rb
}

// PrimaryRejections set the PrimaryRejections property for IndexingPressureMemorySummaryBuilder.
func (rb *IndexingPressureMemorySummaryBuilder) PrimaryRejections(primaryrejections int64) *IndexingPressureMemorySummaryBuilder {
	rb.v.PrimaryRejections = &primaryrejections
	return rb
}

// ReplicaInBytes set the ReplicaInBytes property for IndexingPressureMemorySummaryBuilder.
func (rb *IndexingPressureMemorySummaryBuilder) ReplicaInBytes(replicainbytes int64) *IndexingPressureMemorySummaryBuilder {
	rb.v.ReplicaInBytes = replicainbytes
	return rb
}

// ReplicaRejections set the ReplicaRejections property for IndexingPressureMemorySummaryBuilder.
func (rb *IndexingPressureMemorySummaryBuilder) ReplicaRejections(replicarejections int64) *IndexingPressureMemorySummaryBuilder {
	rb.v.ReplicaRejections = &replicarejections
	return rb
}
