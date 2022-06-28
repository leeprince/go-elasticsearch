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

// ShapeQuery type.
type ShapeQuery struct {
	Boost          *float32                  `json:"boost,omitempty"`
	IgnoreUnmapped *bool                     `json:"ignore_unmapped,omitempty"`
	QueryName_     *string                   `json:"_name,omitempty"`
	ShapeQuery     map[Field]ShapeFieldQuery `json:"ShapeQuery,omitempty"`
}

// ShapeQueryBuilder holds ShapeQuery struct and provides a builder API.
type ShapeQueryBuilder struct {
	v *ShapeQuery
}

// NewShapeQuery provides a builder for the ShapeQuery struct.
func NewShapeQuery() *ShapeQueryBuilder {
	r := ShapeQueryBuilder{
		&ShapeQuery{
			ShapeQuery: make(map[Field]ShapeFieldQuery, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ShapeQuery struct
func (rb *ShapeQueryBuilder) Build() ShapeQuery {
	return *rb.v
}

// Boost set the Boost property for ShapeQueryBuilder.
func (rb *ShapeQueryBuilder) Boost(boost float32) *ShapeQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// IgnoreUnmapped set the IgnoreUnmapped property for ShapeQueryBuilder.
func (rb *ShapeQueryBuilder) IgnoreUnmapped(ignoreunmapped bool) *ShapeQueryBuilder {
	rb.v.IgnoreUnmapped = &ignoreunmapped
	return rb
}

// QueryName_ set the QueryName_ property for ShapeQueryBuilder.
func (rb *ShapeQueryBuilder) QueryName_(queryname_ string) *ShapeQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// ShapeQuery set the ShapeQuery property for ShapeQueryBuilder.
func (rb *ShapeQueryBuilder) ShapeQuery(value map[Field]ShapeFieldQuery) *ShapeQueryBuilder {
	rb.v.ShapeQuery = value
	return rb
}
