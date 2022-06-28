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

// BucketsMultiTermsBucket holds the union for the following types:
//     map[string]MultiTermsBucket
//     []MultiTermsBucket
type BucketsMultiTermsBucket interface{}

// BucketsMultiTermsBucketBuilder holds BucketsMultiTermsBucket struct and provides a builder API.
type BucketsMultiTermsBucketBuilder struct {
	v BucketsMultiTermsBucket
}

// NewBucketsMultiTermsBucket provides a builder for the BucketsMultiTermsBucket struct.
func NewBucketsMultiTermsBucket() *BucketsMultiTermsBucketBuilder {
	return &BucketsMultiTermsBucketBuilder{}
}

// Build finalize the chain and returns the BucketsMultiTermsBucket struct
func (u *BucketsMultiTermsBucketBuilder) Build() BucketsMultiTermsBucket {
	return u.v
}

// Map set the Map property for BucketsMultiTermsBucketBuilder.
func (u *BucketsMultiTermsBucketBuilder) Map(v map[string]MultiTermsBucket) *BucketsMultiTermsBucketBuilder {
	u.v = v
	return u
}

// MultiTermsBuckets set the MultiTermsBuckets property for BucketsMultiTermsBucketBuilder.
func (u *BucketsMultiTermsBucketBuilder) MultiTermsBuckets(v []MultiTermsBucket) *BucketsMultiTermsBucketBuilder {
	u.v = v
	return u
}
