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

// EwmaMovingAverageAggregation type.
type EwmaMovingAverageAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath         `json:"buckets_path,omitempty"`
	Format      *string              `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Meta        *Metadata            `json:"meta,omitempty"`
	Minimize    *bool                `json:"minimize,omitempty"`
	Model       string               `json:"model,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Predict     *int                 `json:"predict,omitempty"`
	Settings    EwmaModelSettings    `json:"settings"`
	Window      *int                 `json:"window,omitempty"`
}

// EwmaMovingAverageAggregationBuilder holds EwmaMovingAverageAggregation struct and provides a builder API.
type EwmaMovingAverageAggregationBuilder struct {
	v *EwmaMovingAverageAggregation
}

// NewEwmaMovingAverageAggregation provides a builder for the EwmaMovingAverageAggregation struct.
func NewEwmaMovingAverageAggregation() *EwmaMovingAverageAggregationBuilder {
	r := EwmaMovingAverageAggregationBuilder{
		&EwmaMovingAverageAggregation{},
	}

	r.v.Model = "ewma"

	return &r
}

// Build finalize the chain and returns the EwmaMovingAverageAggregation struct
func (rb *EwmaMovingAverageAggregationBuilder) Build() EwmaMovingAverageAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.
func (rb *EwmaMovingAverageAggregationBuilder) BucketsPath(bucketspath BucketsPath) *EwmaMovingAverageAggregationBuilder {
	rb.v.BucketsPath = &bucketspath
	return rb
}

// Format set the Format property for EwmaMovingAverageAggregationBuilder.
func (rb *EwmaMovingAverageAggregationBuilder) Format(format string) *EwmaMovingAverageAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// GapPolicy set the GapPolicy property for EwmaMovingAverageAggregationBuilder.
func (rb *EwmaMovingAverageAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *EwmaMovingAverageAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

// Meta set the Meta property for EwmaMovingAverageAggregationBuilder.
func (rb *EwmaMovingAverageAggregationBuilder) Meta(meta Metadata) *EwmaMovingAverageAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Minimize set the Minimize property for EwmaMovingAverageAggregationBuilder.
func (rb *EwmaMovingAverageAggregationBuilder) Minimize(minimize bool) *EwmaMovingAverageAggregationBuilder {
	rb.v.Minimize = &minimize
	return rb
}

// Model set the Model property for EwmaMovingAverageAggregationBuilder.

// Name set the Name property for EwmaMovingAverageAggregationBuilder.
func (rb *EwmaMovingAverageAggregationBuilder) Name(name string) *EwmaMovingAverageAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Predict set the Predict property for EwmaMovingAverageAggregationBuilder.
func (rb *EwmaMovingAverageAggregationBuilder) Predict(predict int) *EwmaMovingAverageAggregationBuilder {
	rb.v.Predict = &predict
	return rb
}

// Settings set the Settings property for EwmaMovingAverageAggregationBuilder.
func (rb *EwmaMovingAverageAggregationBuilder) Settings(settings EwmaModelSettings) *EwmaMovingAverageAggregationBuilder {
	rb.v.Settings = settings
	return rb
}

// Window set the Window property for EwmaMovingAverageAggregationBuilder.
func (rb *EwmaMovingAverageAggregationBuilder) Window(window int) *EwmaMovingAverageAggregationBuilder {
	rb.v.Window = &window
	return rb
}
