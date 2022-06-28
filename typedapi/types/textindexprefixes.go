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

// TextIndexPrefixes type.
type TextIndexPrefixes struct {
	MaxChars int `json:"max_chars"`
	MinChars int `json:"min_chars"`
}

// TextIndexPrefixesBuilder holds TextIndexPrefixes struct and provides a builder API.
type TextIndexPrefixesBuilder struct {
	v *TextIndexPrefixes
}

// NewTextIndexPrefixes provides a builder for the TextIndexPrefixes struct.
func NewTextIndexPrefixes() *TextIndexPrefixesBuilder {
	r := TextIndexPrefixesBuilder{
		&TextIndexPrefixes{},
	}

	return &r
}

// Build finalize the chain and returns the TextIndexPrefixes struct
func (rb *TextIndexPrefixesBuilder) Build() TextIndexPrefixes {
	return *rb.v
}

// MaxChars set the MaxChars property for TextIndexPrefixesBuilder.
func (rb *TextIndexPrefixesBuilder) MaxChars(maxchars int) *TextIndexPrefixesBuilder {
	rb.v.MaxChars = maxchars
	return rb
}

// MinChars set the MinChars property for TextIndexPrefixesBuilder.
func (rb *TextIndexPrefixesBuilder) MinChars(minchars int) *TextIndexPrefixesBuilder {
	rb.v.MinChars = minchars
	return rb
}
