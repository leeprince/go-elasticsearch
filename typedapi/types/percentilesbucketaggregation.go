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

// PercentilesBucketAggregation type.
type PercentilesBucketAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath         `json:"buckets_path,omitempty"`
	Format      *string              `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Meta        *Metadata            `json:"meta,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Percents    []float64            `json:"percents,omitempty"`
}

// PercentilesBucketAggregationBuilder holds PercentilesBucketAggregation struct and provides a builder API.
type PercentilesBucketAggregationBuilder struct {
	v *PercentilesBucketAggregation
}

// NewPercentilesBucketAggregation provides a builder for the PercentilesBucketAggregation struct.
func NewPercentilesBucketAggregation() *PercentilesBucketAggregationBuilder {
	r := PercentilesBucketAggregationBuilder{
		&PercentilesBucketAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the PercentilesBucketAggregation struct
func (rb *PercentilesBucketAggregationBuilder) Build() PercentilesBucketAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.
func (rb *PercentilesBucketAggregationBuilder) BucketsPath(bucketspath BucketsPath) *PercentilesBucketAggregationBuilder {
	rb.v.BucketsPath = &bucketspath
	return rb
}

// Format set the Format property for PercentilesBucketAggregationBuilder.
func (rb *PercentilesBucketAggregationBuilder) Format(format string) *PercentilesBucketAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// GapPolicy set the GapPolicy property for PercentilesBucketAggregationBuilder.
func (rb *PercentilesBucketAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *PercentilesBucketAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

// Meta set the Meta property for PercentilesBucketAggregationBuilder.
func (rb *PercentilesBucketAggregationBuilder) Meta(meta Metadata) *PercentilesBucketAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for PercentilesBucketAggregationBuilder.
func (rb *PercentilesBucketAggregationBuilder) Name(name string) *PercentilesBucketAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Percents set the Percents property for PercentilesBucketAggregationBuilder.
func (rb *PercentilesBucketAggregationBuilder) Percents(percents ...float64) *PercentilesBucketAggregationBuilder {
	rb.v.Percents = append(rb.v.Percents, percents...)
	return rb
}
