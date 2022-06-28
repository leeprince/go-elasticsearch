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

// CharGroupTokenizer type.
type CharGroupTokenizer struct {
	MaxTokenLength  *int           `json:"max_token_length,omitempty"`
	TokenizeOnChars []string       `json:"tokenize_on_chars"`
	Type            string         `json:"type,omitempty"`
	Version         *VersionString `json:"version,omitempty"`
}

// CharGroupTokenizerBuilder holds CharGroupTokenizer struct and provides a builder API.
type CharGroupTokenizerBuilder struct {
	v *CharGroupTokenizer
}

// NewCharGroupTokenizer provides a builder for the CharGroupTokenizer struct.
func NewCharGroupTokenizer() *CharGroupTokenizerBuilder {
	r := CharGroupTokenizerBuilder{
		&CharGroupTokenizer{},
	}

	r.v.Type = "char_group"

	return &r
}

// Build finalize the chain and returns the CharGroupTokenizer struct
func (rb *CharGroupTokenizerBuilder) Build() CharGroupTokenizer {
	return *rb.v
}

// MaxTokenLength set the MaxTokenLength property for CharGroupTokenizerBuilder.
func (rb *CharGroupTokenizerBuilder) MaxTokenLength(maxtokenlength int) *CharGroupTokenizerBuilder {
	rb.v.MaxTokenLength = &maxtokenlength
	return rb
}

// TokenizeOnChars set the TokenizeOnChars property for CharGroupTokenizerBuilder.
func (rb *CharGroupTokenizerBuilder) TokenizeOnChars(tokenize_on_chars ...string) *CharGroupTokenizerBuilder {
	rb.v.TokenizeOnChars = append(rb.v.TokenizeOnChars, tokenize_on_chars...)
	return rb
}

// Type set the Type property for CharGroupTokenizerBuilder.

// Version set the Version property for CharGroupTokenizerBuilder.
func (rb *CharGroupTokenizerBuilder) Version(version VersionString) *CharGroupTokenizerBuilder {
	rb.v.Version = &version
	return rb
}
