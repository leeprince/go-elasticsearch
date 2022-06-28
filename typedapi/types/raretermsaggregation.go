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

// RareTermsAggregation type.
type RareTermsAggregation struct {
	Exclude     *TermsExclude `json:"exclude,omitempty"`
	Field       *Field        `json:"field,omitempty"`
	Include     *TermsInclude `json:"include,omitempty"`
	MaxDocCount *int64        `json:"max_doc_count,omitempty"`
	Meta        *Metadata     `json:"meta,omitempty"`
	Missing     *Missing      `json:"missing,omitempty"`
	Name        *string       `json:"name,omitempty"`
	Precision   *float64      `json:"precision,omitempty"`
	ValueType   *string       `json:"value_type,omitempty"`
}

// RareTermsAggregationBuilder holds RareTermsAggregation struct and provides a builder API.
type RareTermsAggregationBuilder struct {
	v *RareTermsAggregation
}

// NewRareTermsAggregation provides a builder for the RareTermsAggregation struct.
func NewRareTermsAggregation() *RareTermsAggregationBuilder {
	r := RareTermsAggregationBuilder{
		&RareTermsAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the RareTermsAggregation struct
func (rb *RareTermsAggregationBuilder) Build() RareTermsAggregation {
	return *rb.v
}

// Exclude set the Exclude property for RareTermsAggregationBuilder.
func (rb *RareTermsAggregationBuilder) Exclude(exclude TermsExclude) *RareTermsAggregationBuilder {
	rb.v.Exclude = &exclude
	return rb
}

// Field set the Field property for RareTermsAggregationBuilder.
func (rb *RareTermsAggregationBuilder) Field(field Field) *RareTermsAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Include set the Include property for RareTermsAggregationBuilder.
func (rb *RareTermsAggregationBuilder) Include(include TermsInclude) *RareTermsAggregationBuilder {
	rb.v.Include = &include
	return rb
}

// MaxDocCount set the MaxDocCount property for RareTermsAggregationBuilder.
func (rb *RareTermsAggregationBuilder) MaxDocCount(maxdoccount int64) *RareTermsAggregationBuilder {
	rb.v.MaxDocCount = &maxdoccount
	return rb
}

// Meta set the Meta property for RareTermsAggregationBuilder.
func (rb *RareTermsAggregationBuilder) Meta(meta Metadata) *RareTermsAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Missing set the Missing property for RareTermsAggregationBuilder.
func (rb *RareTermsAggregationBuilder) Missing(missing Missing) *RareTermsAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Name set the Name property for RareTermsAggregationBuilder.
func (rb *RareTermsAggregationBuilder) Name(name string) *RareTermsAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Precision set the Precision property for RareTermsAggregationBuilder.
func (rb *RareTermsAggregationBuilder) Precision(precision float64) *RareTermsAggregationBuilder {
	rb.v.Precision = &precision
	return rb
}

// ValueType set the ValueType property for RareTermsAggregationBuilder.
func (rb *RareTermsAggregationBuilder) ValueType(valuetype string) *RareTermsAggregationBuilder {
	rb.v.ValueType = &valuetype
	return rb
}
