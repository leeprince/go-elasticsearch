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

// CharFilterTypes type.
type CharFilterTypes struct {
	AnalyzerTypes      []FieldTypes `json:"analyzer_types"`
	BuiltInAnalyzers   []FieldTypes `json:"built_in_analyzers"`
	BuiltInCharFilters []FieldTypes `json:"built_in_char_filters"`
	BuiltInFilters     []FieldTypes `json:"built_in_filters"`
	BuiltInTokenizers  []FieldTypes `json:"built_in_tokenizers"`
	CharFilterTypes    []FieldTypes `json:"char_filter_types"`
	FilterTypes        []FieldTypes `json:"filter_types"`
	TokenizerTypes     []FieldTypes `json:"tokenizer_types"`
}

// CharFilterTypesBuilder holds CharFilterTypes struct and provides a builder API.
type CharFilterTypesBuilder struct {
	v *CharFilterTypes
}

// NewCharFilterTypes provides a builder for the CharFilterTypes struct.
func NewCharFilterTypes() *CharFilterTypesBuilder {
	r := CharFilterTypesBuilder{
		&CharFilterTypes{},
	}

	return &r
}

// Build finalize the chain and returns the CharFilterTypes struct
func (rb *CharFilterTypesBuilder) Build() CharFilterTypes {
	return *rb.v
}

// AnalyzerTypes set the AnalyzerTypes property for CharFilterTypesBuilder.
func (rb *CharFilterTypesBuilder) AnalyzerTypes(analyzer_types ...FieldTypes) *CharFilterTypesBuilder {
	rb.v.AnalyzerTypes = append(rb.v.AnalyzerTypes, analyzer_types...)
	return rb
}

// BuiltInAnalyzers set the BuiltInAnalyzers property for CharFilterTypesBuilder.
func (rb *CharFilterTypesBuilder) BuiltInAnalyzers(built_in_analyzers ...FieldTypes) *CharFilterTypesBuilder {
	rb.v.BuiltInAnalyzers = append(rb.v.BuiltInAnalyzers, built_in_analyzers...)
	return rb
}

// BuiltInCharFilters set the BuiltInCharFilters property for CharFilterTypesBuilder.
func (rb *CharFilterTypesBuilder) BuiltInCharFilters(built_in_char_filters ...FieldTypes) *CharFilterTypesBuilder {
	rb.v.BuiltInCharFilters = append(rb.v.BuiltInCharFilters, built_in_char_filters...)
	return rb
}

// BuiltInFilters set the BuiltInFilters property for CharFilterTypesBuilder.
func (rb *CharFilterTypesBuilder) BuiltInFilters(built_in_filters ...FieldTypes) *CharFilterTypesBuilder {
	rb.v.BuiltInFilters = append(rb.v.BuiltInFilters, built_in_filters...)
	return rb
}

// BuiltInTokenizers set the BuiltInTokenizers property for CharFilterTypesBuilder.
func (rb *CharFilterTypesBuilder) BuiltInTokenizers(built_in_tokenizers ...FieldTypes) *CharFilterTypesBuilder {
	rb.v.BuiltInTokenizers = append(rb.v.BuiltInTokenizers, built_in_tokenizers...)
	return rb
}

// CharFilterTypes set the CharFilterTypes property for CharFilterTypesBuilder.
func (rb *CharFilterTypesBuilder) CharFilterTypes(char_filter_types ...FieldTypes) *CharFilterTypesBuilder {
	rb.v.CharFilterTypes = append(rb.v.CharFilterTypes, char_filter_types...)
	return rb
}

// FilterTypes set the FilterTypes property for CharFilterTypesBuilder.
func (rb *CharFilterTypesBuilder) FilterTypes(filter_types ...FieldTypes) *CharFilterTypesBuilder {
	rb.v.FilterTypes = append(rb.v.FilterTypes, filter_types...)
	return rb
}

// TokenizerTypes set the TokenizerTypes property for CharFilterTypesBuilder.
func (rb *CharFilterTypesBuilder) TokenizerTypes(tokenizer_types ...FieldTypes) *CharFilterTypesBuilder {
	rb.v.TokenizerTypes = append(rb.v.TokenizerTypes, tokenizer_types...)
	return rb
}
