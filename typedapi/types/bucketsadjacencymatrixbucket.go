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

// BucketsAdjacencyMatrixBucket holds the union for the following types:
//     []AdjacencyMatrixBucket
//     map[string]AdjacencyMatrixBucket
type BucketsAdjacencyMatrixBucket interface{}

// BucketsAdjacencyMatrixBucketBuilder holds BucketsAdjacencyMatrixBucket struct and provides a builder API.
type BucketsAdjacencyMatrixBucketBuilder struct {
	v BucketsAdjacencyMatrixBucket
}

// NewBucketsAdjacencyMatrixBucket provides a builder for the BucketsAdjacencyMatrixBucket struct.
func NewBucketsAdjacencyMatrixBucket() *BucketsAdjacencyMatrixBucketBuilder {
	return &BucketsAdjacencyMatrixBucketBuilder{}
}

// Build finalize the chain and returns the BucketsAdjacencyMatrixBucket struct
func (u *BucketsAdjacencyMatrixBucketBuilder) Build() BucketsAdjacencyMatrixBucket {
	return u.v
}

// AdjacencyMatrixBuckets set the AdjacencyMatrixBuckets property for BucketsAdjacencyMatrixBucketBuilder.
func (u *BucketsAdjacencyMatrixBucketBuilder) AdjacencyMatrixBuckets(v []AdjacencyMatrixBucket) *BucketsAdjacencyMatrixBucketBuilder {
	u.v = v
	return u
}

// Map set the Map property for BucketsAdjacencyMatrixBucketBuilder.
func (u *BucketsAdjacencyMatrixBucketBuilder) Map(v map[string]AdjacencyMatrixBucket) *BucketsAdjacencyMatrixBucketBuilder {
	u.v = v
	return u
}
