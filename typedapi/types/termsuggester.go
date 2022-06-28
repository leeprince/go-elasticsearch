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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/stringdistance"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/suggestmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/suggestsort"
)

// TermSuggester type.
type TermSuggester struct {
	Analyzer       *string                        `json:"analyzer,omitempty"`
	Field          Field                          `json:"field"`
	LowercaseTerms *bool                          `json:"lowercase_terms,omitempty"`
	MaxEdits       *int                           `json:"max_edits,omitempty"`
	MaxInspections *int                           `json:"max_inspections,omitempty"`
	MaxTermFreq    *float32                       `json:"max_term_freq,omitempty"`
	MinDocFreq     *float32                       `json:"min_doc_freq,omitempty"`
	MinWordLength  *int                           `json:"min_word_length,omitempty"`
	PrefixLength   *int                           `json:"prefix_length,omitempty"`
	ShardSize      *int                           `json:"shard_size,omitempty"`
	Size           *int                           `json:"size,omitempty"`
	Sort           *suggestsort.SuggestSort       `json:"sort,omitempty"`
	StringDistance *stringdistance.StringDistance `json:"string_distance,omitempty"`
	SuggestMode    *suggestmode.SuggestMode       `json:"suggest_mode,omitempty"`
	Text           *string                        `json:"text,omitempty"`
}

// TermSuggesterBuilder holds TermSuggester struct and provides a builder API.
type TermSuggesterBuilder struct {
	v *TermSuggester
}

// NewTermSuggester provides a builder for the TermSuggester struct.
func NewTermSuggester() *TermSuggesterBuilder {
	r := TermSuggesterBuilder{
		&TermSuggester{},
	}

	return &r
}

// Build finalize the chain and returns the TermSuggester struct
func (rb *TermSuggesterBuilder) Build() TermSuggester {
	return *rb.v
}

// Analyzer set the Analyzer property for TermSuggesterBuilder.
func (rb *TermSuggesterBuilder) Analyzer(analyzer string) *TermSuggesterBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

// Field set the Field property for TermSuggesterBuilder.
func (rb *TermSuggesterBuilder) Field(field Field) *TermSuggesterBuilder {
	rb.v.Field = field
	return rb
}

// LowercaseTerms set the LowercaseTerms property for TermSuggesterBuilder.
func (rb *TermSuggesterBuilder) LowercaseTerms(lowercaseterms bool) *TermSuggesterBuilder {
	rb.v.LowercaseTerms = &lowercaseterms
	return rb
}

// MaxEdits set the MaxEdits property for TermSuggesterBuilder.
func (rb *TermSuggesterBuilder) MaxEdits(maxedits int) *TermSuggesterBuilder {
	rb.v.MaxEdits = &maxedits
	return rb
}

// MaxInspections set the MaxInspections property for TermSuggesterBuilder.
func (rb *TermSuggesterBuilder) MaxInspections(maxinspections int) *TermSuggesterBuilder {
	rb.v.MaxInspections = &maxinspections
	return rb
}

// MaxTermFreq set the MaxTermFreq property for TermSuggesterBuilder.
func (rb *TermSuggesterBuilder) MaxTermFreq(maxtermfreq float32) *TermSuggesterBuilder {
	rb.v.MaxTermFreq = &maxtermfreq
	return rb
}

// MinDocFreq set the MinDocFreq property for TermSuggesterBuilder.
func (rb *TermSuggesterBuilder) MinDocFreq(mindocfreq float32) *TermSuggesterBuilder {
	rb.v.MinDocFreq = &mindocfreq
	return rb
}

// MinWordLength set the MinWordLength property for TermSuggesterBuilder.
func (rb *TermSuggesterBuilder) MinWordLength(minwordlength int) *TermSuggesterBuilder {
	rb.v.MinWordLength = &minwordlength
	return rb
}

// PrefixLength set the PrefixLength property for TermSuggesterBuilder.
func (rb *TermSuggesterBuilder) PrefixLength(prefixlength int) *TermSuggesterBuilder {
	rb.v.PrefixLength = &prefixlength
	return rb
}

// ShardSize set the ShardSize property for TermSuggesterBuilder.
func (rb *TermSuggesterBuilder) ShardSize(shardsize int) *TermSuggesterBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}

// Size set the Size property for TermSuggesterBuilder.
func (rb *TermSuggesterBuilder) Size(size int) *TermSuggesterBuilder {
	rb.v.Size = &size
	return rb
}

// Sort set the Sort property for TermSuggesterBuilder.
func (rb *TermSuggesterBuilder) Sort(sort suggestsort.SuggestSort) *TermSuggesterBuilder {
	rb.v.Sort = &sort
	return rb
}

// StringDistance set the StringDistance property for TermSuggesterBuilder.
func (rb *TermSuggesterBuilder) StringDistance(stringdistance stringdistance.StringDistance) *TermSuggesterBuilder {
	rb.v.StringDistance = &stringdistance
	return rb
}

// SuggestMode set the SuggestMode property for TermSuggesterBuilder.
func (rb *TermSuggesterBuilder) SuggestMode(suggestmode suggestmode.SuggestMode) *TermSuggesterBuilder {
	rb.v.SuggestMode = &suggestmode
	return rb
}

// Text set the Text property for TermSuggesterBuilder.
func (rb *TermSuggesterBuilder) Text(text string) *TermSuggesterBuilder {
	rb.v.Text = &text
	return rb
}
