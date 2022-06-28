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

// KuromojiPartOfSpeechTokenFilter type.
type KuromojiPartOfSpeechTokenFilter struct {
	Stoptags []string       `json:"stoptags"`
	Type     string         `json:"type,omitempty"`
	Version  *VersionString `json:"version,omitempty"`
}

// KuromojiPartOfSpeechTokenFilterBuilder holds KuromojiPartOfSpeechTokenFilter struct and provides a builder API.
type KuromojiPartOfSpeechTokenFilterBuilder struct {
	v *KuromojiPartOfSpeechTokenFilter
}

// NewKuromojiPartOfSpeechTokenFilter provides a builder for the KuromojiPartOfSpeechTokenFilter struct.
func NewKuromojiPartOfSpeechTokenFilter() *KuromojiPartOfSpeechTokenFilterBuilder {
	r := KuromojiPartOfSpeechTokenFilterBuilder{
		&KuromojiPartOfSpeechTokenFilter{},
	}

	r.v.Type = "kuromoji_part_of_speech"

	return &r
}

// Build finalize the chain and returns the KuromojiPartOfSpeechTokenFilter struct
func (rb *KuromojiPartOfSpeechTokenFilterBuilder) Build() KuromojiPartOfSpeechTokenFilter {
	return *rb.v
}

// Stoptags set the Stoptags property for KuromojiPartOfSpeechTokenFilterBuilder.
func (rb *KuromojiPartOfSpeechTokenFilterBuilder) Stoptags(stoptags ...string) *KuromojiPartOfSpeechTokenFilterBuilder {
	rb.v.Stoptags = append(rb.v.Stoptags, stoptags...)
	return rb
}

// Type set the Type property for KuromojiPartOfSpeechTokenFilterBuilder.

// Version set the Version property for KuromojiPartOfSpeechTokenFilterBuilder.
func (rb *KuromojiPartOfSpeechTokenFilterBuilder) Version(version VersionString) *KuromojiPartOfSpeechTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
