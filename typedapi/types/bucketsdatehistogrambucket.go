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

// BucketsDateHistogramBucket holds the union for the following types:
//     []DateHistogramBucket
//     map[string]DateHistogramBucket
type BucketsDateHistogramBucket interface{}

// BucketsDateHistogramBucketBuilder holds BucketsDateHistogramBucket struct and provides a builder API.
type BucketsDateHistogramBucketBuilder struct {
	v BucketsDateHistogramBucket
}

// NewBucketsDateHistogramBucket provides a builder for the BucketsDateHistogramBucket struct.
func NewBucketsDateHistogramBucket() *BucketsDateHistogramBucketBuilder {
	return &BucketsDateHistogramBucketBuilder{}
}

// Build finalize the chain and returns the BucketsDateHistogramBucket struct
func (u *BucketsDateHistogramBucketBuilder) Build() BucketsDateHistogramBucket {
	return u.v
}

// DateHistogramBuckets set the DateHistogramBuckets property for BucketsDateHistogramBucketBuilder.
func (u *BucketsDateHistogramBucketBuilder) DateHistogramBuckets(v []DateHistogramBucket) *BucketsDateHistogramBucketBuilder {
	u.v = v
	return u
}

// Map set the Map property for BucketsDateHistogramBucketBuilder.
func (u *BucketsDateHistogramBucketBuilder) Map(v map[string]DateHistogramBucket) *BucketsDateHistogramBucketBuilder {
	u.v = v
	return u
}
