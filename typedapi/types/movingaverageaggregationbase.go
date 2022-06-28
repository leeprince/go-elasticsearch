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

// MovingAverageAggregationBase type.
type MovingAverageAggregationBase struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath         `json:"buckets_path,omitempty"`
	Format      *string              `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Meta        *Metadata            `json:"meta,omitempty"`
	Minimize    *bool                `json:"minimize,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Predict     *int                 `json:"predict,omitempty"`
	Window      *int                 `json:"window,omitempty"`
}

// MovingAverageAggregationBaseBuilder holds MovingAverageAggregationBase struct and provides a builder API.
type MovingAverageAggregationBaseBuilder struct {
	v *MovingAverageAggregationBase
}

// NewMovingAverageAggregationBase provides a builder for the MovingAverageAggregationBase struct.
func NewMovingAverageAggregationBase() *MovingAverageAggregationBaseBuilder {
	r := MovingAverageAggregationBaseBuilder{
		&MovingAverageAggregationBase{},
	}

	return &r
}

// Build finalize the chain and returns the MovingAverageAggregationBase struct
func (rb *MovingAverageAggregationBaseBuilder) Build() MovingAverageAggregationBase {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.
func (rb *MovingAverageAggregationBaseBuilder) BucketsPath(bucketspath BucketsPath) *MovingAverageAggregationBaseBuilder {
	rb.v.BucketsPath = &bucketspath
	return rb
}

// Format set the Format property for MovingAverageAggregationBaseBuilder.
func (rb *MovingAverageAggregationBaseBuilder) Format(format string) *MovingAverageAggregationBaseBuilder {
	rb.v.Format = &format
	return rb
}

// GapPolicy set the GapPolicy property for MovingAverageAggregationBaseBuilder.
func (rb *MovingAverageAggregationBaseBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *MovingAverageAggregationBaseBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

// Meta set the Meta property for MovingAverageAggregationBaseBuilder.
func (rb *MovingAverageAggregationBaseBuilder) Meta(meta Metadata) *MovingAverageAggregationBaseBuilder {
	rb.v.Meta = &meta
	return rb
}

// Minimize set the Minimize property for MovingAverageAggregationBaseBuilder.
func (rb *MovingAverageAggregationBaseBuilder) Minimize(minimize bool) *MovingAverageAggregationBaseBuilder {
	rb.v.Minimize = &minimize
	return rb
}

// Name set the Name property for MovingAverageAggregationBaseBuilder.
func (rb *MovingAverageAggregationBaseBuilder) Name(name string) *MovingAverageAggregationBaseBuilder {
	rb.v.Name = &name
	return rb
}

// Predict set the Predict property for MovingAverageAggregationBaseBuilder.
func (rb *MovingAverageAggregationBaseBuilder) Predict(predict int) *MovingAverageAggregationBaseBuilder {
	rb.v.Predict = &predict
	return rb
}

// Window set the Window property for MovingAverageAggregationBaseBuilder.
func (rb *MovingAverageAggregationBaseBuilder) Window(window int) *MovingAverageAggregationBaseBuilder {
	rb.v.Window = &window
	return rb
}
