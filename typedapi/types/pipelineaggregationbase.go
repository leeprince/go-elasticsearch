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

// PipelineAggregationBase type.
type PipelineAggregationBase struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath         `json:"buckets_path,omitempty"`
	Format      *string              `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Meta        *Metadata            `json:"meta,omitempty"`
	Name        *string              `json:"name,omitempty"`
}

// PipelineAggregationBaseBuilder holds PipelineAggregationBase struct and provides a builder API.
type PipelineAggregationBaseBuilder struct {
	v *PipelineAggregationBase
}

// NewPipelineAggregationBase provides a builder for the PipelineAggregationBase struct.
func NewPipelineAggregationBase() *PipelineAggregationBaseBuilder {
	r := PipelineAggregationBaseBuilder{
		&PipelineAggregationBase{},
	}

	return &r
}

// Build finalize the chain and returns the PipelineAggregationBase struct
func (rb *PipelineAggregationBaseBuilder) Build() PipelineAggregationBase {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.
func (rb *PipelineAggregationBaseBuilder) BucketsPath(bucketspath BucketsPath) *PipelineAggregationBaseBuilder {
	rb.v.BucketsPath = &bucketspath
	return rb
}

// Format set the Format property for PipelineAggregationBaseBuilder.
func (rb *PipelineAggregationBaseBuilder) Format(format string) *PipelineAggregationBaseBuilder {
	rb.v.Format = &format
	return rb
}

// GapPolicy set the GapPolicy property for PipelineAggregationBaseBuilder.
func (rb *PipelineAggregationBaseBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *PipelineAggregationBaseBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

// Meta set the Meta property for PipelineAggregationBaseBuilder.
func (rb *PipelineAggregationBaseBuilder) Meta(meta Metadata) *PipelineAggregationBaseBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for PipelineAggregationBaseBuilder.
func (rb *PipelineAggregationBaseBuilder) Name(name string) *PipelineAggregationBaseBuilder {
	rb.v.Name = &name
	return rb
}
