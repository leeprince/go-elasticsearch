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

// SettingsSimilarity type.
type SettingsSimilarity struct {
	Bm25          *SettingsSimilarityBm25          `json:"bm25,omitempty"`
	Dfi           *SettingsSimilarityDfi           `json:"dfi,omitempty"`
	Dfr           *SettingsSimilarityDfr           `json:"dfr,omitempty"`
	Ib            *SettingsSimilarityIb            `json:"ib,omitempty"`
	Lmd           *SettingsSimilarityLmd           `json:"lmd,omitempty"`
	Lmj           *SettingsSimilarityLmj           `json:"lmj,omitempty"`
	ScriptedTfidf *SettingsSimilarityScriptedTfidf `json:"scripted_tfidf,omitempty"`
}

// SettingsSimilarityBuilder holds SettingsSimilarity struct and provides a builder API.
type SettingsSimilarityBuilder struct {
	v *SettingsSimilarity
}

// NewSettingsSimilarity provides a builder for the SettingsSimilarity struct.
func NewSettingsSimilarity() *SettingsSimilarityBuilder {
	r := SettingsSimilarityBuilder{
		&SettingsSimilarity{},
	}

	return &r
}

// Build finalize the chain and returns the SettingsSimilarity struct
func (rb *SettingsSimilarityBuilder) Build() SettingsSimilarity {
	return *rb.v
}

// Bm25 set the Bm25 property for SettingsSimilarityBuilder.
func (rb *SettingsSimilarityBuilder) Bm25(bm25 SettingsSimilarityBm25) *SettingsSimilarityBuilder {
	rb.v.Bm25 = &bm25
	return rb
}

// Dfi set the Dfi property for SettingsSimilarityBuilder.
func (rb *SettingsSimilarityBuilder) Dfi(dfi SettingsSimilarityDfi) *SettingsSimilarityBuilder {
	rb.v.Dfi = &dfi
	return rb
}

// Dfr set the Dfr property for SettingsSimilarityBuilder.
func (rb *SettingsSimilarityBuilder) Dfr(dfr SettingsSimilarityDfr) *SettingsSimilarityBuilder {
	rb.v.Dfr = &dfr
	return rb
}

// Ib set the Ib property for SettingsSimilarityBuilder.
func (rb *SettingsSimilarityBuilder) Ib(ib SettingsSimilarityIb) *SettingsSimilarityBuilder {
	rb.v.Ib = &ib
	return rb
}

// Lmd set the Lmd property for SettingsSimilarityBuilder.
func (rb *SettingsSimilarityBuilder) Lmd(lmd SettingsSimilarityLmd) *SettingsSimilarityBuilder {
	rb.v.Lmd = &lmd
	return rb
}

// Lmj set the Lmj property for SettingsSimilarityBuilder.
func (rb *SettingsSimilarityBuilder) Lmj(lmj SettingsSimilarityLmj) *SettingsSimilarityBuilder {
	rb.v.Lmj = &lmj
	return rb
}

// ScriptedTfidf set the ScriptedTfidf property for SettingsSimilarityBuilder.
func (rb *SettingsSimilarityBuilder) ScriptedTfidf(scriptedtfidf SettingsSimilarityScriptedTfidf) *SettingsSimilarityBuilder {
	rb.v.ScriptedTfidf = &scriptedtfidf
	return rb
}
