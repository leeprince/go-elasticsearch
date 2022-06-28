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

// TokenizerDefinition holds the union for the following types:
//     CharGroupTokenizer
//     EdgeNGramTokenizer
//     IcuTokenizer
//     KeywordTokenizer
//     KuromojiTokenizer
//     LetterTokenizer
//     LowercaseTokenizer
//     NGramTokenizer
//     NoriTokenizer
//     PathHierarchyTokenizer
//     PatternTokenizer
//     StandardTokenizer
//     UaxEmailUrlTokenizer
//     WhitespaceTokenizer
type TokenizerDefinition interface{}

// TokenizerDefinitionBuilder holds TokenizerDefinition struct and provides a builder API.
type TokenizerDefinitionBuilder struct {
	v TokenizerDefinition
}

// NewTokenizerDefinition provides a builder for the TokenizerDefinition struct.
func NewTokenizerDefinition() *TokenizerDefinitionBuilder {
	return &TokenizerDefinitionBuilder{}
}

// Build finalize the chain and returns the TokenizerDefinition struct
func (u *TokenizerDefinitionBuilder) Build() TokenizerDefinition {
	return u.v
}

// CharGroupTokenizer set the CharGroupTokenizer property for TokenizerDefinitionBuilder.
func (u *TokenizerDefinitionBuilder) CharGroupTokenizer(v CharGroupTokenizer) *TokenizerDefinitionBuilder {
	u.v = v
	return u
}

// EdgeNGramTokenizer set the EdgeNGramTokenizer property for TokenizerDefinitionBuilder.
func (u *TokenizerDefinitionBuilder) EdgeNGramTokenizer(v EdgeNGramTokenizer) *TokenizerDefinitionBuilder {
	u.v = v
	return u
}

// IcuTokenizer set the IcuTokenizer property for TokenizerDefinitionBuilder.
func (u *TokenizerDefinitionBuilder) IcuTokenizer(v IcuTokenizer) *TokenizerDefinitionBuilder {
	u.v = v
	return u
}

// KeywordTokenizer set the KeywordTokenizer property for TokenizerDefinitionBuilder.
func (u *TokenizerDefinitionBuilder) KeywordTokenizer(v KeywordTokenizer) *TokenizerDefinitionBuilder {
	u.v = v
	return u
}

// KuromojiTokenizer set the KuromojiTokenizer property for TokenizerDefinitionBuilder.
func (u *TokenizerDefinitionBuilder) KuromojiTokenizer(v KuromojiTokenizer) *TokenizerDefinitionBuilder {
	u.v = v
	return u
}

// LetterTokenizer set the LetterTokenizer property for TokenizerDefinitionBuilder.
func (u *TokenizerDefinitionBuilder) LetterTokenizer(v LetterTokenizer) *TokenizerDefinitionBuilder {
	u.v = v
	return u
}

// LowercaseTokenizer set the LowercaseTokenizer property for TokenizerDefinitionBuilder.
func (u *TokenizerDefinitionBuilder) LowercaseTokenizer(v LowercaseTokenizer) *TokenizerDefinitionBuilder {
	u.v = v
	return u
}

// NGramTokenizer set the NGramTokenizer property for TokenizerDefinitionBuilder.
func (u *TokenizerDefinitionBuilder) NGramTokenizer(v NGramTokenizer) *TokenizerDefinitionBuilder {
	u.v = v
	return u
}

// NoriTokenizer set the NoriTokenizer property for TokenizerDefinitionBuilder.
func (u *TokenizerDefinitionBuilder) NoriTokenizer(v NoriTokenizer) *TokenizerDefinitionBuilder {
	u.v = v
	return u
}

// PathHierarchyTokenizer set the PathHierarchyTokenizer property for TokenizerDefinitionBuilder.
func (u *TokenizerDefinitionBuilder) PathHierarchyTokenizer(v PathHierarchyTokenizer) *TokenizerDefinitionBuilder {
	u.v = v
	return u
}

// PatternTokenizer set the PatternTokenizer property for TokenizerDefinitionBuilder.
func (u *TokenizerDefinitionBuilder) PatternTokenizer(v PatternTokenizer) *TokenizerDefinitionBuilder {
	u.v = v
	return u
}

// StandardTokenizer set the StandardTokenizer property for TokenizerDefinitionBuilder.
func (u *TokenizerDefinitionBuilder) StandardTokenizer(v StandardTokenizer) *TokenizerDefinitionBuilder {
	u.v = v
	return u
}

// UaxEmailUrlTokenizer set the UaxEmailUrlTokenizer property for TokenizerDefinitionBuilder.
func (u *TokenizerDefinitionBuilder) UaxEmailUrlTokenizer(v UaxEmailUrlTokenizer) *TokenizerDefinitionBuilder {
	u.v = v
	return u
}

// WhitespaceTokenizer set the WhitespaceTokenizer property for TokenizerDefinitionBuilder.
func (u *TokenizerDefinitionBuilder) WhitespaceTokenizer(v WhitespaceTokenizer) *TokenizerDefinitionBuilder {
	u.v = v
	return u
}
