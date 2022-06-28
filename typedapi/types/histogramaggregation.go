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

// HistogramAggregation type.
type HistogramAggregation struct {
	ExtendedBounds *ExtendedBoundsdouble `json:"extended_bounds,omitempty"`
	Field          *Field                `json:"field,omitempty"`
	Format         *string               `json:"format,omitempty"`
	HardBounds     *ExtendedBoundsdouble `json:"hard_bounds,omitempty"`
	Interval       *float64              `json:"interval,omitempty"`
	Keyed          *bool                 `json:"keyed,omitempty"`
	Meta           *Metadata             `json:"meta,omitempty"`
	MinDocCount    *int                  `json:"min_doc_count,omitempty"`
	Missing        *float64              `json:"missing,omitempty"`
	Name           *string               `json:"name,omitempty"`
	Offset         *float64              `json:"offset,omitempty"`
	Order          *HistogramOrder       `json:"order,omitempty"`
	Script         *Script               `json:"script,omitempty"`
}

// HistogramAggregationBuilder holds HistogramAggregation struct and provides a builder API.
type HistogramAggregationBuilder struct {
	v *HistogramAggregation
}

// NewHistogramAggregation provides a builder for the HistogramAggregation struct.
func NewHistogramAggregation() *HistogramAggregationBuilder {
	r := HistogramAggregationBuilder{
		&HistogramAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the HistogramAggregation struct
func (rb *HistogramAggregationBuilder) Build() HistogramAggregation {
	return *rb.v
}

// ExtendedBounds set the ExtendedBounds property for HistogramAggregationBuilder.
func (rb *HistogramAggregationBuilder) ExtendedBounds(extendedbounds ExtendedBoundsdouble) *HistogramAggregationBuilder {
	rb.v.ExtendedBounds = &extendedbounds
	return rb
}

// Field set the Field property for HistogramAggregationBuilder.
func (rb *HistogramAggregationBuilder) Field(field Field) *HistogramAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Format set the Format property for HistogramAggregationBuilder.
func (rb *HistogramAggregationBuilder) Format(format string) *HistogramAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// HardBounds set the HardBounds property for HistogramAggregationBuilder.
func (rb *HistogramAggregationBuilder) HardBounds(hardbounds ExtendedBoundsdouble) *HistogramAggregationBuilder {
	rb.v.HardBounds = &hardbounds
	return rb
}

// Interval set the Interval property for HistogramAggregationBuilder.
func (rb *HistogramAggregationBuilder) Interval(interval float64) *HistogramAggregationBuilder {
	rb.v.Interval = &interval
	return rb
}

// Keyed set the Keyed property for HistogramAggregationBuilder.
func (rb *HistogramAggregationBuilder) Keyed(keyed bool) *HistogramAggregationBuilder {
	rb.v.Keyed = &keyed
	return rb
}

// Meta set the Meta property for HistogramAggregationBuilder.
func (rb *HistogramAggregationBuilder) Meta(meta Metadata) *HistogramAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// MinDocCount set the MinDocCount property for HistogramAggregationBuilder.
func (rb *HistogramAggregationBuilder) MinDocCount(mindoccount int) *HistogramAggregationBuilder {
	rb.v.MinDocCount = &mindoccount
	return rb
}

// Missing set the Missing property for HistogramAggregationBuilder.
func (rb *HistogramAggregationBuilder) Missing(missing float64) *HistogramAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Name set the Name property for HistogramAggregationBuilder.
func (rb *HistogramAggregationBuilder) Name(name string) *HistogramAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Offset set the Offset property for HistogramAggregationBuilder.
func (rb *HistogramAggregationBuilder) Offset(offset float64) *HistogramAggregationBuilder {
	rb.v.Offset = &offset
	return rb
}

// Order set the Order property for HistogramAggregationBuilder.
func (rb *HistogramAggregationBuilder) Order(order HistogramOrder) *HistogramAggregationBuilder {
	rb.v.Order = &order
	return rb
}

// Script set the Script property for HistogramAggregationBuilder.
func (rb *HistogramAggregationBuilder) Script(script Script) *HistogramAggregationBuilder {
	rb.v.Script = &script
	return rb
}
