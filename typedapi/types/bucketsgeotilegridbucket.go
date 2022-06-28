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

// BucketsGeoTileGridBucket holds the union for the following types:
//     []GeoTileGridBucket
//     map[string]GeoTileGridBucket
type BucketsGeoTileGridBucket interface{}

// BucketsGeoTileGridBucketBuilder holds BucketsGeoTileGridBucket struct and provides a builder API.
type BucketsGeoTileGridBucketBuilder struct {
	v BucketsGeoTileGridBucket
}

// NewBucketsGeoTileGridBucket provides a builder for the BucketsGeoTileGridBucket struct.
func NewBucketsGeoTileGridBucket() *BucketsGeoTileGridBucketBuilder {
	return &BucketsGeoTileGridBucketBuilder{}
}

// Build finalize the chain and returns the BucketsGeoTileGridBucket struct
func (u *BucketsGeoTileGridBucketBuilder) Build() BucketsGeoTileGridBucket {
	return u.v
}

// GeoTileGridBuckets set the GeoTileGridBuckets property for BucketsGeoTileGridBucketBuilder.
func (u *BucketsGeoTileGridBucketBuilder) GeoTileGridBuckets(v []GeoTileGridBucket) *BucketsGeoTileGridBucketBuilder {
	u.v = v
	return u
}

// Map set the Map property for BucketsGeoTileGridBucketBuilder.
func (u *BucketsGeoTileGridBucketBuilder) Map(v map[string]GeoTileGridBucket) *BucketsGeoTileGridBucketBuilder {
	u.v = v
	return u
}
