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

// KuromojiStemmerTokenFilter type.
type KuromojiStemmerTokenFilter struct {
	MinimumLength int            `json:"minimum_length"`
	Type          string         `json:"type,omitempty"`
	Version       *VersionString `json:"version,omitempty"`
}

// KuromojiStemmerTokenFilterBuilder holds KuromojiStemmerTokenFilter struct and provides a builder API.
type KuromojiStemmerTokenFilterBuilder struct {
	v *KuromojiStemmerTokenFilter
}

// NewKuromojiStemmerTokenFilter provides a builder for the KuromojiStemmerTokenFilter struct.
func NewKuromojiStemmerTokenFilter() *KuromojiStemmerTokenFilterBuilder {
	r := KuromojiStemmerTokenFilterBuilder{
		&KuromojiStemmerTokenFilter{},
	}

	r.v.Type = "kuromoji_stemmer"

	return &r
}

// Build finalize the chain and returns the KuromojiStemmerTokenFilter struct
func (rb *KuromojiStemmerTokenFilterBuilder) Build() KuromojiStemmerTokenFilter {
	return *rb.v
}

// MinimumLength set the MinimumLength property for KuromojiStemmerTokenFilterBuilder.
func (rb *KuromojiStemmerTokenFilterBuilder) MinimumLength(minimumlength int) *KuromojiStemmerTokenFilterBuilder {
	rb.v.MinimumLength = minimumlength
	return rb
}

// Type set the Type property for KuromojiStemmerTokenFilterBuilder.

// Version set the Version property for KuromojiStemmerTokenFilterBuilder.
func (rb *KuromojiStemmerTokenFilterBuilder) Version(version VersionString) *KuromojiStemmerTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
