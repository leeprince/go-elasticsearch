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

// DerivativeAggregate type.
type DerivativeAggregate struct {
	Meta                    *Metadata `json:"meta,omitempty"`
	NormalizedValue         *float64  `json:"normalized_value,omitempty"`
	NormalizedValueAsString *string   `json:"normalized_value_as_string,omitempty"`
	// Value The metric value. A missing value generally means that there was no data to
	// aggregate,
	// unless specified otherwise.
	Value         float64 `json:"value,omitempty"`
	ValueAsString *string `json:"value_as_string,omitempty"`
}

// DerivativeAggregateBuilder holds DerivativeAggregate struct and provides a builder API.
type DerivativeAggregateBuilder struct {
	v *DerivativeAggregate
}

// NewDerivativeAggregate provides a builder for the DerivativeAggregate struct.
func NewDerivativeAggregate() *DerivativeAggregateBuilder {
	r := DerivativeAggregateBuilder{
		&DerivativeAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the DerivativeAggregate struct
func (rb *DerivativeAggregateBuilder) Build() DerivativeAggregate {
	return *rb.v
}

// Meta set the Meta property for DerivativeAggregateBuilder.
func (rb *DerivativeAggregateBuilder) Meta(meta Metadata) *DerivativeAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}

// NormalizedValue set the NormalizedValue property for DerivativeAggregateBuilder.
func (rb *DerivativeAggregateBuilder) NormalizedValue(normalizedvalue float64) *DerivativeAggregateBuilder {
	rb.v.NormalizedValue = &normalizedvalue
	return rb
}

// NormalizedValueAsString set the NormalizedValueAsString property for DerivativeAggregateBuilder.
func (rb *DerivativeAggregateBuilder) NormalizedValueAsString(normalizedvalueasstring string) *DerivativeAggregateBuilder {
	rb.v.NormalizedValueAsString = &normalizedvalueasstring
	return rb
}

// Value The metric value. A missing value generally means that there was no data to
// aggregate,
// unless specified otherwise.
func (rb *DerivativeAggregateBuilder) Value(value float64) *DerivativeAggregateBuilder {
	rb.v.Value = value
	return rb
}

// ValueAsString set the ValueAsString property for DerivativeAggregateBuilder.
func (rb *DerivativeAggregateBuilder) ValueAsString(valueasstring string) *DerivativeAggregateBuilder {
	rb.v.ValueAsString = &valueasstring
	return rb
}
