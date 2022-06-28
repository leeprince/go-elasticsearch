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

// CardinalityAggregation type.
type CardinalityAggregation struct {
	Field              *Field   `json:"field,omitempty"`
	Missing            *Missing `json:"missing,omitempty"`
	PrecisionThreshold *int     `json:"precision_threshold,omitempty"`
	Rehash             *bool    `json:"rehash,omitempty"`
	Script             *Script  `json:"script,omitempty"`
}

// CardinalityAggregationBuilder holds CardinalityAggregation struct and provides a builder API.
type CardinalityAggregationBuilder struct {
	v *CardinalityAggregation
}

// NewCardinalityAggregation provides a builder for the CardinalityAggregation struct.
func NewCardinalityAggregation() *CardinalityAggregationBuilder {
	r := CardinalityAggregationBuilder{
		&CardinalityAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the CardinalityAggregation struct
func (rb *CardinalityAggregationBuilder) Build() CardinalityAggregation {
	return *rb.v
}

// Field set the Field property for CardinalityAggregationBuilder.
func (rb *CardinalityAggregationBuilder) Field(field Field) *CardinalityAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Missing set the Missing property for CardinalityAggregationBuilder.
func (rb *CardinalityAggregationBuilder) Missing(missing Missing) *CardinalityAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// PrecisionThreshold set the PrecisionThreshold property for CardinalityAggregationBuilder.
func (rb *CardinalityAggregationBuilder) PrecisionThreshold(precisionthreshold int) *CardinalityAggregationBuilder {
	rb.v.PrecisionThreshold = &precisionthreshold
	return rb
}

// Rehash set the Rehash property for CardinalityAggregationBuilder.
func (rb *CardinalityAggregationBuilder) Rehash(rehash bool) *CardinalityAggregationBuilder {
	rb.v.Rehash = &rehash
	return rb
}

// Script set the Script property for CardinalityAggregationBuilder.
func (rb *CardinalityAggregationBuilder) Script(script Script) *CardinalityAggregationBuilder {
	rb.v.Script = &script
	return rb
}
