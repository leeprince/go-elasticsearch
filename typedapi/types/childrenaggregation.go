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

// ChildrenAggregation type.
type ChildrenAggregation struct {
	Meta *Metadata     `json:"meta,omitempty"`
	Name *string       `json:"name,omitempty"`
	Type *RelationName `json:"type,omitempty"`
}

// ChildrenAggregationBuilder holds ChildrenAggregation struct and provides a builder API.
type ChildrenAggregationBuilder struct {
	v *ChildrenAggregation
}

// NewChildrenAggregation provides a builder for the ChildrenAggregation struct.
func NewChildrenAggregation() *ChildrenAggregationBuilder {
	r := ChildrenAggregationBuilder{
		&ChildrenAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the ChildrenAggregation struct
func (rb *ChildrenAggregationBuilder) Build() ChildrenAggregation {
	return *rb.v
}

// Meta set the Meta property for ChildrenAggregationBuilder.
func (rb *ChildrenAggregationBuilder) Meta(meta Metadata) *ChildrenAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for ChildrenAggregationBuilder.
func (rb *ChildrenAggregationBuilder) Name(name string) *ChildrenAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Type set the Type property for ChildrenAggregationBuilder.
func (rb *ChildrenAggregationBuilder) Type_(type_ RelationName) *ChildrenAggregationBuilder {
	rb.v.Type = &type_
	return rb
}
