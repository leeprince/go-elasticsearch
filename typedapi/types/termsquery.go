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

// TermsQuery type.
type TermsQuery struct {
	Boost      *float32                  `json:"boost,omitempty"`
	QueryName_ *string                   `json:"_name,omitempty"`
	TermsQuery map[Field]TermsQueryField `json:"TermsQuery,omitempty"`
}

// TermsQueryBuilder holds TermsQuery struct and provides a builder API.
type TermsQueryBuilder struct {
	v *TermsQuery
}

// NewTermsQuery provides a builder for the TermsQuery struct.
func NewTermsQuery() *TermsQueryBuilder {
	r := TermsQueryBuilder{
		&TermsQuery{
			TermsQuery: make(map[Field]TermsQueryField, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TermsQuery struct
func (rb *TermsQueryBuilder) Build() TermsQuery {
	return *rb.v
}

// Boost set the Boost property for TermsQueryBuilder.
func (rb *TermsQueryBuilder) Boost(boost float32) *TermsQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// QueryName_ set the QueryName_ property for TermsQueryBuilder.
func (rb *TermsQueryBuilder) QueryName_(queryname_ string) *TermsQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// TermsQuery set the TermsQuery property for TermsQueryBuilder.
func (rb *TermsQueryBuilder) TermsQuery(value map[Field]TermsQueryField) *TermsQueryBuilder {
	rb.v.TermsQuery = value
	return rb
}
