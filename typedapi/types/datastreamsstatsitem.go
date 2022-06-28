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

// DataStreamsStatsItem type.
type DataStreamsStatsItem struct {
	BackingIndices   int       `json:"backing_indices"`
	DataStream       Name      `json:"data_stream"`
	MaximumTimestamp int64     `json:"maximum_timestamp"`
	StoreSize        *ByteSize `json:"store_size,omitempty"`
	StoreSizeBytes   int       `json:"store_size_bytes"`
}

// DataStreamsStatsItemBuilder holds DataStreamsStatsItem struct and provides a builder API.
type DataStreamsStatsItemBuilder struct {
	v *DataStreamsStatsItem
}

// NewDataStreamsStatsItem provides a builder for the DataStreamsStatsItem struct.
func NewDataStreamsStatsItem() *DataStreamsStatsItemBuilder {
	r := DataStreamsStatsItemBuilder{
		&DataStreamsStatsItem{},
	}

	return &r
}

// Build finalize the chain and returns the DataStreamsStatsItem struct
func (rb *DataStreamsStatsItemBuilder) Build() DataStreamsStatsItem {
	return *rb.v
}

// BackingIndices set the BackingIndices property for DataStreamsStatsItemBuilder.
func (rb *DataStreamsStatsItemBuilder) BackingIndices(backingindices int) *DataStreamsStatsItemBuilder {
	rb.v.BackingIndices = backingindices
	return rb
}

// DataStream set the DataStream property for DataStreamsStatsItemBuilder.
func (rb *DataStreamsStatsItemBuilder) DataStream(datastream Name) *DataStreamsStatsItemBuilder {
	rb.v.DataStream = datastream
	return rb
}

// MaximumTimestamp set the MaximumTimestamp property for DataStreamsStatsItemBuilder.
func (rb *DataStreamsStatsItemBuilder) MaximumTimestamp(maximumtimestamp int64) *DataStreamsStatsItemBuilder {
	rb.v.MaximumTimestamp = maximumtimestamp
	return rb
}

// StoreSize set the StoreSize property for DataStreamsStatsItemBuilder.
func (rb *DataStreamsStatsItemBuilder) StoreSize(storesize ByteSize) *DataStreamsStatsItemBuilder {
	rb.v.StoreSize = &storesize
	return rb
}

// StoreSizeBytes set the StoreSizeBytes property for DataStreamsStatsItemBuilder.
func (rb *DataStreamsStatsItemBuilder) StoreSizeBytes(storesizebytes int) *DataStreamsStatsItemBuilder {
	rb.v.StoreSizeBytes = storesizebytes
	return rb
}
