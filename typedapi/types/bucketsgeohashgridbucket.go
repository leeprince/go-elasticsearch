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

// BucketsGeoHashGridBucket holds the union for the following types:
//     []GeoHashGridBucket
//     map[string]GeoHashGridBucket
type BucketsGeoHashGridBucket interface{}

// BucketsGeoHashGridBucketBuilder holds BucketsGeoHashGridBucket struct and provides a builder API.
type BucketsGeoHashGridBucketBuilder struct {
	v BucketsGeoHashGridBucket
}

// NewBucketsGeoHashGridBucket provides a builder for the BucketsGeoHashGridBucket struct.
func NewBucketsGeoHashGridBucket() *BucketsGeoHashGridBucketBuilder {
	return &BucketsGeoHashGridBucketBuilder{}
}

// Build finalize the chain and returns the BucketsGeoHashGridBucket struct
func (u *BucketsGeoHashGridBucketBuilder) Build() BucketsGeoHashGridBucket {
	return u.v
}

// GeoHashGridBuckets set the GeoHashGridBuckets property for BucketsGeoHashGridBucketBuilder.
func (u *BucketsGeoHashGridBucketBuilder) GeoHashGridBuckets(v []GeoHashGridBucket) *BucketsGeoHashGridBucketBuilder {
	u.v = v
	return u
}

// Map set the Map property for BucketsGeoHashGridBucketBuilder.
func (u *BucketsGeoHashGridBucketBuilder) Map(v map[string]GeoHashGridBucket) *BucketsGeoHashGridBucketBuilder {
	u.v = v
	return u
}
