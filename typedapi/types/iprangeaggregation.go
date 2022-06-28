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

// IpRangeAggregation type.
type IpRangeAggregation struct {
	Field  *Field                    `json:"field,omitempty"`
	Meta   *Metadata                 `json:"meta,omitempty"`
	Name   *string                   `json:"name,omitempty"`
	Ranges []IpRangeAggregationRange `json:"ranges,omitempty"`
}

// IpRangeAggregationBuilder holds IpRangeAggregation struct and provides a builder API.
type IpRangeAggregationBuilder struct {
	v *IpRangeAggregation
}

// NewIpRangeAggregation provides a builder for the IpRangeAggregation struct.
func NewIpRangeAggregation() *IpRangeAggregationBuilder {
	r := IpRangeAggregationBuilder{
		&IpRangeAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the IpRangeAggregation struct
func (rb *IpRangeAggregationBuilder) Build() IpRangeAggregation {
	return *rb.v
}

// Field set the Field property for IpRangeAggregationBuilder.
func (rb *IpRangeAggregationBuilder) Field(field Field) *IpRangeAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Meta set the Meta property for IpRangeAggregationBuilder.
func (rb *IpRangeAggregationBuilder) Meta(meta Metadata) *IpRangeAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for IpRangeAggregationBuilder.
func (rb *IpRangeAggregationBuilder) Name(name string) *IpRangeAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Ranges set the Ranges property for IpRangeAggregationBuilder.
func (rb *IpRangeAggregationBuilder) Ranges(ranges ...IpRangeAggregationRange) *IpRangeAggregationBuilder {
	rb.v.Ranges = append(rb.v.Ranges, ranges...)
	return rb
}
