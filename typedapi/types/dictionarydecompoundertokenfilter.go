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

// DictionaryDecompounderTokenFilter type.
type DictionaryDecompounderTokenFilter struct {
	HyphenationPatternsPath *string        `json:"hyphenation_patterns_path,omitempty"`
	MaxSubwordSize          *int           `json:"max_subword_size,omitempty"`
	MinSubwordSize          *int           `json:"min_subword_size,omitempty"`
	MinWordSize             *int           `json:"min_word_size,omitempty"`
	OnlyLongestMatch        *bool          `json:"only_longest_match,omitempty"`
	Type                    string         `json:"type,omitempty"`
	Version                 *VersionString `json:"version,omitempty"`
	WordList                []string       `json:"word_list,omitempty"`
	WordListPath            *string        `json:"word_list_path,omitempty"`
}

// DictionaryDecompounderTokenFilterBuilder holds DictionaryDecompounderTokenFilter struct and provides a builder API.
type DictionaryDecompounderTokenFilterBuilder struct {
	v *DictionaryDecompounderTokenFilter
}

// NewDictionaryDecompounderTokenFilter provides a builder for the DictionaryDecompounderTokenFilter struct.
func NewDictionaryDecompounderTokenFilter() *DictionaryDecompounderTokenFilterBuilder {
	r := DictionaryDecompounderTokenFilterBuilder{
		&DictionaryDecompounderTokenFilter{},
	}

	r.v.Type = "dictionary_decompounder"

	return &r
}

// Build finalize the chain and returns the DictionaryDecompounderTokenFilter struct
func (rb *DictionaryDecompounderTokenFilterBuilder) Build() DictionaryDecompounderTokenFilter {
	return *rb.v
}

// HyphenationPatternsPath set the HyphenationPatternsPath property for DictionaryDecompounderTokenFilterBuilder.
func (rb *DictionaryDecompounderTokenFilterBuilder) HyphenationPatternsPath(hyphenationpatternspath string) *DictionaryDecompounderTokenFilterBuilder {
	rb.v.HyphenationPatternsPath = &hyphenationpatternspath
	return rb
}

// MaxSubwordSize set the MaxSubwordSize property for DictionaryDecompounderTokenFilterBuilder.
func (rb *DictionaryDecompounderTokenFilterBuilder) MaxSubwordSize(maxsubwordsize int) *DictionaryDecompounderTokenFilterBuilder {
	rb.v.MaxSubwordSize = &maxsubwordsize
	return rb
}

// MinSubwordSize set the MinSubwordSize property for DictionaryDecompounderTokenFilterBuilder.
func (rb *DictionaryDecompounderTokenFilterBuilder) MinSubwordSize(minsubwordsize int) *DictionaryDecompounderTokenFilterBuilder {
	rb.v.MinSubwordSize = &minsubwordsize
	return rb
}

// MinWordSize set the MinWordSize property for DictionaryDecompounderTokenFilterBuilder.
func (rb *DictionaryDecompounderTokenFilterBuilder) MinWordSize(minwordsize int) *DictionaryDecompounderTokenFilterBuilder {
	rb.v.MinWordSize = &minwordsize
	return rb
}

// OnlyLongestMatch set the OnlyLongestMatch property for DictionaryDecompounderTokenFilterBuilder.
func (rb *DictionaryDecompounderTokenFilterBuilder) OnlyLongestMatch(onlylongestmatch bool) *DictionaryDecompounderTokenFilterBuilder {
	rb.v.OnlyLongestMatch = &onlylongestmatch
	return rb
}

// Type set the Type property for DictionaryDecompounderTokenFilterBuilder.

// Version set the Version property for DictionaryDecompounderTokenFilterBuilder.
func (rb *DictionaryDecompounderTokenFilterBuilder) Version(version VersionString) *DictionaryDecompounderTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}

// WordList set the WordList property for DictionaryDecompounderTokenFilterBuilder.
func (rb *DictionaryDecompounderTokenFilterBuilder) WordList(word_list ...string) *DictionaryDecompounderTokenFilterBuilder {
	rb.v.WordList = append(rb.v.WordList, word_list...)
	return rb
}

// WordListPath set the WordListPath property for DictionaryDecompounderTokenFilterBuilder.
func (rb *DictionaryDecompounderTokenFilterBuilder) WordListPath(wordlistpath string) *DictionaryDecompounderTokenFilterBuilder {
	rb.v.WordListPath = &wordlistpath
	return rb
}
