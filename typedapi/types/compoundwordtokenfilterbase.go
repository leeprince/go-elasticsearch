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

// CompoundWordTokenFilterBase type.
type CompoundWordTokenFilterBase struct {
	HyphenationPatternsPath *string        `json:"hyphenation_patterns_path,omitempty"`
	MaxSubwordSize          *int           `json:"max_subword_size,omitempty"`
	MinSubwordSize          *int           `json:"min_subword_size,omitempty"`
	MinWordSize             *int           `json:"min_word_size,omitempty"`
	OnlyLongestMatch        *bool          `json:"only_longest_match,omitempty"`
	Version                 *VersionString `json:"version,omitempty"`
	WordList                []string       `json:"word_list,omitempty"`
	WordListPath            *string        `json:"word_list_path,omitempty"`
}

// CompoundWordTokenFilterBaseBuilder holds CompoundWordTokenFilterBase struct and provides a builder API.
type CompoundWordTokenFilterBaseBuilder struct {
	v *CompoundWordTokenFilterBase
}

// NewCompoundWordTokenFilterBase provides a builder for the CompoundWordTokenFilterBase struct.
func NewCompoundWordTokenFilterBase() *CompoundWordTokenFilterBaseBuilder {
	r := CompoundWordTokenFilterBaseBuilder{
		&CompoundWordTokenFilterBase{},
	}

	return &r
}

// Build finalize the chain and returns the CompoundWordTokenFilterBase struct
func (rb *CompoundWordTokenFilterBaseBuilder) Build() CompoundWordTokenFilterBase {
	return *rb.v
}

// HyphenationPatternsPath set the HyphenationPatternsPath property for CompoundWordTokenFilterBaseBuilder.
func (rb *CompoundWordTokenFilterBaseBuilder) HyphenationPatternsPath(hyphenationpatternspath string) *CompoundWordTokenFilterBaseBuilder {
	rb.v.HyphenationPatternsPath = &hyphenationpatternspath
	return rb
}

// MaxSubwordSize set the MaxSubwordSize property for CompoundWordTokenFilterBaseBuilder.
func (rb *CompoundWordTokenFilterBaseBuilder) MaxSubwordSize(maxsubwordsize int) *CompoundWordTokenFilterBaseBuilder {
	rb.v.MaxSubwordSize = &maxsubwordsize
	return rb
}

// MinSubwordSize set the MinSubwordSize property for CompoundWordTokenFilterBaseBuilder.
func (rb *CompoundWordTokenFilterBaseBuilder) MinSubwordSize(minsubwordsize int) *CompoundWordTokenFilterBaseBuilder {
	rb.v.MinSubwordSize = &minsubwordsize
	return rb
}

// MinWordSize set the MinWordSize property for CompoundWordTokenFilterBaseBuilder.
func (rb *CompoundWordTokenFilterBaseBuilder) MinWordSize(minwordsize int) *CompoundWordTokenFilterBaseBuilder {
	rb.v.MinWordSize = &minwordsize
	return rb
}

// OnlyLongestMatch set the OnlyLongestMatch property for CompoundWordTokenFilterBaseBuilder.
func (rb *CompoundWordTokenFilterBaseBuilder) OnlyLongestMatch(onlylongestmatch bool) *CompoundWordTokenFilterBaseBuilder {
	rb.v.OnlyLongestMatch = &onlylongestmatch
	return rb
}

// Version set the Version property for CompoundWordTokenFilterBaseBuilder.
func (rb *CompoundWordTokenFilterBaseBuilder) Version(version VersionString) *CompoundWordTokenFilterBaseBuilder {
	rb.v.Version = &version
	return rb
}

// WordList set the WordList property for CompoundWordTokenFilterBaseBuilder.
func (rb *CompoundWordTokenFilterBaseBuilder) WordList(word_list ...string) *CompoundWordTokenFilterBaseBuilder {
	rb.v.WordList = append(rb.v.WordList, word_list...)
	return rb
}

// WordListPath set the WordListPath property for CompoundWordTokenFilterBaseBuilder.
func (rb *CompoundWordTokenFilterBaseBuilder) WordListPath(wordlistpath string) *CompoundWordTokenFilterBaseBuilder {
	rb.v.WordListPath = &wordlistpath
	return rb
}
