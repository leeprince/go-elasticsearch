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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// MoreLikeThisQuery type.
type MoreLikeThisQuery struct {
	Analyzer               *string                  `json:"analyzer,omitempty"`
	Boost                  *float32                 `json:"boost,omitempty"`
	BoostTerms             *float64                 `json:"boost_terms,omitempty"`
	FailOnUnsupportedField *bool                    `json:"fail_on_unsupported_field,omitempty"`
	Fields                 []Field                  `json:"fields,omitempty"`
	Include                *bool                    `json:"include,omitempty"`
	Like                   []Like                   `json:"like"`
	MaxDocFreq             *int                     `json:"max_doc_freq,omitempty"`
	MaxQueryTerms          *int                     `json:"max_query_terms,omitempty"`
	MaxWordLength          *int                     `json:"max_word_length,omitempty"`
	MinDocFreq             *int                     `json:"min_doc_freq,omitempty"`
	MinTermFreq            *int                     `json:"min_term_freq,omitempty"`
	MinWordLength          *int                     `json:"min_word_length,omitempty"`
	MinimumShouldMatch     *MinimumShouldMatch      `json:"minimum_should_match,omitempty"`
	PerFieldAnalyzer       map[Field]string         `json:"per_field_analyzer,omitempty"`
	QueryName_             *string                  `json:"_name,omitempty"`
	Routing                *Routing                 `json:"routing,omitempty"`
	StopWords              *StopWords               `json:"stop_words,omitempty"`
	Unlike                 []Like                   `json:"unlike,omitempty"`
	Version                *VersionNumber           `json:"version,omitempty"`
	VersionType            *versiontype.VersionType `json:"version_type,omitempty"`
}

// MoreLikeThisQueryBuilder holds MoreLikeThisQuery struct and provides a builder API.
type MoreLikeThisQueryBuilder struct {
	v *MoreLikeThisQuery
}

// NewMoreLikeThisQuery provides a builder for the MoreLikeThisQuery struct.
func NewMoreLikeThisQuery() *MoreLikeThisQueryBuilder {
	r := MoreLikeThisQueryBuilder{
		&MoreLikeThisQuery{
			PerFieldAnalyzer: make(map[Field]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the MoreLikeThisQuery struct
func (rb *MoreLikeThisQueryBuilder) Build() MoreLikeThisQuery {
	return *rb.v
}

// Analyzer set the Analyzer property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) Analyzer(analyzer string) *MoreLikeThisQueryBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

// Boost set the Boost property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) Boost(boost float32) *MoreLikeThisQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// BoostTerms set the BoostTerms property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) BoostTerms(boostterms float64) *MoreLikeThisQueryBuilder {
	rb.v.BoostTerms = &boostterms
	return rb
}

// FailOnUnsupportedField set the FailOnUnsupportedField property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) FailOnUnsupportedField(failonunsupportedfield bool) *MoreLikeThisQueryBuilder {
	rb.v.FailOnUnsupportedField = &failonunsupportedfield
	return rb
}

// Fields set the Fields property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) Fields(fields ...Field) *MoreLikeThisQueryBuilder {
	rb.v.Fields = append(rb.v.Fields, fields...)
	return rb
}

// Include set the Include property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) Include(include bool) *MoreLikeThisQueryBuilder {
	rb.v.Include = &include
	return rb
}

// Like set the Like property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) Like(arg []Like) *MoreLikeThisQueryBuilder {
	rb.v.Like = arg
	return rb
}

// MaxDocFreq set the MaxDocFreq property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) MaxDocFreq(maxdocfreq int) *MoreLikeThisQueryBuilder {
	rb.v.MaxDocFreq = &maxdocfreq
	return rb
}

// MaxQueryTerms set the MaxQueryTerms property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) MaxQueryTerms(maxqueryterms int) *MoreLikeThisQueryBuilder {
	rb.v.MaxQueryTerms = &maxqueryterms
	return rb
}

// MaxWordLength set the MaxWordLength property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) MaxWordLength(maxwordlength int) *MoreLikeThisQueryBuilder {
	rb.v.MaxWordLength = &maxwordlength
	return rb
}

// MinDocFreq set the MinDocFreq property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) MinDocFreq(mindocfreq int) *MoreLikeThisQueryBuilder {
	rb.v.MinDocFreq = &mindocfreq
	return rb
}

// MinTermFreq set the MinTermFreq property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) MinTermFreq(mintermfreq int) *MoreLikeThisQueryBuilder {
	rb.v.MinTermFreq = &mintermfreq
	return rb
}

// MinWordLength set the MinWordLength property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) MinWordLength(minwordlength int) *MoreLikeThisQueryBuilder {
	rb.v.MinWordLength = &minwordlength
	return rb
}

// MinimumShouldMatch set the MinimumShouldMatch property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) MinimumShouldMatch(minimumshouldmatch MinimumShouldMatch) *MoreLikeThisQueryBuilder {
	rb.v.MinimumShouldMatch = &minimumshouldmatch
	return rb
}

// PerFieldAnalyzer set the PerFieldAnalyzer property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) PerFieldAnalyzer(value map[Field]string) *MoreLikeThisQueryBuilder {
	rb.v.PerFieldAnalyzer = value
	return rb
}

// QueryName_ set the QueryName_ property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) QueryName_(queryname_ string) *MoreLikeThisQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// Routing set the Routing property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) Routing(routing Routing) *MoreLikeThisQueryBuilder {
	rb.v.Routing = &routing
	return rb
}

// StopWords set the StopWords property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) StopWords(stopwords StopWords) *MoreLikeThisQueryBuilder {
	rb.v.StopWords = &stopwords
	return rb
}

// Unlike set the Unlike property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) Unlike(arg []Like) *MoreLikeThisQueryBuilder {
	rb.v.Unlike = arg
	return rb
}

// Version set the Version property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) Version(version VersionNumber) *MoreLikeThisQueryBuilder {
	rb.v.Version = &version
	return rb
}

// VersionType set the VersionType property for MoreLikeThisQueryBuilder.
func (rb *MoreLikeThisQueryBuilder) VersionType(versiontype versiontype.VersionType) *MoreLikeThisQueryBuilder {
	rb.v.VersionType = &versiontype
	return rb
}
