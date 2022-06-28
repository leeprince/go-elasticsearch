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

// BucketsLongRareTermsBucket holds the union for the following types:
//     []LongRareTermsBucket
//     map[string]LongRareTermsBucket
type BucketsLongRareTermsBucket interface{}

// BucketsLongRareTermsBucketBuilder holds BucketsLongRareTermsBucket struct and provides a builder API.
type BucketsLongRareTermsBucketBuilder struct {
	v BucketsLongRareTermsBucket
}

// NewBucketsLongRareTermsBucket provides a builder for the BucketsLongRareTermsBucket struct.
func NewBucketsLongRareTermsBucket() *BucketsLongRareTermsBucketBuilder {
	return &BucketsLongRareTermsBucketBuilder{}
}

// Build finalize the chain and returns the BucketsLongRareTermsBucket struct
func (u *BucketsLongRareTermsBucketBuilder) Build() BucketsLongRareTermsBucket {
	return u.v
}

// LongRareTermsBuckets set the LongRareTermsBuckets property for BucketsLongRareTermsBucketBuilder.
func (u *BucketsLongRareTermsBucketBuilder) LongRareTermsBuckets(v []LongRareTermsBucket) *BucketsLongRareTermsBucketBuilder {
	u.v = v
	return u
}

// Map set the Map property for BucketsLongRareTermsBucketBuilder.
func (u *BucketsLongRareTermsBucketBuilder) Map(v map[string]LongRareTermsBucket) *BucketsLongRareTermsBucketBuilder {
	u.v = v
	return u
}
