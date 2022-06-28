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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
)

// MatchPhrasePrefixQuery type.
type MatchPhrasePrefixQuery struct {
	Analyzer       *string                        `json:"analyzer,omitempty"`
	Boost          *float32                       `json:"boost,omitempty"`
	MaxExpansions  *int                           `json:"max_expansions,omitempty"`
	Query          string                         `json:"query"`
	QueryName_     *string                        `json:"_name,omitempty"`
	Slop           *int                           `json:"slop,omitempty"`
	ZeroTermsQuery *zerotermsquery.ZeroTermsQuery `json:"zero_terms_query,omitempty"`
}

// MatchPhrasePrefixQueryBuilder holds MatchPhrasePrefixQuery struct and provides a builder API.
type MatchPhrasePrefixQueryBuilder struct {
	v *MatchPhrasePrefixQuery
}

// NewMatchPhrasePrefixQuery provides a builder for the MatchPhrasePrefixQuery struct.
func NewMatchPhrasePrefixQuery() *MatchPhrasePrefixQueryBuilder {
	r := MatchPhrasePrefixQueryBuilder{
		&MatchPhrasePrefixQuery{},
	}

	return &r
}

// Build finalize the chain and returns the MatchPhrasePrefixQuery struct
func (rb *MatchPhrasePrefixQueryBuilder) Build() MatchPhrasePrefixQuery {
	return *rb.v
}

// Analyzer set the Analyzer property for MatchPhrasePrefixQueryBuilder.
func (rb *MatchPhrasePrefixQueryBuilder) Analyzer(analyzer string) *MatchPhrasePrefixQueryBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

// Boost set the Boost property for MatchPhrasePrefixQueryBuilder.
func (rb *MatchPhrasePrefixQueryBuilder) Boost(boost float32) *MatchPhrasePrefixQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// MaxExpansions set the MaxExpansions property for MatchPhrasePrefixQueryBuilder.
func (rb *MatchPhrasePrefixQueryBuilder) MaxExpansions(maxexpansions int) *MatchPhrasePrefixQueryBuilder {
	rb.v.MaxExpansions = &maxexpansions
	return rb
}

// Query set the Query property for MatchPhrasePrefixQueryBuilder.
func (rb *MatchPhrasePrefixQueryBuilder) Query(query string) *MatchPhrasePrefixQueryBuilder {
	rb.v.Query = query
	return rb
}

// QueryName_ set the QueryName_ property for MatchPhrasePrefixQueryBuilder.
func (rb *MatchPhrasePrefixQueryBuilder) QueryName_(queryname_ string) *MatchPhrasePrefixQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// Slop set the Slop property for MatchPhrasePrefixQueryBuilder.
func (rb *MatchPhrasePrefixQueryBuilder) Slop(slop int) *MatchPhrasePrefixQueryBuilder {
	rb.v.Slop = &slop
	return rb
}

// ZeroTermsQuery set the ZeroTermsQuery property for MatchPhrasePrefixQueryBuilder.
func (rb *MatchPhrasePrefixQueryBuilder) ZeroTermsQuery(zerotermsquery zerotermsquery.ZeroTermsQuery) *MatchPhrasePrefixQueryBuilder {
	rb.v.ZeroTermsQuery = &zerotermsquery
	return rb
}
