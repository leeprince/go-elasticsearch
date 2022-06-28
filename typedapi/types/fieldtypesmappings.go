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

// FieldTypesMappings type.
type FieldTypesMappings struct {
	FieldTypes        []FieldTypes        `json:"field_types"`
	RuntimeFieldTypes []RuntimeFieldTypes `json:"runtime_field_types,omitempty"`
}

// FieldTypesMappingsBuilder holds FieldTypesMappings struct and provides a builder API.
type FieldTypesMappingsBuilder struct {
	v *FieldTypesMappings
}

// NewFieldTypesMappings provides a builder for the FieldTypesMappings struct.
func NewFieldTypesMappings() *FieldTypesMappingsBuilder {
	r := FieldTypesMappingsBuilder{
		&FieldTypesMappings{},
	}

	return &r
}

// Build finalize the chain and returns the FieldTypesMappings struct
func (rb *FieldTypesMappingsBuilder) Build() FieldTypesMappings {
	return *rb.v
}

// FieldTypes set the FieldTypes property for FieldTypesMappingsBuilder.
func (rb *FieldTypesMappingsBuilder) FieldTypes(field_types ...FieldTypes) *FieldTypesMappingsBuilder {
	rb.v.FieldTypes = append(rb.v.FieldTypes, field_types...)
	return rb
}

// RuntimeFieldTypes set the RuntimeFieldTypes property for FieldTypesMappingsBuilder.
func (rb *FieldTypesMappingsBuilder) RuntimeFieldTypes(runtime_field_types ...RuntimeFieldTypes) *FieldTypesMappingsBuilder {
	rb.v.RuntimeFieldTypes = append(rb.v.RuntimeFieldTypes, runtime_field_types...)
	return rb
}
