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

// GeoBoundsAggregate type.
type GeoBoundsAggregate struct {
	Bounds GeoBounds `json:"bounds"`
	Meta   *Metadata `json:"meta,omitempty"`
}

// GeoBoundsAggregateBuilder holds GeoBoundsAggregate struct and provides a builder API.
type GeoBoundsAggregateBuilder struct {
	v *GeoBoundsAggregate
}

// NewGeoBoundsAggregate provides a builder for the GeoBoundsAggregate struct.
func NewGeoBoundsAggregate() *GeoBoundsAggregateBuilder {
	r := GeoBoundsAggregateBuilder{
		&GeoBoundsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the GeoBoundsAggregate struct
func (rb *GeoBoundsAggregateBuilder) Build() GeoBoundsAggregate {
	return *rb.v
}

// Bounds set the Bounds property for GeoBoundsAggregateBuilder.
func (rb *GeoBoundsAggregateBuilder) Bounds(bounds GeoBounds) *GeoBoundsAggregateBuilder {
	rb.v.Bounds = bounds
	return rb
}

// Meta set the Meta property for GeoBoundsAggregateBuilder.
func (rb *GeoBoundsAggregateBuilder) Meta(meta Metadata) *GeoBoundsAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}
