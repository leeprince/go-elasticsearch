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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
)

// SimpleQueryStringQuery type.
type SimpleQueryStringQuery struct {
	AnalyzeWildcard                 *bool                   `json:"analyze_wildcard,omitempty"`
	Analyzer                        *string                 `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery *bool                   `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost                           *float32                `json:"boost,omitempty"`
	DefaultOperator                 *operator.Operator      `json:"default_operator,omitempty"`
	Fields                          []Field                 `json:"fields,omitempty"`
	Flags                           *SimpleQueryStringFlags `json:"flags,omitempty"`
	FuzzyMaxExpansions              *int                    `json:"fuzzy_max_expansions,omitempty"`
	FuzzyPrefixLength               *int                    `json:"fuzzy_prefix_length,omitempty"`
	FuzzyTranspositions             *bool                   `json:"fuzzy_transpositions,omitempty"`
	Lenient                         *bool                   `json:"lenient,omitempty"`
	MinimumShouldMatch              *MinimumShouldMatch     `json:"minimum_should_match,omitempty"`
	Query                           string                  `json:"query"`
	QueryName_                      *string                 `json:"_name,omitempty"`
	QuoteFieldSuffix                *string                 `json:"quote_field_suffix,omitempty"`
}

// SimpleQueryStringQueryBuilder holds SimpleQueryStringQuery struct and provides a builder API.
type SimpleQueryStringQueryBuilder struct {
	v *SimpleQueryStringQuery
}

// NewSimpleQueryStringQuery provides a builder for the SimpleQueryStringQuery struct.
func NewSimpleQueryStringQuery() *SimpleQueryStringQueryBuilder {
	r := SimpleQueryStringQueryBuilder{
		&SimpleQueryStringQuery{},
	}

	return &r
}

// Build finalize the chain and returns the SimpleQueryStringQuery struct
func (rb *SimpleQueryStringQueryBuilder) Build() SimpleQueryStringQuery {
	return *rb.v
}

// AnalyzeWildcard set the AnalyzeWildcard property for SimpleQueryStringQueryBuilder.
func (rb *SimpleQueryStringQueryBuilder) AnalyzeWildcard(analyzewildcard bool) *SimpleQueryStringQueryBuilder {
	rb.v.AnalyzeWildcard = &analyzewildcard
	return rb
}

// Analyzer set the Analyzer property for SimpleQueryStringQueryBuilder.
func (rb *SimpleQueryStringQueryBuilder) Analyzer(analyzer string) *SimpleQueryStringQueryBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

// AutoGenerateSynonymsPhraseQuery set the AutoGenerateSynonymsPhraseQuery property for SimpleQueryStringQueryBuilder.
func (rb *SimpleQueryStringQueryBuilder) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *SimpleQueryStringQueryBuilder {
	rb.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery
	return rb
}

// Boost set the Boost property for SimpleQueryStringQueryBuilder.
func (rb *SimpleQueryStringQueryBuilder) Boost(boost float32) *SimpleQueryStringQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// DefaultOperator set the DefaultOperator property for SimpleQueryStringQueryBuilder.
func (rb *SimpleQueryStringQueryBuilder) DefaultOperator(defaultoperator operator.Operator) *SimpleQueryStringQueryBuilder {
	rb.v.DefaultOperator = &defaultoperator
	return rb
}

// Fields set the Fields property for SimpleQueryStringQueryBuilder.
func (rb *SimpleQueryStringQueryBuilder) Fields(fields ...Field) *SimpleQueryStringQueryBuilder {
	rb.v.Fields = append(rb.v.Fields, fields...)
	return rb
}

// Flags set the Flags property for SimpleQueryStringQueryBuilder.
func (rb *SimpleQueryStringQueryBuilder) Flags(flags SimpleQueryStringFlags) *SimpleQueryStringQueryBuilder {
	rb.v.Flags = &flags
	return rb
}

// FuzzyMaxExpansions set the FuzzyMaxExpansions property for SimpleQueryStringQueryBuilder.
func (rb *SimpleQueryStringQueryBuilder) FuzzyMaxExpansions(fuzzymaxexpansions int) *SimpleQueryStringQueryBuilder {
	rb.v.FuzzyMaxExpansions = &fuzzymaxexpansions
	return rb
}

// FuzzyPrefixLength set the FuzzyPrefixLength property for SimpleQueryStringQueryBuilder.
func (rb *SimpleQueryStringQueryBuilder) FuzzyPrefixLength(fuzzyprefixlength int) *SimpleQueryStringQueryBuilder {
	rb.v.FuzzyPrefixLength = &fuzzyprefixlength
	return rb
}

// FuzzyTranspositions set the FuzzyTranspositions property for SimpleQueryStringQueryBuilder.
func (rb *SimpleQueryStringQueryBuilder) FuzzyTranspositions(fuzzytranspositions bool) *SimpleQueryStringQueryBuilder {
	rb.v.FuzzyTranspositions = &fuzzytranspositions
	return rb
}

// Lenient set the Lenient property for SimpleQueryStringQueryBuilder.
func (rb *SimpleQueryStringQueryBuilder) Lenient(lenient bool) *SimpleQueryStringQueryBuilder {
	rb.v.Lenient = &lenient
	return rb
}

// MinimumShouldMatch set the MinimumShouldMatch property for SimpleQueryStringQueryBuilder.
func (rb *SimpleQueryStringQueryBuilder) MinimumShouldMatch(minimumshouldmatch MinimumShouldMatch) *SimpleQueryStringQueryBuilder {
	rb.v.MinimumShouldMatch = &minimumshouldmatch
	return rb
}

// Query set the Query property for SimpleQueryStringQueryBuilder.
func (rb *SimpleQueryStringQueryBuilder) Query(query string) *SimpleQueryStringQueryBuilder {
	rb.v.Query = query
	return rb
}

// QueryName_ set the QueryName_ property for SimpleQueryStringQueryBuilder.
func (rb *SimpleQueryStringQueryBuilder) QueryName_(queryname_ string) *SimpleQueryStringQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// QuoteFieldSuffix set the QuoteFieldSuffix property for SimpleQueryStringQueryBuilder.
func (rb *SimpleQueryStringQueryBuilder) QuoteFieldSuffix(quotefieldsuffix string) *SimpleQueryStringQueryBuilder {
	rb.v.QuoteFieldSuffix = &quotefieldsuffix
	return rb
}
