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

// BucketPathAggregation type.
type BucketPathAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath `json:"buckets_path,omitempty"`
	Meta        *Metadata    `json:"meta,omitempty"`
	Name        *string      `json:"name,omitempty"`
}

// BucketPathAggregationBuilder holds BucketPathAggregation struct and provides a builder API.
type BucketPathAggregationBuilder struct {
	v *BucketPathAggregation
}

// NewBucketPathAggregation provides a builder for the BucketPathAggregation struct.
func NewBucketPathAggregation() *BucketPathAggregationBuilder {
	r := BucketPathAggregationBuilder{
		&BucketPathAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the BucketPathAggregation struct
func (rb *BucketPathAggregationBuilder) Build() BucketPathAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.
func (rb *BucketPathAggregationBuilder) BucketsPath(bucketspath BucketsPath) *BucketPathAggregationBuilder {
	rb.v.BucketsPath = &bucketspath
	return rb
}

// Meta set the Meta property for BucketPathAggregationBuilder.
func (rb *BucketPathAggregationBuilder) Meta(meta Metadata) *BucketPathAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for BucketPathAggregationBuilder.
func (rb *BucketPathAggregationBuilder) Name(name string) *BucketPathAggregationBuilder {
	rb.v.Name = &name
	return rb
}
