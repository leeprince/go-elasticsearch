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

// SpanNearQuery type.
type SpanNearQuery struct {
	Boost      *float32    `json:"boost,omitempty"`
	Clauses    []SpanQuery `json:"clauses"`
	InOrder    *bool       `json:"in_order,omitempty"`
	QueryName_ *string     `json:"_name,omitempty"`
	Slop       *int        `json:"slop,omitempty"`
}

// SpanNearQueryBuilder holds SpanNearQuery struct and provides a builder API.
type SpanNearQueryBuilder struct {
	v *SpanNearQuery
}

// NewSpanNearQuery provides a builder for the SpanNearQuery struct.
func NewSpanNearQuery() *SpanNearQueryBuilder {
	r := SpanNearQueryBuilder{
		&SpanNearQuery{},
	}

	return &r
}

// Build finalize the chain and returns the SpanNearQuery struct
func (rb *SpanNearQueryBuilder) Build() SpanNearQuery {
	return *rb.v
}

// Boost set the Boost property for SpanNearQueryBuilder.
func (rb *SpanNearQueryBuilder) Boost(boost float32) *SpanNearQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// Clauses set the Clauses property for SpanNearQueryBuilder.
func (rb *SpanNearQueryBuilder) Clauses(clauses ...SpanQuery) *SpanNearQueryBuilder {
	rb.v.Clauses = append(rb.v.Clauses, clauses...)
	return rb
}

// InOrder set the InOrder property for SpanNearQueryBuilder.
func (rb *SpanNearQueryBuilder) InOrder(inorder bool) *SpanNearQueryBuilder {
	rb.v.InOrder = &inorder
	return rb
}

// QueryName_ set the QueryName_ property for SpanNearQueryBuilder.
func (rb *SpanNearQueryBuilder) QueryName_(queryname_ string) *SpanNearQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// Slop set the Slop property for SpanNearQueryBuilder.
func (rb *SpanNearQueryBuilder) Slop(slop int) *SpanNearQueryBuilder {
	rb.v.Slop = &slop
	return rb
}
