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

// ShardFileSizeInfo type.
type ShardFileSizeInfo struct {
	AverageSizeInBytes *int64 `json:"average_size_in_bytes,omitempty"`
	Count              *int64 `json:"count,omitempty"`
	Description        string `json:"description"`
	MaxSizeInBytes     *int64 `json:"max_size_in_bytes,omitempty"`
	MinSizeInBytes     *int64 `json:"min_size_in_bytes,omitempty"`
	SizeInBytes        int64  `json:"size_in_bytes"`
}

// ShardFileSizeInfoBuilder holds ShardFileSizeInfo struct and provides a builder API.
type ShardFileSizeInfoBuilder struct {
	v *ShardFileSizeInfo
}

// NewShardFileSizeInfo provides a builder for the ShardFileSizeInfo struct.
func NewShardFileSizeInfo() *ShardFileSizeInfoBuilder {
	r := ShardFileSizeInfoBuilder{
		&ShardFileSizeInfo{},
	}

	return &r
}

// Build finalize the chain and returns the ShardFileSizeInfo struct
func (rb *ShardFileSizeInfoBuilder) Build() ShardFileSizeInfo {
	return *rb.v
}

// AverageSizeInBytes set the AverageSizeInBytes property for ShardFileSizeInfoBuilder.
func (rb *ShardFileSizeInfoBuilder) AverageSizeInBytes(averagesizeinbytes int64) *ShardFileSizeInfoBuilder {
	rb.v.AverageSizeInBytes = &averagesizeinbytes
	return rb
}

// Count set the Count property for ShardFileSizeInfoBuilder.
func (rb *ShardFileSizeInfoBuilder) Count(count int64) *ShardFileSizeInfoBuilder {
	rb.v.Count = &count
	return rb
}

// Description set the Description property for ShardFileSizeInfoBuilder.
func (rb *ShardFileSizeInfoBuilder) Description(description string) *ShardFileSizeInfoBuilder {
	rb.v.Description = description
	return rb
}

// MaxSizeInBytes set the MaxSizeInBytes property for ShardFileSizeInfoBuilder.
func (rb *ShardFileSizeInfoBuilder) MaxSizeInBytes(maxsizeinbytes int64) *ShardFileSizeInfoBuilder {
	rb.v.MaxSizeInBytes = &maxsizeinbytes
	return rb
}

// MinSizeInBytes set the MinSizeInBytes property for ShardFileSizeInfoBuilder.
func (rb *ShardFileSizeInfoBuilder) MinSizeInBytes(minsizeinbytes int64) *ShardFileSizeInfoBuilder {
	rb.v.MinSizeInBytes = &minsizeinbytes
	return rb
}

// SizeInBytes set the SizeInBytes property for ShardFileSizeInfoBuilder.
func (rb *ShardFileSizeInfoBuilder) SizeInBytes(sizeinbytes int64) *ShardFileSizeInfoBuilder {
	rb.v.SizeInBytes = sizeinbytes
	return rb
}
