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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/gappolicy"
)

// BucketSortAggregation type.
type BucketSortAggregation struct {
	From      *int                 `json:"from,omitempty"`
	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Meta      *Metadata            `json:"meta,omitempty"`
	Name      *string              `json:"name,omitempty"`
	Size      *int                 `json:"size,omitempty"`
	Sort      *Sort                `json:"sort,omitempty"`
}

// BucketSortAggregationBuilder holds BucketSortAggregation struct and provides a builder API.
type BucketSortAggregationBuilder struct {
	v *BucketSortAggregation
}

// NewBucketSortAggregation provides a builder for the BucketSortAggregation struct.
func NewBucketSortAggregation() *BucketSortAggregationBuilder {
	r := BucketSortAggregationBuilder{
		&BucketSortAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the BucketSortAggregation struct
func (rb *BucketSortAggregationBuilder) Build() BucketSortAggregation {
	return *rb.v
}

// From set the From property for BucketSortAggregationBuilder.
func (rb *BucketSortAggregationBuilder) From(from int) *BucketSortAggregationBuilder {
	rb.v.From = &from
	return rb
}

// GapPolicy set the GapPolicy property for BucketSortAggregationBuilder.
func (rb *BucketSortAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *BucketSortAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

// Meta set the Meta property for BucketSortAggregationBuilder.
func (rb *BucketSortAggregationBuilder) Meta(meta Metadata) *BucketSortAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for BucketSortAggregationBuilder.
func (rb *BucketSortAggregationBuilder) Name(name string) *BucketSortAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Size set the Size property for BucketSortAggregationBuilder.
func (rb *BucketSortAggregationBuilder) Size(size int) *BucketSortAggregationBuilder {
	rb.v.Size = &size
	return rb
}

// Sort set the Sort property for BucketSortAggregationBuilder.
func (rb *BucketSortAggregationBuilder) Sort(sort Sort) *BucketSortAggregationBuilder {
	rb.v.Sort = &sort
	return rb
}
