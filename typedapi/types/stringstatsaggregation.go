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

// StringStatsAggregation type.
type StringStatsAggregation struct {
	Field            *Field   `json:"field,omitempty"`
	Missing          *Missing `json:"missing,omitempty"`
	Script           *Script  `json:"script,omitempty"`
	ShowDistribution *bool    `json:"show_distribution,omitempty"`
}

// StringStatsAggregationBuilder holds StringStatsAggregation struct and provides a builder API.
type StringStatsAggregationBuilder struct {
	v *StringStatsAggregation
}

// NewStringStatsAggregation provides a builder for the StringStatsAggregation struct.
func NewStringStatsAggregation() *StringStatsAggregationBuilder {
	r := StringStatsAggregationBuilder{
		&StringStatsAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the StringStatsAggregation struct
func (rb *StringStatsAggregationBuilder) Build() StringStatsAggregation {
	return *rb.v
}

// Field set the Field property for StringStatsAggregationBuilder.
func (rb *StringStatsAggregationBuilder) Field(field Field) *StringStatsAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Missing set the Missing property for StringStatsAggregationBuilder.
func (rb *StringStatsAggregationBuilder) Missing(missing Missing) *StringStatsAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Script set the Script property for StringStatsAggregationBuilder.
func (rb *StringStatsAggregationBuilder) Script(script Script) *StringStatsAggregationBuilder {
	rb.v.Script = &script
	return rb
}

// ShowDistribution set the ShowDistribution property for StringStatsAggregationBuilder.
func (rb *StringStatsAggregationBuilder) ShowDistribution(showdistribution bool) *StringStatsAggregationBuilder {
	rb.v.ShowDistribution = &showdistribution
	return rb
}
