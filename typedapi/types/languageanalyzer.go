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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/language"
)

// LanguageAnalyzer type.
type LanguageAnalyzer struct {
	Language      language.Language `json:"language"`
	StemExclusion []string          `json:"stem_exclusion"`
	Stopwords     *StopWords        `json:"stopwords,omitempty"`
	StopwordsPath *string           `json:"stopwords_path,omitempty"`
	Type          string            `json:"type,omitempty"`
	Version       *VersionString    `json:"version,omitempty"`
}

// LanguageAnalyzerBuilder holds LanguageAnalyzer struct and provides a builder API.
type LanguageAnalyzerBuilder struct {
	v *LanguageAnalyzer
}

// NewLanguageAnalyzer provides a builder for the LanguageAnalyzer struct.
func NewLanguageAnalyzer() *LanguageAnalyzerBuilder {
	r := LanguageAnalyzerBuilder{
		&LanguageAnalyzer{},
	}

	r.v.Type = "language"

	return &r
}

// Build finalize the chain and returns the LanguageAnalyzer struct
func (rb *LanguageAnalyzerBuilder) Build() LanguageAnalyzer {
	return *rb.v
}

// Language set the Language property for LanguageAnalyzerBuilder.
func (rb *LanguageAnalyzerBuilder) Language(language language.Language) *LanguageAnalyzerBuilder {
	rb.v.Language = language
	return rb
}

// StemExclusion set the StemExclusion property for LanguageAnalyzerBuilder.
func (rb *LanguageAnalyzerBuilder) StemExclusion(stem_exclusion ...string) *LanguageAnalyzerBuilder {
	rb.v.StemExclusion = append(rb.v.StemExclusion, stem_exclusion...)
	return rb
}

// Stopwords set the Stopwords property for LanguageAnalyzerBuilder.
func (rb *LanguageAnalyzerBuilder) Stopwords(stopwords StopWords) *LanguageAnalyzerBuilder {
	rb.v.Stopwords = &stopwords
	return rb
}

// StopwordsPath set the StopwordsPath property for LanguageAnalyzerBuilder.
func (rb *LanguageAnalyzerBuilder) StopwordsPath(stopwordspath string) *LanguageAnalyzerBuilder {
	rb.v.StopwordsPath = &stopwordspath
	return rb
}

// Type set the Type property for LanguageAnalyzerBuilder.

// Version set the Version property for LanguageAnalyzerBuilder.
func (rb *LanguageAnalyzerBuilder) Version(version VersionString) *LanguageAnalyzerBuilder {
	rb.v.Version = &version
	return rb
}
