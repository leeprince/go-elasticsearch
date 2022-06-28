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

// IndexMappingRecord type.
type IndexMappingRecord struct {
	Item     *TypeMapping `json:"item,omitempty"`
	Mappings TypeMapping  `json:"mappings"`
}

// IndexMappingRecordBuilder holds IndexMappingRecord struct and provides a builder API.
type IndexMappingRecordBuilder struct {
	v *IndexMappingRecord
}

// NewIndexMappingRecord provides a builder for the IndexMappingRecord struct.
func NewIndexMappingRecord() *IndexMappingRecordBuilder {
	r := IndexMappingRecordBuilder{
		&IndexMappingRecord{},
	}

	return &r
}

// Build finalize the chain and returns the IndexMappingRecord struct
func (rb *IndexMappingRecordBuilder) Build() IndexMappingRecord {
	return *rb.v
}

// Item set the Item property for IndexMappingRecordBuilder.
func (rb *IndexMappingRecordBuilder) Item(item TypeMapping) *IndexMappingRecordBuilder {
	rb.v.Item = &item
	return rb
}

// Mappings set the Mappings property for IndexMappingRecordBuilder.
func (rb *IndexMappingRecordBuilder) Mappings(mappings TypeMapping) *IndexMappingRecordBuilder {
	rb.v.Mappings = mappings
	return rb
}
