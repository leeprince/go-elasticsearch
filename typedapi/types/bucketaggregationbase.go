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

// BucketAggregationBase type.
type BucketAggregationBase struct {
	Meta *Metadata `json:"meta,omitempty"`
	Name *string   `json:"name,omitempty"`
}

// BucketAggregationBaseBuilder holds BucketAggregationBase struct and provides a builder API.
type BucketAggregationBaseBuilder struct {
	v *BucketAggregationBase
}

// NewBucketAggregationBase provides a builder for the BucketAggregationBase struct.
func NewBucketAggregationBase() *BucketAggregationBaseBuilder {
	r := BucketAggregationBaseBuilder{
		&BucketAggregationBase{},
	}

	return &r
}

// Build finalize the chain and returns the BucketAggregationBase struct
func (rb *BucketAggregationBaseBuilder) Build() BucketAggregationBase {
	return *rb.v
}

// Meta set the Meta property for BucketAggregationBaseBuilder.
func (rb *BucketAggregationBaseBuilder) Meta(meta Metadata) *BucketAggregationBaseBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for BucketAggregationBaseBuilder.
func (rb *BucketAggregationBaseBuilder) Name(name string) *BucketAggregationBaseBuilder {
	rb.v.Name = &name
	return rb
}
