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

// MedianAbsoluteDeviationAggregation type.
type MedianAbsoluteDeviationAggregation struct {
	Compression *float64 `json:"compression,omitempty"`
	Field       *Field   `json:"field,omitempty"`
	Format      *string  `json:"format,omitempty"`
	Missing     *Missing `json:"missing,omitempty"`
	Script      *Script  `json:"script,omitempty"`
}

// MedianAbsoluteDeviationAggregationBuilder holds MedianAbsoluteDeviationAggregation struct and provides a builder API.
type MedianAbsoluteDeviationAggregationBuilder struct {
	v *MedianAbsoluteDeviationAggregation
}

// NewMedianAbsoluteDeviationAggregation provides a builder for the MedianAbsoluteDeviationAggregation struct.
func NewMedianAbsoluteDeviationAggregation() *MedianAbsoluteDeviationAggregationBuilder {
	r := MedianAbsoluteDeviationAggregationBuilder{
		&MedianAbsoluteDeviationAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the MedianAbsoluteDeviationAggregation struct
func (rb *MedianAbsoluteDeviationAggregationBuilder) Build() MedianAbsoluteDeviationAggregation {
	return *rb.v
}

// Compression set the Compression property for MedianAbsoluteDeviationAggregationBuilder.
func (rb *MedianAbsoluteDeviationAggregationBuilder) Compression(compression float64) *MedianAbsoluteDeviationAggregationBuilder {
	rb.v.Compression = &compression
	return rb
}

// Field set the Field property for MedianAbsoluteDeviationAggregationBuilder.
func (rb *MedianAbsoluteDeviationAggregationBuilder) Field(field Field) *MedianAbsoluteDeviationAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Format set the Format property for MedianAbsoluteDeviationAggregationBuilder.
func (rb *MedianAbsoluteDeviationAggregationBuilder) Format(format string) *MedianAbsoluteDeviationAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// Missing set the Missing property for MedianAbsoluteDeviationAggregationBuilder.
func (rb *MedianAbsoluteDeviationAggregationBuilder) Missing(missing Missing) *MedianAbsoluteDeviationAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Script set the Script property for MedianAbsoluteDeviationAggregationBuilder.
func (rb *MedianAbsoluteDeviationAggregationBuilder) Script(script Script) *MedianAbsoluteDeviationAggregationBuilder {
	rb.v.Script = &script
	return rb
}
