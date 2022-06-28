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

// Suggest holds the union for the following types:
//     CompletionSuggest
//     PhraseSuggest
//     TermSuggest
type Suggest interface{}

// SuggestBuilder holds Suggest struct and provides a builder API.
type SuggestBuilder struct {
	v Suggest
}

// NewSuggest provides a builder for the Suggest struct.
func NewSuggest() *SuggestBuilder {
	return &SuggestBuilder{}
}

// Build finalize the chain and returns the Suggest struct
func (u *SuggestBuilder) Build() Suggest {
	return u.v
}

// CompletionSuggest set the CompletionSuggest property for SuggestBuilder.
func (u *SuggestBuilder) CompletionSuggest(v CompletionSuggest) *SuggestBuilder {
	u.v = v
	return u
}

// PhraseSuggest set the PhraseSuggest property for SuggestBuilder.
func (u *SuggestBuilder) PhraseSuggest(v PhraseSuggest) *SuggestBuilder {
	u.v = v
	return u
}

// TermSuggest set the TermSuggest property for SuggestBuilder.
func (u *SuggestBuilder) TermSuggest(v TermSuggest) *SuggestBuilder {
	u.v = v
	return u
}
