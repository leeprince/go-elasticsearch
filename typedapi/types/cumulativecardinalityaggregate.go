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

// CumulativeCardinalityAggregate type.
type CumulativeCardinalityAggregate struct {
	Meta          *Metadata `json:"meta,omitempty"`
	Value         int64     `json:"value"`
	ValueAsString *string   `json:"value_as_string,omitempty"`
}

// CumulativeCardinalityAggregateBuilder holds CumulativeCardinalityAggregate struct and provides a builder API.
type CumulativeCardinalityAggregateBuilder struct {
	v *CumulativeCardinalityAggregate
}

// NewCumulativeCardinalityAggregate provides a builder for the CumulativeCardinalityAggregate struct.
func NewCumulativeCardinalityAggregate() *CumulativeCardinalityAggregateBuilder {
	r := CumulativeCardinalityAggregateBuilder{
		&CumulativeCardinalityAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the CumulativeCardinalityAggregate struct
func (rb *CumulativeCardinalityAggregateBuilder) Build() CumulativeCardinalityAggregate {
	return *rb.v
}

// Meta set the Meta property for CumulativeCardinalityAggregateBuilder.
func (rb *CumulativeCardinalityAggregateBuilder) Meta(meta Metadata) *CumulativeCardinalityAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}

// Value set the Value property for CumulativeCardinalityAggregateBuilder.
func (rb *CumulativeCardinalityAggregateBuilder) Value(value int64) *CumulativeCardinalityAggregateBuilder {
	rb.v.Value = value
	return rb
}

// ValueAsString set the ValueAsString property for CumulativeCardinalityAggregateBuilder.
func (rb *CumulativeCardinalityAggregateBuilder) ValueAsString(valueasstring string) *CumulativeCardinalityAggregateBuilder {
	rb.v.ValueAsString = &valueasstring
	return rb
}
