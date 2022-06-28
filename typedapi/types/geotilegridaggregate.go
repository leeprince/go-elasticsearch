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

// GeoTileGridAggregate type.
type GeoTileGridAggregate struct {
	Buckets BucketsGeoTileGridBucket `json:"buckets"`
	Meta    *Metadata                `json:"meta,omitempty"`
}

// GeoTileGridAggregateBuilder holds GeoTileGridAggregate struct and provides a builder API.
type GeoTileGridAggregateBuilder struct {
	v *GeoTileGridAggregate
}

// NewGeoTileGridAggregate provides a builder for the GeoTileGridAggregate struct.
func NewGeoTileGridAggregate() *GeoTileGridAggregateBuilder {
	r := GeoTileGridAggregateBuilder{
		&GeoTileGridAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the GeoTileGridAggregate struct
func (rb *GeoTileGridAggregateBuilder) Build() GeoTileGridAggregate {
	return *rb.v
}

// Buckets set the Buckets property for GeoTileGridAggregateBuilder.
func (rb *GeoTileGridAggregateBuilder) Buckets(buckets BucketsGeoTileGridBucket) *GeoTileGridAggregateBuilder {
	rb.v.Buckets = buckets
	return rb
}

// Meta set the Meta property for GeoTileGridAggregateBuilder.
func (rb *GeoTileGridAggregateBuilder) Meta(meta Metadata) *GeoTileGridAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}
