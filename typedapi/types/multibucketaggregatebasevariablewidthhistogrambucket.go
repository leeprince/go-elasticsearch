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

// MultiBucketAggregateBaseVariableWidthHistogramBucket type.
type MultiBucketAggregateBaseVariableWidthHistogramBucket struct {
	Buckets BucketsVariableWidthHistogramBucket `json:"buckets"`
	Meta    *Metadata                           `json:"meta,omitempty"`
}

// MultiBucketAggregateBaseVariableWidthHistogramBucketBuilder holds MultiBucketAggregateBaseVariableWidthHistogramBucket struct and provides a builder API.
type MultiBucketAggregateBaseVariableWidthHistogramBucketBuilder struct {
	v *MultiBucketAggregateBaseVariableWidthHistogramBucket
}

// NewMultiBucketAggregateBaseVariableWidthHistogramBucket provides a builder for the MultiBucketAggregateBaseVariableWidthHistogramBucket struct.
func NewMultiBucketAggregateBaseVariableWidthHistogramBucket() *MultiBucketAggregateBaseVariableWidthHistogramBucketBuilder {
	r := MultiBucketAggregateBaseVariableWidthHistogramBucketBuilder{
		&MultiBucketAggregateBaseVariableWidthHistogramBucket{},
	}

	return &r
}

// Build finalize the chain and returns the MultiBucketAggregateBaseVariableWidthHistogramBucket struct
func (rb *MultiBucketAggregateBaseVariableWidthHistogramBucketBuilder) Build() MultiBucketAggregateBaseVariableWidthHistogramBucket {
	return *rb.v
}

// Buckets set the Buckets property for MultiBucketAggregateBaseVariableWidthHistogramBucketBuilder.
func (rb *MultiBucketAggregateBaseVariableWidthHistogramBucketBuilder) Buckets(buckets BucketsVariableWidthHistogramBucket) *MultiBucketAggregateBaseVariableWidthHistogramBucketBuilder {
	rb.v.Buckets = buckets
	return rb
}

// Meta set the Meta property for MultiBucketAggregateBaseVariableWidthHistogramBucketBuilder.
func (rb *MultiBucketAggregateBaseVariableWidthHistogramBucketBuilder) Meta(meta Metadata) *MultiBucketAggregateBaseVariableWidthHistogramBucketBuilder {
	rb.v.Meta = &meta
	return rb
}
