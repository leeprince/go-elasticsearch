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

// PhraseSuggester type.
type PhraseSuggester struct {
	Analyzer                *string                  `json:"analyzer,omitempty"`
	Collate                 *PhraseSuggestCollate    `json:"collate,omitempty"`
	Confidence              *float64                 `json:"confidence,omitempty"`
	DirectGenerator         []DirectGenerator        `json:"direct_generator,omitempty"`
	Field                   Field                    `json:"field"`
	ForceUnigrams           *bool                    `json:"force_unigrams,omitempty"`
	GramSize                *int                     `json:"gram_size,omitempty"`
	Highlight               *PhraseSuggestHighlight  `json:"highlight,omitempty"`
	MaxErrors               *float64                 `json:"max_errors,omitempty"`
	RealWordErrorLikelihood *float64                 `json:"real_word_error_likelihood,omitempty"`
	Separator               *string                  `json:"separator,omitempty"`
	ShardSize               *int                     `json:"shard_size,omitempty"`
	Size                    *int                     `json:"size,omitempty"`
	Smoothing               *SmoothingModelContainer `json:"smoothing,omitempty"`
	Text                    *string                  `json:"text,omitempty"`
	TokenLimit              *int                     `json:"token_limit,omitempty"`
}

// PhraseSuggesterBuilder holds PhraseSuggester struct and provides a builder API.
type PhraseSuggesterBuilder struct {
	v *PhraseSuggester
}

// NewPhraseSuggester provides a builder for the PhraseSuggester struct.
func NewPhraseSuggester() *PhraseSuggesterBuilder {
	r := PhraseSuggesterBuilder{
		&PhraseSuggester{},
	}

	return &r
}

// Build finalize the chain and returns the PhraseSuggester struct
func (rb *PhraseSuggesterBuilder) Build() PhraseSuggester {
	return *rb.v
}

// Analyzer set the Analyzer property for PhraseSuggesterBuilder.
func (rb *PhraseSuggesterBuilder) Analyzer(analyzer string) *PhraseSuggesterBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

// Collate set the Collate property for PhraseSuggesterBuilder.
func (rb *PhraseSuggesterBuilder) Collate(collate PhraseSuggestCollate) *PhraseSuggesterBuilder {
	rb.v.Collate = &collate
	return rb
}

// Confidence set the Confidence property for PhraseSuggesterBuilder.
func (rb *PhraseSuggesterBuilder) Confidence(confidence float64) *PhraseSuggesterBuilder {
	rb.v.Confidence = &confidence
	return rb
}

// DirectGenerator set the DirectGenerator property for PhraseSuggesterBuilder.
func (rb *PhraseSuggesterBuilder) DirectGenerator(direct_generator ...DirectGenerator) *PhraseSuggesterBuilder {
	rb.v.DirectGenerator = append(rb.v.DirectGenerator, direct_generator...)
	return rb
}

// Field set the Field property for PhraseSuggesterBuilder.
func (rb *PhraseSuggesterBuilder) Field(field Field) *PhraseSuggesterBuilder {
	rb.v.Field = field
	return rb
}

// ForceUnigrams set the ForceUnigrams property for PhraseSuggesterBuilder.
func (rb *PhraseSuggesterBuilder) ForceUnigrams(forceunigrams bool) *PhraseSuggesterBuilder {
	rb.v.ForceUnigrams = &forceunigrams
	return rb
}

// GramSize set the GramSize property for PhraseSuggesterBuilder.
func (rb *PhraseSuggesterBuilder) GramSize(gramsize int) *PhraseSuggesterBuilder {
	rb.v.GramSize = &gramsize
	return rb
}

// Highlight set the Highlight property for PhraseSuggesterBuilder.
func (rb *PhraseSuggesterBuilder) Highlight(highlight PhraseSuggestHighlight) *PhraseSuggesterBuilder {
	rb.v.Highlight = &highlight
	return rb
}

// MaxErrors set the MaxErrors property for PhraseSuggesterBuilder.
func (rb *PhraseSuggesterBuilder) MaxErrors(maxerrors float64) *PhraseSuggesterBuilder {
	rb.v.MaxErrors = &maxerrors
	return rb
}

// RealWordErrorLikelihood set the RealWordErrorLikelihood property for PhraseSuggesterBuilder.
func (rb *PhraseSuggesterBuilder) RealWordErrorLikelihood(realworderrorlikelihood float64) *PhraseSuggesterBuilder {
	rb.v.RealWordErrorLikelihood = &realworderrorlikelihood
	return rb
}

// Separator set the Separator property for PhraseSuggesterBuilder.
func (rb *PhraseSuggesterBuilder) Separator(separator string) *PhraseSuggesterBuilder {
	rb.v.Separator = &separator
	return rb
}

// ShardSize set the ShardSize property for PhraseSuggesterBuilder.
func (rb *PhraseSuggesterBuilder) ShardSize(shardsize int) *PhraseSuggesterBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}

// Size set the Size property for PhraseSuggesterBuilder.
func (rb *PhraseSuggesterBuilder) Size(size int) *PhraseSuggesterBuilder {
	rb.v.Size = &size
	return rb
}

// Smoothing set the Smoothing property for PhraseSuggesterBuilder.
func (rb *PhraseSuggesterBuilder) Smoothing(smoothing SmoothingModelContainer) *PhraseSuggesterBuilder {
	rb.v.Smoothing = &smoothing
	return rb
}

// Text set the Text property for PhraseSuggesterBuilder.
func (rb *PhraseSuggesterBuilder) Text(text string) *PhraseSuggesterBuilder {
	rb.v.Text = &text
	return rb
}

// TokenLimit set the TokenLimit property for PhraseSuggesterBuilder.
func (rb *PhraseSuggesterBuilder) TokenLimit(tokenlimit int) *PhraseSuggesterBuilder {
	rb.v.TokenLimit = &tokenlimit
	return rb
}
