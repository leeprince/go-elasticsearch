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

// StringStatsAggregate type.
type StringStatsAggregate struct {
	AvgLength         float64            `json:"avg_length,omitempty"`
	AvgLengthAsString *string            `json:"avg_length_as_string,omitempty"`
	Count             int64              `json:"count"`
	Distribution      map[string]float64 `json:"distribution,omitempty"`
	Entropy           float64            `json:"entropy,omitempty"`
	MaxLength         int                `json:"max_length,omitempty"`
	MaxLengthAsString *string            `json:"max_length_as_string,omitempty"`
	Meta              *Metadata          `json:"meta,omitempty"`
	MinLength         int                `json:"min_length,omitempty"`
	MinLengthAsString *string            `json:"min_length_as_string,omitempty"`
}

// StringStatsAggregateBuilder holds StringStatsAggregate struct and provides a builder API.
type StringStatsAggregateBuilder struct {
	v *StringStatsAggregate
}

// NewStringStatsAggregate provides a builder for the StringStatsAggregate struct.
func NewStringStatsAggregate() *StringStatsAggregateBuilder {
	r := StringStatsAggregateBuilder{
		&StringStatsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the StringStatsAggregate struct
func (rb *StringStatsAggregateBuilder) Build() StringStatsAggregate {
	return *rb.v
}

// AvgLength set the AvgLength property for StringStatsAggregateBuilder.
func (rb *StringStatsAggregateBuilder) AvgLength(avglength float64) *StringStatsAggregateBuilder {
	rb.v.AvgLength = avglength
	return rb
}

// AvgLengthAsString set the AvgLengthAsString property for StringStatsAggregateBuilder.
func (rb *StringStatsAggregateBuilder) AvgLengthAsString(avglengthasstring string) *StringStatsAggregateBuilder {
	rb.v.AvgLengthAsString = &avglengthasstring
	return rb
}

// Count set the Count property for StringStatsAggregateBuilder.
func (rb *StringStatsAggregateBuilder) Count(count int64) *StringStatsAggregateBuilder {
	rb.v.Count = count
	return rb
}

// Distribution set the Distribution property for StringStatsAggregateBuilder.
func (rb *StringStatsAggregateBuilder) Distribution(distribution map[string]float64) *StringStatsAggregateBuilder {
	rb.v.Distribution = distribution
	return rb
}

// Entropy set the Entropy property for StringStatsAggregateBuilder.
func (rb *StringStatsAggregateBuilder) Entropy(entropy float64) *StringStatsAggregateBuilder {
	rb.v.Entropy = entropy
	return rb
}

// MaxLength set the MaxLength property for StringStatsAggregateBuilder.
func (rb *StringStatsAggregateBuilder) MaxLength(maxlength int) *StringStatsAggregateBuilder {
	rb.v.MaxLength = maxlength
	return rb
}

// MaxLengthAsString set the MaxLengthAsString property for StringStatsAggregateBuilder.
func (rb *StringStatsAggregateBuilder) MaxLengthAsString(maxlengthasstring string) *StringStatsAggregateBuilder {
	rb.v.MaxLengthAsString = &maxlengthasstring
	return rb
}

// Meta set the Meta property for StringStatsAggregateBuilder.
func (rb *StringStatsAggregateBuilder) Meta(meta Metadata) *StringStatsAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}

// MinLength set the MinLength property for StringStatsAggregateBuilder.
func (rb *StringStatsAggregateBuilder) MinLength(minlength int) *StringStatsAggregateBuilder {
	rb.v.MinLength = minlength
	return rb
}

// MinLengthAsString set the MinLengthAsString property for StringStatsAggregateBuilder.
func (rb *StringStatsAggregateBuilder) MinLengthAsString(minlengthasstring string) *StringStatsAggregateBuilder {
	rb.v.MinLengthAsString = &minlengthasstring
	return rb
}
