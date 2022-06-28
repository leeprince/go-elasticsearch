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

// SumBucketAggregation type.
type SumBucketAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath         `json:"buckets_path,omitempty"`
	Format      *string              `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Meta        *Metadata            `json:"meta,omitempty"`
	Name        *string              `json:"name,omitempty"`
}

// SumBucketAggregationBuilder holds SumBucketAggregation struct and provides a builder API.
type SumBucketAggregationBuilder struct {
	v *SumBucketAggregation
}

// NewSumBucketAggregation provides a builder for the SumBucketAggregation struct.
func NewSumBucketAggregation() *SumBucketAggregationBuilder {
	r := SumBucketAggregationBuilder{
		&SumBucketAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the SumBucketAggregation struct
func (rb *SumBucketAggregationBuilder) Build() SumBucketAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.
func (rb *SumBucketAggregationBuilder) BucketsPath(bucketspath BucketsPath) *SumBucketAggregationBuilder {
	rb.v.BucketsPath = &bucketspath
	return rb
}

// Format set the Format property for SumBucketAggregationBuilder.
func (rb *SumBucketAggregationBuilder) Format(format string) *SumBucketAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// GapPolicy set the GapPolicy property for SumBucketAggregationBuilder.
func (rb *SumBucketAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *SumBucketAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

// Meta set the Meta property for SumBucketAggregationBuilder.
func (rb *SumBucketAggregationBuilder) Meta(meta Metadata) *SumBucketAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for SumBucketAggregationBuilder.
func (rb *SumBucketAggregationBuilder) Name(name string) *SumBucketAggregationBuilder {
	rb.v.Name = &name
	return rb
}
