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

// StatsAggregation type.
type StatsAggregation struct {
	Field   *Field   `json:"field,omitempty"`
	Format  *string  `json:"format,omitempty"`
	Missing *Missing `json:"missing,omitempty"`
	Script  *Script  `json:"script,omitempty"`
}

// StatsAggregationBuilder holds StatsAggregation struct and provides a builder API.
type StatsAggregationBuilder struct {
	v *StatsAggregation
}

// NewStatsAggregation provides a builder for the StatsAggregation struct.
func NewStatsAggregation() *StatsAggregationBuilder {
	r := StatsAggregationBuilder{
		&StatsAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the StatsAggregation struct
func (rb *StatsAggregationBuilder) Build() StatsAggregation {
	return *rb.v
}

// Field set the Field property for StatsAggregationBuilder.
func (rb *StatsAggregationBuilder) Field(field Field) *StatsAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Format set the Format property for StatsAggregationBuilder.
func (rb *StatsAggregationBuilder) Format(format string) *StatsAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// Missing set the Missing property for StatsAggregationBuilder.
func (rb *StatsAggregationBuilder) Missing(missing Missing) *StatsAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Script set the Script property for StatsAggregationBuilder.
func (rb *StatsAggregationBuilder) Script(script Script) *StatsAggregationBuilder {
	rb.v.Script = &script
	return rb
}
