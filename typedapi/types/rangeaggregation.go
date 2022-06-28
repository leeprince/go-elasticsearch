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

// RangeAggregation type.
type RangeAggregation struct {
	Field   *Field             `json:"field,omitempty"`
	Format  *string            `json:"format,omitempty"`
	Keyed   *bool              `json:"keyed,omitempty"`
	Meta    *Metadata          `json:"meta,omitempty"`
	Missing *int               `json:"missing,omitempty"`
	Name    *string            `json:"name,omitempty"`
	Ranges  []AggregationRange `json:"ranges,omitempty"`
	Script  *Script            `json:"script,omitempty"`
}

// RangeAggregationBuilder holds RangeAggregation struct and provides a builder API.
type RangeAggregationBuilder struct {
	v *RangeAggregation
}

// NewRangeAggregation provides a builder for the RangeAggregation struct.
func NewRangeAggregation() *RangeAggregationBuilder {
	r := RangeAggregationBuilder{
		&RangeAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the RangeAggregation struct
func (rb *RangeAggregationBuilder) Build() RangeAggregation {
	return *rb.v
}

// Field set the Field property for RangeAggregationBuilder.
func (rb *RangeAggregationBuilder) Field(field Field) *RangeAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Format set the Format property for RangeAggregationBuilder.
func (rb *RangeAggregationBuilder) Format(format string) *RangeAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// Keyed set the Keyed property for RangeAggregationBuilder.
func (rb *RangeAggregationBuilder) Keyed(keyed bool) *RangeAggregationBuilder {
	rb.v.Keyed = &keyed
	return rb
}

// Meta set the Meta property for RangeAggregationBuilder.
func (rb *RangeAggregationBuilder) Meta(meta Metadata) *RangeAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Missing set the Missing property for RangeAggregationBuilder.
func (rb *RangeAggregationBuilder) Missing(missing int) *RangeAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Name set the Name property for RangeAggregationBuilder.
func (rb *RangeAggregationBuilder) Name(name string) *RangeAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Ranges set the Ranges property for RangeAggregationBuilder.
func (rb *RangeAggregationBuilder) Ranges(ranges ...AggregationRange) *RangeAggregationBuilder {
	rb.v.Ranges = append(rb.v.Ranges, ranges...)
	return rb
}

// Script set the Script property for RangeAggregationBuilder.
func (rb *RangeAggregationBuilder) Script(script Script) *RangeAggregationBuilder {
	rb.v.Script = &script
	return rb
}
