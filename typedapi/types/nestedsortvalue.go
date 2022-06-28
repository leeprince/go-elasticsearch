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

// NestedSortValue type.
type NestedSortValue struct {
	Filter      *QueryContainer  `json:"filter,omitempty"`
	MaxChildren *int             `json:"max_children,omitempty"`
	Nested      *NestedSortValue `json:"nested,omitempty"`
	Path        Field            `json:"path"`
}

// NestedSortValueBuilder holds NestedSortValue struct and provides a builder API.
type NestedSortValueBuilder struct {
	v *NestedSortValue
}

// NewNestedSortValue provides a builder for the NestedSortValue struct.
func NewNestedSortValue() *NestedSortValueBuilder {
	r := NestedSortValueBuilder{
		&NestedSortValue{},
	}

	return &r
}

// Build finalize the chain and returns the NestedSortValue struct
func (rb *NestedSortValueBuilder) Build() NestedSortValue {
	return *rb.v
}

// Filter set the Filter property for NestedSortValueBuilder.
func (rb *NestedSortValueBuilder) Filter(filter QueryContainer) *NestedSortValueBuilder {
	rb.v.Filter = &filter
	return rb
}

// MaxChildren set the MaxChildren property for NestedSortValueBuilder.
func (rb *NestedSortValueBuilder) MaxChildren(maxchildren int) *NestedSortValueBuilder {
	rb.v.MaxChildren = &maxchildren
	return rb
}

// Nested set the Nested property for NestedSortValueBuilder.
func (rb *NestedSortValueBuilder) Nested(nested NestedSortValue) *NestedSortValueBuilder {
	rb.v.Nested = &nested
	return rb
}

// Path set the Path property for NestedSortValueBuilder.
func (rb *NestedSortValueBuilder) Path(path Field) *NestedSortValueBuilder {
	rb.v.Path = path
	return rb
}
