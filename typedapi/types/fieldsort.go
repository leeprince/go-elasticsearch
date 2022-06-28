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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fieldsortnumerictype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fieldtype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// FieldSort type.
type FieldSort struct {
	Format       *string                                    `json:"format,omitempty"`
	Missing      *Missing                                   `json:"missing,omitempty"`
	Mode         *sortmode.SortMode                         `json:"mode,omitempty"`
	Nested       *NestedSortValue                           `json:"nested,omitempty"`
	NumericType  *fieldsortnumerictype.FieldSortNumericType `json:"numeric_type,omitempty"`
	Order        *sortorder.SortOrder                       `json:"order,omitempty"`
	UnmappedType *fieldtype.FieldType                       `json:"unmapped_type,omitempty"`
}

// FieldSortBuilder holds FieldSort struct and provides a builder API.
type FieldSortBuilder struct {
	v *FieldSort
}

// NewFieldSort provides a builder for the FieldSort struct.
func NewFieldSort() *FieldSortBuilder {
	r := FieldSortBuilder{
		&FieldSort{},
	}

	return &r
}

// Build finalize the chain and returns the FieldSort struct
func (rb *FieldSortBuilder) Build() FieldSort {
	return *rb.v
}

// Format set the Format property for FieldSortBuilder.
func (rb *FieldSortBuilder) Format(format string) *FieldSortBuilder {
	rb.v.Format = &format
	return rb
}

// Missing set the Missing property for FieldSortBuilder.
func (rb *FieldSortBuilder) Missing(missing Missing) *FieldSortBuilder {
	rb.v.Missing = &missing
	return rb
}

// Mode set the Mode property for FieldSortBuilder.
func (rb *FieldSortBuilder) Mode(mode sortmode.SortMode) *FieldSortBuilder {
	rb.v.Mode = &mode
	return rb
}

// Nested set the Nested property for FieldSortBuilder.
func (rb *FieldSortBuilder) Nested(nested NestedSortValue) *FieldSortBuilder {
	rb.v.Nested = &nested
	return rb
}

// NumericType set the NumericType property for FieldSortBuilder.
func (rb *FieldSortBuilder) NumericType(numerictype fieldsortnumerictype.FieldSortNumericType) *FieldSortBuilder {
	rb.v.NumericType = &numerictype
	return rb
}

// Order set the Order property for FieldSortBuilder.
func (rb *FieldSortBuilder) Order(order sortorder.SortOrder) *FieldSortBuilder {
	rb.v.Order = &order
	return rb
}

// UnmappedType set the UnmappedType property for FieldSortBuilder.
func (rb *FieldSortBuilder) UnmappedType(unmappedtype fieldtype.FieldType) *FieldSortBuilder {
	rb.v.UnmappedType = &unmappedtype
	return rb
}
