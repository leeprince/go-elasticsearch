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

// IndexSettingsAnalysis type.
type IndexSettingsAnalysis struct {
	Analyzer   map[string]Analyzer    `json:"analyzer,omitempty"`
	CharFilter map[string]CharFilter  `json:"char_filter,omitempty"`
	Filter     map[string]TokenFilter `json:"filter,omitempty"`
	Normalizer map[string]Normalizer  `json:"normalizer,omitempty"`
	Tokenizer  map[string]Tokenizer   `json:"tokenizer,omitempty"`
}

// IndexSettingsAnalysisBuilder holds IndexSettingsAnalysis struct and provides a builder API.
type IndexSettingsAnalysisBuilder struct {
	v *IndexSettingsAnalysis
}

// NewIndexSettingsAnalysis provides a builder for the IndexSettingsAnalysis struct.
func NewIndexSettingsAnalysis() *IndexSettingsAnalysisBuilder {
	r := IndexSettingsAnalysisBuilder{
		&IndexSettingsAnalysis{
			Analyzer:   make(map[string]Analyzer, 0),
			CharFilter: make(map[string]CharFilter, 0),
			Filter:     make(map[string]TokenFilter, 0),
			Normalizer: make(map[string]Normalizer, 0),
			Tokenizer:  make(map[string]Tokenizer, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the IndexSettingsAnalysis struct
func (rb *IndexSettingsAnalysisBuilder) Build() IndexSettingsAnalysis {
	return *rb.v
}

// Analyzer set the Analyzer property for IndexSettingsAnalysisBuilder.
func (rb *IndexSettingsAnalysisBuilder) Analyzer(value map[string]Analyzer) *IndexSettingsAnalysisBuilder {
	rb.v.Analyzer = value
	return rb
}

// CharFilter set the CharFilter property for IndexSettingsAnalysisBuilder.
func (rb *IndexSettingsAnalysisBuilder) CharFilter(value map[string]CharFilter) *IndexSettingsAnalysisBuilder {
	rb.v.CharFilter = value
	return rb
}

// Filter set the Filter property for IndexSettingsAnalysisBuilder.
func (rb *IndexSettingsAnalysisBuilder) Filter(value map[string]TokenFilter) *IndexSettingsAnalysisBuilder {
	rb.v.Filter = value
	return rb
}

// Normalizer set the Normalizer property for IndexSettingsAnalysisBuilder.
func (rb *IndexSettingsAnalysisBuilder) Normalizer(value map[string]Normalizer) *IndexSettingsAnalysisBuilder {
	rb.v.Normalizer = value
	return rb
}

// Tokenizer set the Tokenizer property for IndexSettingsAnalysisBuilder.
func (rb *IndexSettingsAnalysisBuilder) Tokenizer(value map[string]Tokenizer) *IndexSettingsAnalysisBuilder {
	rb.v.Tokenizer = value
	return rb
}
