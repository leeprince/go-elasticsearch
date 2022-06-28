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

// BucketsCompositeBucket holds the union for the following types:
//     []CompositeBucket
//     map[string]CompositeBucket
type BucketsCompositeBucket interface{}

// BucketsCompositeBucketBuilder holds BucketsCompositeBucket struct and provides a builder API.
type BucketsCompositeBucketBuilder struct {
	v BucketsCompositeBucket
}

// NewBucketsCompositeBucket provides a builder for the BucketsCompositeBucket struct.
func NewBucketsCompositeBucket() *BucketsCompositeBucketBuilder {
	return &BucketsCompositeBucketBuilder{}
}

// Build finalize the chain and returns the BucketsCompositeBucket struct
func (u *BucketsCompositeBucketBuilder) Build() BucketsCompositeBucket {
	return u.v
}

// CompositeBuckets set the CompositeBuckets property for BucketsCompositeBucketBuilder.
func (u *BucketsCompositeBucketBuilder) CompositeBuckets(v []CompositeBucket) *BucketsCompositeBucketBuilder {
	u.v = v
	return u
}

// Map set the Map property for BucketsCompositeBucketBuilder.
func (u *BucketsCompositeBucketBuilder) Map(v map[string]CompositeBucket) *BucketsCompositeBucketBuilder {
	u.v = v
	return u
}
