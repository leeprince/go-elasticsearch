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

// CommonGramsTokenFilter type.
type CommonGramsTokenFilter struct {
	CommonWords     []string       `json:"common_words,omitempty"`
	CommonWordsPath *string        `json:"common_words_path,omitempty"`
	IgnoreCase      *bool          `json:"ignore_case,omitempty"`
	QueryMode       *bool          `json:"query_mode,omitempty"`
	Type            string         `json:"type,omitempty"`
	Version         *VersionString `json:"version,omitempty"`
}

// CommonGramsTokenFilterBuilder holds CommonGramsTokenFilter struct and provides a builder API.
type CommonGramsTokenFilterBuilder struct {
	v *CommonGramsTokenFilter
}

// NewCommonGramsTokenFilter provides a builder for the CommonGramsTokenFilter struct.
func NewCommonGramsTokenFilter() *CommonGramsTokenFilterBuilder {
	r := CommonGramsTokenFilterBuilder{
		&CommonGramsTokenFilter{},
	}

	r.v.Type = "common_grams"

	return &r
}

// Build finalize the chain and returns the CommonGramsTokenFilter struct
func (rb *CommonGramsTokenFilterBuilder) Build() CommonGramsTokenFilter {
	return *rb.v
}

// CommonWords set the CommonWords property for CommonGramsTokenFilterBuilder.
func (rb *CommonGramsTokenFilterBuilder) CommonWords(common_words ...string) *CommonGramsTokenFilterBuilder {
	rb.v.CommonWords = append(rb.v.CommonWords, common_words...)
	return rb
}

// CommonWordsPath set the CommonWordsPath property for CommonGramsTokenFilterBuilder.
func (rb *CommonGramsTokenFilterBuilder) CommonWordsPath(commonwordspath string) *CommonGramsTokenFilterBuilder {
	rb.v.CommonWordsPath = &commonwordspath
	return rb
}

// IgnoreCase set the IgnoreCase property for CommonGramsTokenFilterBuilder.
func (rb *CommonGramsTokenFilterBuilder) IgnoreCase(ignorecase bool) *CommonGramsTokenFilterBuilder {
	rb.v.IgnoreCase = &ignorecase
	return rb
}

// QueryMode set the QueryMode property for CommonGramsTokenFilterBuilder.
func (rb *CommonGramsTokenFilterBuilder) QueryMode(querymode bool) *CommonGramsTokenFilterBuilder {
	rb.v.QueryMode = &querymode
	return rb
}

// Type set the Type property for CommonGramsTokenFilterBuilder.

// Version set the Version property for CommonGramsTokenFilterBuilder.
func (rb *CommonGramsTokenFilterBuilder) Version(version VersionString) *CommonGramsTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
