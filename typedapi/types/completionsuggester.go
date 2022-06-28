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

// CompletionSuggester type.
type CompletionSuggester struct {
	Analyzer       *string                       `json:"analyzer,omitempty"`
	Contexts       map[Field][]CompletionContext `json:"contexts,omitempty"`
	Field          Field                         `json:"field"`
	Fuzzy          *SuggestFuzziness             `json:"fuzzy,omitempty"`
	Prefix         *string                       `json:"prefix,omitempty"`
	Regex          *string                       `json:"regex,omitempty"`
	Size           *int                          `json:"size,omitempty"`
	SkipDuplicates *bool                         `json:"skip_duplicates,omitempty"`
}

// CompletionSuggesterBuilder holds CompletionSuggester struct and provides a builder API.
type CompletionSuggesterBuilder struct {
	v *CompletionSuggester
}

// NewCompletionSuggester provides a builder for the CompletionSuggester struct.
func NewCompletionSuggester() *CompletionSuggesterBuilder {
	r := CompletionSuggesterBuilder{
		&CompletionSuggester{
			Contexts: make(map[Field][]CompletionContext, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the CompletionSuggester struct
func (rb *CompletionSuggesterBuilder) Build() CompletionSuggester {
	return *rb.v
}

// Analyzer set the Analyzer property for CompletionSuggesterBuilder.
func (rb *CompletionSuggesterBuilder) Analyzer(analyzer string) *CompletionSuggesterBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

// Contexts set the Contexts property for CompletionSuggesterBuilder.
//TODO: Union in map Not implemented yet: UnionOf
func (rb *CompletionSuggesterBuilder) Contexts(value map[Field][]CompletionContext) *CompletionSuggesterBuilder {
	rb.v.Contexts = value
	return rb
}

// Field set the Field property for CompletionSuggesterBuilder.
func (rb *CompletionSuggesterBuilder) Field(field Field) *CompletionSuggesterBuilder {
	rb.v.Field = field
	return rb
}

// Fuzzy set the Fuzzy property for CompletionSuggesterBuilder.
func (rb *CompletionSuggesterBuilder) Fuzzy(fuzzy SuggestFuzziness) *CompletionSuggesterBuilder {
	rb.v.Fuzzy = &fuzzy
	return rb
}

// Prefix set the Prefix property for CompletionSuggesterBuilder.
func (rb *CompletionSuggesterBuilder) Prefix(prefix string) *CompletionSuggesterBuilder {
	rb.v.Prefix = &prefix
	return rb
}

// Regex set the Regex property for CompletionSuggesterBuilder.
func (rb *CompletionSuggesterBuilder) Regex(regex string) *CompletionSuggesterBuilder {
	rb.v.Regex = &regex
	return rb
}

// Size set the Size property for CompletionSuggesterBuilder.
func (rb *CompletionSuggesterBuilder) Size(size int) *CompletionSuggesterBuilder {
	rb.v.Size = &size
	return rb
}

// SkipDuplicates set the SkipDuplicates property for CompletionSuggesterBuilder.
func (rb *CompletionSuggesterBuilder) SkipDuplicates(skipduplicates bool) *CompletionSuggesterBuilder {
	rb.v.SkipDuplicates = &skipduplicates
	return rb
}
