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

// ExtendedStatsAggregation type.
type ExtendedStatsAggregation struct {
	Field   *Field   `json:"field,omitempty"`
	Format  *string  `json:"format,omitempty"`
	Missing *Missing `json:"missing,omitempty"`
	Script  *Script  `json:"script,omitempty"`
	Sigma   *float64 `json:"sigma,omitempty"`
}

// ExtendedStatsAggregationBuilder holds ExtendedStatsAggregation struct and provides a builder API.
type ExtendedStatsAggregationBuilder struct {
	v *ExtendedStatsAggregation
}

// NewExtendedStatsAggregation provides a builder for the ExtendedStatsAggregation struct.
func NewExtendedStatsAggregation() *ExtendedStatsAggregationBuilder {
	r := ExtendedStatsAggregationBuilder{
		&ExtendedStatsAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the ExtendedStatsAggregation struct
func (rb *ExtendedStatsAggregationBuilder) Build() ExtendedStatsAggregation {
	return *rb.v
}

// Field set the Field property for ExtendedStatsAggregationBuilder.
func (rb *ExtendedStatsAggregationBuilder) Field(field Field) *ExtendedStatsAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Format set the Format property for ExtendedStatsAggregationBuilder.
func (rb *ExtendedStatsAggregationBuilder) Format(format string) *ExtendedStatsAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// Missing set the Missing property for ExtendedStatsAggregationBuilder.
func (rb *ExtendedStatsAggregationBuilder) Missing(missing Missing) *ExtendedStatsAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Script set the Script property for ExtendedStatsAggregationBuilder.
func (rb *ExtendedStatsAggregationBuilder) Script(script Script) *ExtendedStatsAggregationBuilder {
	rb.v.Script = &script
	return rb
}

// Sigma set the Sigma property for ExtendedStatsAggregationBuilder.
func (rb *ExtendedStatsAggregationBuilder) Sigma(sigma float64) *ExtendedStatsAggregationBuilder {
	rb.v.Sigma = &sigma
	return rb
}
