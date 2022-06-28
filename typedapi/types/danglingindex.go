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

// DanglingIndex type.
type DanglingIndex struct {
	CreationDateMillis EpochMillis `json:"creation_date_millis"`
	IndexName          string      `json:"index_name"`
	IndexUuid          string      `json:"index_uuid"`
	NodeIds            Ids         `json:"node_ids"`
}

// DanglingIndexBuilder holds DanglingIndex struct and provides a builder API.
type DanglingIndexBuilder struct {
	v *DanglingIndex
}

// NewDanglingIndex provides a builder for the DanglingIndex struct.
func NewDanglingIndex() *DanglingIndexBuilder {
	r := DanglingIndexBuilder{
		&DanglingIndex{},
	}

	return &r
}

// Build finalize the chain and returns the DanglingIndex struct
func (rb *DanglingIndexBuilder) Build() DanglingIndex {
	return *rb.v
}

// CreationDateMillis set the CreationDateMillis property for DanglingIndexBuilder.
func (rb *DanglingIndexBuilder) CreationDateMillis(creationdatemillis EpochMillis) *DanglingIndexBuilder {
	rb.v.CreationDateMillis = creationdatemillis
	return rb
}

// IndexName set the IndexName property for DanglingIndexBuilder.
func (rb *DanglingIndexBuilder) IndexName(indexname string) *DanglingIndexBuilder {
	rb.v.IndexName = indexname
	return rb
}

// IndexUuid set the IndexUuid property for DanglingIndexBuilder.
func (rb *DanglingIndexBuilder) IndexUuid(indexuuid string) *DanglingIndexBuilder {
	rb.v.IndexUuid = indexuuid
	return rb
}

// NodeIds set the NodeIds property for DanglingIndexBuilder.
func (rb *DanglingIndexBuilder) NodeIds(nodeids Ids) *DanglingIndexBuilder {
	rb.v.NodeIds = nodeids
	return rb
}
