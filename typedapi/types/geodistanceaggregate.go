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

// GeoDistanceAggregate type.
type GeoDistanceAggregate struct {
	Buckets BucketsRangeBucket `json:"buckets"`
	Meta    *Metadata          `json:"meta,omitempty"`
}

// GeoDistanceAggregateBuilder holds GeoDistanceAggregate struct and provides a builder API.
type GeoDistanceAggregateBuilder struct {
	v *GeoDistanceAggregate
}

// NewGeoDistanceAggregate provides a builder for the GeoDistanceAggregate struct.
func NewGeoDistanceAggregate() *GeoDistanceAggregateBuilder {
	r := GeoDistanceAggregateBuilder{
		&GeoDistanceAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the GeoDistanceAggregate struct
func (rb *GeoDistanceAggregateBuilder) Build() GeoDistanceAggregate {
	return *rb.v
}

// Buckets set the Buckets property for GeoDistanceAggregateBuilder.
func (rb *GeoDistanceAggregateBuilder) Buckets(buckets BucketsRangeBucket) *GeoDistanceAggregateBuilder {
	rb.v.Buckets = buckets
	return rb
}

// Meta set the Meta property for GeoDistanceAggregateBuilder.
func (rb *GeoDistanceAggregateBuilder) Meta(meta Metadata) *GeoDistanceAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}
