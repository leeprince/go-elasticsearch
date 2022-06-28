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

// FieldCollapse type.
type FieldCollapse struct {
	Collapse                   *FieldCollapse `json:"collapse,omitempty"`
	Field                      Field          `json:"field"`
	InnerHits                  []InnerHits    `json:"inner_hits,omitempty"`
	MaxConcurrentGroupSearches *int           `json:"max_concurrent_group_searches,omitempty"`
}

// FieldCollapseBuilder holds FieldCollapse struct and provides a builder API.
type FieldCollapseBuilder struct {
	v *FieldCollapse
}

// NewFieldCollapse provides a builder for the FieldCollapse struct.
func NewFieldCollapse() *FieldCollapseBuilder {
	r := FieldCollapseBuilder{
		&FieldCollapse{},
	}

	return &r
}

// Build finalize the chain and returns the FieldCollapse struct
func (rb *FieldCollapseBuilder) Build() FieldCollapse {
	return *rb.v
}

// Collapse set the Collapse property for FieldCollapseBuilder.
func (rb *FieldCollapseBuilder) Collapse(collapse FieldCollapse) *FieldCollapseBuilder {
	rb.v.Collapse = &collapse
	return rb
}

// Field set the Field property for FieldCollapseBuilder.
func (rb *FieldCollapseBuilder) Field(field Field) *FieldCollapseBuilder {
	rb.v.Field = field
	return rb
}

// InnerHits set the InnerHits property for FieldCollapseBuilder.
func (rb *FieldCollapseBuilder) InnerHits(arg []InnerHits) *FieldCollapseBuilder {
	rb.v.InnerHits = arg
	return rb
}

// MaxConcurrentGroupSearches set the MaxConcurrentGroupSearches property for FieldCollapseBuilder.
func (rb *FieldCollapseBuilder) MaxConcurrentGroupSearches(maxconcurrentgroupsearches int) *FieldCollapseBuilder {
	rb.v.MaxConcurrentGroupSearches = &maxconcurrentgroupsearches
	return rb
}
