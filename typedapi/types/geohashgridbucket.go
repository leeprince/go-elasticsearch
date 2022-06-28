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

// GeoHashGridBucket type.
type GeoHashGridBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount     int64                       `json:"doc_count"`
	Key          GeoHash                     `json:"key"`
}

// GeoHashGridBucketBuilder holds GeoHashGridBucket struct and provides a builder API.
type GeoHashGridBucketBuilder struct {
	v *GeoHashGridBucket
}

// NewGeoHashGridBucket provides a builder for the GeoHashGridBucket struct.
func NewGeoHashGridBucket() *GeoHashGridBucketBuilder {
	r := GeoHashGridBucketBuilder{
		&GeoHashGridBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the GeoHashGridBucket struct
func (rb *GeoHashGridBucketBuilder) Build() GeoHashGridBucket {
	return *rb.v
}

// Aggregations set the Aggregations property for GeoHashGridBucketBuilder.
func (rb *GeoHashGridBucketBuilder) Aggregations(value map[AggregateName]Aggregate) *GeoHashGridBucketBuilder {
	rb.v.Aggregations = value
	return rb
}

// DocCount set the DocCount property for GeoHashGridBucketBuilder.
func (rb *GeoHashGridBucketBuilder) DocCount(doccount int64) *GeoHashGridBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

// Key set the Key property for GeoHashGridBucketBuilder.
func (rb *GeoHashGridBucketBuilder) Key(key GeoHash) *GeoHashGridBucketBuilder {
	rb.v.Key = key
	return rb
}
