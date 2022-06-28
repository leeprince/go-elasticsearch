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

// SegmentsStats type.
type SegmentsStats struct {
	Count                       int                          `json:"count"`
	DocValuesMemory             *ByteSize                    `json:"doc_values_memory,omitempty"`
	DocValuesMemoryInBytes      int                          `json:"doc_values_memory_in_bytes"`
	FileSizes                   map[string]ShardFileSizeInfo `json:"file_sizes"`
	FixedBitSet                 *ByteSize                    `json:"fixed_bit_set,omitempty"`
	FixedBitSetMemoryInBytes    int                          `json:"fixed_bit_set_memory_in_bytes"`
	IndexWriterMaxMemoryInBytes *int                         `json:"index_writer_max_memory_in_bytes,omitempty"`
	IndexWriterMemory           *ByteSize                    `json:"index_writer_memory,omitempty"`
	IndexWriterMemoryInBytes    int                          `json:"index_writer_memory_in_bytes"`
	MaxUnsafeAutoIdTimestamp    int64                        `json:"max_unsafe_auto_id_timestamp"`
	Memory                      *ByteSize                    `json:"memory,omitempty"`
	MemoryInBytes               int                          `json:"memory_in_bytes"`
	NormsMemory                 *ByteSize                    `json:"norms_memory,omitempty"`
	NormsMemoryInBytes          int                          `json:"norms_memory_in_bytes"`
	PointsMemory                *ByteSize                    `json:"points_memory,omitempty"`
	PointsMemoryInBytes         int                          `json:"points_memory_in_bytes"`
	StoredFieldsMemoryInBytes   int                          `json:"stored_fields_memory_in_bytes"`
	StoredMemory                *ByteSize                    `json:"stored_memory,omitempty"`
	TermVectorsMemoryInBytes    int                          `json:"term_vectors_memory_in_bytes"`
	TermVectoryMemory           *ByteSize                    `json:"term_vectory_memory,omitempty"`
	TermsMemory                 *ByteSize                    `json:"terms_memory,omitempty"`
	TermsMemoryInBytes          int                          `json:"terms_memory_in_bytes"`
	VersionMapMemory            *ByteSize                    `json:"version_map_memory,omitempty"`
	VersionMapMemoryInBytes     int                          `json:"version_map_memory_in_bytes"`
}

// SegmentsStatsBuilder holds SegmentsStats struct and provides a builder API.
type SegmentsStatsBuilder struct {
	v *SegmentsStats
}

// NewSegmentsStats provides a builder for the SegmentsStats struct.
func NewSegmentsStats() *SegmentsStatsBuilder {
	r := SegmentsStatsBuilder{
		&SegmentsStats{
			FileSizes: make(map[string]ShardFileSizeInfo, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the SegmentsStats struct
func (rb *SegmentsStatsBuilder) Build() SegmentsStats {
	return *rb.v
}

// Count set the Count property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) Count(count int) *SegmentsStatsBuilder {
	rb.v.Count = count
	return rb
}

// DocValuesMemory set the DocValuesMemory property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) DocValuesMemory(docvaluesmemory ByteSize) *SegmentsStatsBuilder {
	rb.v.DocValuesMemory = &docvaluesmemory
	return rb
}

// DocValuesMemoryInBytes set the DocValuesMemoryInBytes property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) DocValuesMemoryInBytes(docvaluesmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.DocValuesMemoryInBytes = docvaluesmemoryinbytes
	return rb
}

// FileSizes set the FileSizes property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) FileSizes(value map[string]ShardFileSizeInfo) *SegmentsStatsBuilder {
	rb.v.FileSizes = value
	return rb
}

// FixedBitSet set the FixedBitSet property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) FixedBitSet(fixedbitset ByteSize) *SegmentsStatsBuilder {
	rb.v.FixedBitSet = &fixedbitset
	return rb
}

// FixedBitSetMemoryInBytes set the FixedBitSetMemoryInBytes property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) FixedBitSetMemoryInBytes(fixedbitsetmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.FixedBitSetMemoryInBytes = fixedbitsetmemoryinbytes
	return rb
}

// IndexWriterMaxMemoryInBytes set the IndexWriterMaxMemoryInBytes property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) IndexWriterMaxMemoryInBytes(indexwritermaxmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.IndexWriterMaxMemoryInBytes = &indexwritermaxmemoryinbytes
	return rb
}

// IndexWriterMemory set the IndexWriterMemory property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) IndexWriterMemory(indexwritermemory ByteSize) *SegmentsStatsBuilder {
	rb.v.IndexWriterMemory = &indexwritermemory
	return rb
}

// IndexWriterMemoryInBytes set the IndexWriterMemoryInBytes property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) IndexWriterMemoryInBytes(indexwritermemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.IndexWriterMemoryInBytes = indexwritermemoryinbytes
	return rb
}

// MaxUnsafeAutoIdTimestamp set the MaxUnsafeAutoIdTimestamp property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) MaxUnsafeAutoIdTimestamp(maxunsafeautoidtimestamp int64) *SegmentsStatsBuilder {
	rb.v.MaxUnsafeAutoIdTimestamp = maxunsafeautoidtimestamp
	return rb
}

// Memory set the Memory property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) Memory(memory ByteSize) *SegmentsStatsBuilder {
	rb.v.Memory = &memory
	return rb
}

// MemoryInBytes set the MemoryInBytes property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) MemoryInBytes(memoryinbytes int) *SegmentsStatsBuilder {
	rb.v.MemoryInBytes = memoryinbytes
	return rb
}

// NormsMemory set the NormsMemory property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) NormsMemory(normsmemory ByteSize) *SegmentsStatsBuilder {
	rb.v.NormsMemory = &normsmemory
	return rb
}

// NormsMemoryInBytes set the NormsMemoryInBytes property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) NormsMemoryInBytes(normsmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.NormsMemoryInBytes = normsmemoryinbytes
	return rb
}

// PointsMemory set the PointsMemory property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) PointsMemory(pointsmemory ByteSize) *SegmentsStatsBuilder {
	rb.v.PointsMemory = &pointsmemory
	return rb
}

// PointsMemoryInBytes set the PointsMemoryInBytes property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) PointsMemoryInBytes(pointsmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.PointsMemoryInBytes = pointsmemoryinbytes
	return rb
}

// StoredFieldsMemoryInBytes set the StoredFieldsMemoryInBytes property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) StoredFieldsMemoryInBytes(storedfieldsmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.StoredFieldsMemoryInBytes = storedfieldsmemoryinbytes
	return rb
}

// StoredMemory set the StoredMemory property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) StoredMemory(storedmemory ByteSize) *SegmentsStatsBuilder {
	rb.v.StoredMemory = &storedmemory
	return rb
}

// TermVectorsMemoryInBytes set the TermVectorsMemoryInBytes property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) TermVectorsMemoryInBytes(termvectorsmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.TermVectorsMemoryInBytes = termvectorsmemoryinbytes
	return rb
}

// TermVectoryMemory set the TermVectoryMemory property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) TermVectoryMemory(termvectorymemory ByteSize) *SegmentsStatsBuilder {
	rb.v.TermVectoryMemory = &termvectorymemory
	return rb
}

// TermsMemory set the TermsMemory property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) TermsMemory(termsmemory ByteSize) *SegmentsStatsBuilder {
	rb.v.TermsMemory = &termsmemory
	return rb
}

// TermsMemoryInBytes set the TermsMemoryInBytes property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) TermsMemoryInBytes(termsmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.TermsMemoryInBytes = termsmemoryinbytes
	return rb
}

// VersionMapMemory set the VersionMapMemory property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) VersionMapMemory(versionmapmemory ByteSize) *SegmentsStatsBuilder {
	rb.v.VersionMapMemory = &versionmapmemory
	return rb
}

// VersionMapMemoryInBytes set the VersionMapMemoryInBytes property for SegmentsStatsBuilder.
func (rb *SegmentsStatsBuilder) VersionMapMemoryInBytes(versionmapmemoryinbytes int) *SegmentsStatsBuilder {
	rb.v.VersionMapMemoryInBytes = versionmapmemoryinbytes
	return rb
}
