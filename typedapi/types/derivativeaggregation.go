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

// DerivativeAggregation type.
type DerivativeAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath         `json:"buckets_path,omitempty"`
	Format      *string              `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Meta        *Metadata            `json:"meta,omitempty"`
	Name        *string              `json:"name,omitempty"`
}

// DerivativeAggregationBuilder holds DerivativeAggregation struct and provides a builder API.
type DerivativeAggregationBuilder struct {
	v *DerivativeAggregation
}

// NewDerivativeAggregation provides a builder for the DerivativeAggregation struct.
func NewDerivativeAggregation() *DerivativeAggregationBuilder {
	r := DerivativeAggregationBuilder{
		&DerivativeAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the DerivativeAggregation struct
func (rb *DerivativeAggregationBuilder) Build() DerivativeAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.
func (rb *DerivativeAggregationBuilder) BucketsPath(bucketspath BucketsPath) *DerivativeAggregationBuilder {
	rb.v.BucketsPath = &bucketspath
	return rb
}

// Format set the Format property for DerivativeAggregationBuilder.
func (rb *DerivativeAggregationBuilder) Format(format string) *DerivativeAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// GapPolicy set the GapPolicy property for DerivativeAggregationBuilder.
func (rb *DerivativeAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *DerivativeAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

// Meta set the Meta property for DerivativeAggregationBuilder.
func (rb *DerivativeAggregationBuilder) Meta(meta Metadata) *DerivativeAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for DerivativeAggregationBuilder.
func (rb *DerivativeAggregationBuilder) Name(name string) *DerivativeAggregationBuilder {
	rb.v.Name = &name
	return rb
}
