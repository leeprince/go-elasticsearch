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

// AnalyzeDetail type.
type AnalyzeDetail struct {
	Analyzer       *AnalyzerDetail    `json:"analyzer,omitempty"`
	Charfilters    []CharFilterDetail `json:"charfilters,omitempty"`
	CustomAnalyzer bool               `json:"custom_analyzer"`
	Tokenfilters   []TokenDetail      `json:"tokenfilters,omitempty"`
	Tokenizer      *TokenDetail       `json:"tokenizer,omitempty"`
}

// AnalyzeDetailBuilder holds AnalyzeDetail struct and provides a builder API.
type AnalyzeDetailBuilder struct {
	v *AnalyzeDetail
}

// NewAnalyzeDetail provides a builder for the AnalyzeDetail struct.
func NewAnalyzeDetail() *AnalyzeDetailBuilder {
	r := AnalyzeDetailBuilder{
		&AnalyzeDetail{},
	}

	return &r
}

// Build finalize the chain and returns the AnalyzeDetail struct
func (rb *AnalyzeDetailBuilder) Build() AnalyzeDetail {
	return *rb.v
}

// Analyzer set the Analyzer property for AnalyzeDetailBuilder.
func (rb *AnalyzeDetailBuilder) Analyzer(analyzer AnalyzerDetail) *AnalyzeDetailBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

// Charfilters set the Charfilters property for AnalyzeDetailBuilder.
func (rb *AnalyzeDetailBuilder) Charfilters(charfilters ...CharFilterDetail) *AnalyzeDetailBuilder {
	rb.v.Charfilters = append(rb.v.Charfilters, charfilters...)
	return rb
}

// CustomAnalyzer set the CustomAnalyzer property for AnalyzeDetailBuilder.
func (rb *AnalyzeDetailBuilder) CustomAnalyzer(customanalyzer bool) *AnalyzeDetailBuilder {
	rb.v.CustomAnalyzer = customanalyzer
	return rb
}

// Tokenfilters set the Tokenfilters property for AnalyzeDetailBuilder.
func (rb *AnalyzeDetailBuilder) Tokenfilters(tokenfilters ...TokenDetail) *AnalyzeDetailBuilder {
	rb.v.Tokenfilters = append(rb.v.Tokenfilters, tokenfilters...)
	return rb
}

// Tokenizer set the Tokenizer property for AnalyzeDetailBuilder.
func (rb *AnalyzeDetailBuilder) Tokenizer(tokenizer TokenDetail) *AnalyzeDetailBuilder {
	rb.v.Tokenizer = &tokenizer
	return rb
}
