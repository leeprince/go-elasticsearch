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

// MaxAggregation type.
type MaxAggregation struct {
	Field   *Field   `json:"field,omitempty"`
	Format  *string  `json:"format,omitempty"`
	Missing *Missing `json:"missing,omitempty"`
	Script  *Script  `json:"script,omitempty"`
}

// MaxAggregationBuilder holds MaxAggregation struct and provides a builder API.
type MaxAggregationBuilder struct {
	v *MaxAggregation
}

// NewMaxAggregation provides a builder for the MaxAggregation struct.
func NewMaxAggregation() *MaxAggregationBuilder {
	r := MaxAggregationBuilder{
		&MaxAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the MaxAggregation struct
func (rb *MaxAggregationBuilder) Build() MaxAggregation {
	return *rb.v
}

// Field set the Field property for MaxAggregationBuilder.
func (rb *MaxAggregationBuilder) Field(field Field) *MaxAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Format set the Format property for MaxAggregationBuilder.
func (rb *MaxAggregationBuilder) Format(format string) *MaxAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// Missing set the Missing property for MaxAggregationBuilder.
func (rb *MaxAggregationBuilder) Missing(missing Missing) *MaxAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Script set the Script property for MaxAggregationBuilder.
func (rb *MaxAggregationBuilder) Script(script Script) *MaxAggregationBuilder {
	rb.v.Script = &script
	return rb
}
