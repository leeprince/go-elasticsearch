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

// ParentIdQuery type.
type ParentIdQuery struct {
	Boost          *float32      `json:"boost,omitempty"`
	Id             *Id           `json:"id,omitempty"`
	IgnoreUnmapped *bool         `json:"ignore_unmapped,omitempty"`
	QueryName_     *string       `json:"_name,omitempty"`
	Type           *RelationName `json:"type,omitempty"`
}

// ParentIdQueryBuilder holds ParentIdQuery struct and provides a builder API.
type ParentIdQueryBuilder struct {
	v *ParentIdQuery
}

// NewParentIdQuery provides a builder for the ParentIdQuery struct.
func NewParentIdQuery() *ParentIdQueryBuilder {
	r := ParentIdQueryBuilder{
		&ParentIdQuery{},
	}

	return &r
}

// Build finalize the chain and returns the ParentIdQuery struct
func (rb *ParentIdQueryBuilder) Build() ParentIdQuery {
	return *rb.v
}

// Boost set the Boost property for ParentIdQueryBuilder.
func (rb *ParentIdQueryBuilder) Boost(boost float32) *ParentIdQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// Id set the Id property for ParentIdQueryBuilder.
func (rb *ParentIdQueryBuilder) Id(id Id) *ParentIdQueryBuilder {
	rb.v.Id = &id
	return rb
}

// IgnoreUnmapped set the IgnoreUnmapped property for ParentIdQueryBuilder.
func (rb *ParentIdQueryBuilder) IgnoreUnmapped(ignoreunmapped bool) *ParentIdQueryBuilder {
	rb.v.IgnoreUnmapped = &ignoreunmapped
	return rb
}

// QueryName_ set the QueryName_ property for ParentIdQueryBuilder.
func (rb *ParentIdQueryBuilder) QueryName_(queryname_ string) *ParentIdQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// Type set the Type property for ParentIdQueryBuilder.
func (rb *ParentIdQueryBuilder) Type_(type_ RelationName) *ParentIdQueryBuilder {
	rb.v.Type = &type_
	return rb
}
