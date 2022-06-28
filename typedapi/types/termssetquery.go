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

// TermsSetQuery type.
type TermsSetQuery struct {
	Boost                    *float32 `json:"boost,omitempty"`
	MinimumShouldMatchField  *Field   `json:"minimum_should_match_field,omitempty"`
	MinimumShouldMatchScript *Script  `json:"minimum_should_match_script,omitempty"`
	QueryName_               *string  `json:"_name,omitempty"`
	Terms                    []string `json:"terms"`
}

// TermsSetQueryBuilder holds TermsSetQuery struct and provides a builder API.
type TermsSetQueryBuilder struct {
	v *TermsSetQuery
}

// NewTermsSetQuery provides a builder for the TermsSetQuery struct.
func NewTermsSetQuery() *TermsSetQueryBuilder {
	r := TermsSetQueryBuilder{
		&TermsSetQuery{},
	}

	return &r
}

// Build finalize the chain and returns the TermsSetQuery struct
func (rb *TermsSetQueryBuilder) Build() TermsSetQuery {
	return *rb.v
}

// Boost set the Boost property for TermsSetQueryBuilder.
func (rb *TermsSetQueryBuilder) Boost(boost float32) *TermsSetQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// MinimumShouldMatchField set the MinimumShouldMatchField property for TermsSetQueryBuilder.
func (rb *TermsSetQueryBuilder) MinimumShouldMatchField(minimumshouldmatchfield Field) *TermsSetQueryBuilder {
	rb.v.MinimumShouldMatchField = &minimumshouldmatchfield
	return rb
}

// MinimumShouldMatchScript set the MinimumShouldMatchScript property for TermsSetQueryBuilder.
func (rb *TermsSetQueryBuilder) MinimumShouldMatchScript(minimumshouldmatchscript Script) *TermsSetQueryBuilder {
	rb.v.MinimumShouldMatchScript = &minimumshouldmatchscript
	return rb
}

// QueryName_ set the QueryName_ property for TermsSetQueryBuilder.
func (rb *TermsSetQueryBuilder) QueryName_(queryname_ string) *TermsSetQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// Terms set the Terms property for TermsSetQueryBuilder.
func (rb *TermsSetQueryBuilder) Terms(terms ...string) *TermsSetQueryBuilder {
	rb.v.Terms = append(rb.v.Terms, terms...)
	return rb
}
