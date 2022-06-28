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

// MedianAbsoluteDeviationAggregate type.
type MedianAbsoluteDeviationAggregate struct {
	Meta *Metadata `json:"meta,omitempty"`
	// Value The metric value. A missing value generally means that there was no data to
	// aggregate,
	// unless specified otherwise.
	Value         float64 `json:"value,omitempty"`
	ValueAsString *string `json:"value_as_string,omitempty"`
}

// MedianAbsoluteDeviationAggregateBuilder holds MedianAbsoluteDeviationAggregate struct and provides a builder API.
type MedianAbsoluteDeviationAggregateBuilder struct {
	v *MedianAbsoluteDeviationAggregate
}

// NewMedianAbsoluteDeviationAggregate provides a builder for the MedianAbsoluteDeviationAggregate struct.
func NewMedianAbsoluteDeviationAggregate() *MedianAbsoluteDeviationAggregateBuilder {
	r := MedianAbsoluteDeviationAggregateBuilder{
		&MedianAbsoluteDeviationAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the MedianAbsoluteDeviationAggregate struct
func (rb *MedianAbsoluteDeviationAggregateBuilder) Build() MedianAbsoluteDeviationAggregate {
	return *rb.v
}

// Meta set the Meta property for MedianAbsoluteDeviationAggregateBuilder.
func (rb *MedianAbsoluteDeviationAggregateBuilder) Meta(meta Metadata) *MedianAbsoluteDeviationAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}

// Value The metric value. A missing value generally means that there was no data to
// aggregate,
// unless specified otherwise.
func (rb *MedianAbsoluteDeviationAggregateBuilder) Value(value float64) *MedianAbsoluteDeviationAggregateBuilder {
	rb.v.Value = value
	return rb
}

// ValueAsString set the ValueAsString property for MedianAbsoluteDeviationAggregateBuilder.
func (rb *MedianAbsoluteDeviationAggregateBuilder) ValueAsString(valueasstring string) *MedianAbsoluteDeviationAggregateBuilder {
	rb.v.ValueAsString = &valueasstring
	return rb
}
