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

// InferenceAggregation type.
type InferenceAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath     *BucketsPath              `json:"buckets_path,omitempty"`
	Format          *string                   `json:"format,omitempty"`
	GapPolicy       *gappolicy.GapPolicy      `json:"gap_policy,omitempty"`
	InferenceConfig *InferenceConfigContainer `json:"inference_config,omitempty"`
	Meta            *Metadata                 `json:"meta,omitempty"`
	ModelId         Name                      `json:"model_id"`
	Name            *string                   `json:"name,omitempty"`
}

// InferenceAggregationBuilder holds InferenceAggregation struct and provides a builder API.
type InferenceAggregationBuilder struct {
	v *InferenceAggregation
}

// NewInferenceAggregation provides a builder for the InferenceAggregation struct.
func NewInferenceAggregation() *InferenceAggregationBuilder {
	r := InferenceAggregationBuilder{
		&InferenceAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the InferenceAggregation struct
func (rb *InferenceAggregationBuilder) Build() InferenceAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.
func (rb *InferenceAggregationBuilder) BucketsPath(bucketspath BucketsPath) *InferenceAggregationBuilder {
	rb.v.BucketsPath = &bucketspath
	return rb
}

// Format set the Format property for InferenceAggregationBuilder.
func (rb *InferenceAggregationBuilder) Format(format string) *InferenceAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// GapPolicy set the GapPolicy property for InferenceAggregationBuilder.
func (rb *InferenceAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *InferenceAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

// InferenceConfig set the InferenceConfig property for InferenceAggregationBuilder.
func (rb *InferenceAggregationBuilder) InferenceConfig(inferenceconfig InferenceConfigContainer) *InferenceAggregationBuilder {
	rb.v.InferenceConfig = &inferenceconfig
	return rb
}

// Meta set the Meta property for InferenceAggregationBuilder.
func (rb *InferenceAggregationBuilder) Meta(meta Metadata) *InferenceAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// ModelId set the ModelId property for InferenceAggregationBuilder.
func (rb *InferenceAggregationBuilder) ModelId(modelid Name) *InferenceAggregationBuilder {
	rb.v.ModelId = modelid
	return rb
}

// Name set the Name property for InferenceAggregationBuilder.
func (rb *InferenceAggregationBuilder) Name(name string) *InferenceAggregationBuilder {
	rb.v.Name = &name
	return rb
}
