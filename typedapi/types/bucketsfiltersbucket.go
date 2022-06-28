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

// BucketsFiltersBucket holds the union for the following types:
//     []FiltersBucket
//     map[string]FiltersBucket
type BucketsFiltersBucket interface{}

// BucketsFiltersBucketBuilder holds BucketsFiltersBucket struct and provides a builder API.
type BucketsFiltersBucketBuilder struct {
	v BucketsFiltersBucket
}

// NewBucketsFiltersBucket provides a builder for the BucketsFiltersBucket struct.
func NewBucketsFiltersBucket() *BucketsFiltersBucketBuilder {
	return &BucketsFiltersBucketBuilder{}
}

// Build finalize the chain and returns the BucketsFiltersBucket struct
func (u *BucketsFiltersBucketBuilder) Build() BucketsFiltersBucket {
	return u.v
}

// FiltersBuckets set the FiltersBuckets property for BucketsFiltersBucketBuilder.
func (u *BucketsFiltersBucketBuilder) FiltersBuckets(v []FiltersBucket) *BucketsFiltersBucketBuilder {
	u.v = v
	return u
}

// Map set the Map property for BucketsFiltersBucketBuilder.
func (u *BucketsFiltersBucketBuilder) Map(v map[string]FiltersBucket) *BucketsFiltersBucketBuilder {
	u.v = v
	return u
}
