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

// SerializedClusterStateDetail type.
type SerializedClusterStateDetail struct {
	CompressedSize          *string `json:"compressed_size,omitempty"`
	CompressedSizeInBytes   *int64  `json:"compressed_size_in_bytes,omitempty"`
	Count                   *int64  `json:"count,omitempty"`
	UncompressedSize        *string `json:"uncompressed_size,omitempty"`
	UncompressedSizeInBytes *int64  `json:"uncompressed_size_in_bytes,omitempty"`
}

// SerializedClusterStateDetailBuilder holds SerializedClusterStateDetail struct and provides a builder API.
type SerializedClusterStateDetailBuilder struct {
	v *SerializedClusterStateDetail
}

// NewSerializedClusterStateDetail provides a builder for the SerializedClusterStateDetail struct.
func NewSerializedClusterStateDetail() *SerializedClusterStateDetailBuilder {
	r := SerializedClusterStateDetailBuilder{
		&SerializedClusterStateDetail{},
	}

	return &r
}

// Build finalize the chain and returns the SerializedClusterStateDetail struct
func (rb *SerializedClusterStateDetailBuilder) Build() SerializedClusterStateDetail {
	return *rb.v
}

// CompressedSize set the CompressedSize property for SerializedClusterStateDetailBuilder.
func (rb *SerializedClusterStateDetailBuilder) CompressedSize(compressedsize string) *SerializedClusterStateDetailBuilder {
	rb.v.CompressedSize = &compressedsize
	return rb
}

// CompressedSizeInBytes set the CompressedSizeInBytes property for SerializedClusterStateDetailBuilder.
func (rb *SerializedClusterStateDetailBuilder) CompressedSizeInBytes(compressedsizeinbytes int64) *SerializedClusterStateDetailBuilder {
	rb.v.CompressedSizeInBytes = &compressedsizeinbytes
	return rb
}

// Count set the Count property for SerializedClusterStateDetailBuilder.
func (rb *SerializedClusterStateDetailBuilder) Count(count int64) *SerializedClusterStateDetailBuilder {
	rb.v.Count = &count
	return rb
}

// UncompressedSize set the UncompressedSize property for SerializedClusterStateDetailBuilder.
func (rb *SerializedClusterStateDetailBuilder) UncompressedSize(uncompressedsize string) *SerializedClusterStateDetailBuilder {
	rb.v.UncompressedSize = &uncompressedsize
	return rb
}

// UncompressedSizeInBytes set the UncompressedSizeInBytes property for SerializedClusterStateDetailBuilder.
func (rb *SerializedClusterStateDetailBuilder) UncompressedSizeInBytes(uncompressedsizeinbytes int64) *SerializedClusterStateDetailBuilder {
	rb.v.UncompressedSizeInBytes = &uncompressedsizeinbytes
	return rb
}
