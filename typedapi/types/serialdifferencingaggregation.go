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

// SerialDifferencingAggregation type.
type SerialDifferencingAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath         `json:"buckets_path,omitempty"`
	Format      *string              `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Lag         *int                 `json:"lag,omitempty"`
	Meta        *Metadata            `json:"meta,omitempty"`
	Name        *string              `json:"name,omitempty"`
}

// SerialDifferencingAggregationBuilder holds SerialDifferencingAggregation struct and provides a builder API.
type SerialDifferencingAggregationBuilder struct {
	v *SerialDifferencingAggregation
}

// NewSerialDifferencingAggregation provides a builder for the SerialDifferencingAggregation struct.
func NewSerialDifferencingAggregation() *SerialDifferencingAggregationBuilder {
	r := SerialDifferencingAggregationBuilder{
		&SerialDifferencingAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the SerialDifferencingAggregation struct
func (rb *SerialDifferencingAggregationBuilder) Build() SerialDifferencingAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.
func (rb *SerialDifferencingAggregationBuilder) BucketsPath(bucketspath BucketsPath) *SerialDifferencingAggregationBuilder {
	rb.v.BucketsPath = &bucketspath
	return rb
}

// Format set the Format property for SerialDifferencingAggregationBuilder.
func (rb *SerialDifferencingAggregationBuilder) Format(format string) *SerialDifferencingAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// GapPolicy set the GapPolicy property for SerialDifferencingAggregationBuilder.
func (rb *SerialDifferencingAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *SerialDifferencingAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

// Lag set the Lag property for SerialDifferencingAggregationBuilder.
func (rb *SerialDifferencingAggregationBuilder) Lag(lag int) *SerialDifferencingAggregationBuilder {
	rb.v.Lag = &lag
	return rb
}

// Meta set the Meta property for SerialDifferencingAggregationBuilder.
func (rb *SerialDifferencingAggregationBuilder) Meta(meta Metadata) *SerialDifferencingAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for SerialDifferencingAggregationBuilder.
func (rb *SerialDifferencingAggregationBuilder) Name(name string) *SerialDifferencingAggregationBuilder {
	rb.v.Name = &name
	return rb
}
