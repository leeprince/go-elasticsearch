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

// GeoCentroidAggregate type.
type GeoCentroidAggregate struct {
	Count    int64        `json:"count"`
	Location *GeoLocation `json:"location,omitempty"`
	Meta     *Metadata    `json:"meta,omitempty"`
}

// GeoCentroidAggregateBuilder holds GeoCentroidAggregate struct and provides a builder API.
type GeoCentroidAggregateBuilder struct {
	v *GeoCentroidAggregate
}

// NewGeoCentroidAggregate provides a builder for the GeoCentroidAggregate struct.
func NewGeoCentroidAggregate() *GeoCentroidAggregateBuilder {
	r := GeoCentroidAggregateBuilder{
		&GeoCentroidAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the GeoCentroidAggregate struct
func (rb *GeoCentroidAggregateBuilder) Build() GeoCentroidAggregate {
	return *rb.v
}

// Count set the Count property for GeoCentroidAggregateBuilder.
func (rb *GeoCentroidAggregateBuilder) Count(count int64) *GeoCentroidAggregateBuilder {
	rb.v.Count = count
	return rb
}

// Location set the Location property for GeoCentroidAggregateBuilder.
func (rb *GeoCentroidAggregateBuilder) Location(location GeoLocation) *GeoCentroidAggregateBuilder {
	rb.v.Location = &location
	return rb
}

// Meta set the Meta property for GeoCentroidAggregateBuilder.
func (rb *GeoCentroidAggregateBuilder) Meta(meta Metadata) *GeoCentroidAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}
