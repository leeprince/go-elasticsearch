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

// KuromojiIterationMarkCharFilter type.
type KuromojiIterationMarkCharFilter struct {
	NormalizeKana  bool           `json:"normalize_kana"`
	NormalizeKanji bool           `json:"normalize_kanji"`
	Type           string         `json:"type,omitempty"`
	Version        *VersionString `json:"version,omitempty"`
}

// KuromojiIterationMarkCharFilterBuilder holds KuromojiIterationMarkCharFilter struct and provides a builder API.
type KuromojiIterationMarkCharFilterBuilder struct {
	v *KuromojiIterationMarkCharFilter
}

// NewKuromojiIterationMarkCharFilter provides a builder for the KuromojiIterationMarkCharFilter struct.
func NewKuromojiIterationMarkCharFilter() *KuromojiIterationMarkCharFilterBuilder {
	r := KuromojiIterationMarkCharFilterBuilder{
		&KuromojiIterationMarkCharFilter{},
	}

	r.v.Type = "kuromoji_iteration_mark"

	return &r
}

// Build finalize the chain and returns the KuromojiIterationMarkCharFilter struct
func (rb *KuromojiIterationMarkCharFilterBuilder) Build() KuromojiIterationMarkCharFilter {
	return *rb.v
}

// NormalizeKana set the NormalizeKana property for KuromojiIterationMarkCharFilterBuilder.
func (rb *KuromojiIterationMarkCharFilterBuilder) NormalizeKana(normalizekana bool) *KuromojiIterationMarkCharFilterBuilder {
	rb.v.NormalizeKana = normalizekana
	return rb
}

// NormalizeKanji set the NormalizeKanji property for KuromojiIterationMarkCharFilterBuilder.
func (rb *KuromojiIterationMarkCharFilterBuilder) NormalizeKanji(normalizekanji bool) *KuromojiIterationMarkCharFilterBuilder {
	rb.v.NormalizeKanji = normalizekanji
	return rb
}

// Type set the Type property for KuromojiIterationMarkCharFilterBuilder.

// Version set the Version property for KuromojiIterationMarkCharFilterBuilder.
func (rb *KuromojiIterationMarkCharFilterBuilder) Version(version VersionString) *KuromojiIterationMarkCharFilterBuilder {
	rb.v.Version = &version
	return rb
}
