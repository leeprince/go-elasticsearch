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

// DateRangeAggregation type.
type DateRangeAggregation struct {
	Field    *Field                `json:"field,omitempty"`
	Format   *string               `json:"format,omitempty"`
	Keyed    *bool                 `json:"keyed,omitempty"`
	Meta     *Metadata             `json:"meta,omitempty"`
	Missing  *Missing              `json:"missing,omitempty"`
	Name     *string               `json:"name,omitempty"`
	Ranges   []DateRangeExpression `json:"ranges,omitempty"`
	TimeZone *string               `json:"time_zone,omitempty"`
}

// DateRangeAggregationBuilder holds DateRangeAggregation struct and provides a builder API.
type DateRangeAggregationBuilder struct {
	v *DateRangeAggregation
}

// NewDateRangeAggregation provides a builder for the DateRangeAggregation struct.
func NewDateRangeAggregation() *DateRangeAggregationBuilder {
	r := DateRangeAggregationBuilder{
		&DateRangeAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the DateRangeAggregation struct
func (rb *DateRangeAggregationBuilder) Build() DateRangeAggregation {
	return *rb.v
}

// Field set the Field property for DateRangeAggregationBuilder.
func (rb *DateRangeAggregationBuilder) Field(field Field) *DateRangeAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Format set the Format property for DateRangeAggregationBuilder.
func (rb *DateRangeAggregationBuilder) Format(format string) *DateRangeAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// Keyed set the Keyed property for DateRangeAggregationBuilder.
func (rb *DateRangeAggregationBuilder) Keyed(keyed bool) *DateRangeAggregationBuilder {
	rb.v.Keyed = &keyed
	return rb
}

// Meta set the Meta property for DateRangeAggregationBuilder.
func (rb *DateRangeAggregationBuilder) Meta(meta Metadata) *DateRangeAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Missing set the Missing property for DateRangeAggregationBuilder.
func (rb *DateRangeAggregationBuilder) Missing(missing Missing) *DateRangeAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Name set the Name property for DateRangeAggregationBuilder.
func (rb *DateRangeAggregationBuilder) Name(name string) *DateRangeAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Ranges set the Ranges property for DateRangeAggregationBuilder.
func (rb *DateRangeAggregationBuilder) Ranges(ranges ...DateRangeExpression) *DateRangeAggregationBuilder {
	rb.v.Ranges = append(rb.v.Ranges, ranges...)
	return rb
}

// TimeZone set the TimeZone property for DateRangeAggregationBuilder.
func (rb *DateRangeAggregationBuilder) TimeZone(timezone string) *DateRangeAggregationBuilder {
	rb.v.TimeZone = &timezone
	return rb
}
