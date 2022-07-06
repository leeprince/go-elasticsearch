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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/segmentsortmissing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/segmentsortmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/segmentsortorder"
)

// IndexSegmentSort type.
type IndexSegmentSort struct {
	Field   *Fields                                 `json:"field,omitempty"`
	Missing []segmentsortmissing.SegmentSortMissing `json:"missing,omitempty"`
	Mode    []segmentsortmode.SegmentSortMode       `json:"mode,omitempty"`
	Order   []segmentsortorder.SegmentSortOrder     `json:"order,omitempty"`
}

// IndexSegmentSortBuilder holds IndexSegmentSort struct and provides a builder API.
type IndexSegmentSortBuilder struct {
	v *IndexSegmentSort
}

// NewIndexSegmentSort provides a builder for the IndexSegmentSort struct.
func NewIndexSegmentSort() *IndexSegmentSortBuilder {
	r := IndexSegmentSortBuilder{
		&IndexSegmentSort{},
	}

	return &r
}

// Build finalize the chain and returns the IndexSegmentSort struct
func (rb *IndexSegmentSortBuilder) Build() IndexSegmentSort {
	return *rb.v
}

// Field set the Field property for IndexSegmentSortBuilder.
func (rb *IndexSegmentSortBuilder) Field(field Fields) *IndexSegmentSortBuilder {
	rb.v.Field = &field
	return rb
}

// Missing set the Missing property for IndexSegmentSortBuilder.
func (rb *IndexSegmentSortBuilder) Missing(arg []segmentsortmissing.SegmentSortMissing) *IndexSegmentSortBuilder {
	rb.v.Missing = arg
	return rb
}

// Mode set the Mode property for IndexSegmentSortBuilder.
func (rb *IndexSegmentSortBuilder) Mode(arg []segmentsortmode.SegmentSortMode) *IndexSegmentSortBuilder {
	rb.v.Mode = arg
	return rb
}

// Order set the Order property for IndexSegmentSortBuilder.
func (rb *IndexSegmentSortBuilder) Order(arg []segmentsortorder.SegmentSortOrder) *IndexSegmentSortBuilder {
	rb.v.Order = arg
	return rb
}
