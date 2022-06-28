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

// StoreStats type.
type StoreStats struct {
	Reserved                *ByteSize `json:"reserved,omitempty"`
	ReservedInBytes         int       `json:"reserved_in_bytes"`
	Size                    *ByteSize `json:"size,omitempty"`
	SizeInBytes             int       `json:"size_in_bytes"`
	TotalDataSetSize        *ByteSize `json:"total_data_set_size,omitempty"`
	TotalDataSetSizeInBytes *int      `json:"total_data_set_size_in_bytes,omitempty"`
}

// StoreStatsBuilder holds StoreStats struct and provides a builder API.
type StoreStatsBuilder struct {
	v *StoreStats
}

// NewStoreStats provides a builder for the StoreStats struct.
func NewStoreStats() *StoreStatsBuilder {
	r := StoreStatsBuilder{
		&StoreStats{},
	}

	return &r
}

// Build finalize the chain and returns the StoreStats struct
func (rb *StoreStatsBuilder) Build() StoreStats {
	return *rb.v
}

// Reserved set the Reserved property for StoreStatsBuilder.
func (rb *StoreStatsBuilder) Reserved(reserved ByteSize) *StoreStatsBuilder {
	rb.v.Reserved = &reserved
	return rb
}

// ReservedInBytes set the ReservedInBytes property for StoreStatsBuilder.
func (rb *StoreStatsBuilder) ReservedInBytes(reservedinbytes int) *StoreStatsBuilder {
	rb.v.ReservedInBytes = reservedinbytes
	return rb
}

// Size set the Size property for StoreStatsBuilder.
func (rb *StoreStatsBuilder) Size(size ByteSize) *StoreStatsBuilder {
	rb.v.Size = &size
	return rb
}

// SizeInBytes set the SizeInBytes property for StoreStatsBuilder.
func (rb *StoreStatsBuilder) SizeInBytes(sizeinbytes int) *StoreStatsBuilder {
	rb.v.SizeInBytes = sizeinbytes
	return rb
}

// TotalDataSetSize set the TotalDataSetSize property for StoreStatsBuilder.
func (rb *StoreStatsBuilder) TotalDataSetSize(totaldatasetsize ByteSize) *StoreStatsBuilder {
	rb.v.TotalDataSetSize = &totaldatasetsize
	return rb
}

// TotalDataSetSizeInBytes set the TotalDataSetSizeInBytes property for StoreStatsBuilder.
func (rb *StoreStatsBuilder) TotalDataSetSizeInBytes(totaldatasetsizeinbytes int) *StoreStatsBuilder {
	rb.v.TotalDataSetSizeInBytes = &totaldatasetsizeinbytes
	return rb
}
