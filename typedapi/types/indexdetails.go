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

// IndexDetails type.
type IndexDetails struct {
	MaxSegmentsPerShard int64     `json:"max_segments_per_shard"`
	ShardCount          int       `json:"shard_count"`
	Size                *ByteSize `json:"size,omitempty"`
	SizeInBytes         int64     `json:"size_in_bytes"`
}

// IndexDetailsBuilder holds IndexDetails struct and provides a builder API.
type IndexDetailsBuilder struct {
	v *IndexDetails
}

// NewIndexDetails provides a builder for the IndexDetails struct.
func NewIndexDetails() *IndexDetailsBuilder {
	r := IndexDetailsBuilder{
		&IndexDetails{},
	}

	return &r
}

// Build finalize the chain and returns the IndexDetails struct
func (rb *IndexDetailsBuilder) Build() IndexDetails {
	return *rb.v
}

// MaxSegmentsPerShard set the MaxSegmentsPerShard property for IndexDetailsBuilder.
func (rb *IndexDetailsBuilder) MaxSegmentsPerShard(maxsegmentspershard int64) *IndexDetailsBuilder {
	rb.v.MaxSegmentsPerShard = maxsegmentspershard
	return rb
}

// ShardCount set the ShardCount property for IndexDetailsBuilder.
func (rb *IndexDetailsBuilder) ShardCount(shardcount int) *IndexDetailsBuilder {
	rb.v.ShardCount = shardcount
	return rb
}

// Size set the Size property for IndexDetailsBuilder.
func (rb *IndexDetailsBuilder) Size(size ByteSize) *IndexDetailsBuilder {
	rb.v.Size = &size
	return rb
}

// SizeInBytes set the SizeInBytes property for IndexDetailsBuilder.
func (rb *IndexDetailsBuilder) SizeInBytes(sizeinbytes int64) *IndexDetailsBuilder {
	rb.v.SizeInBytes = sizeinbytes
	return rb
}
