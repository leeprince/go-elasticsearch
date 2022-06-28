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

// CumulativeCardinalityAggregation type.
type CumulativeCardinalityAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath         `json:"buckets_path,omitempty"`
	Format      *string              `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Meta        *Metadata            `json:"meta,omitempty"`
	Name        *string              `json:"name,omitempty"`
}

// CumulativeCardinalityAggregationBuilder holds CumulativeCardinalityAggregation struct and provides a builder API.
type CumulativeCardinalityAggregationBuilder struct {
	v *CumulativeCardinalityAggregation
}

// NewCumulativeCardinalityAggregation provides a builder for the CumulativeCardinalityAggregation struct.
func NewCumulativeCardinalityAggregation() *CumulativeCardinalityAggregationBuilder {
	r := CumulativeCardinalityAggregationBuilder{
		&CumulativeCardinalityAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the CumulativeCardinalityAggregation struct
func (rb *CumulativeCardinalityAggregationBuilder) Build() CumulativeCardinalityAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.
func (rb *CumulativeCardinalityAggregationBuilder) BucketsPath(bucketspath BucketsPath) *CumulativeCardinalityAggregationBuilder {
	rb.v.BucketsPath = &bucketspath
	return rb
}

// Format set the Format property for CumulativeCardinalityAggregationBuilder.
func (rb *CumulativeCardinalityAggregationBuilder) Format(format string) *CumulativeCardinalityAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// GapPolicy set the GapPolicy property for CumulativeCardinalityAggregationBuilder.
func (rb *CumulativeCardinalityAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *CumulativeCardinalityAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

// Meta set the Meta property for CumulativeCardinalityAggregationBuilder.
func (rb *CumulativeCardinalityAggregationBuilder) Meta(meta Metadata) *CumulativeCardinalityAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for CumulativeCardinalityAggregationBuilder.
func (rb *CumulativeCardinalityAggregationBuilder) Name(name string) *CumulativeCardinalityAggregationBuilder {
	rb.v.Name = &name
	return rb
}
