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

// AverageAggregation type.
type AverageAggregation struct {
	Field   *Field   `json:"field,omitempty"`
	Format  *string  `json:"format,omitempty"`
	Missing *Missing `json:"missing,omitempty"`
	Script  *Script  `json:"script,omitempty"`
}

// AverageAggregationBuilder holds AverageAggregation struct and provides a builder API.
type AverageAggregationBuilder struct {
	v *AverageAggregation
}

// NewAverageAggregation provides a builder for the AverageAggregation struct.
func NewAverageAggregation() *AverageAggregationBuilder {
	r := AverageAggregationBuilder{
		&AverageAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the AverageAggregation struct
func (rb *AverageAggregationBuilder) Build() AverageAggregation {
	return *rb.v
}

// Field set the Field property for AverageAggregationBuilder.
func (rb *AverageAggregationBuilder) Field(field Field) *AverageAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Format set the Format property for AverageAggregationBuilder.
func (rb *AverageAggregationBuilder) Format(format string) *AverageAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// Missing set the Missing property for AverageAggregationBuilder.
func (rb *AverageAggregationBuilder) Missing(missing Missing) *AverageAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Script set the Script property for AverageAggregationBuilder.
func (rb *AverageAggregationBuilder) Script(script Script) *AverageAggregationBuilder {
	rb.v.Script = &script
	return rb
}
