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

// SumAggregation type.
type SumAggregation struct {
	Field   *Field   `json:"field,omitempty"`
	Format  *string  `json:"format,omitempty"`
	Missing *Missing `json:"missing,omitempty"`
	Script  *Script  `json:"script,omitempty"`
}

// SumAggregationBuilder holds SumAggregation struct and provides a builder API.
type SumAggregationBuilder struct {
	v *SumAggregation
}

// NewSumAggregation provides a builder for the SumAggregation struct.
func NewSumAggregation() *SumAggregationBuilder {
	r := SumAggregationBuilder{
		&SumAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the SumAggregation struct
func (rb *SumAggregationBuilder) Build() SumAggregation {
	return *rb.v
}

// Field set the Field property for SumAggregationBuilder.
func (rb *SumAggregationBuilder) Field(field Field) *SumAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Format set the Format property for SumAggregationBuilder.
func (rb *SumAggregationBuilder) Format(format string) *SumAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// Missing set the Missing property for SumAggregationBuilder.
func (rb *SumAggregationBuilder) Missing(missing Missing) *SumAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Script set the Script property for SumAggregationBuilder.
func (rb *SumAggregationBuilder) Script(script Script) *SumAggregationBuilder {
	rb.v.Script = &script
	return rb
}
