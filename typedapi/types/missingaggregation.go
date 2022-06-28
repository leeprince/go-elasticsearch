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

// MissingAggregation type.
type MissingAggregation struct {
	Field   *Field    `json:"field,omitempty"`
	Meta    *Metadata `json:"meta,omitempty"`
	Missing *Missing  `json:"missing,omitempty"`
	Name    *string   `json:"name,omitempty"`
}

// MissingAggregationBuilder holds MissingAggregation struct and provides a builder API.
type MissingAggregationBuilder struct {
	v *MissingAggregation
}

// NewMissingAggregation provides a builder for the MissingAggregation struct.
func NewMissingAggregation() *MissingAggregationBuilder {
	r := MissingAggregationBuilder{
		&MissingAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the MissingAggregation struct
func (rb *MissingAggregationBuilder) Build() MissingAggregation {
	return *rb.v
}

// Field set the Field property for MissingAggregationBuilder.
func (rb *MissingAggregationBuilder) Field(field Field) *MissingAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Meta set the Meta property for MissingAggregationBuilder.
func (rb *MissingAggregationBuilder) Meta(meta Metadata) *MissingAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Missing set the Missing property for MissingAggregationBuilder.
func (rb *MissingAggregationBuilder) Missing(missing Missing) *MissingAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Name set the Name property for MissingAggregationBuilder.
func (rb *MissingAggregationBuilder) Name(name string) *MissingAggregationBuilder {
	rb.v.Name = &name
	return rb
}
