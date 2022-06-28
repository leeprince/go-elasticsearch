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

// BucketsVariableWidthHistogramBucket holds the union for the following types:
//     map[string]VariableWidthHistogramBucket
//     []VariableWidthHistogramBucket
type BucketsVariableWidthHistogramBucket interface{}

// BucketsVariableWidthHistogramBucketBuilder holds BucketsVariableWidthHistogramBucket struct and provides a builder API.
type BucketsVariableWidthHistogramBucketBuilder struct {
	v BucketsVariableWidthHistogramBucket
}

// NewBucketsVariableWidthHistogramBucket provides a builder for the BucketsVariableWidthHistogramBucket struct.
func NewBucketsVariableWidthHistogramBucket() *BucketsVariableWidthHistogramBucketBuilder {
	return &BucketsVariableWidthHistogramBucketBuilder{}
}

// Build finalize the chain and returns the BucketsVariableWidthHistogramBucket struct
func (u *BucketsVariableWidthHistogramBucketBuilder) Build() BucketsVariableWidthHistogramBucket {
	return u.v
}

// Map set the Map property for BucketsVariableWidthHistogramBucketBuilder.
func (u *BucketsVariableWidthHistogramBucketBuilder) Map(v map[string]VariableWidthHistogramBucket) *BucketsVariableWidthHistogramBucketBuilder {
	u.v = v
	return u
}

// VariableWidthHistogramBuckets set the VariableWidthHistogramBuckets property for BucketsVariableWidthHistogramBucketBuilder.
func (u *BucketsVariableWidthHistogramBucketBuilder) VariableWidthHistogramBuckets(v []VariableWidthHistogramBucket) *BucketsVariableWidthHistogramBucketBuilder {
	u.v = v
	return u
}
