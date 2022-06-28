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

// MovingPercentilesAggregation type.
type MovingPercentilesAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath         `json:"buckets_path,omitempty"`
	Format      *string              `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Keyed       *bool                `json:"keyed,omitempty"`
	Meta        *Metadata            `json:"meta,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Shift       *int                 `json:"shift,omitempty"`
	Window      *int                 `json:"window,omitempty"`
}

// MovingPercentilesAggregationBuilder holds MovingPercentilesAggregation struct and provides a builder API.
type MovingPercentilesAggregationBuilder struct {
	v *MovingPercentilesAggregation
}

// NewMovingPercentilesAggregation provides a builder for the MovingPercentilesAggregation struct.
func NewMovingPercentilesAggregation() *MovingPercentilesAggregationBuilder {
	r := MovingPercentilesAggregationBuilder{
		&MovingPercentilesAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the MovingPercentilesAggregation struct
func (rb *MovingPercentilesAggregationBuilder) Build() MovingPercentilesAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.
func (rb *MovingPercentilesAggregationBuilder) BucketsPath(bucketspath BucketsPath) *MovingPercentilesAggregationBuilder {
	rb.v.BucketsPath = &bucketspath
	return rb
}

// Format set the Format property for MovingPercentilesAggregationBuilder.
func (rb *MovingPercentilesAggregationBuilder) Format(format string) *MovingPercentilesAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// GapPolicy set the GapPolicy property for MovingPercentilesAggregationBuilder.
func (rb *MovingPercentilesAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *MovingPercentilesAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

// Keyed set the Keyed property for MovingPercentilesAggregationBuilder.
func (rb *MovingPercentilesAggregationBuilder) Keyed(keyed bool) *MovingPercentilesAggregationBuilder {
	rb.v.Keyed = &keyed
	return rb
}

// Meta set the Meta property for MovingPercentilesAggregationBuilder.
func (rb *MovingPercentilesAggregationBuilder) Meta(meta Metadata) *MovingPercentilesAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for MovingPercentilesAggregationBuilder.
func (rb *MovingPercentilesAggregationBuilder) Name(name string) *MovingPercentilesAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Shift set the Shift property for MovingPercentilesAggregationBuilder.
func (rb *MovingPercentilesAggregationBuilder) Shift(shift int) *MovingPercentilesAggregationBuilder {
	rb.v.Shift = &shift
	return rb
}

// Window set the Window property for MovingPercentilesAggregationBuilder.
func (rb *MovingPercentilesAggregationBuilder) Window(window int) *MovingPercentilesAggregationBuilder {
	rb.v.Window = &window
	return rb
}
