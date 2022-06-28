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

// MultiBucketAggregateBaseAdjacencyMatrixBucket type.
type MultiBucketAggregateBaseAdjacencyMatrixBucket struct {
	Buckets BucketsAdjacencyMatrixBucket `json:"buckets"`
	Meta    *Metadata                    `json:"meta,omitempty"`
}

// MultiBucketAggregateBaseAdjacencyMatrixBucketBuilder holds MultiBucketAggregateBaseAdjacencyMatrixBucket struct and provides a builder API.
type MultiBucketAggregateBaseAdjacencyMatrixBucketBuilder struct {
	v *MultiBucketAggregateBaseAdjacencyMatrixBucket
}

// NewMultiBucketAggregateBaseAdjacencyMatrixBucket provides a builder for the MultiBucketAggregateBaseAdjacencyMatrixBucket struct.
func NewMultiBucketAggregateBaseAdjacencyMatrixBucket() *MultiBucketAggregateBaseAdjacencyMatrixBucketBuilder {
	r := MultiBucketAggregateBaseAdjacencyMatrixBucketBuilder{
		&MultiBucketAggregateBaseAdjacencyMatrixBucket{},
	}

	return &r
}

// Build finalize the chain and returns the MultiBucketAggregateBaseAdjacencyMatrixBucket struct
func (rb *MultiBucketAggregateBaseAdjacencyMatrixBucketBuilder) Build() MultiBucketAggregateBaseAdjacencyMatrixBucket {
	return *rb.v
}

// Buckets set the Buckets property for MultiBucketAggregateBaseAdjacencyMatrixBucketBuilder.
func (rb *MultiBucketAggregateBaseAdjacencyMatrixBucketBuilder) Buckets(buckets BucketsAdjacencyMatrixBucket) *MultiBucketAggregateBaseAdjacencyMatrixBucketBuilder {
	rb.v.Buckets = buckets
	return rb
}

// Meta set the Meta property for MultiBucketAggregateBaseAdjacencyMatrixBucketBuilder.
func (rb *MultiBucketAggregateBaseAdjacencyMatrixBucketBuilder) Meta(meta Metadata) *MultiBucketAggregateBaseAdjacencyMatrixBucketBuilder {
	rb.v.Meta = &meta
	return rb
}
