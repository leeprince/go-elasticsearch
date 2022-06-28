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

// MovingFunctionAggregation type.
type MovingFunctionAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath         `json:"buckets_path,omitempty"`
	Format      *string              `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Meta        *Metadata            `json:"meta,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Script      *string              `json:"script,omitempty"`
	Shift       *int                 `json:"shift,omitempty"`
	Window      *int                 `json:"window,omitempty"`
}

// MovingFunctionAggregationBuilder holds MovingFunctionAggregation struct and provides a builder API.
type MovingFunctionAggregationBuilder struct {
	v *MovingFunctionAggregation
}

// NewMovingFunctionAggregation provides a builder for the MovingFunctionAggregation struct.
func NewMovingFunctionAggregation() *MovingFunctionAggregationBuilder {
	r := MovingFunctionAggregationBuilder{
		&MovingFunctionAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the MovingFunctionAggregation struct
func (rb *MovingFunctionAggregationBuilder) Build() MovingFunctionAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.
func (rb *MovingFunctionAggregationBuilder) BucketsPath(bucketspath BucketsPath) *MovingFunctionAggregationBuilder {
	rb.v.BucketsPath = &bucketspath
	return rb
}

// Format set the Format property for MovingFunctionAggregationBuilder.
func (rb *MovingFunctionAggregationBuilder) Format(format string) *MovingFunctionAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// GapPolicy set the GapPolicy property for MovingFunctionAggregationBuilder.
func (rb *MovingFunctionAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *MovingFunctionAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

// Meta set the Meta property for MovingFunctionAggregationBuilder.
func (rb *MovingFunctionAggregationBuilder) Meta(meta Metadata) *MovingFunctionAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for MovingFunctionAggregationBuilder.
func (rb *MovingFunctionAggregationBuilder) Name(name string) *MovingFunctionAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Script set the Script property for MovingFunctionAggregationBuilder.
func (rb *MovingFunctionAggregationBuilder) Script(script string) *MovingFunctionAggregationBuilder {
	rb.v.Script = &script
	return rb
}

// Shift set the Shift property for MovingFunctionAggregationBuilder.
func (rb *MovingFunctionAggregationBuilder) Shift(shift int) *MovingFunctionAggregationBuilder {
	rb.v.Shift = &shift
	return rb
}

// Window set the Window property for MovingFunctionAggregationBuilder.
func (rb *MovingFunctionAggregationBuilder) Window(window int) *MovingFunctionAggregationBuilder {
	rb.v.Window = &window
	return rb
}
