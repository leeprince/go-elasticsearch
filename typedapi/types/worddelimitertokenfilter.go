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

// WordDelimiterTokenFilter type.
type WordDelimiterTokenFilter struct {
	CatenateAll           *bool          `json:"catenate_all,omitempty"`
	CatenateNumbers       *bool          `json:"catenate_numbers,omitempty"`
	CatenateWords         *bool          `json:"catenate_words,omitempty"`
	GenerateNumberParts   *bool          `json:"generate_number_parts,omitempty"`
	GenerateWordParts     *bool          `json:"generate_word_parts,omitempty"`
	PreserveOriginal      *bool          `json:"preserve_original,omitempty"`
	ProtectedWords        []string       `json:"protected_words,omitempty"`
	ProtectedWordsPath    *string        `json:"protected_words_path,omitempty"`
	SplitOnCaseChange     *bool          `json:"split_on_case_change,omitempty"`
	SplitOnNumerics       *bool          `json:"split_on_numerics,omitempty"`
	StemEnglishPossessive *bool          `json:"stem_english_possessive,omitempty"`
	Type                  string         `json:"type,omitempty"`
	TypeTable             []string       `json:"type_table,omitempty"`
	TypeTablePath         *string        `json:"type_table_path,omitempty"`
	Version               *VersionString `json:"version,omitempty"`
}

// WordDelimiterTokenFilterBuilder holds WordDelimiterTokenFilter struct and provides a builder API.
type WordDelimiterTokenFilterBuilder struct {
	v *WordDelimiterTokenFilter
}

// NewWordDelimiterTokenFilter provides a builder for the WordDelimiterTokenFilter struct.
func NewWordDelimiterTokenFilter() *WordDelimiterTokenFilterBuilder {
	r := WordDelimiterTokenFilterBuilder{
		&WordDelimiterTokenFilter{},
	}

	r.v.Type = "word_delimiter"

	return &r
}

// Build finalize the chain and returns the WordDelimiterTokenFilter struct
func (rb *WordDelimiterTokenFilterBuilder) Build() WordDelimiterTokenFilter {
	return *rb.v
}

// CatenateAll set the CatenateAll property for WordDelimiterTokenFilterBuilder.
func (rb *WordDelimiterTokenFilterBuilder) CatenateAll(catenateall bool) *WordDelimiterTokenFilterBuilder {
	rb.v.CatenateAll = &catenateall
	return rb
}

// CatenateNumbers set the CatenateNumbers property for WordDelimiterTokenFilterBuilder.
func (rb *WordDelimiterTokenFilterBuilder) CatenateNumbers(catenatenumbers bool) *WordDelimiterTokenFilterBuilder {
	rb.v.CatenateNumbers = &catenatenumbers
	return rb
}

// CatenateWords set the CatenateWords property for WordDelimiterTokenFilterBuilder.
func (rb *WordDelimiterTokenFilterBuilder) CatenateWords(catenatewords bool) *WordDelimiterTokenFilterBuilder {
	rb.v.CatenateWords = &catenatewords
	return rb
}

// GenerateNumberParts set the GenerateNumberParts property for WordDelimiterTokenFilterBuilder.
func (rb *WordDelimiterTokenFilterBuilder) GenerateNumberParts(generatenumberparts bool) *WordDelimiterTokenFilterBuilder {
	rb.v.GenerateNumberParts = &generatenumberparts
	return rb
}

// GenerateWordParts set the GenerateWordParts property for WordDelimiterTokenFilterBuilder.
func (rb *WordDelimiterTokenFilterBuilder) GenerateWordParts(generatewordparts bool) *WordDelimiterTokenFilterBuilder {
	rb.v.GenerateWordParts = &generatewordparts
	return rb
}

// PreserveOriginal set the PreserveOriginal property for WordDelimiterTokenFilterBuilder.
func (rb *WordDelimiterTokenFilterBuilder) PreserveOriginal(preserveoriginal bool) *WordDelimiterTokenFilterBuilder {
	rb.v.PreserveOriginal = &preserveoriginal
	return rb
}

// ProtectedWords set the ProtectedWords property for WordDelimiterTokenFilterBuilder.
func (rb *WordDelimiterTokenFilterBuilder) ProtectedWords(protected_words ...string) *WordDelimiterTokenFilterBuilder {
	rb.v.ProtectedWords = append(rb.v.ProtectedWords, protected_words...)
	return rb
}

// ProtectedWordsPath set the ProtectedWordsPath property for WordDelimiterTokenFilterBuilder.
func (rb *WordDelimiterTokenFilterBuilder) ProtectedWordsPath(protectedwordspath string) *WordDelimiterTokenFilterBuilder {
	rb.v.ProtectedWordsPath = &protectedwordspath
	return rb
}

// SplitOnCaseChange set the SplitOnCaseChange property for WordDelimiterTokenFilterBuilder.
func (rb *WordDelimiterTokenFilterBuilder) SplitOnCaseChange(splitoncasechange bool) *WordDelimiterTokenFilterBuilder {
	rb.v.SplitOnCaseChange = &splitoncasechange
	return rb
}

// SplitOnNumerics set the SplitOnNumerics property for WordDelimiterTokenFilterBuilder.
func (rb *WordDelimiterTokenFilterBuilder) SplitOnNumerics(splitonnumerics bool) *WordDelimiterTokenFilterBuilder {
	rb.v.SplitOnNumerics = &splitonnumerics
	return rb
}

// StemEnglishPossessive set the StemEnglishPossessive property for WordDelimiterTokenFilterBuilder.
func (rb *WordDelimiterTokenFilterBuilder) StemEnglishPossessive(stemenglishpossessive bool) *WordDelimiterTokenFilterBuilder {
	rb.v.StemEnglishPossessive = &stemenglishpossessive
	return rb
}

// Type set the Type property for WordDelimiterTokenFilterBuilder.

// TypeTable set the TypeTable property for WordDelimiterTokenFilterBuilder.
func (rb *WordDelimiterTokenFilterBuilder) TypeTable(type_table ...string) *WordDelimiterTokenFilterBuilder {
	rb.v.TypeTable = append(rb.v.TypeTable, type_table...)
	return rb
}

// TypeTablePath set the TypeTablePath property for WordDelimiterTokenFilterBuilder.
func (rb *WordDelimiterTokenFilterBuilder) TypeTablePath(typetablepath string) *WordDelimiterTokenFilterBuilder {
	rb.v.TypeTablePath = &typetablepath
	return rb
}

// Version set the Version property for WordDelimiterTokenFilterBuilder.
func (rb *WordDelimiterTokenFilterBuilder) Version(version VersionString) *WordDelimiterTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
