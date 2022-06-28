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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
)

// MultiMatchQuery type.
type MultiMatchQuery struct {
	Analyzer                        *string                        `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery *bool                          `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost                           *float32                       `json:"boost,omitempty"`
	CutoffFrequency                 *float64                       `json:"cutoff_frequency,omitempty"`
	Fields                          *Fields                        `json:"fields,omitempty"`
	Fuzziness                       *Fuzziness                     `json:"fuzziness,omitempty"`
	FuzzyRewrite                    *MultiTermQueryRewrite         `json:"fuzzy_rewrite,omitempty"`
	FuzzyTranspositions             *bool                          `json:"fuzzy_transpositions,omitempty"`
	Lenient                         *bool                          `json:"lenient,omitempty"`
	MaxExpansions                   *int                           `json:"max_expansions,omitempty"`
	MinimumShouldMatch              *MinimumShouldMatch            `json:"minimum_should_match,omitempty"`
	Operator                        *operator.Operator             `json:"operator,omitempty"`
	PrefixLength                    *int                           `json:"prefix_length,omitempty"`
	Query                           string                         `json:"query"`
	QueryName_                      *string                        `json:"_name,omitempty"`
	Slop                            *int                           `json:"slop,omitempty"`
	TieBreaker                      *float64                       `json:"tie_breaker,omitempty"`
	Type                            *textquerytype.TextQueryType   `json:"type,omitempty"`
	ZeroTermsQuery                  *zerotermsquery.ZeroTermsQuery `json:"zero_terms_query,omitempty"`
}

// MultiMatchQueryBuilder holds MultiMatchQuery struct and provides a builder API.
type MultiMatchQueryBuilder struct {
	v *MultiMatchQuery
}

// NewMultiMatchQuery provides a builder for the MultiMatchQuery struct.
func NewMultiMatchQuery() *MultiMatchQueryBuilder {
	r := MultiMatchQueryBuilder{
		&MultiMatchQuery{},
	}

	return &r
}

// Build finalize the chain and returns the MultiMatchQuery struct
func (rb *MultiMatchQueryBuilder) Build() MultiMatchQuery {
	return *rb.v
}

// Analyzer set the Analyzer property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) Analyzer(analyzer string) *MultiMatchQueryBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

// AutoGenerateSynonymsPhraseQuery set the AutoGenerateSynonymsPhraseQuery property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *MultiMatchQueryBuilder {
	rb.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery
	return rb
}

// Boost set the Boost property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) Boost(boost float32) *MultiMatchQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// CutoffFrequency set the CutoffFrequency property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) CutoffFrequency(cutofffrequency float64) *MultiMatchQueryBuilder {
	rb.v.CutoffFrequency = &cutofffrequency
	return rb
}

// Fields set the Fields property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) Fields(fields Fields) *MultiMatchQueryBuilder {
	rb.v.Fields = &fields
	return rb
}

// Fuzziness set the Fuzziness property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) Fuzziness(fuzziness Fuzziness) *MultiMatchQueryBuilder {
	rb.v.Fuzziness = &fuzziness
	return rb
}

// FuzzyRewrite set the FuzzyRewrite property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) FuzzyRewrite(fuzzyrewrite MultiTermQueryRewrite) *MultiMatchQueryBuilder {
	rb.v.FuzzyRewrite = &fuzzyrewrite
	return rb
}

// FuzzyTranspositions set the FuzzyTranspositions property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) FuzzyTranspositions(fuzzytranspositions bool) *MultiMatchQueryBuilder {
	rb.v.FuzzyTranspositions = &fuzzytranspositions
	return rb
}

// Lenient set the Lenient property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) Lenient(lenient bool) *MultiMatchQueryBuilder {
	rb.v.Lenient = &lenient
	return rb
}

// MaxExpansions set the MaxExpansions property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) MaxExpansions(maxexpansions int) *MultiMatchQueryBuilder {
	rb.v.MaxExpansions = &maxexpansions
	return rb
}

// MinimumShouldMatch set the MinimumShouldMatch property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) MinimumShouldMatch(minimumshouldmatch MinimumShouldMatch) *MultiMatchQueryBuilder {
	rb.v.MinimumShouldMatch = &minimumshouldmatch
	return rb
}

// Operator set the Operator property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) Operator(operator operator.Operator) *MultiMatchQueryBuilder {
	rb.v.Operator = &operator
	return rb
}

// PrefixLength set the PrefixLength property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) PrefixLength(prefixlength int) *MultiMatchQueryBuilder {
	rb.v.PrefixLength = &prefixlength
	return rb
}

// Query set the Query property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) Query(query string) *MultiMatchQueryBuilder {
	rb.v.Query = query
	return rb
}

// QueryName_ set the QueryName_ property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) QueryName_(queryname_ string) *MultiMatchQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// Slop set the Slop property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) Slop(slop int) *MultiMatchQueryBuilder {
	rb.v.Slop = &slop
	return rb
}

// TieBreaker set the TieBreaker property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) TieBreaker(tiebreaker float64) *MultiMatchQueryBuilder {
	rb.v.TieBreaker = &tiebreaker
	return rb
}

// Type set the Type property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) Type_(type_ textquerytype.TextQueryType) *MultiMatchQueryBuilder {
	rb.v.Type = &type_
	return rb
}

// ZeroTermsQuery set the ZeroTermsQuery property for MultiMatchQueryBuilder.
func (rb *MultiMatchQueryBuilder) ZeroTermsQuery(zerotermsquery zerotermsquery.ZeroTermsQuery) *MultiMatchQueryBuilder {
	rb.v.ZeroTermsQuery = &zerotermsquery
	return rb
}
