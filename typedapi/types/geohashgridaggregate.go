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

// GeoHashGridAggregate type.
type GeoHashGridAggregate struct {
	Buckets BucketsGeoHashGridBucket `json:"buckets"`
	Meta    *Metadata                `json:"meta,omitempty"`
}

// GeoHashGridAggregateBuilder holds GeoHashGridAggregate struct and provides a builder API.
type GeoHashGridAggregateBuilder struct {
	v *GeoHashGridAggregate
}

// NewGeoHashGridAggregate provides a builder for the GeoHashGridAggregate struct.
func NewGeoHashGridAggregate() *GeoHashGridAggregateBuilder {
	r := GeoHashGridAggregateBuilder{
		&GeoHashGridAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the GeoHashGridAggregate struct
func (rb *GeoHashGridAggregateBuilder) Build() GeoHashGridAggregate {
	return *rb.v
}

// Buckets set the Buckets property for GeoHashGridAggregateBuilder.
func (rb *GeoHashGridAggregateBuilder) Buckets(buckets BucketsGeoHashGridBucket) *GeoHashGridAggregateBuilder {
	rb.v.Buckets = buckets
	return rb
}

// Meta set the Meta property for GeoHashGridAggregateBuilder.
func (rb *GeoHashGridAggregateBuilder) Meta(meta Metadata) *GeoHashGridAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}
